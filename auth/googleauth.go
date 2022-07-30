package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var GOOGLE_ACCOUNT_DOMAIN = []string{"accounts.google.com", "https://accounts.google.com"}

var GOOGLE_TOKEN_VALIDATION_URL = "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token="

func ReadJsonData(content []uint8) map[string]interface{} {

	var payload map[string]interface{}
	err := json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal: ", err)
	}

	return payload
}

func VerifyIss(google_account_domain []string, iss string) bool {
	for _, data := range google_account_domain {
		if iss == data {
			return true
		}
	}

	return false
}

func GoogleTokenValidation(token, token_url string) (bool, map[string]interface{}) {
	url := fmt.Sprint(token_url + token)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	payload := ReadJsonData(data)
	iss := payload["iss"]

	if iss == nil {
		err := make(map[string]interface{})
		err["error"] = "Invalid Token"
		return false, err
	}

	valid := VerifyIss(GOOGLE_ACCOUNT_DOMAIN, iss.(string))

	return valid, payload
}

func GoogleAuth(context *fiber.Ctx) error {

	// payload := struct {
	// 	Token string `json:"token"`
	// }{}

	// if err := context.BodyParser(&payload); err != nil {
	// 	fmt.Println(err)
	// }

	// if payload.Token == "" {
	// 	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"detail": "token is empty",
	// 	})
	// }

	// valid, data := GoogleTokenValidation(payload.Token, GOOGLE_TOKEN_VALIDATION_URL)
	// if !valid {
	// 	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"detail": "Invalid Token",
	// 	})
	// }

	// fmt.Println(valid)
	// fmt.Println(data)

	// request header
	token := string(context.Request().Header.Peek("Authorization"))

	if token == "" {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "token is empty",
		})
	}

	valid, data := GoogleTokenValidation(token, GOOGLE_TOKEN_VALIDATION_URL)
	if !valid {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "Invalid Token",
		})
	}

	fmt.Println(valid)
	fmt.Println(data)

	return nil
}

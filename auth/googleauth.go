package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	token_info := strings.Split(token, " ")
	if len(token_info) != 2 {
		err := make(map[string]interface{})
		err["detail"] = "length of the token should be 2"
		return false, err
	}

	if token_info[0] != "Bearer" {
		err := make(map[string]interface{})
		err["detail"] = "Token should be Bearer"
		return false, err
	}

	url := fmt.Sprint(token_url + token_info[1])

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
		err["detail"] = "Invalid Token"
		return false, err
	}

	valid := VerifyIss(GOOGLE_ACCOUNT_DOMAIN, iss.(string))

	return valid, payload
}

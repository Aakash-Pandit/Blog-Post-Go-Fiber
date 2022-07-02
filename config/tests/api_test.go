package tests

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// ToDo
func TestHomePage(t *testing.T) {
	app := fiber.New()
	url := "http://localhost:8080/api/v1"

	request := httptest.NewRequest(fiber.MethodGet, url, nil)
	fmt.Println("request", request)
	resp, err := app.Test(request)
	fmt.Println("response", resp)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode, "Checking Home Page API Status code")
	// utils.AssertEqual(t, "POST, OPTIONS", resp.Header.Get(fiber.HeaderAllow))

}

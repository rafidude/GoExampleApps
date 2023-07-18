package main

import (
	"contact/database"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Assume we have a function to initialize our Fiber app
func setupApp() *fiber.App {
	app := fiber.New()

	app.Post("/save", database.SaveContact)

	return app
}

func TestSaveRoute(t *testing.T) {
	app := setupApp()

	// Construct form data
	formData := url.Values{
		"email":   []string{"test@email.com"},
		"message": []string{"This is a message."},
	}

	req := httptest.NewRequest("POST", "/save", strings.NewReader(formData.Encode()))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)

	resp, err := app.Test(req)

	assert.Nil(t, err)                              // Assert that there was no error
	assert.Equal(t, http.StatusOK, resp.StatusCode) // Assert the HTTP status code is 200

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err) // Assert that there was no error

	assert.Equal(t, "Form submitted successfully!", string(body)) // Assert that the response body matches our expectation
}

package middleware

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLoggerSuccess(t *testing.T) {

	app := fiber.New()

	app.Use(Logger)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	req := httptest.NewRequest(
		"GET",
		"/",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestLoggerError(t *testing.T) {

	app := fiber.New()

	app.Use(Logger)

	app.Get("/", func(c *fiber.Ctx) error {
		return errors.New("handler error")
	})

	req := httptest.NewRequest(
		"GET",
		"/",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.StatusCode != 500 {
		t.Errorf("expected 500, got %d", resp.StatusCode)
	}
}

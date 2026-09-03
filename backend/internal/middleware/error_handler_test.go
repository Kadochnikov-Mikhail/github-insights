package middleware

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github-insights/internal/apperror"

	"github.com/gofiber/fiber/v2"
)

func TestErrorHandlerUsernameRequired(t *testing.T) {

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return apperror.ErrUsernameRequired
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

	if resp.StatusCode != 400 {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}

	body := make(map[string]string)

	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["error"] != apperror.ErrUsernameRequired.Error() {
		t.Errorf(
			"expected %s, got %s",
			apperror.ErrUsernameRequired.Error(),
			body["error"],
		)
	}
}

func TestErrorHandlerUserNotFound(t *testing.T) {

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return apperror.ErrUserNotFound
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

	if resp.StatusCode != 404 {
		t.Errorf("expected 404, got %d", resp.StatusCode)
	}

	body := make(map[string]string)

	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["error"] != apperror.ErrUserNotFound.Error() {
		t.Errorf(
			"expected %s, got %s",
			apperror.ErrUserNotFound.Error(),
			body["error"],
		)
	}
}

func TestErrorHandlerInternalServerError(t *testing.T) {

	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return errors.New("something went wrong")
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

	body := make(map[string]string)

	err = json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body["error"] != "something went wrong" {
		t.Errorf(
			"expected something went wrong, got %s",
			body["error"],
		)
	}
}

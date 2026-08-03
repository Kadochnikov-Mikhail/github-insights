package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github-insights/internal/models"
	"github-insights/internal/apperror"
)

func TestGetUserSuccess(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path != "/users/torvalds" {
				t.Errorf(
					"expected path /users/torvalds, got %s",
					r.URL.Path,
				)
			}

			user := models.GitHubUser{
				Login: "torvalds",
			}

			json.NewEncoder(w).Encode(user)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	user, err := client.GetUser("torvalds")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.Login != "torvalds" {
		t.Errorf(
			"expected torvalds, got %s",
			user.Login,
		)
	}
}

func TestGetUserNotFound(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusNotFound)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUser("unknown")

	if err != apperror.ErrUserNotFound {
		t.Errorf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestGetUserServerError(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.WriteHeader(http.StatusInternalServerError)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUser("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expected := "github API returned status: 500"

	if err.Error() != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			err.Error(),
		)
	}
}

func TestGetUserInvalidJSON(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set(
				"Content-Type",
				"application/json",
			)

			w.Write([]byte(`{invalid json`))
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUser("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetUserReposSuccess(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.URL.Path != "/users/torvalds/repos" {
				t.Errorf(
					"expected path /users/torvalds/repos, got %s",
					r.URL.Path,
				)
			}

			repos := []models.GitHubRepo{
				{
					Name:     "linux",
					Stars:    100,
					Language: "C",
				},
				{
					Name:     "project",
					Stars:    50,
					Language: "Go",
				},
			}

			json.NewEncoder(w).Encode(repos)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	repos, err := client.GetUserRepos("torvalds")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(repos) != 2 {
		t.Errorf(
			"expected 2 repos, got %d",
			len(repos),
		)
	}

	if repos[0].Name != "linux" {
		t.Errorf(
			"expected linux, got %s",
			repos[0].Name,
		)
	}
}

func TestGetUserReposNotFound(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUserRepos("unknown")

	if err != apperror.ErrUserNotFound {
		t.Errorf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestGetUserReposServerError(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUserRepos("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expected := "github API returned status: 500"

	if err.Error() != expected {
		t.Errorf(
			"expected %s, got %s",
			expected,
			err.Error(),
		)
	}
}

func TestGetUserReposInvalidJSON(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`{invalid json`))
		}),
	)

	defer server.Close()

	client := GitHubAPIClient{
		BaseURL: server.URL,
	}

	_, err := client.GetUserRepos("torvalds")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestNewGitHubAPIClient(t *testing.T) {

	client := NewGitHubAPIClient()

	if client.BaseURL != "https://api.github.com" {
		t.Errorf(
			"expected https://api.github.com, got %s",
			client.BaseURL,
		)
	}
}
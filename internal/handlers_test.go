package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetAll(t *testing.T) {
	t.Run("get all", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env", nil)
		response := httptest.NewRecorder()

		envHandler(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		gotStr := response.Body.String()
		want, _ := json.Marshal(os.Environ())

		if gotStr != string(want) {
			t.Errorf("got %q want %q", gotStr, string(want))
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("get existing key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/PATH", nil)
		response := httptest.NewRecorder()

		envHandler(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		gotStr := response.Body.String()
		want, _ := json.Marshal(os.Getenv("PATH"))

		if gotStr != string(want) {
			t.Errorf("got %q want %q", gotStr, string(want))
		}
	})

	t.Run("get not existing key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/notfound", nil)
		response := httptest.NewRecorder()

		envHandler(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		gotStr := response.Body.String()
		want, _ := json.Marshal(os.Getenv("notfound"))

		if gotStr != string(want) {
			t.Errorf("got %q want %q", gotStr, string(want))
		}
	})

	t.Run("get empty key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/", nil)
		response := httptest.NewRecorder()

		envHandler(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
	})
}

func Test404(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/envNotfound", nil)
	response := httptest.NewRecorder()

	envHandler(response, request)

	assertStatus(t, response.Code, http.StatusNotFound)
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("did not get correct status, got %d, want %d", got, want)
	}
}

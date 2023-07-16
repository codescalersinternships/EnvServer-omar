package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"
)

type subEnvVariables struct {
	variables map[string]string
}

func (s *subEnvVariables) getAll() []string {
	values := []string{}
	for k, v := range s.variables {
		values = append(values, fmt.Sprintf("%v=%v", k, v))
	}
	return values
}

func (s *subEnvVariables) get(key string) string {
	return s.variables[key]
}

func TestGetAll(t *testing.T) {
	server := EnvServer{Env: &subEnvVariables{
		variables: map[string]string{"key1": "value1", "key2": "value2"},
	}}
	values := []string{"key1=value1", "key2=value2"}

	t.Run("get all", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/", nil)
		response := httptest.NewRecorder()

		server.HandleEnv(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		gotStr := response.Body.String()
		got := strings.Split(gotStr[1:len(gotStr)-1], " ")
		sort.Strings(got)

		if !reflect.DeepEqual(got, values) {
			t.Errorf("got %v want %v", got, values)
		}
	})
}

func TestGet(t *testing.T) {
	server := EnvServer{Env: &subEnvVariables{
		variables: map[string]string{"key1": "value1", "key2": "value2"},
	}}

	t.Run("get existing key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/key1", nil)
		response := httptest.NewRecorder()

		server.HandleEnvKey(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		got := response.Body.String()
		want := "value1"

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get not existing key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/env/notfound", nil)
		response := httptest.NewRecorder()

		server.HandleEnvKey(response, request)

		assertStatus(t, response.Code, http.StatusOK)

		got := response.Body.String()
		want := ""

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

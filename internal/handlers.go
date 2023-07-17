package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func envHandler(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")[1:]
	partsCount := len(urlParts)

	switch {
	case r.Method == http.MethodGet && (partsCount == 1 && urlParts[0] == "env"):
		getAllEnv(w)
	case r.Method == http.MethodGet && partsCount == 2 && urlParts[0] == "env":
		getEnv(w, urlParts[1])
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func getAllEnv(w http.ResponseWriter) {
	envJson, err := json.Marshal(os.Environ())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(envJson); err != nil {
		log.Println(err)
	}
}

func getEnv(w http.ResponseWriter, key string) {
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	envJson, err := json.Marshal(os.Getenv(key))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(envJson); err != nil {
		log.Println(err)
	}
}

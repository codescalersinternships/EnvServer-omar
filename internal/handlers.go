package internal

import (
	"fmt"
	"net/http"
	"strings"
)

type variables interface {
	getAll() []string
	get(key string) string
}

// EnvServer is a HTTP interface for getting environment variables.
type EnvServer struct {
	Env variables
}

func (e *EnvServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")[1:]
	partsCount := len(urlParts)

	switch {
	case r.Method == http.MethodGet &&
		((partsCount == 1 && urlParts[0] == "env") ||
			(partsCount == 2 && urlParts[0] == "env" && urlParts[1] == "")):
		fmt.Fprint(w, e.Env.getAll())
	case r.Method == http.MethodGet && partsCount == 2 && urlParts[0] == "env":
		fmt.Fprint(w, e.Env.get(urlParts[1]))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

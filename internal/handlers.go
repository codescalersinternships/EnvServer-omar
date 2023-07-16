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

// HandleEnv is an endpoint gets all the environment variables on the host.
func (e *EnvServer) HandleEnv(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, e.Env.getAll())
}

// HandleEnvKey is an endpoint gets value of environment variable on the host.
func (e *EnvServer) HandleEnvKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/env/")
	fmt.Fprint(w, e.Env.get(key))
}

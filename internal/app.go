package internal

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

var ErrInvalidPort = errors.New("port must be between 1 and 65535")

type App struct {
	port int
}

func NewApp(port int) App {
	return App{port}
}

func (a *App) Run() error {
	if a.port < 1 || a.port > 65535 {
		return ErrInvalidPort
	}

	server := &EnvServer{Env: &HostEnvVariables{}}

	println("Server is listening on port " + strconv.Itoa(a.port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(a.port), server))

	return nil
}

package internal

import (
	"errors"
	"net/http"
	"strconv"
)

var ErrInvalidPort = errors.New("port must be between 1 and 65535")

// App is struct handle the endpoint of the api.
type App struct {
	port int
}

// NewApp is a factory return app instance with the given port.
func NewApp(port int) (App, error) {
	if port < 1 || port > 65535 {
		return App{}, ErrInvalidPort
	}
	return App{port}, nil
}

// Run is the entry point running the api.
func (a *App) Run() error {
	http.HandleFunc("/", envHandler)

	println("Server is listening on port " + strconv.Itoa(a.port))
	err := http.ListenAndServe(":"+strconv.Itoa(a.port), nil)

	return err
}

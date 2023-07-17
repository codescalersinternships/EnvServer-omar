package internal

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

var ErrInvalidPort = errors.New("port must be between 1 and 65535")

// App is struct handle the endpoint of the api.
type App struct {
	port int
}

// NewApp is a factory return app instance with the given port.
func NewApp(port int) App {
	return App{port}
}

// Run is the entry point running the api.
func (a *App) Run() error {
	if a.port < 1 || a.port > 65535 {
		return ErrInvalidPort
	}

	http.HandleFunc("/", envHandler)

	println("Server is listening on port " + strconv.Itoa(a.port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(a.port), nil))

	return nil
}

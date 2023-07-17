package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/EnvServer-omar/internal"
)

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "specify port number")

	flag.Parse()

	app := internal.NewApp(port)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

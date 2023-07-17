package main

import (
	"log"
	"net/http"

	"github.com/codescalersinternships/EnvServer-omar/internal"
)

func main() {
	server := &internal.EnvServer{Env: &internal.HostEnvVariables{}}

	port := "8080"
	println("Server is listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

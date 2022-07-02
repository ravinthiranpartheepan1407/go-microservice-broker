package main

import (
	"fmt"
	"net/http"
)

const port = "80"

type Config struct {
}

func main() {

	app := Config{}

	fmt.Printf("Starting Broker Service on Port %s", port)

	serverInt := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	err := serverInt.ListenAndServe()

	if err != nil {
		fmt.Printf("Error Found: %v", err)
	}

}

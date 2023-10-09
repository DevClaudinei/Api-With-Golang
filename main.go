package main

import (
	"api/src/configuration"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	configuration.Loading()
	r := router.Generate()

	fmt.Printf("Listening on port %d", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), r))
}

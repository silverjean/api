package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.OnLoad();

	
	route := router.Generate()
	
	fmt.Printf("Running and listen on port %d 🚀", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), route))
}
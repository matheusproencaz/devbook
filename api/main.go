package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	config.Load()

	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Listening in port %s ✨\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

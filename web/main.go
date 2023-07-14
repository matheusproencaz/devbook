package main

import (
	"fmt"
	"log"
	"net/http"

	"web/src/config"
	"web/src/router"
	"web/utils"
)

func main() {
	r := router.Generate()
	utils.LoadTemplates()
	config.Load()

	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Listening in port%s ✨\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"web/src/config"
	"web/src/cookies"
	"web/src/router"
	"web/utils"
)

// func init() {
// 	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
// 	fmt.Println(hashKey)
// 	fmt.Println(blockKey)
// }

func main() {
	utils.LoadTemplates()
	config.Load()
	cookies.Load()
	r := router.Generate()

	port := fmt.Sprintf(":%d", config.Port)
	fmt.Printf("Listening in port%s ✨\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

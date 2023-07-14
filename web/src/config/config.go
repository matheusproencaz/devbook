package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Port  é a porta onde o webapp vai estar rodando.
	Port = 0
)

// Load loads the port number from the environment variable "WEBAPP_PORT".
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("WEBAPP_PORT"))
	if erro != nil {
		Port = 9000
	}
}

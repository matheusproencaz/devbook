package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// API URL representa a URL para comunicação com a API
	APIURL = ""
	//Port  é a porta onde o webapp vai estar rodando.
	Port = 0
	//HashKey é utilizado para autenticar o cookie
	HashKey []byte
	//BlockKey é utilizado para criptografar o cookie
	BlockKey []byte
)

// Load inicializa as variavéis de ambiente
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("WEBAPP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}

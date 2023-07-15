package middlewares

import (
	"log"
	"net/http"
	"web/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Método: %s, URL: %s, Host: %s", r.Method, r.URL, r.Host)
		next(w, r)
	}
}

// Authentication verifica a existência dos cookies de dados de autenticação
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Read(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next(w, r)
	}
}

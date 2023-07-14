package controllers

import (
	"net/http"
	"web/utils"
)

// LoadLoginScreen renderiza a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadSignupScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

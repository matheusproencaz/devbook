package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/responses"
)

// Login utiliza o e-mail e senha do usuário para autenticar na aplicação.
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(user))
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	// O professor fez o jwt token retornar pelo body, eu no caso utilizei os headers, token.
	// Como seria caso eu tivesse feito pelo body:
	// token, _ = ioutil.ReadAll(response.Body) //ioutil retorna slice de byte([]byte)
	// string(token) //converter para string
	token := response.Header.Get("Token")
	fmt.Println(token)
}

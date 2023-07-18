package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"
	"web/src/config"
	"web/src/cookies"
	"web/src/requests"
	"web/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser chama a API para cadastrar um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// Unfollow chama a API para parar de seguir um usuário
func Unfollow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, erro := requests.RequestWithAuth(r, http.MethodPost, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// Unfollow chama a API para seguir um usuário
func Follow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, erro := requests.RequestWithAuth(r, http.MethodPost, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// EditUser chama a API para atualizar o usuário
func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
	})
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, erro := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(user))
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePassword chama a API para atualizar a senha do usuário
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passwords, erro := json.Marshal(map[string]string{
		"new":     r.FormValue("new"),
		"current": r.FormValue("current"),
	})
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/updatepassword", config.APIURL, userId)

	response, erro := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(passwords))
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

// UpdatePassword chama a API para deletar o usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, erro := requests.RequestWithAuth(r, http.MethodDelete, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

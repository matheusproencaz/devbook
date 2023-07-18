package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/utils"

	"github.com/gorilla/mux"
)

// LoadLoginScreen renderiza a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadSignupScreen renderiza a tela de cadastro
func LoadSignupScreen(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

// LoadSignupScreen renderiza a tela de home com as publicações
func LoadHomeScreen(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	var posts []models.Posts
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Posts
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

// LoadSignupScreen renderiza a tela de edição de publicação
func LoadEditionPostPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, erro := strconv.ParseUint(vars["postId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postId)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	var post models.Posts

	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "edition.html", post)
}

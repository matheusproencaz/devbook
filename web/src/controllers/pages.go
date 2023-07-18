package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"web/src/config"
	"web/src/cookies"
	"web/src/models"
	"web/src/requests"
	"web/src/responses"
	"web/utils"

	"github.com/gorilla/mux"
)

// LoadLoginPage renderiza a tela de login
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {

	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadSignupPage renderiza a tela de cadastro
func LoadSignupPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}

// LoadSignupPage renderiza a tela de home com as publicações
func LoadHomePage(w http.ResponseWriter, r *http.Request) {
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

	var posts []models.Post
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

// LoadSignupPage renderiza a tela de edição de publicação
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

	var post models.Post

	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "edition.html", post)
}

// LoadSignupPage renderiza a tela de de usuários que atendem o filtro passado.
func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?username=%s", config.APIURL, nameOrNick)

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

	var users []models.User
	if erro = json.NewDecoder(response.Body).Decode(&users); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "users.html", users)
}

// LoadSignupPage renderiza a tela de perfil de usuário
func LoadUserProfilePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedInUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == loggedInUserId {
		http.Redirect(w, r, "/profile", 302)
		return
	}

	user, erro := models.GetCompleteInfoUser(userID, r)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "user.html", struct {
		User           models.User
		LoggedInUserID uint64
	}{
		User:           user,
		LoggedInUserID: loggedInUserId,
	})
}

// LoadLoggedInUserProfilePage renderiza a tela de perfil do usuário logado
func LoadLoggedInUserProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, erro := models.GetCompleteInfoUser(userID, r)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecuteTemplate(w, "profile.html", user)
}

// LoadUserEditionPage renderiza a tela de edição do usuário
func LoadUserEditionPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.GetUserData(channel, userID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: "Erro ao buscar usuário!"})
		return
	}

	utils.ExecuteTemplate(w, "editUser.html", user)
}

// LoadUserEditionPage renderiza a tela de alteração da senha
func LoadUserPasswordEditionPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "editUserPassword.html", nil)
}

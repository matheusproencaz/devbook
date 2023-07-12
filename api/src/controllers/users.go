package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser chama o repositório para criar um usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := user.Prepare(models.SingUp); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers chama o repositório para chamar todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("username"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	users, erro := repository.GetUsers(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

// GetUserById chama o repositório para buscar um usuário pelo seu ID.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, erro := repository.GetUserById(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser chama o repositório para atualizar um usuário
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o próprio"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal([]byte(body), &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(models.Update); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if erro = repository.UpdateUser(userID, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser chama o repositório para deletar um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar um usuário que não seja o próprio"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	if erro = repository.DeleteUser(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser permite um usuário seguir o outro
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível seguir o próprio usuário"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if erro = repository.FollowUser(userID, followerID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser permite um usuário parar seguir o outro
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	followerID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir o próprio usuário"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	if erro = repository.UnFollowUser(userID, followerID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers permite um usuário acessar quem segue ele, e também quem segue outros usuários
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	followers, erro := repository.GetFollowers(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// Following permite um usuário ver quem ele segue, e tabém quem os outros usuários seguem.
func GetFollowing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	followers, erro := repository.GetFollowing(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// UpdatePassword permite alterar a senha de um usuário
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	vars := mux.Vars(r)
	userID, erro := strconv.ParseUint(vars["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if userID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de um usuário que não seja o seu"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	var password models.Password
	if erro = json.Unmarshal(body, &password); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	passwordInDatabase, erro := repository.GetPassword(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(passwordInDatabase, password.Current); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com a que está salva no banco de dados"))
		return
	}

	passwordHash, erro := security.Hash(password.New)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePassword(userID, string(passwordHash)); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost cria uma publicação
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Posts

	if erro = json.Unmarshal(body, &post); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	post.AuthorID = userID

	if erro = post.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	post.ID, erro = repository.CreatePost(post)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts traz as publicações que apareceriam no feed do usuário
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, erro := repository.GetPosts(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// GetPostByID pega uma única publicação
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postID, erro := strconv.ParseUint(vars["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)

	post, erro := repository.GetPostByID(postID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// UpdatePost atualiza os dados de uma publicação
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
	}

	vars := mux.Vars(r)
	postID, erro := strconv.ParseUint(vars["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	postInDatabase, erro := repository.GetPostByID(postID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postInDatabase.AuthorID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var post models.Posts
	if erro = json.Unmarshal(body, &post); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = post.Prepare(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePost(postID, post); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost deleta uma públicação
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := authentication.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
	}

	vars := mux.Vars(r)
	postID, erro := strconv.ParseUint(vars["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	postInDatabase, erro := repository.GetPostByID(postID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if postInDatabase.AuthorID != userIDToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repository.DeletePost(postID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetPostByUser pega todas as publicações de um usuário
func GetPostByUser(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewPostsRepository(db)
	posts, erro := repository.GetPostByUser(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// LikePost aumenta o número de likes de uma publicação
func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, erro := strconv.ParseUint(vars["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	if erro = repository.LikePost(postID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DislikePost aumenta o número de likes de uma publicação
func DislikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, erro := strconv.ParseUint(vars["postId"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	if erro = repository.DislikePost(postID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

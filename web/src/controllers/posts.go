package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/requests"
	"web/src/responses"

	"github.com/gorilla/mux"
)

// CreatPost chama a API para criar uma publicação no banco de dados.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)

	response, erro := requests.RequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))
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

// LikePost chama a API para curtir uma publicação no banco de dados.
func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId, erro := strconv.ParseUint(vars["postId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.APIURL, postId)
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

// LikePost chama a API para descurtir uma publicação no banco de dados.
func DislikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId, erro := strconv.ParseUint(vars["postId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.APIURL, postId)
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

// LikePost chama a API para atualizar uma publicação no banco de dados.
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId, erro := strconv.ParseUint(vars["postId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	r.ParseForm()
	updatedPost, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postId)

	response, erro := requests.RequestWithAuth(r, http.MethodPut, url, bytes.NewBuffer(updatedPost))
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

// LikePost chama a API para deletar uma publicação no banco de dados.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId, erro := strconv.ParseUint(vars["postId"], 10, 64)
	if erro != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postId)

	response, erro := requests.RequestWithAuth(r, http.MethodDelete, url, nil)
	if erro != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
	}

	responses.JSON(w, response.StatusCode, nil)
}

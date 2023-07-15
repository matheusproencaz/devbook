package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/requests"
	"web/src/responses"
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

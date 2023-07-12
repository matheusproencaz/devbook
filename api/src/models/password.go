package models

// Password representa o formado da requisição da alteração de senha
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}

package models

// DataAuth contém o id e o token do usuário autenticado
type DataAuth struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

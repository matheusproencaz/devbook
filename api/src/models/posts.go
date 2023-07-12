package models

import (
	"errors"
	"strings"
	"time"
)

// Posts representa uma publicação feita por um usuário
type Posts struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatAt    time.Time `json:"create_at,omitempty"`
}

// Prepare vai chamar os métodos validate e format para validar os campos e formatalos da forma correta.
func (post *Posts) Prepare() error {
	post.format()
	if erro := post.validate(); erro != nil {
		return erro
	}
	return nil
}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}
	if len(post.Title) > 50 {
		return errors.New("O título não pode ser maior do que 50 caracteres!")
	}
	if post.Content == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}
	if len(post.Content) > 300 {
		return errors.New("O conteúdo não pode ser maior do que 300 caracteres!")
	}
	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}

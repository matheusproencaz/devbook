package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário utilizado a rede social
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

func (user *User) validate(step Step) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}
	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco.")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O email inserido é inválido")
	}
	if step == Step(SingUp) && user.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco.")
	}
	return nil
}

func (user *User) format(step Step) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == SingUp {
		passwordWithHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}

// Prepare valida e formata o usuário recebido.
func (user *User) Prepare(step Step) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.format(step); erro != nil {
		return erro
	}

	return nil
}

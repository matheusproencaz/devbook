package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// NewUserRepository cria um repositório de usuários
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insere um usuário no banco de dados
func (repository users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}
	lastIDInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(lastIDInsertedId), nil
}

// GetUsers traz todos os usuários atendem um friltro de nome ou nick.
func (repository users) GetUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, create_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserById traz um usuário pelo seu ID do banco de dados
func (repository users) GetUserById(userID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"SELECT id, name, nick, email, create_at FROM users WHERE id = ?",
		userID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

// UpdateUser atualiza um usuário no banco de dados
func (repository users) UpdateUser(userID uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"UPDATE users SET NAME = ?, NICK = ?, EMAIL = ? WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userID); erro != nil {
		return erro
	}
	return nil
}

// DeleteUser deleta um usuário no banco de dados
func (repository users) DeleteUser(userID uint64) error {

	statement, erro := repository.db.Prepare("DELETE FROM users WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID); erro != nil {
		return erro
	}

	return nil
}

// GetByEmail buscar um usuário pelo seu email no banco de dados
func (repository users) GetByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

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

// FollowUser segue um usuário no banco de dados
func (repository users) FollowUser(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare("INSERT IGNORE INTO followers VALUES (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// UnFollowUser deixa de seguir um usuário no banco de dados
func (repository users) UnFollowUser(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

// GetFollowers Busca os seguidores do usuário do ID passado.
func (repository users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.create_at FROM users 
		u INNER JOIN followers f ON u.id = f.follower_id 
		WHERE f.user_id = ?`, userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	followers := make([]models.User, 0)

	for lines.Next() {
		var follower models.User

		if erro = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateAt,
		); erro != nil {
			return nil, erro
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

// GetFollowing Busca com base no ID do usuário passado quem ele segue
func (repository users) GetFollowing(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(
		`SELECT u.id, u.name, u.nick, u.email, u.create_at FROM users u 
		INNER JOIN followers f ON u.id = f.user_id 
		WHERE f.follower_id = ?`, userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	followers := make([]models.User, 0)

	for lines.Next() {
		var follower models.User
		if erro = lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateAt,
		); erro != nil {
			return nil, erro
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

// GetPassword pega a senha salva no banco de dados do usuário passado.
func (repository users) GetPassword(userID uint64) (string, error) {
	lines, erro := repository.db.Query("SELECT password FROM users WHERE id = ?", userID)
	if erro != nil {
		return "", erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.Password,
		); erro != nil {
			return "", erro
		}
	}
	return user.Password, nil
}

// UpdatePassword altera a senha de um usuário
func (repository users) UpdatePassword(userID uint64, password string) error {
	statement, erro := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}

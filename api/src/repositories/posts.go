package repositories

import (
	"api/src/models"
	"database/sql"
)

type posts struct {
	db *sql.DB
}

// NewPostsRepository cria um repositório de publicações
func NewPostsRepository(db *sql.DB) *posts {
	return &posts{db}
}

// CreatePost insera uma publicação no banco de dados
func (repository posts) CreatePost(posts models.Posts) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultDB, erro := statement.Exec(posts.Title, posts.Content, posts.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastInsertId, erro := resultDB.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

/** GetPosts busca todas as publicações dos usuários que o usuário logado
* 	segue e também suas próprias publicações no banco de dados
 */
func (repository posts) GetPosts(userID uint64) ([]models.Posts, error) {
	// Usando LEFT JOIN
	// lines, erro := repository.db.Query(
	// 	`SELECT DISTINCT p.*, u.nick FROM posts p
	// 	INNER JOIN users u ON u.id = p.author_id
	// 	LEFT JOIN followers f ON p.author_id = f.user_id
	// 	WHERE u.id = ? OR f.follower_id = ?`,
	// 	userID, userID,
	// )

	// Não usando LEFT JOIN
	lines, erro := repository.db.Query(
		`SELECT DISTINCT p.*, u.nick
		FROM posts p INNER JOIN users u ON u.id = p.author_id
		INNER JOIN followers f ON f.follower_id = ?
		WHERE p.author_id = ? OR p.author_id = f.user_id
		ORDER BY 1 DESC`,
		userID, userID,
	)

	if erro != nil {
		return nil, erro
	}

	var posts []models.Posts

	for lines.Next() {
		var post models.Posts
		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		posts = append(posts, post)
	}
	return posts, nil
}

// GetPostByID traz uma única publicação do banco de dados
func (repository posts) GetPostByID(postID uint64) (models.Posts, error) {
	lines, erro := repository.db.Query(
		`SELECT p.*, u.nick FROM posts p 
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = ?`,
		postID,
	)
	if erro != nil {
		return models.Posts{}, erro
	}

	var post models.Posts

	if lines.Next() {
		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatAt,
			&post.AuthorNick,
		); erro != nil {
			return models.Posts{}, erro
		}
	}
	return post, nil
}

// UpdatePost altera os dados de uma publicação no banco de dados
func (repository posts) UpdatePost(postID uint64, post models.Posts) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postID); erro != nil {
		return erro
	}

	return nil
}

// DeletePost deleta uma publicação no banco de dados
func (repository posts) DeletePost(postID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

// GetPostByUser busca as publicações de um usuário no banco de dados
func (repository posts) GetPostByUser(postID uint64) ([]models.Posts, error) {
	lines, erro := repository.db.Query(
		`SELECT p.*, u.nick FROM posts p 
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.author_id = ?`,
		postID,
	)
	if erro != nil {
		return nil, erro
	}

	var posts []models.Posts

	for lines.Next() {
		var post models.Posts

		if erro = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatAt,
			&post.AuthorNick,
		); erro != nil {
			return nil, erro
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// LikePost aumenta o número de curtidas na publicação
func (repository posts) LikePost(postID uint64) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

// DislikePost diminui o número de curtidas na publicação
func (repository posts) DislikePost(postID uint64) error {
	statement, erro := repository.db.Prepare(`UPDATE posts SET likes = 
	CASE 
		WHEN likes > 0 THEN likes - 1
		ELSE likes 
	END 
	WHERE id = ?`)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}

package models

import "time"

// Post representa uma publicação feita por um usuário
type Post struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreateAt   time.Time `json:"create_at,omitempty"`
}

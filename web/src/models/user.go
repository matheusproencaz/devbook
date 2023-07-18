package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"web/src/config"
	"web/src/requests"
)

// User representa uma pessoa utilizando a rede social
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreateAt  time.Time `json:"create_at"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

// GetCompleteInfoUser faz 4 requisições na API para montar o usuário completo
func GetCompleteInfoUser(userId uint64, r *http.Request) (User, error) {
	channelUser := make(chan User)
	channelFollowers := make(chan []User)
	channelFollowings := make(chan []User)
	channelPosts := make(chan []Post)

	go GetUserData(channelUser, userId, r)
	go GetFollowers(channelFollowers, userId, r)
	go GetFollowing(channelFollowings, userId, r)
	go GetUserPosts(channelPosts, userId, r)

	var (
		user       User
		followers  []User
		followings []User
		posts      []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case completedUser := <-channelUser:
			if completedUser.ID == 0 {
				return User{}, errors.New("Erro ao buscar o usuário!")
			}

			user = completedUser

		case completedFollowers := <-channelFollowers:
			if completedFollowers == nil {
				return User{}, errors.New("Erro ao buscar os seguidores do usuário!")
			}

			followers = completedFollowers

		case completedFollowings := <-channelFollowings:
			if completedFollowings == nil {
				return User{}, errors.New("Erro ao buscar quem o usuário segue!")
			}

			followings = completedFollowings

		case completedPosts := <-channelPosts:
			if completedPosts == nil {
				return User{}, errors.New("Erro ao buscar os posts do usuário!")
			}

			posts = completedPosts
		}
	}

	user.Followers = followers
	user.Following = followings
	user.Posts = posts

	return user, nil
}

// GetUserData chama a API para buscar os dados base do usuário
func GetUserData(canal chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userId)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if erro = json.NewDecoder(response.Body).Decode(&user); erro != nil {
		canal <- User{}
		return
	}

	canal <- user
}

// GetFollowers chama a API para buscar os seguidores do usuário
func GetFollowers(canal chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userId)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var followers []User

	if erro = json.NewDecoder(response.Body).Decode(&followers); erro != nil {
		canal <- nil
		return
	}

	if followers == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- followers
}

// GetFollowing chama a API para buscar quem o usuário segue
func GetFollowing(canal chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userId)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var following []User

	if erro = json.NewDecoder(response.Body).Decode(&following); erro != nil {
		canal <- nil
		return
	}

	if following == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- following
}

// GetUserPosts chama a API para buscar os posts do usuário
func GetUserPosts(canal chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userId)
	response, erro := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post

	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		canal <- nil
		return
	}

	if posts == nil {
		canal <- make([]Post, 0)
		return
	}

	canal <- posts
}

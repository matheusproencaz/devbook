package cookies

import (
	"net/http"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Load utiliza as variáveis de ambiente para a criação do SecureCookie
func Load() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	dataEncoded, erro := s.Encode("data", data)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    dataEncoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read retorna os valores armazenados no cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("data")
	if erro != nil {
		return nil, erro
	}

	values := make(map[string]string)
	if erro = s.Decode("data", cookie.Value, &values); erro != nil {
		return nil, erro
	}

	return values, nil
}

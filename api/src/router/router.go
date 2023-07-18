package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate cria um mux router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}

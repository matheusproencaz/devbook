package router

import (
	"web/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorna um router mux
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configuration(r)
}

package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route é um struc que representa uma rota da Aplicação Web.
type Route struct {
	URI       string
	Method    string
	Handler   func(http.ResponseWriter, *http.Request)
	NeedsAuth bool
}

// Configuration coloca todas as rotas dentro do router
func Configuration(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

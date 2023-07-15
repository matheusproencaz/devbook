package routes

import (
	"net/http"
	"web/src/middlewares"

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
	routes = append(routes, homeRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {
		if route.NeedsAuth {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Handler)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Handler),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

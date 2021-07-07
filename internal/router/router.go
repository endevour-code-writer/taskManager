package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router Router

type Router struct {
	RoutesMultiplexer *chi.Mux
}

func Init() Router {
	return router
}

func init() {
	mux := chi.NewRouter()
	router = Router{mux}
	router.RoutesMultiplexer.Use(middleware.Recoverer)
}

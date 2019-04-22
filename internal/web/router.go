package http

import (
	"github.com/go-chi/chi"
)

func route() *chi.Route {
	r := chi.NewRouter()
	r.Route("/s", func(chi.Router) {
		r.Get("/", )
	})
}

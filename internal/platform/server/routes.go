package server

import "github.com/go-chi/chi/v5"

func (s *Server) routes() {
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", s.userHandler.HandleGetUsers())
			r.Post("/", s.userHandler.HandleCreateUser())
		})
	})
}

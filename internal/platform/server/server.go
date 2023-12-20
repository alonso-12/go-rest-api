package server

import (
	"matryer/internal/platform/server/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router      *chi.Mux
	logger      *logrus.Logger
	userHandler handlers.IUserHandler
}

func New(router *chi.Mux, logger *logrus.Logger, userHandler handlers.IUserHandler) *Server {
	s := &Server{
		router:      router,
		logger:      logger,
		userHandler: userHandler,
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

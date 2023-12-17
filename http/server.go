package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	server *http.Server
	router *chi.Mux
}

func NewServer(addr string) *Server {
	s := &Server{
		server: &http.Server{
			Addr: addr,
		},
		router: chi.NewRouter(),
	}

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wondrous Baazar *.*"))
	})

	s.server.Handler = s.router

	return s
}

func (s *Server) Listen() error {
	err := s.server.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx := context.Background()
	return s.server.Shutdown(ctx)
}

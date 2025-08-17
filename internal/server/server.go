package server

import (
	"log"
	"net/http"
	"time"

	"github.com/bugagych84/go1fl-sprint6-final-tpl/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	HTTP   *http.Server
}

func New(l *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.Upload)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: l,
		HTTP:   srv,
	}
}

func (s *Server) Start() error {
	s.Logger.Printf("listening on %s", s.HTTP.Addr)
	return s.HTTP.ListenAndServe()
}

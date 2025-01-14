package main

import (
	"log"
	"net/http"

	"github.com/vhall1/expense-tracker/service.snowflake/config"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()

	if err := config.RegisterRoutes(mux); err != nil {
		return err
	}

	log.Printf("[service.snowflake] listening for requests on port %s", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, mux)
}

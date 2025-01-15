package main

import (
	"log"
	"net/http"

	handler "github.com/vhall1/expense-tracker/service.snowflake/handler/snowflake"
	"github.com/vhall1/expense-tracker/service.snowflake/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	snowflakeService, err := service.NewSnowflakeService(int64(0))
	if err != nil {
		return err
	}

	snowflakeHandler := handler.NewHttpSnowflakeService(snowflakeService)
	snowflakeHandler.RegisterRoutes(router)

	log.Println("[service.snowflake] Starting http server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

package server

import (
	"fmt"
	"myapp/internal/config"
	"myapp/internal/routers"
	"myapp/pkg/log"
	"net/http"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() {
	router := routers.NewRouter()
	logger, err := log.NewLogger(s.cfg)
	if err != nil {
		panic(err)
	}
	logger.Info(fmt.Sprintf("Starting server on port %s", s.cfg.ServerPort))
	if err := http.ListenAndServe(":"+s.cfg.ServerPort, router); err != nil {
		logger.Error(fmt.Sprintf("could not start server: %v", err))
	}
}

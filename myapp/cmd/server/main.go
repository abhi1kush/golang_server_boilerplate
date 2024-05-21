package main

import (
	stdLog "log"
	"myapp/internal/config"
	"myapp/internal/server"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		stdLog.Fatalf("failed to load config: %v", err)
	}

	// Create and start the server
	s := server.NewServer(cfg)
	s.Start()
}

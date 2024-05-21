package routers

import (
	"myapp/internal/handlers"
	"myapp/internal/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	userHandler := handlers.NewUserHandler()
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")

	// Apply middleware
	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.AuthMiddleware)

	return router
}

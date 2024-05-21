package routers

import (
	"myapp/internal/handlers"
	"myapp/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthHandler).Methods("GET")
	apiRouter := router.PathPrefix("/api").Subrouter()
	// Apply middleware
	apiRouter.Use(middleware.LoggerMiddleware)
	apiRouter.Use(middleware.AuthenticationMiddleware)
	userHandler := handlers.NewUserHandler()
	apiRouter.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	return router
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

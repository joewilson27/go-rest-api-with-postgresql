package router

import (
	"go-rest-api-with-postgresql/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// Add data
	router.HandleFunc("/api/book", controller.AddBook).Methods("POST", "OPTIONS")
	// Get a single data
	router.HandleFunc("/api/book/{id}", controller.GetBook).Methods("GET", "OPTIONS")
	// Get all data
	router.HandleFunc("/api/book", controller.GetAllBooks).Methods("GET", "OPTIONS")
	// Update data
	router.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT", "OPTIONS")
	// Delete data
	router.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}

package router

import (
	"go-mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controller.MarkAsRead).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controller.DeleteABook).Methods("DELETE")
	router.HandleFunc("/api/books/{id}", controller.DeleteAllBooks).Methods("DELETE")

	return router
}

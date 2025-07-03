package router

import (
	"github.com/gorilla/mux"
	"github.com/ritesh-karankal/go-postgres/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return r
}

package main

import (
	"github.com/gorilla/mux"
	"github.com/iluxaorlov/godo/internal/app/godo"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/task", godo.ReadAllHandler).Methods(http.MethodGet)
	r.HandleFunc("/task", godo.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/task/{id}", godo.ReadOneHandler).Methods(http.MethodGet)
	r.HandleFunc("/task/{id}", godo.UpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/task/{id}", godo.DeleteHandler).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
package main

import (
	"github.com/gorilla/mux"
	"github.com/iluxaorlov/todo/internal/app/todo"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/task", todo.ReadAllHandler).Methods(http.MethodGet)
	r.HandleFunc("/task", todo.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/task/{id}", todo.ReadOneHandler).Methods(http.MethodGet)
	r.HandleFunc("/task/{id}", todo.UpdateHandler).Methods(http.MethodPut)
	r.HandleFunc("/task/{id}", todo.DeleteHandler).Methods(http.MethodDelete)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err.Error())
	}
}
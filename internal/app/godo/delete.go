package godo

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/db")
	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)

	row := conn.QueryRow(context.Background(), "DELETE FROM task WHERE id = $1 RETURNING id, text, status", vars["id"])

	var task Task

	w.Header().Set("Content-Type", "application/json")

	err = row.Scan(&task.Id, &task.Text, &task.Status)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(&task)
	if err != nil {
		panic(err.Error())
	}
}
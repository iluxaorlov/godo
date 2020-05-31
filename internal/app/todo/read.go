package todo

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"net/http"
)

func ReadOneHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/db")
	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)

	row := conn.QueryRow(context.Background(), "SELECT id, text, status FROM task WHERE id = $1", vars["id"])

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

func ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/db")
	if err != nil {
		panic(err.Error())
	}

	rows, err := conn.Query(context.Background(), "SELECT id, text, status FROM task")
	if err != nil {
		panic(err.Error())
	}

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		err := rows.Scan(&task.Id, &task.Text, &task.Status)
		if err != nil {
			panic(err.Error())
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&tasks)
	if err != nil {
		panic(err.Error())
	}
}
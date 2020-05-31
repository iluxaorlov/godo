package todo

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"io/ioutil"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/db")
	if err != nil {
		panic(err.Error())
	}

	var task Task

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		panic(err.Error())
	}

	vars := mux.Vars(r)

	row := conn.QueryRow(context.Background(), "UPDATE task SET text = $1, status = $2 WHERE id = $3 RETURNING id, text, status", task.Text, task.Status, vars["id"])

	task = Task{}

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
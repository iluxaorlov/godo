package godo

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx"
	"io/ioutil"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	var task Task

	err = json.Unmarshal(body, &task)
	if err != nil {
		panic(err.Error())
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://user:password@localhost:5432/db")
	if err != nil {
		panic(err.Error())
	}

	row := conn.QueryRow(context.Background(), "INSERT INTO task (text) VALUES ($1) RETURNING id", &task.Text)

	err = row.Scan(&task.Id)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(&task)
	if err != nil {
		panic(err.Error())
	}
}
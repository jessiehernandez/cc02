package main

import (
	"database/sql"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/testing/example/handler"
	"github.com/testing/example/repository"
	"github.com/testing/example/service"
)

func main() {
	db, err := sql.Open("pgx", "postgres://jessie.hernandez@localhost:5432/testing?sslmode=disable")

	if err != nil {
		panic("could not connect to the database")
	}

	http.Handle("/preference", handler.NewUserPreference(
		service.NewUserPreference(
			repository.NewUserPreference(
				db,
			),
		),
	))

	http.ListenAndServe(":8080", nil)
}

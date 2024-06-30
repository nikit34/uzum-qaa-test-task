package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/nikit34/uzum-qaa-test-task/integration-tests-rest/book"
	"github.com/nikit34/uzum-qaa-test-task/integration-tests-rest/rest"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type jsonError struct {
	Code string `json:"code"`
	Msg string `json:"msg"`
}

type Book struct {
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Image string `json:"image"`
	Genre string `json:"genre"`
	YearPublished int `json:"year_published"`
}

type health struct {
	Status string `json:"status"`
	Messages []string `json:"messages"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error making DB connected: %s", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
	    log.Fatalf("Error making DB driver: %s", err.Error())
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/",
		"music",
		driver,
	)
	if err != nil {
		log.Fatalf("Error making migration engine: %s", err.Error())
	}

	if err := migrator.Steps(2); err != nil {
		log.Fatalf("Error making migration step: %s", err.Error())
	}

	r := mux.NewRouter()

	retriver := book.NewRetriever(db)
	r.Handle("/book/{isbn}", rest.NewGetBookHandler(retriver))

	r.HandleFunc(
		"/healthcheck",
		func(w http.ResponseWriter, r *http.Request) {
			h := health{
				Status:   "OK",
				Messages: []string{},
			}

			b, _ := json.Marshal(h)

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		},
	)

	s := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	s.ListenAndServe()
}

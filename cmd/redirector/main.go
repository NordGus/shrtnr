package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NordGus/shrtnr/domain/url/find"
	"github.com/NordGus/shrtnr/domain/url/storage"
	"github.com/NordGus/shrtnr/fileserver"
	"github.com/NordGus/shrtnr/redirect"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	environment *string
	port        *int
	dbPath      *string
)

func init() {
	environment = flag.String("env", "development", "defines services environment")
	port = flag.Int("port", 4269, "port where the app will listened")
	dbPath = flag.String("db-file-path", "./data/shrtnr.db", "path to SQLite DB file")

	flag.Parse()
}

func main() {
	var (
		db     *sqlx.DB
		ctx    = context.Background()
		router = chi.NewRouter()
	)

	db, err := sqlx.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	// URL domain storage initialization. Must be started first.
	err = storage.Start(db)
	if err != nil {
		log.Fatalln(err)
	}

	// URL domain find initialization. Must be started after storage initialization.
	err = find.Start(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// redirect service initialization
	err = redirect.Start(*environment)
	if err != nil {
		log.Fatalln(err)
	}

	router.Use(middleware.RequestID, middleware.Logger)

	redirect.Routes(router)
	fileserver.PublicRoutes(router)

	err = http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalln(err)
	}
}

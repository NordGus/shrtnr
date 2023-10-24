package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NordGus/shrtnr/domain"
	"github.com/NordGus/shrtnr/domain/shared/middleware"
	"github.com/NordGus/shrtnr/fileserver"
	hypermedia "github.com/NordGus/shrtnr/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	environment          *string
	port                 *int
	dbPath               *string
	urlLimit             *uint
	searchTermLimits     *int
	maxSearchConcurrency *uint
	redirectHost         *string
)

func init() {
	environment = flag.String("env", "development", "defines application environment")
	port = flag.Int("port", 4269, "port where the app will listened")
	dbPath = flag.String("db-file-path", "./data/shrtnr.db", "path to SQLite DB file")
	urlLimit = flag.Uint("capacity", 2500, "limit of URLs that the service can contain")
	searchTermLimits = flag.Int("search-term-limit", 10, "the limit of terms that the search cache returns when called")
	maxSearchConcurrency = flag.Uint("search-concurrency", 30, "limits the amount of concurrent processes when checking trie cache for searching functionality")
	redirectHost = flag.String("redirect-service-url", "http://localhost:4269/r", "url to the redirection service")

	flag.Parse()
}

func main() {
	var (
		db  *sqlx.DB
		ctx = context.Background()
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

	err = domain.Start(ctx, *environment, db, *urlLimit, *maxSearchConcurrency, *searchTermLimits, *redirectHost)
	if err != nil {
		log.Fatalln(err)
	}

	err = hypermedia.Start(*environment, *redirectHost)
	if err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()
	router.Use(chimiddleware.RequestID, chimiddleware.Logger)

	if *environment == "development" {
		router.Use(middleware.CORS)
	}

	hypermedia.Routes(router)
	fileserver.PublicRoutes(router)
	fileserver.PrivateRoutes(router)

	err = http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

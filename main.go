package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"github.com/NordGus/shrtnr/domain/find"
	"log"
	"net/http"

	"github.com/NordGus/shrtnr/domain/create"
	"github.com/NordGus/shrtnr/domain/fileserver"
	"github.com/NordGus/shrtnr/domain/messagebus"
	"github.com/NordGus/shrtnr/domain/redirect"
	"github.com/NordGus/shrtnr/domain/search"
	"github.com/NordGus/shrtnr/domain/shared/middleware"
	"github.com/NordGus/shrtnr/domain/storage"
	hypermedia "github.com/NordGus/shrtnr/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

var (
	environment          *string
	redirectHost         *string
	port                 *int
	urlLimit             *uint
	searchTermLimits     *int
	maxSearchConcurrency *uint

	//go:embed dist private
	content embed.FS
)

func init() {
	environment = flag.String("env", "development", "defines application environment")
	redirectHost = flag.String("redirect-host", "l.hst:4269", "defines the short redirect host")
	port = flag.Int("port", 4269, "port where the app will listened")
	urlLimit = flag.Uint("capacity", 2500, "limit of URLs that the service can contain")
	searchTermLimits = flag.Int("search-term-limit", 10, "the limit of terms that the search cache returns when called")
	maxSearchConcurrency = flag.Uint("search-concurrency", 30, "limits the amount of concurrent processes when checking trie cache for searching functionality")

	flag.Parse()
}

func main() {
	ctx := context.Background()

	// Infrastructure initialization
	messagebus.Start(ctx)
	if err := storage.Start(*environment); err != nil {
		log.Fatalln(err)
	}
	fileserver.Start(content)

	// Domain initialization
	if err := redirect.Start(ctx, *environment, *redirectHost); err != nil {
		log.Fatalln(err)
	}
	if err := find.Start(ctx); err != nil {
		log.Fatalln(err)
	}
	create.Start(ctx, *urlLimit)
	search.Start(ctx, *maxSearchConcurrency, *searchTermLimits)

	// Api initialization
	if err := hypermedia.Start(*environment); err != nil {
		log.Fatalln(err)
	}

	router := chi.NewRouter()
	router.Use(chimiddleware.Logger, redirect.Middleware)

	if *environment == "development" {
		router.Use(middleware.CORS)
	}

	hypermedia.Routes(router)
	fileserver.Routes(router)

	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

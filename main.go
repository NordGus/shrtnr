package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NordGus/shrtnr/server/create"
	"github.com/NordGus/shrtnr/server/fileserver"
	"github.com/NordGus/shrtnr/server/messagebus"
	"github.com/NordGus/shrtnr/server/redirect"
	"github.com/NordGus/shrtnr/server/search"
	"github.com/NordGus/shrtnr/server/shared/middleware"
	"github.com/NordGus/shrtnr/server/storage"

	api "github.com/NordGus/shrtnr/server/http"

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
	storage.Start(*environment)
	fileserver.Start(content)

	// Domain initialization
	if err := redirect.Start(ctx, *environment, *redirectHost); err != nil {
		log.Fatalln(err)
	}
	create.Start(ctx, *urlLimit)
	search.Start(ctx, *maxSearchConcurrency, *searchTermLimits)

	// Api initialization
	api.Start(*environment)

	router := chi.NewRouter()
	router.Use(chimiddleware.Logger, redirect.Middleware)

	if *environment == "development" {
		router.Use(middleware.CORS)
	}

	api.Routes(router)
	fileserver.Routes(router)

	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

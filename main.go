package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NordGus/shrtnr/server/create"
	"github.com/NordGus/shrtnr/server/messagebus"
	"github.com/NordGus/shrtnr/server/redirect"
	"github.com/NordGus/shrtnr/server/search"
	"github.com/NordGus/shrtnr/server/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	api "github.com/NordGus/shrtnr/server/http"
)

var (
	environment          *string
	redirectHost         *string
	port                 *int
	urlLimit             *uint
	maxSearchConcurrency *uint
)

func init() {
	environment = flag.String("env", "development", "defines application environment")
	redirectHost = flag.String("redirect-host", "l.hst:4269", "defines the short redirect host")
	port = flag.Int("port", 4269, "port where the app will listened")
	urlLimit = flag.Uint("capacity", 2500, "limit of URLs that the service can contain")
	maxSearchConcurrency = flag.Uint("search-concurrency", 30, "limits the amount of concurrent processes when checking trie cache for searching functionality")

	flag.Parse()
}

func main() {
	ctx := context.Background()
	router := chi.NewRouter()

	messagebus.Start(ctx)
	storage.Start(*environment)
	redirect.Start(ctx, *environment, *redirectHost)
	create.Start(ctx, *urlLimit)
	search.Start(ctx, *urlLimit, *maxSearchConcurrency)
	api.Start(*environment)

	router.Use(middleware.Logger, redirect.Middleware)

	if *environment == "development" {
		router.Use(devCORSMiddleware)
	}

	api.Routes(router)

	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), router)
	if err != nil {
		log.Fatalf("something went wrong initalizing http server: %v\n", err)
	}
}

func devCORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		next.ServeHTTP(writer, request)
	})
}

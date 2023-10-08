package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NordGus/rom-stack/server/ingest"
	"github.com/NordGus/rom-stack/server/messagebus"
	"github.com/NordGus/rom-stack/server/redirect"
	"github.com/NordGus/rom-stack/server/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

var (
	environment  *string
	redirectHost *string
	port         *int
	urlLimit     *uint
)

func init() {
	environment = flag.String("env", "development", "defines application environment")
	redirectHost = flag.String("redirect-host", "l.hst:4269", "defines the short redirect host")
	port = flag.Int("port", 4269, "port where the app will listened")
	urlLimit = flag.Uint("capacity", 2500, "limit of URLs that the service can contain")

	flag.Parse()
}

func main() {
	ctx := context.Background()
	router := chi.NewRouter()

	messagebus.Start(ctx)
	storage.Start(*environment)
	redirect.Start(*redirectHost)
	ingest.Start(ctx, *urlLimit)

	router.Use(middleware.Logger, redirect.Middleware)

	if *environment == "development" {
		router.Use(devCORSMiddleware)
	}

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

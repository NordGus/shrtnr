package main

import (
	"flag"
	"fmt"
	"github.com/NordGus/shrtnr/database"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPath *string
)

func init() {
	dbPath = flag.String("db-file-path", "./data/shrtnr.db", "path to SQLite DB file")

	flag.Parse()
}

func main() {
	f, err := os.OpenFile(*dbPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS application_migrations (version INTEGER PRIMARY KEY NOT NULL) STRICT;")
	if err != nil {
		log.Fatalln("failed to create application_migrations table")
	}

	err = database.Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("shrtnr: database migrations done!")
}

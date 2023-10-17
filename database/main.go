// Package database is a quick and dirty implementation of a Rails-like migration tool
package database

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

//go:embed migrations
var migrations embed.FS

func Migrate(db *sqlx.DB) error {
	mgrtns, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, mgrtn := range mgrtns {
		if err = runMigration(db, mgrtn); err != nil {
			return err
		}
	}

	return nil
}

func runMigration(db *sqlx.DB, mgrtn fs.DirEntry) error {
	f, err := migrations.Open(fmt.Sprintf("migrations/%v", mgrtn.Name()))
	if err != nil {
		return err
	}

	defer func(f fs.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	ver, err := strconv.ParseInt(strings.Split(mgrtn.Name(), "_")[0], 10, 64)
	if err != nil {
		return err
	}

	err = db.Get(&ver, "SELECT version FROM application_migrations WHERE version = ?;", ver)
	if err == nil {
		return nil
	}

	info, err := f.Stat()
	if err != nil {
		return err
	}

	contents := make([]byte, info.Size())

	_, err = f.Read(contents)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(string(contents))
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalln(err)
		}

		return err
	}

	_, err = tx.Exec("INSERT INTO application_migrations (version) VALUES (?);", ver)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalln(err)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	log.Printf("database: migration %s executed \n", mgrtn.Name())

	return nil
}

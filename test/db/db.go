package db

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DBFunc is a function which takes a database connection
// and returns an error type.
type DBFunc func(*sql.DB) error

// DB takes a database connection and some DBFunc types and returns a
// function, which returns an error.
// This can be used as a setup or teardown test.Action in a Suite.
func DB(db *sql.DB, funcs ...DBFunc) func() error {
	return func() error {
		for _, fn := range funcs {
			if err := fn(db); err != nil {
				return err
			}
		}
		return nil
	}
}

// Seed is a DBFunc which looks for a folder in the current working directory called fixtures
// If it locates a fixtures folder, it locates all files with a .sql prefix, reads the SQL
// queries and executes them against the provided database connection.
func Seed(db *sql.DB) error {
	return filepath.Walk("fixtures", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() != "fixtures" {
			return filepath.SkipDir
		}

		if ext := filepath.Ext(path); ext == ".sql" {
			query, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if _, err := db.Exec(string(query)); err != nil {
				return err
			}
		}

		return nil
	})
}

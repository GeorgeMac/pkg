package test

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DBFunc is a function which takes a database connection
// and returns an error type.
type DBFunc func(*sql.DB) error

// DBConfig is used to configured the setup / teardown of a database
// It is configured using DBOption types
type DBConfig struct {
	setup, teardown []DBFunc
}

// DBOption configures a pointer to a DBConfig
type DBOption func(*DBConfig)

// DB takes a database connection and returns a SuiteOption
// The SuiteOption will ensure setup / teardown around execution
// of the entire packages test suite.
func DB(db *sql.DB, opts ...DBOption) SuiteOption {
	return func(conf *SuiteConfig) {
		conf.db = db
		for _, opt := range opts {
			opt(&conf.dbConf)
		}
	}
}

// Setup takes a DBFunc and ensures it is called before
// a packages tests are performed.
func Setup(fn DBFunc) DBOption {
	return func(conf *DBConfig) {
		conf.setup = append(conf.setup, fn)
	}
}

// Teardown takes DBFunc and ensures it is called after
// a packages tests are completed.
func Teardown(fn DBFunc) DBOption {
	return func(conf *DBConfig) {
		conf.teardown = append(conf.teardown, fn)
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

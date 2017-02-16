package test

import (
	"database/sql"
	"os"
	"testing"
)

// SuiteConfig is used to configure a Suite
// This is done so by passing to SuiteOption types
type SuiteConfig struct {
	db     *sql.DB
	dbConf DBConfig
}

// SuiteOption is used to manipulate and SuiteConfig type
type SuiteOption func(*SuiteConfig)

// Suite is a helper function for performing common setup/tear-down
// around a packages test. It takes a pointer to a testing.M, runs the
// entire packages tests and captures the exit code, then in truncates
// the test database.
func Suite(m *testing.M, opts ...SuiteOption) {
	var conf SuiteConfig
	for _, opt := range opts {
		opt(&conf)
	}

	for _, setup := range conf.dbConf.setup {
		if err := setup(conf.db); err != nil {
			panic(err)
		}
	}

	status := m.Run()

	for _, tear := range conf.dbConf.teardown {
		if err := tear(conf.db); err != nil {
			panic(err)
		}
	}

	os.Exit(status)
}

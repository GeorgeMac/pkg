package test

import (
	"os"
	"testing"
)

// SuiteConfig is used to configure a Suite
// This is done so by passing to SuiteOption types
type SuiteConfig struct {
	setup, teardown Actions
}

// SuiteOption is used to manipulate and SuiteConfig type
type SuiteOption func(*SuiteConfig)

// Setup takes an Action and ensures it is called before
// a packages tests are performed.
func Setup(action Action) SuiteOption {
	return action.AsSetup
}

// Teardown takes Action and ensures it is called after
// a packages tests are completed.
func Teardown(action Action) SuiteOption {
	return action.AsTeardown
}

// Suite is a helper function for performing common setup/tear-down
// around a packages test. It takes a pointer to a testing.M, runs the
// entire packages tests and captures the exit code, then in truncates
// the test database.
func Suite(m *testing.M, opts ...SuiteOption) {
	var conf SuiteConfig
	for _, opt := range opts {
		opt(&conf)
	}

	// run all setup Actions
	conf.setup.runAll()

	// run test suite
	status := m.Run()

	// run all teardown Actions
	conf.teardown.runAll()

	os.Exit(status)
}

// Action is something that can take place before or after tests
// and may return an error
type Action func() error

// AsSetup is a SuiteOption which appends the Action as a setup step
func (a Action) AsSetup(conf *SuiteConfig) { conf.setup.add(a) }

// AsTeardown is a SuiteOption which appends the Action as a teardown step
func (a Action) AsTeardown(conf *SuiteConfig) { conf.teardown.add(a) }

// Actions is a slice of Action
type Actions []Action

func (a *Actions) add(Action Action) {
	*a = append(*a, Action)
}

func (a *Actions) runAll() {
	for _, action := range *a {
		if err := action(); err != nil {
			panic(err)
		}
	}
}

//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/aserto-dev/mage-loot/common"
	"github.com/aserto-dev/mage-loot/deps"
)

func init() {
	// Set private repositories
	os.Setenv("GOPRIVATE", "github.com/aserto-dev")
}

// Lint runs linting for the entire project.
func Lint() error {
	return common.Lint()
}

// Test runs all tests and generates a code coverage report.
func Test() error {
	return common.Test()
}

// Build builds all binaries in ./cmd.
func Build() error {
	return common.BuildReleaser()
}

// Release releases the project.
func Release() error {
	return common.Release()
}

// Deps installs all dependencies required to build the project.
func Deps() {
	deps.GetAllDeps()
}

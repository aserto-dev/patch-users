//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/aserto-dev/mage-loot/common"
	"github.com/aserto-dev/mage-loot/deps"
	"github.com/pkg/errors"
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
	if os.Getenv("GITHUB_TOKEN") == "" {
		return fmt.Errorf("GITHUB_TOKEN environment variable is undefined")
	}

	if os.Getenv("HOMEBREW_TAP") == "" {
		return fmt.Errorf("HOMEBREW_TAP environment variable is undefined")
	}

	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		return fmt.Errorf("GOOGLE_APPLICATION_CREDENTIALS environment variable is undefined")
	}

	if err := writeVersion(); err != nil {
		return err
	}

	return common.Release()
}

// Deps installs all dependencies required to build the project.
func Deps() {
	deps.GetAllDeps()
}

func writeVersion() error {
	version, err := exec.Command("git", "describe", "--tags").Output()
	if err != nil {
		return errors.Wrap(err, "failed to get current git tag")
	}

	file, err := os.Create("VERSION.txt")
	if err != nil {
		return errors.Wrap(err, "failed to create version file")
	}

	defer file.Close()

	if _, err := file.Write(version); err != nil {
		return errors.Wrap(err, "failed to write to version file")
	}

	return file.Sync()
}

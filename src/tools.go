//go:build tools
// +build tools

package src

// When adding a new tool, add the import for it here to make sure it gets a
// version in the go.mod file. Also add a go:generate statement here so that we
// can install all the tools for CI checks and local development just by running
// go generate tools.go.

//go:generate go install golang.org/x/tools/cmd/stringer

import (
	_ "golang.org/x/tools/cmd/stringer"
)

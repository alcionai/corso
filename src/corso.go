package main

import (
	"github.com/alcionai/corso/src/cli"
	prof "github.com/alcionai/corso/src/internal/profile"
)

func main() {
	prof.Profiler()
	cli.Handle()
}

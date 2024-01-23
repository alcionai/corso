package main

import (
	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/cli/debug"
)

func main() {
	// pprof and memstats hooks
	debug.SetupMemoryProfile()

	cli.Handle()
}

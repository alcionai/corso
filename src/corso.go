package main

import (
	"context"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/pkg/profile"
)

var profileTicker = time.NewTicker(60 * time.Second)
var printTicker = time.NewTicker(10 * time.Second)
var profileCounter = 0

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	go func() {
		for {
			select {
			case <-profileTicker.C:
				filename := "mem." + strconv.Itoa(profileCounter) + ".pprof"

				f, _ := os.Create(filename)
				if err := pprof.WriteHeapProfile(f); err != nil {
					log.Fatal("could not write memory profile: ", err)
				}

				f.Close()

				profileCounter++

			}
		}
	}()

	go func() {
		for {
			select {
			case <-printTicker.C:
				PrintMemUsage()
			}
		}
	}()

	cli.Handle()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	ctx := context.Background()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	logger.Ctx(ctx).Info("Alloc = ", bToMb(m.Alloc), " MB")
	logger.Ctx(ctx).Info("Heap = ", bToMb(m.HeapAlloc), " MB")
	logger.Ctx(ctx).Info("TotalAlloc = ", bToMb(m.TotalAlloc), " MB")
	logger.Ctx(ctx).Info("Sys = ", bToMb(m.Sys), " MB")
	logger.Ctx(ctx).Info("NumGC = ", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

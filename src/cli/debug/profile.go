package debug

import (
	"context"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/pkg/profile"

	"github.com/alcionai/corso/src/pkg/logger"
)

var (
	profileTicker    = time.NewTicker(1 * time.Second)
	timeSinceRefresh = time.Now()
	printTicker      = time.NewTicker(1 * time.Second)
	profileCounter   = 0
)

func SetupMemoryProfile() {
	defer profile.Start(profile.MemProfile).Stop()

	// debug.SetMemoryLimit(0.5 * 1024 * 1024 * 1024)

	go func() {
		//nolint:gosimple
		for {
			select {
			case <-profileTicker.C:
				var m runtime.MemStats

				runtime.ReadMemStats(&m)

				// If it's been 3 mins since last pprof capture, take another one.
				if time.Since(timeSinceRefresh) > 3*time.Minute {
					filename := "mem." + strconv.Itoa(profileCounter) + ".pprof"

					f, _ := os.Create(filename)
					if err := pprof.WriteHeapProfile(f); err != nil {
						log.Fatal("could not write memory profile: ", err)
					}

					f.Close()

					profileCounter++

					timeSinceRefresh = time.Now()
				}
			}
		}
	}()

	go func() {
		//nolint:gosimple
		for {
			select {
			case <-printTicker.C:
				PrintMemUsage()
			}
		}
	}()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	ctx := context.Background()

	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	logger.Ctx(ctx).Info("HeapAlloc = ", bToMb(m.HeapAlloc), " MB") // same as Alloc

	logger.Ctx(ctx).Info("HeapReleased = ", bToMb(m.HeapReleased), " MB")
	logger.Ctx(ctx).Info("HeapObjects = ", bToMb(m.HeapObjects), " MB")
	logger.Ctx(ctx).Info("HeapSys = ", bToMb(m.HeapSys), " MB")
	logger.Ctx(ctx).Info("HeapIdle = ", bToMb(m.HeapIdle), " MB")
	logger.Ctx(ctx).Info("HeapInuse = ", bToMb(m.HeapInuse), " MB")

	logger.Ctx(ctx).Info("NumGC = ", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

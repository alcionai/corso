package common

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

const (
	NumItems = 300000
	ItemSize = 1024
	FileName = "input.json"
)

type FooArray struct {
	Entries []*Foo `json:"entries"`
}

type Foo struct {
	ID      string            `json:"id"`
	Labels  map[string]string `json:"labels"`
	ModTime time.Time         `json:"modified"`
	Deleted bool              `json:"deleted,omitempty"`
	Content json.RawMessage   `json:"data"`
}

type Content struct {
	ID   string `json:"id"`
	Data []byte `json:"data"`
}

func PrintMemUsage() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

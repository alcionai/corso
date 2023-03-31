package common

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"time"
)

const (
	NumItems   = 300000
	ItemSize   = 1024
	GzipSuffix = ".gz"
	FileName   = "input.json"
)

var ManifestFileName = fmt.Sprintf("manifest-input.%d.json", NumItems)

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

type Manifest struct {
	Entries []*ManifestEntry `json:"entries"`
}

type ManifestEntry struct {
	ID      string            `json:"id"`
	Labels  map[string]string `json:"labels"`
	ModTime time.Time         `json:"modified"`
	Deleted bool              `json:"deleted,omitempty"`
	Content json.RawMessage   `json:"data"`
}

type SnapManifest struct {
	ID               string            `json:"id"`
	Source           SourceInfo        `json:"source"`
	Description      string            `json:"description"`
	StartTime        int64             `json:"startTime"`
	EndTime          int64             `json:"endTime"`
	Stats            StatsS            `json:"stats,omitempty"`
	IncompleteReason string            `json:"incomplete,omitempty"`
	RootEntry        *DirEntry         `json:"rootEntry"`
	Tags             map[string]string `json:"tags,omitempty"`
}

type SourceInfo struct {
	Host     string `json:"host"`
	UserName string `json:"userName"`
	Path     string `json:"path"`
}

type StatsS struct {
	TotalFileSize         int64 `json:"totalSize"`
	ExcludedTotalFileSize int64 `json:"excludedTotalSize"`
	TotalFileCount        int32 `json:"fileCount"`
	CachedFiles           int32 `json:"cachedFiles"`
	NonCachedFiles        int32 `json:"nonCachedFiles"`
	TotalDirectoryCount   int32 `json:"dirCount"`
	ExcludedFileCount     int32 `json:"excludedFileCount"`
	ExcludedDirCount      int32 `json:"excludedDirCount"`
	IgnoredErrorCount     int32 `json:"ignoredErrorCount"`
	ErrorCount            int32 `json:"errorCount"`
}

type DirEntry struct {
	Name        string            `json:"name,omitempty"`
	EntryType   string            `json:"type,omitempty"`
	Permissions int               `json:"mode,omitempty"`
	FileSize    int64             `json:"size,omitempty"`
	ModTime     int64             `json:"mtime,omitempty"`
	UserID      int32             `json:"uid,omitempty"`
	GroupID     int32             `json:"gid,omitempty"`
	ObjectID    string            `json:"obj,omitempty"`
	DirSummary  *DirectorySummary `json:"summ,omitempty"`
}

type DirectorySummary struct {
	TotalFileSize     int64  `json:"size"`
	TotalFileCount    int64  `json:"files"`
	TotalSymlinkCount int64  `json:"symlinks"`
	TotalDirCount     int64  `json:"dirs"`
	MaxModTime        int64  `json:"maxTime"`
	IncompleteReason  string `json:"incomplete,omitempty"`
	FatalErrorCount   int    `json:"numFailed"`
	IgnoredErrorCount int    `json:"numIgnoredErrors,omitempty"`
}

type Decoder interface {
	ManifestDecoder
	ByteManifestDecoder
}

type ManifestDecoder interface {
	Decode(r io.Reader, gcStats bool) error
}

type ByteManifestDecoder interface {
	DecodeBytes(data []byte, gcStats bool) error
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

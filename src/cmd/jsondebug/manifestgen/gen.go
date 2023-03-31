package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cmd/jsondebug/common"
)

const (
	hostName = "host-name"
	userName = "user-name"
)

func generateManifestEntry() (*common.ManifestEntry, error) {
	snapMan := generateSnapManifest()

	// Base tag set for all snapshots.
	tags := map[string]string{
		"type":     "snapshot",
		"hostname": snapMan.Source.Host,
		"username": snapMan.Source.UserName,
		"path":     snapMan.Source.Path,
	}

	maps.Copy(tags, snapMan.Tags)

	serializedSnapMan, err := json.Marshal(snapMan)
	if err != nil {
		return nil, errors.Wrap(err, "serializing inner struct")
	}

	res := &common.ManifestEntry{
		ID:      randStringLen(32),
		ModTime: time.Now(),
		Deleted: rand.Uint32()&1 != 0,
		Labels:  tags,
		Content: serializedSnapMan,
	}

	return res, nil
}

func generateSnapManifest() common.SnapManifest {
	var incomplete string

	// Roughly 1/4 incomplete.
	if rand.Intn(100) < 25 {
		incomplete = "checkpoint"
	}

	path := randString()

	res := common.SnapManifest{
		Source: common.SourceInfo{
			Host:     hostName,
			UserName: userName,
			Path:     path,
		},
		StartTime: rand.Int63(),
		EndTime:   rand.Int63(),
		Stats: common.StatsS{
			TotalFileSize:         rand.Int63(),
			ExcludedTotalFileSize: int64(rand.Uint32()),
			TotalFileCount:        rand.Int31(),
			CachedFiles:           rand.Int31(),
			NonCachedFiles:        rand.Int31(),
			TotalDirectoryCount:   rand.Int31(),
			ExcludedFileCount:     rand.Int31(),
			ExcludedDirCount:      rand.Int31(),
			IgnoredErrorCount:     rand.Int31(),
			ErrorCount:            rand.Int31(),
		},
		IncompleteReason: incomplete,
		RootEntry: &common.DirEntry{
			Name:        path,
			EntryType:   randStringLen(1),
			Permissions: rand.Intn(512),
			FileSize:    rand.Int63(),
			ModTime:     rand.Int63(),
			UserID:      rand.Int31(),
			GroupID:     rand.Int31(),
			ObjectID:    randStringLen(32),
			DirSummary: &common.DirectorySummary{
				TotalFileSize:     rand.Int63(),
				TotalFileCount:    rand.Int63(),
				TotalSymlinkCount: rand.Int63(),
				TotalDirCount:     rand.Int63(),
				MaxModTime:        rand.Int63(),
				IncompleteReason:  incomplete,
				FatalErrorCount:   rand.Int(),
				IgnoredErrorCount: rand.Int(),
			},
		},
		Tags: map[string]string{
			// User stand-in.
			"tag:" + randStringLen(40): "0",
			"tag:backup-id":            randStringLen(36),
			"tag:is-canon-backup":      "0",
			// Service/data type stand-in.
			"tag:" + randStringLen(20): "0",
		},
	}

	return res
}

var charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func randString() string {
	// String lengths between [10, 128] bytes.
	return randStringLen(rand.Intn(119) + 10)
}

func randStringLen(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(res)
}

func main() {
	data := &common.Manifest{
		Entries: make([]*common.ManifestEntry, 0, common.NumItems),
	}

	for i := 0; i < common.NumItems; i++ {
		entry, err := generateManifestEntry()
		if err != nil {
			fmt.Printf("Error generating random entry: %v\n", err)
			return
		}

		data.Entries = append(data.Entries, entry)
	}

	f, err := os.Create(common.ManifestFileName)
	if err != nil {
		fmt.Printf("Error making regular output file: %v\n", err)
		return
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	if err := enc.Encode(data); err != nil {
		fmt.Printf("Error writing json to regular file: %v\n", err)
		return
	}

	fgz, err := os.Create(common.ManifestFileName + common.GzipSuffix)
	if err != nil {
		fmt.Printf("Error making gzip output file: %v\n", err)
		return
	}

	defer fgz.Close()

	gz := gzip.NewWriter(fgz)
	defer gz.Close()

	enc = json.NewEncoder(gz)
	if err := enc.Encode(data); err != nil {
		fmt.Printf("Error writing json to regular file: %v\n", err)
		return
	}
}

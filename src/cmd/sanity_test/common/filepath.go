package common

import (
	"os"
	"path/filepath"
	"time"

	"github.com/alcionai/clues"
)

func FilepathWalker(
	folderName string,
	exportFileSizes map[string]int64,
	startTime time.Time,
) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return clues.Stack(err)
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(folderName, path)
		if err != nil {
			return clues.Stack(err)
		}

		exportFileSizes[relPath] = info.Size()

		if startTime.After(info.ModTime()) {
			startTime = info.ModTime()
		}

		return nil
	}
}

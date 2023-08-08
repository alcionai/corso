//go:build !windows && !darwin
// +build !windows,!darwin

package logger

import (
	"os"
	"path/filepath"
)

func init() {
	if os.Getenv("XDG_CACHE_HOME") != "" {
		userLogsDir = os.Getenv("XDG_CACHE_HOME")
	} else {
		userLogsDir = filepath.Join(os.Getenv("HOME"), ".cache")
	}
}

package logger

import (
	"os"
	"path/filepath"
)

func init() {
	userLogsDir = filepath.Join(os.Getenv("HOME"), "Library", "Logs")
}

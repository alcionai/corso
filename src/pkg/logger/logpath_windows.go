package logger

import (
	"os"
)

func init() {
	userLogsDir = os.Getenv("LOCALAPPDATA")
}

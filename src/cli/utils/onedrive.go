package utils

import (
	"errors"
)

// ValidateOneDriveRestoreFlags checks common flags for correctness and interdependencies
func ValidateOneDriveRestoreFlags(backupID string) error {
	if len(backupID) == 0 {
		return errors.New("a backup ID is required")
	}

	return nil
}

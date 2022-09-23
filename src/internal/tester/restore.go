package tester

import (
	"github.com/alcionai/corso/src/internal/common"
)

const (
	defaultRestoreContainerPrefix = "Corso_Restore_"
)

func GetDefaultRestoreContainer() string {
	return defaultRestoreContainerPrefix +
		common.FormatNow(common.SimpleDateTimeFormatOneDrive)
}

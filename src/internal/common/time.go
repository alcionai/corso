package common

import (
	"errors"
	"time"
)

// FormatTime produces the standard format for corso time values.
// Always formats into the UTC timezone.
func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// FormatTimeForExchangeFolder produces standard format for
// GraphConnector. Format used on CI testing and default folder
// creation during the restore process
func FormatSimpleDateTime(t time.Time) string {
	timeFolderFormat := "02-Jan-2006_15:04:05"
	return t.UTC().Format(timeFolderFormat)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
func ParseTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot interpret an empty string as time.Time")
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}

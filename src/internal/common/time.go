package common

import (
	"errors"
	"time"
)

const (
	SimpleDateTimeFormat = "02-Jan-2006_15:04:05"
)

// FormatTime produces the standard format for corso time values.
// Always formats into the UTC timezone.
func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// FormatSimpleDateTime produces a simple datetime of the format
// "02-Jan-2006_15:04:05"
func FormatSimpleDateTime(t time.Time) string {
	return t.UTC().Format(SimpleDateTimeFormat)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
func ParseTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot interpret an empty string as time.Time")
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err == nil {
		return t.UTC(), nil
	}
	t, err = time.Parse(SimpleDateTimeFormat, s)
	if err == nil {
		return t.UTC(), nil
	}
	return time.Time{}, errors.New("unable to format time string: " + s)
}

package common

import (
	"errors"
	"time"
)

const (
	StandardTimeFormat                = time.RFC3339Nano
	SimpleDateTimeFormat              = "02-Jan-2006_15:04:05"
	SingleValueExtendedPropertyFormat = time.RFC3339
)

// FormatNow produces the current time in UTC using the provided
// time format.
func FormatNow(fmt string) string {
	return time.Now().UTC().Format(fmt)
}

// FormatTime produces the standard format for corso time values.
// Always formats into the UTC timezone.
func FormatTime(t time.Time) string {
	return t.UTC().Format(StandardTimeFormat)
}

// FormatSimpleDateTime produces a simple datetime of the format
// "02-Jan-2006_15:04:05"
func FormatSimpleDateTime(t time.Time) string {
	return t.UTC().Format(SimpleDateTimeFormat)
}

// FormatLegacyTime produces standard format for string values
// that are placed in SingleValueExtendedProperty tags
func FormatLegacyTime(t time.Time) string {
	return t.UTC().Format(SingleValueExtendedPropertyFormat)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
func ParseTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot interpret an empty string as time.Time")
	}

	t, err := time.Parse(StandardTimeFormat, s)
	if err == nil {
		return t.UTC(), nil
	}

	t, err = time.Parse(SimpleDateTimeFormat, s)
	if err == nil {
		return t.UTC(), nil
	}

	return time.Time{}, errors.New("unable to format time string: " + s)
}

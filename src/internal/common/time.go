package common

import "time"

// FormatTime produces the standard format for corso time values.
// Always formats into the UTC timezone.
func FormatTime(t time.Time) string {
	return t.UTC().Format(time.RFC3339Nano)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
// If it is unable to do so, it returns the epoch time instead.
func ParseTime(s string) time.Time {
	if len(s) == 0 {
		return time.Time{}
	}
	t, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return time.Time{}
	}
	return t.UTC()
}

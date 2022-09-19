package common

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

const (
	// the clipped format occurs when m365 removes the :00 second suffix
	ClippedSimpleTimeFormat = "02-Jan-2006_15:04"
	LegacyTimeFormat        = time.RFC3339
	SimpleDateTimeFormat    = "02-Jan-2006_15:04:05"
	// SimpleDateTimeFormatOneDrive is similar to `SimpleDateTimeFormat`
	// but uses `-` instead of `:` which is a reserved character in
	// OneDrive
	SimpleDateTimeFormatOneDrive = "02-Jan-2006_15-04-05"
	StandardTimeFormat           = time.RFC3339Nano
	TabularOutputTimeFormat      = "2006-01-02T15:04:05Z"
)

var (
	clippedSimpleTimeRE = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}:\d{2}).*`)
	legacyTimeRE        = regexp.MustCompile(
		`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}?([Zz]|[a-zA-Z]{2}|([\+|\-]([01]\d|2[0-3])))).*`)
	simpleDateTimeRE = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}:\d{2}:\d{2}).*`)
	standardTimeRE   = regexp.MustCompile(
		`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?([Zz]|[a-zA-Z]{2}|([\+|\-]([01]\d|2[0-3])))).*`)
	tabularOutputTimeRE = regexp.MustCompile(`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}([Zz]|[a-zA-Z]{2})).*`)
)

var (
	// clipped formats must appear last, else they take priority over the regular Simple format.
	formats = []string{
		StandardTimeFormat, SimpleDateTimeFormat, LegacyTimeFormat, TabularOutputTimeFormat, ClippedSimpleTimeFormat,
	}
	regexes = []*regexp.Regexp{
		standardTimeRE, simpleDateTimeRE, legacyTimeRE, tabularOutputTimeRE, clippedSimpleTimeRE,
	}
)

var ErrNoTimeString = errors.New("no substring contains a known time format")

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

// FormatTabularDisplayTime produces the standard format for displaying
// a timestamp as part of user-readable cli output.
// "2016-01-02T15:04:05Z"
func FormatTabularDisplayTime(t time.Time) string {
	return t.UTC().Format(TabularOutputTimeFormat)
}

// FormatLegacyTime produces standard format for string values
// that are placed in SingleValueExtendedProperty tags
func FormatLegacyTime(t time.Time) string {
	return t.UTC().Format(LegacyTimeFormat)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
func ParseTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot interpret an empty string as time.Time")
	}

	for _, form := range formats {
		t, err := time.Parse(form, s)
		if err == nil {
			return t.UTC(), nil
		}
	}

	return time.Time{}, errors.New("unable to parse time string: " + s)
}

// ExtractTime greedily retrieves a timestamp substring from the provided string.
// returns ErrNoTimeString if no match is found.
func ExtractTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot extract time.Time from an empty string")
	}

	for _, re := range regexes {
		ss := re.FindAllStringSubmatch(s, -1)
		if len(ss) > 0 && len(ss[0]) > 1 {
			return ParseTime(ss[0][1])
		}
	}

	return time.Time{}, errors.Wrap(ErrNoTimeString, s)
}

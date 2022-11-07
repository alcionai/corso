package common

import (
	"regexp"
	"time"

	"github.com/pkg/errors"
)

type TimeFormat string

const (
	// StandardTime is the canonical format used for all data storage in corso
	StandardTime TimeFormat = time.RFC3339Nano

	// DateOnly is accepted by the CLI as a valid input for timestamp-based
	// filters.  Time and timezone are assumed to be 00:00:00 and UTC.
	DateOnly TimeFormat = "2006-01-02"

	// TabularOutput is used when displaying time values to the user in
	// non-json cli outputs.
	TabularOutput TimeFormat = "2006-01-02T15:04:05Z"

	// LegacyTime is used in /exchange/service_restore to comply with certain
	// graphAPI time format requirements.
	LegacyTime TimeFormat = time.RFC3339

	// SimpleDateTime is the default value appended to the root restoration folder name.
	SimpleDateTime TimeFormat = "02-Jan-2006_15:04:05"
	// SimpleDateTimeOneDrive modifies SimpleDateTimeFormat to comply with onedrive folder
	// restrictions: primarily swapping `-` instead of `:` which is a reserved character.
	SimpleDateTimeOneDrive TimeFormat = "02-Jan-2006_15-04-05"

	// m365 will remove the :00 second suffix on folder names, resulting in the following formats.
	ClippedSimple         TimeFormat = "02-Jan-2006_15:04"
	ClippedSimpleOneDrive TimeFormat = "02-Jan-2006_15-04"

	// SimpleTimeTesting is used for testing restore destination folders.
	// Microsecond granularity prevents collisions in parallel package or workflow runs.
	SimpleTimeTesting TimeFormat = SimpleDateTimeOneDrive + ".000000"

	// M365dateTimeTimeZoneTimeFormat is the format used by M365 for datetimetimezone resource
	// https://learn.microsoft.com/en-us/graph/api/resources/datetimetimezone?view=graph-rest-1.0
	M365DateTimeTimeZone TimeFormat = "2006-01-02T15:04:05.000000"
)

// these regexes are used to extract time formats from strings.  Their primary purpose is to
// identify the folders produced in external data during automated testing.  For safety, each
// time format described above should have a matching regexp.
var (
	clippedSimpleRE         = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}:\d{2}).*`)
	clippedSimpleOneDriveRE = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}-\d{2}).*`)
	dateOnlyRE              = regexp.MustCompile(`.*(\d{4}-\d{2}-\d{2}).*`)
	legacyTimeRE            = regexp.MustCompile(
		`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}?([Zz]|[a-zA-Z]{2}|([\+|\-]([01]\d|2[0-3])))).*`)
	simpleDateTimeRE         = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}:\d{2}:\d{2}).*`)
	simpleDateTimeOneDriveRE = regexp.MustCompile(`.*(\d{2}-[a-zA-Z]{3}-\d{4}_\d{2}-\d{2}-\d{2}).*`)
	standardTimeRE           = regexp.MustCompile(
		`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?([Zz]|[a-zA-Z]{2}|([\+|\-]([01]\d|2[0-3])))).*`)
	tabularOutputTimeRE = regexp.MustCompile(`.*(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}([Zz]|[a-zA-Z]{2})).*`)
)

var (
	// shortened formats (clipped*, DateOnly) must follow behind longer formats, otherwise they'll
	// get eagerly chosen as the parsable format, slicing out some data.
	formats = []TimeFormat{
		StandardTime,
		SimpleDateTime,
		SimpleDateTimeOneDrive,
		LegacyTime,
		TabularOutput,
		ClippedSimple,
		ClippedSimpleOneDrive,
		DateOnly,
	}
	regexes = []*regexp.Regexp{
		standardTimeRE,
		simpleDateTimeRE,
		simpleDateTimeOneDriveRE,
		legacyTimeRE,
		tabularOutputTimeRE,
		clippedSimpleRE,
		clippedSimpleOneDriveRE,
		dateOnlyRE,
	}
)

var ErrNoTimeString = errors.New("no substring contains a known time format")

// Now produces the current time as a string in the standard format.
func Now() string {
	return FormatNow(StandardTime)
}

// FormatNow produces the current time in UTC using the provided
// time format.
func FormatNow(fmt TimeFormat) string {
	return FormatTimeWith(time.Now(), fmt)
}

// FormatTimeWith produces the a datetime with the given format.
func FormatTimeWith(t time.Time, fmt TimeFormat) string {
	return t.UTC().Format(string(fmt))
}

// FormatTime produces the standard format for corso time values.
// Always formats into the UTC timezone.
func FormatTime(t time.Time) string {
	return FormatTimeWith(t, StandardTime)
}

// FormatSimpleDateTime produces a simple datetime of the format
// "02-Jan-2006_15:04:05"
func FormatSimpleDateTime(t time.Time) string {
	return FormatTimeWith(t, SimpleDateTime)
}

// FormatTabularDisplayTime produces the standard format for displaying
// a timestamp as part of user-readable cli output.
// "2016-01-02T15:04:05Z"
func FormatTabularDisplayTime(t time.Time) string {
	return FormatTimeWith(t, TabularOutput)
}

// FormatLegacyTime produces standard format for string values
// that are placed in SingleValueExtendedProperty tags
func FormatLegacyTime(t time.Time) string {
	return FormatTimeWith(t, LegacyTime)
}

// ParseTime makes a best attempt to produce a time value from
// the provided string.  Always returns a UTC timezone value.
func ParseTime(s string) (time.Time, error) {
	if len(s) == 0 {
		return time.Time{}, errors.New("cannot interpret an empty string as time.Time")
	}

	for _, form := range formats {
		t, err := time.Parse(string(form), s)
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

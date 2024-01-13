package dttm_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/dttm"
)

type DTTMUnitSuite struct {
	tester.Suite
}

func TestDTTMUnitSuite(t *testing.T) {
	suite.Run(t, &DTTMUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DTTMUnitSuite) TestFormatTime() {
	t := suite.T()
	now := time.Now()
	result := dttm.Format(now)
	assert.Equal(t, now.UTC().Format(time.RFC3339Nano), result)
}

func (suite *DTTMUnitSuite) TestLegacyTime() {
	t := suite.T()
	now := time.Now()
	result := dttm.FormatToLegacy(now)
	assert.Equal(t, now.UTC().Format(time.RFC3339), result)
}

func (suite *DTTMUnitSuite) TestFormatTabularDisplayTime() {
	t := suite.T()
	now := time.Now()
	result := dttm.FormatToTabularDisplay(now)
	assert.Equal(t, now.UTC().Format(string(dttm.TabularOutput)), result)
}

func (suite *DTTMUnitSuite) TestParseTime() {
	t := suite.T()
	now := time.Now()

	nowStr := now.Format(time.RFC3339Nano)
	result, err := dttm.ParseTime(nowStr)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, now.UTC(), result)

	_, err = dttm.ParseTime("")
	require.Error(t, err, clues.ToCore(err))

	_, err = dttm.ParseTime("flablabls")
	require.Error(t, err, clues.ToCore(err))
}

func (suite *DTTMUnitSuite) TestOrNow() {
	var (
		oneMinuteAgo = time.Now().Add(-1 * time.Minute)
		inOneHour    = time.Now().Add(1 * time.Hour)
		empty        time.Time
	)

	table := []struct {
		name   string
		input  time.Time
		expect func(t *testing.T, result time.Time)
	}{
		{
			name:  "a minute ago",
			input: oneMinuteAgo,
			expect: func(t *testing.T, result time.Time) {
				assert.True(t, oneMinuteAgo.Equal(result), "equal")
			},
		},
		{
			name:  "in an hour",
			input: inOneHour,
			expect: func(t *testing.T, result time.Time) {
				assert.True(t, inOneHour.Equal(result), "equal")
			},
		},
		{
			name:  "empty",
			input: empty,
			expect: func(t *testing.T, result time.Time) {
				assert.WithinDuration(t, time.Now(), result, time.Minute)
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), dttm.OrNow(test.input))
		})
	}
}

func (suite *DTTMUnitSuite) TestExtractTime() {
	comparable := func(t *testing.T, tt time.Time, shortFormat dttm.TimeFormat) time.Time {
		ts := dttm.FormatToLegacy(tt.UTC())

		if len(shortFormat) > 0 {
			ts = tt.UTC().Format(string(shortFormat))
		}

		c, err := dttm.ParseTime(ts)

		require.NoError(t, err, clues.ToCore(err))

		return c
	}

	parseT := func(v string) time.Time {
		t, err := time.Parse(time.RFC3339, v)
		require.NoError(suite.T(), err, clues.ToCore(err))

		return t
	}

	inputs := []time.Time{
		time.Now().UTC(),
		time.Now().UTC().Add(-12 * time.Hour),
		parseT("2006-01-02T00:00:00Z"),
		parseT("2006-01-02T12:00:00Z"),
		parseT("2006-01-02T03:01:00Z"),
		parseT("2006-01-02T13:00:02Z"),
		parseT("2006-01-02T03:03:00+01:00"),
		parseT("2006-01-02T03:00:04-01:00"),
	}

	formats := []dttm.TimeFormat{
		dttm.ClippedHuman,
		dttm.ClippedHumanDriveItem,
		dttm.Legacy,
		dttm.HumanReadable,
		dttm.HumanReadableDriveItem,
		dttm.Standard,
		dttm.TabularOutput,
		dttm.SafeForTesting,
		dttm.DateOnly,
	}

	type presuf struct {
		prefix string
		suffix string
	}

	pss := []presuf{
		{"foo", "bar"},
		{"", "bar"},
		{"foo", ""},
		{"", ""},
	}

	type testable struct {
		input         string
		clippedFormat dttm.TimeFormat
		expect        time.Time
	}

	table := []testable{}

	// test matrix: for each input, in each format, with each prefix/suffix, run the test.
	for _, in := range inputs {
		for _, f := range formats {
			shortFormat := f

			if f != dttm.ClippedHuman &&
				f != dttm.ClippedHumanDriveItem &&
				f != dttm.DateOnly {
				shortFormat = ""
			}

			v := dttm.FormatTo(in, f)

			for _, ps := range pss {
				table = append(table, testable{
					input:         ps.prefix + v + ps.suffix,
					expect:        comparable(suite.T(), in, shortFormat),
					clippedFormat: shortFormat,
				})
			}
		}
	}

	for _, test := range table {
		suite.Run(test.input, func() {
			t := suite.T()

			result, err := dttm.ExtractTime(test.input)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect, comparable(t, result, test.clippedFormat))
		})
	}
}

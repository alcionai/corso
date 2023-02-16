package common_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
)

type CommonTimeUnitSuite struct {
	tester.Suite
}

func TestCommonTimeUnitSuite(t *testing.T) {
	s := &CommonTimeUnitSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *CommonTimeUnitSuite) TestFormatTime() {
	t := suite.T()
	now := time.Now()
	result := common.FormatTime(now)
	assert.Equal(t, now.UTC().Format(time.RFC3339Nano), result)
}

func (suite *CommonTimeUnitSuite) TestLegacyTime() {
	t := suite.T()
	now := time.Now()
	result := common.FormatLegacyTime(now)
	assert.Equal(t, now.UTC().Format(time.RFC3339), result)
}

func (suite *CommonTimeUnitSuite) TestFormatTabularDisplayTime() {
	t := suite.T()
	now := time.Now()
	result := common.FormatTabularDisplayTime(now)
	assert.Equal(t, now.UTC().Format(string(common.TabularOutput)), result)
}

func (suite *CommonTimeUnitSuite) TestParseTime() {
	t := suite.T()
	now := time.Now()

	nowStr := now.Format(time.RFC3339Nano)
	result, err := common.ParseTime(nowStr)
	require.NoError(t, err)
	assert.Equal(t, now.UTC(), result)

	_, err = common.ParseTime("")
	require.Error(t, err)

	_, err = common.ParseTime("flablabls")
	require.Error(t, err)
}

func (suite *CommonTimeUnitSuite) TestExtractTime() {
	comparable := func(t *testing.T, tt time.Time, shortFormat common.TimeFormat) time.Time {
		ts := common.FormatLegacyTime(tt.UTC())

		if len(shortFormat) > 0 {
			ts = tt.UTC().Format(string(shortFormat))
		}

		c, err := common.ParseTime(ts)

		require.NoError(t, err)

		return c
	}

	parseT := func(v string) time.Time {
		t, err := time.Parse(time.RFC3339, v)
		require.NoError(suite.T(), err)

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

	formats := []common.TimeFormat{
		common.ClippedSimple,
		common.ClippedSimpleOneDrive,
		common.LegacyTime,
		common.SimpleDateTime,
		common.SimpleDateTimeOneDrive,
		common.StandardTime,
		common.TabularOutput,
		common.SimpleTimeTesting,
		common.DateOnly,
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
		clippedFormat common.TimeFormat
		expect        time.Time
	}

	table := []testable{}

	// test matrix: for each input, in each format, with each prefix/suffix, run the test.
	for _, in := range inputs {
		for _, f := range formats {
			shortFormat := f

			if f != common.ClippedSimple &&
				f != common.ClippedSimpleOneDrive &&
				f != common.DateOnly {
				shortFormat = ""
			}

			v := common.FormatTimeWith(in, f)

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

			result, err := common.ExtractTime(test.input)
			require.NoError(t, err)
			assert.Equal(t, test.expect, comparable(t, result, test.clippedFormat))
		})
	}
}

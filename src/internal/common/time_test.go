package common_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
)

type CommonTimeUnitSuite struct {
	suite.Suite
}

func TestCommonTimeUnitSuite(t *testing.T) {
	suite.Run(t, new(CommonTimeUnitSuite))
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
	assert.Equal(t, now.UTC().Format(common.TabularOutputTimeFormat), result)
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
	clipSimpleTime := func(t string) string {
		return t[:len(t)-3]
	}

	comparable := func(t *testing.T, tt time.Time, clipped bool) time.Time {
		ts := common.FormatLegacyTime(tt.UTC())

		if clipped {
			ts = tt.UTC().Format(common.ClippedSimpleTimeFormat)
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

	type timeFormatter func(time.Time) string

	var (
		clippedF = func(t time.Time) string {
			return clipSimpleTime(common.FormatSimpleDateTime(t))
		}
		legacyF         = common.FormatLegacyTime
		simpleF         = common.FormatSimpleDateTime
		simpleOneDriveF = func(t time.Time) string {
			return common.FormatTimeWith(t, common.SimpleDateTimeFormatOneDrive)
		}
		stdF       = common.FormatTime
		tabularF   = common.FormatTabularDisplayTime
		formatters = []timeFormatter{legacyF, simpleF, simpleOneDriveF, stdF, tabularF, clippedF}
	)

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
		input   string
		expect  time.Time
		clipped bool
	}

	table := []testable{}

	// test matrix: for each input, in each format, with each prefix/suffix, run the test.
	for _, in := range inputs {
		for i, f := range formatters {
			v := f(in)

			for _, ps := range pss {
				table = append(table, testable{
					input:   ps.prefix + v + ps.suffix,
					expect:  comparable(suite.T(), in, i == 4),
					clipped: i == 4,
				})
			}
		}
	}

	for _, test := range table {
		suite.T().Run(test.input, func(t *testing.T) {
			result, err := common.ExtractTime(test.input)
			require.NoError(t, err)
			assert.Equal(t, test.expect, comparable(t, result, test.clipped))
		})
	}
}

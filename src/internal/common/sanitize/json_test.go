package sanitize_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/sanitize"
	"github.com/alcionai/corso/src/internal/tester"
)

type SanitizeJSONUnitSuite struct {
	tester.Suite
}

func TestSanitizeJSONUnitSuite(t *testing.T) {
	suite.Run(t, &SanitizeJSONUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type jsonTest struct {
	name        string
	input       []byte
	expect      []byte
	expectValid assert.BoolAssertionFunc
}

func generateCharacterTests() []jsonTest {
	var (
		res []jsonTest

		baseTestName = "Escape0x%02X"
		baseTestData = `{"foo":"ba%sr"}`
		// The current implementation tranforms most characters to the encoding
		// \u0000 where the 4 0's are hex digits. This is according to the JSON spec
		// in RFC 8259 section 7.
		expect = `{"foo":"ba%s\u00%02Xr"}`
	)

	for i := 0; i < 0x20; i++ {
		// Whitespace characters are tested with manually written tests because they
		// have different handling.
		if i == '\n' || i == '\t' || i == '\r' {
			continue
		}

		res = append(
			res,
			jsonTest{
				name:        fmt.Sprintf(baseTestName, i),
				input:       []byte(fmt.Sprintf(baseTestData, string(rune(i)))),
				expect:      []byte(fmt.Sprintf(expect, "", string(rune(i)))),
				expectValid: assert.True,
			},
			jsonTest{
				name:        fmt.Sprintf(baseTestName, i) + " WithEscapedEscape",
				input:       []byte(fmt.Sprintf(baseTestData, `\\`+string(rune(i)))),
				expect:      []byte(fmt.Sprintf(expect, `\\`, string(rune(i)))),
				expectValid: assert.True,
			})
	}

	return res
}

func (suite *SanitizeJSONUnitSuite) TestJSONBytes() {
	table := []jsonTest{
		{
			name:        "AlreadyValid NoSpecialCharacters",
			input:       []byte(`{"foo":"bar"}`),
			expect:      []byte(`{"foo":"bar"}`),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid EscapedTab",
			input:       []byte("{\"foo\":\"ba\\tr\\\"\"}"),
			expect:      []byte("{\"foo\":\"ba\\tr\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid TabOutsideString",
			input:       []byte("{\"foo\":\t\"bar\\\"\"}"),
			expect:      []byte("{\"foo\":\t\"bar\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid EscapedLinefeed",
			input:       []byte("{\"foo\":\"ba\\nr\\\"\"}"),
			expect:      []byte("{\"foo\":\"ba\\nr\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid LinefeedOutsideString",
			input:       []byte("{\"foo\":\n\"bar\\\"\"}"),
			expect:      []byte("{\"foo\":\n\"bar\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid EscapedCarriageReturn",
			input:       []byte("{\"foo\":\"ba\\rr\\\"\"}"),
			expect:      []byte("{\"foo\":\"ba\\rr\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid CarriageReturnOutsideString",
			input:       []byte("{\"foo\":\r\"bar\\\"\"}"),
			expect:      []byte("{\"foo\":\r\"bar\\\"\"}"),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid UnicodeSequence",
			input:       []byte(`{"foo":"ba\u0008r\""}`),
			expect:      []byte(`{"foo":"ba\u0008r\""}`),
			expectValid: assert.True,
		},
		{
			name:        "AlreadyValid SpecialCharacters",
			input:       []byte(`{"foo":"ba\\r\""}`),
			expect:      []byte(`{"foo":"ba\\r\""}`),
			expectValid: assert.True,
		},
		{
			name: "LF characters in JSON outside quotes",
			input: []byte(`{
				"content":"\n>> ` + "\b\bW" + `"
			}`),
			expect:      nil,
			expectValid: assert.True,
		},
		{
			name:        "No LF characters in JSON",
			input:       []byte(`{"content":"\n>> ` + "\b\bW" + `"}`),
			expect:      nil,
			expectValid: assert.True,
		},
	}

	allTests := append(generateCharacterTests(), table...)

	for _, test := range allTests {
		suite.Run(test.name, func() {
			t := suite.T()

			got := sanitize.JSONBytes(test.input)

			if test.expect != nil {
				assert.Equal(t, test.expect, got)
			}
			test.expectValid(t, json.Valid(got))
		})
	}
}

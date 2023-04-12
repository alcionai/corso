package pii_test

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/pii"
	"github.com/alcionai/corso/src/internal/tester"
)

type URLUnitSuite struct {
	tester.Suite
}

func TestURLUnitSuite(t *testing.T) {
	suite.Run(t, &URLUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *URLUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *URLUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
}

func (suite *URLUnitSuite) TestDoesThings() {
	stubURL := "https://host.com/foo/bar/baz/qux?fnords=smarfs&fnords=brunhilda&beaux=regard"

	table := []struct {
		name      string
		input     string
		expect    string
		safePath  map[string]struct{}
		safeQuery map[string]struct{}
	}{
		{
			name:   "no safety",
			input:  stubURL,
			expect: "https://host.com/***/***/***/***?beaux=***&fnords=***&fnords=***",
		},
		{
			name:     "safe paths",
			input:    stubURL,
			expect:   "https://host.com/foo/***/baz/***?beaux=***&fnords=***&fnords=***",
			safePath: map[string]struct{}{"foo": {}, "baz": {}},
		},
		{
			name:      "safe query",
			input:     stubURL,
			expect:    "https://host.com/***/***/***/***?beaux=regard&fnords=***&fnords=***",
			safeQuery: map[string]struct{}{"beaux": {}},
		},
		{
			name:      "safe path and query",
			input:     stubURL,
			expect:    "https://host.com/foo/***/baz/***?beaux=regard&fnords=***&fnords=***",
			safePath:  map[string]struct{}{"foo": {}, "baz": {}},
			safeQuery: map[string]struct{}{"beaux": {}},
		},
		{
			name:     "empty elements",
			input:    "https://host.com/foo//baz/?fnords=&beaux=",
			expect:   "https://host.com/foo//baz/?beaux=&fnords=",
			safePath: map[string]struct{}{"foo": {}, "baz": {}},
		},
		{
			name:   "no path",
			input:  "https://host.com/",
			expect: "https://host.com/",
		},
		{
			name:   "no path with query",
			input:  "https://host.com/?fnords=smarfs&fnords=brunhilda&beaux=regard",
			expect: "https://host.com/?beaux=***&fnords=***&fnords=***",
		},
		{
			name:   "relative path",
			input:  "/foo/bar/baz/qux?fnords=smarfs&fnords=brunhilda&beaux=regard",
			expect: ":///***/***/***/***?beaux=***&fnords=***&fnords=***",
		},
		{
			name:   "malformed url",
			input:  "i am not a url",
			expect: "://***",
		},
		{
			name:   "empty url",
			input:  "",
			expect: "",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t  = suite.T()
				su = pii.SafeURL{
					URL:           test.input,
					SafePathElems: test.safePath,
					SafeQueryKeys: test.safeQuery,
				}
			)

			result := su.Conceal()
			assert.Equal(t, test.expect, result, "Conceal()")

			result = su.String()
			assert.Equal(t, test.expect, result, "String()")

			result = fmt.Sprintf("%s", su)
			assert.Equal(t, test.expect, result, "fmt %%s")

			result = fmt.Sprintf("%+v", su)
			assert.Equal(t, test.expect, result, "fmt %%+v")
		})
	}
}

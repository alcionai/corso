package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/common"
	"github.com/alcionai/canario/src/internal/tester"
)

type URLUnitSuite struct {
	tester.Suite
}

func TestURLUnitSuite(t *testing.T) {
	suite.Run(t, &URLUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *URLUnitSuite) TestGetQueryParamFromURL() {
	qp := "tempauth"
	table := []struct {
		name           string
		rawURL         string
		queryParam     string
		expectedResult string
		expect         assert.ErrorAssertionFunc
	}{
		{
			name:           "valid",
			rawURL:         "http://localhost:8080?" + qp + "=h.c.s&other=val",
			queryParam:     qp,
			expectedResult: "h.c.s",
			expect:         assert.NoError,
		},
		{
			name:       "query param not found",
			rawURL:     "http://localhost:8080?other=val",
			queryParam: qp,
			expect:     assert.Error,
		},
		{
			name:       "empty query param",
			rawURL:     "http://localhost:8080?" + qp + "=h.c.s&other=val",
			queryParam: "",
			expect:     assert.Error,
		},
		// In case of multiple occurrences, the first occurrence of param is returned.
		{
			name:           "multiple occurrences",
			rawURL:         "http://localhost:8080?" + qp + "=h.c.s&other=val&" + qp + "=h1.c1.s1",
			queryParam:     qp,
			expectedResult: "h.c.s",
			expect:         assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			token, err := common.GetQueryParamFromURL(test.rawURL, test.queryParam)
			test.expect(t, err)

			assert.Equal(t, test.expectedResult, token)
		})
	}
}

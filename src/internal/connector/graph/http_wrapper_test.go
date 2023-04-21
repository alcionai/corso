package graph

import (
	"net/http"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type HTTPWrapperIntgSuite struct {
	tester.Suite
}

func TestHTTPWrapperIntgSuite(t *testing.T) {
	suite.Run(t, &HTTPWrapperIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *HTTPWrapperIntgSuite) TestNewHTTPWrapper() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t  = suite.T()
		hw = NewHTTPWrapper()
	)

	resp, err := hw.Request(
		ctx,
		http.MethodGet,
		"https://www.corsobackup.io",
		nil,
		nil)

	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

package operations

import (
	"context"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/kopia"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type RestoreOpSuite struct {
	suite.Suite
}

func TestRestoreOpSuite(t *testing.T) {
	suite.Run(t, new(RestoreOpSuite))
}

// TODO: after modelStore integration is added, mock the store and/or
// move this to an integration test.
func (suite *RestoreOpSuite) TestRestoreOperation_PersistResults() {
	t := suite.T()
	ctx := context.Background()

	var (
		kw    = &kopia.Wrapper{}
		acct  = account.Account{}
		now   = time.Now()
		stats = restoreStats{
			readErr:  multierror.Append(nil, assert.AnError),
			writeErr: assert.AnError,
			cs:       []connector.DataCollection{&connector.ExchangeDataCollection{}},
			gc: &support.ConnectorOperationStatus{
				ObjectCount: 1,
			},
		}
	)

	op, err := NewRestoreOperation(ctx, Options{}, kw, acct, "foo", nil)
	require.NoError(t, err)

	op.persistResults(now, &stats)

	assert.Equal(t, op.Status, Failed)
	assert.Equal(t, op.Results.ItemsRead, len(stats.cs))
	assert.Equal(t, op.Results.ReadErrors, stats.readErr)
	assert.Equal(t, op.Results.ItemsWritten, stats.gc.ObjectCount)
	assert.Equal(t, op.Results.WriteErrors, stats.writeErr)
	assert.Equal(t, op.Results.StartedAt, now)
	assert.Less(t, now, op.Results.CompletedAt)
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type RestoreOpIntegrationSuite struct {
	suite.Suite
}

func TestRestoreOpIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(ctesting.CorsoCITests); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(RestoreOpIntegrationSuite))
}

func (suite *RestoreOpIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *RestoreOpIntegrationSuite) TestNewRestoreOperation() {
	kw := &kopia.Wrapper{}
	acct, err := ctesting.NewM365Account()
	require.NoError(suite.T(), err)

	table := []struct {
		name     string
		opts     Options
		kw       *kopia.Wrapper
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", Options{}, kw, acct, nil, assert.NoError},
		{"missing kopia", Options{}, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := NewRestoreOperation(
				context.Background(),
				Options{},
				test.kw,
				test.acct,
				"restore-point-id",
				nil)
			test.errCheck(t, err)
		})
	}
}

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
	"github.com/alcionai/corso/internal/kopia"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/account"
)

// ---------------------------------------------------------------------------
// unit
// ---------------------------------------------------------------------------

type BackupOpSuite struct {
	suite.Suite
}

func TestBackupOpSuite(t *testing.T) {
	suite.Run(t, new(BackupOpSuite))
}

// TODO: after modelStore integration is added, mock the store and/or
// move this to an integration test.
func (suite *BackupOpSuite) TestBackupOperation_PersistResults() {
	t := suite.T()
	ctx := context.Background()

	var (
		kw        = &kopia.Wrapper{}
		acct      = account.Account{}
		now       = time.Now()
		cs        = []connector.DataCollection{&connector.ExchangeDataCollection{}}
		readErrs  = multierror.Append(nil, assert.AnError)
		writeErrs = assert.AnError
		stats     = &kopia.BackupStats{
			TotalFileCount: 1,
		}
	)

	op, err := NewBackupOperation(ctx, Options{}, kw, acct, nil)
	require.NoError(t, err)

	op.persistResults(now, cs, stats, readErrs, writeErrs)

	assert.Equal(t, op.Status, Failed)
	assert.Equal(t, op.Results.ItemsRead, len(cs))
	assert.Equal(t, op.Results.ReadErrors, readErrs)
	assert.Equal(t, op.Results.ItemsWritten, stats.TotalFileCount)
	assert.Equal(t, op.Results.WriteErrors, writeErrs)
	assert.Equal(t, op.Results.StartedAt, now)
	assert.Less(t, now, op.Results.CompletedAt)
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	suite.Suite
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(ctesting.CorsoCITests); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(BackupOpIntegrationSuite))
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(
		append(
			ctesting.AWSStorageCredEnvs,
			ctesting.M365AcctCredEnvs...,
		)...,
	)
	require.NoError(suite.T(), err)
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
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
			_, err := NewBackupOperation(
				context.Background(),
				Options{},
				test.kw,
				test.acct,
				nil)
			test.errCheck(t, err)
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run() {
	t := suite.T()
	ctx := context.Background()

	// m365User := "lidiah@8qzvrj.onmicrosoft.com"
	// not the user we want to use, but all the others are
	// suffering from JsonParseNode syndrome
	m365User := "george.martinez@8qzvrj.onmicrosoft.com"
	acct, err := ctesting.NewM365Account()
	require.NoError(t, err)

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	defer k.Close(ctx)

	w, err := kopia.NewWrapper(k)
	require.NoError(t, err)

	bo, err := NewBackupOperation(
		ctx,
		Options{},
		w,
		acct,
		[]string{m365User})
	require.NoError(t, err)

	stats, err := bo.Run(ctx)
	require.NoError(t, err)
	require.NotNil(t, stats)
	assert.Equal(t, bo.Status, Successful)
	assert.Greater(t, stats.TotalFileCount, 0)
	assert.Zero(t, stats.ErrorCount)
}

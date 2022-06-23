package operations_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/operations"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/repository"
)

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
		credentials.TenantID,
		credentials.ClientID,
		credentials.ClientSecret,
	)
	require.NoError(suite.T(), err)
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	kw := &kopia.KopiaWrapper{}
	creds := credentials.GetM365()
	table := []struct {
		name     string
		opts     operations.OperationOpts
		kw       *kopia.KopiaWrapper
		creds    credentials.M365
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", operations.OperationOpts{}, kw, creds, nil, assert.NoError},
		{"missing kopia", operations.OperationOpts{}, nil, creds, nil, assert.Error},
		{"invalid creds", operations.OperationOpts{}, kw, credentials.M365{}, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := operations.NewBackupOperation(
				context.Background(),
				operations.OperationOpts{},
				test.kw,
				test.creds,
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
	m365 := credentials.GetM365()
	acct := repository.Account{
		ClientID:     m365.ClientID,
		ClientSecret: m365.ClientSecret,
		TenantID:     m365.TenantID,
	}

	// need to initialize the repository before we can test connecting to it.
	st, err := ctesting.NewPrefixedS3Storage(t)
	require.NoError(t, err)

	r, err := repository.Initialize(ctx, acct, st)
	require.NoError(t, err)

	bo, err := r.NewBackup(ctx, []string{m365User})
	require.NoError(t, err)

	stats, err := bo.Run(ctx)
	require.NoError(t, err)
	require.NotNil(t, stats)
	assert.Equal(t, bo.Status, operations.Successful)
	assert.Greater(t, stats.TotalFileCount, 0)
	assert.Zero(t, stats.ErrorCount)
}

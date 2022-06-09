package operations_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/operations"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/credentials"
	"github.com/alcionai/corso/pkg/source"
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
	if _, err := ctesting.GetRequiredEnvVars(
		credentials.TenantID,
		credentials.ClientID,
		credentials.ClientSecret,
	); err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	kw := &kopia.KopiaWrapper{}
	creds := credentials.GetM365()
	table := []struct {
		name     string
		opts     operations.OperationOpts
		kw       *kopia.KopiaWrapper
		creds    credentials.M365
		source   *source.Source
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

// todo (rkeepers) - TestBackup_Run()

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
)

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
	_, err := ctesting.GetRequiredEnvVars(
		credentials.TenantID,
		credentials.ClientID,
		credentials.ClientSecret,
	)
	require.NoError(suite.T(), err)
}

func (suite *RestoreOpIntegrationSuite) TestNewRestoreOperation() {
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
			_, err := operations.NewRestoreOperation(
				context.Background(),
				operations.OperationOpts{},
				test.kw,
				test.creds,
				"restore-point-id",
				nil)
			test.errCheck(t, err)
		})
	}
}

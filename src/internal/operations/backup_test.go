package operations_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/kopia"
	"github.com/alcionai/corso/internal/operations"
	ctesting "github.com/alcionai/corso/internal/testing"
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

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	table := []struct {
		name    string
		opts    operations.OperationOpts
		kw      *kopia.KopiaWrapper
		targets []string
	}{
		{"good", operations.OperationOpts{}, new(kopia.KopiaWrapper), nil},
		{"missing kopia", operations.OperationOpts{}, nil, nil},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := operations.NewBackupOperation(
				context.Background(),
				operations.OperationOpts{},
				new(kopia.KopiaWrapper),
				nil)
			assert.NoError(t, err)
		})
	}
}

// todo (rkeepers) - TestBackup_Run()

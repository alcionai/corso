package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/pkg/repository"
	"github.com/alcionai/corso/pkg/storage"
)

type RepositorySuite struct {
	suite.Suite
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func (suite *RepositorySuite) TestInitialize() {
	table := []struct {
		name     string
		storage  storage.Storage
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := repository.Initialize(context.Background(), test.account, test.storage)
			test.errCheck(suite.T(), err, "")
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases from
func (suite *RepositorySuite) TestConnect() {
	table := []struct {
		name     string
		storage  storage.Storage
		account  repository.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			storage.NewStorage(storage.ProviderUnknown),
			repository.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			_, err := repository.Connect(context.Background(), test.account, test.storage)
			test.errCheck(suite.T(), err)
		})
	}
}

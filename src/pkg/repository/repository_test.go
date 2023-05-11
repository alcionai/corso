package repository_test

import (
	"os"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	rep "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

// ---------------
// unit tests
// ---------------

type RepositoryUnitSuite struct {
	tester.Suite
}

func TestRepositoryUnitSuite(t *testing.T) {
	suite.Run(t, &RepositoryUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RepositoryUnitSuite) TestInitialize() {
	table := []struct {
		name     string
		storage  func() (storage.Storage, error)
		account  account.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			account.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err, clues.ToCore(err))

			_, err = repository.Initialize(ctx, test.account, st, control.Defaults())
			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases
func (suite *RepositoryUnitSuite) TestConnect() {
	table := []struct {
		name     string
		storage  func() (storage.Storage, error)
		account  account.Account
		errCheck assert.ErrorAssertionFunc
	}{
		{
			storage.ProviderUnknown.String(),
			func() (storage.Storage, error) {
				return storage.NewStorage(storage.ProviderUnknown)
			},
			account.Account{},
			assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err, clues.ToCore(err))

			_, err = repository.Connect(ctx, test.account, st, control.Defaults())
			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}

// ---------------
// integration tests
// ---------------

type RepositoryIntegrationSuite struct {
	tester.Suite
}

func TestRepositoryIntegrationSuite(t *testing.T) {
	suite.Run(t, &RepositoryIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs}),
	})
}

func (suite *RepositoryIntegrationSuite) TestInitialize() {
	ctx, flush := tester.NewContext()
	defer flush()

	table := []struct {
		name     string
		account  account.Account
		storage  func(*testing.T) storage.Storage
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:     "success",
			storage:  tester.NewPrefixedS3Storage,
			errCheck: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			st := test.storage(t)
			r, err := repository.Initialize(ctx, test.account, st, control.Defaults())
			if err == nil {
				defer func() {
					err := r.Close(ctx)
					assert.NoError(t, err, clues.ToCore(err))
				}()
			}

			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}

const (
	roleARNEnvKey = "CORSO_TEST_S3_ROLE"
	roleDuration  = time.Minute * 20
)

func (suite *RepositoryIntegrationSuite) TestInitializeWithRole() {
	if _, ok := os.LookupEnv(roleARNEnvKey); !ok {
		suite.T().Skip(roleARNEnvKey + " not set")
	}

	ctx, flush := tester.NewContext()
	defer flush()

	st := tester.NewPrefixedS3Storage(suite.T())

	st.Role = os.Getenv(roleARNEnvKey)
	st.SessionName = "corso-repository-test"
	st.SessionDuration = roleDuration.String()

	r, err := repository.Initialize(ctx, account.Account{}, st, control.Options{})
	require.NoError(suite.T(), err)

	defer func() {
		r.Close(ctx)
	}()
}

func (suite *RepositoryIntegrationSuite) TestConnect() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	_, err := repository.Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	// now re-connect
	_, err = repository.Connect(ctx, account.Account{}, st, control.Defaults())
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *RepositoryIntegrationSuite) TestConnect_sameID() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	oldID := r.GetID()

	err = r.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// now re-connect
	r, err = repository.Connect(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, oldID, r.GetID())
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	userID := tester.M365UserID(t)

	bo, err := r.NewBackup(ctx, selectors.Selector{DiscreteOwner: userID})
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)
	dest := tester.DefaultTestRestoreDestination("")

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	ro, err := r.NewRestore(ctx, "backup-id", selectors.Selector{DiscreteOwner: "test"}, dest)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, ro)
}

func (suite *RepositoryIntegrationSuite) TestNewMaintenance() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	mo, err := r.NewMaintenance(ctx, rep.Maintenance{})
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, mo)
}

func (suite *RepositoryIntegrationSuite) TestConnect_DisableMetrics() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	_, err := repository.Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err)

	// now re-connect
	r, err := repository.Connect(ctx, account.Account{}, st, control.Options{DisableMetrics: true})
	assert.NoError(t, err)

	assert.Equal(t, "", r.GetID())
}

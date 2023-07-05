package repository

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
	ctrlRepo "github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/extensions"
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

			ctx, flush := tester.NewContext(t)
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err, clues.ToCore(err))

			_, err = Initialize(ctx, test.account, st, control.Defaults())
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

			ctx, flush := tester.NewContext(t)
			defer flush()

			st, err := test.storage()
			assert.NoError(t, err, clues.ToCore(err))

			_, err = Connect(ctx, test.account, st, "not_found", control.Defaults())
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

			ctx, flush := tester.NewContext(t)
			defer flush()

			st := test.storage(t)
			r, err := Initialize(ctx, test.account, st, control.Defaults())
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

	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	st := tester.NewPrefixedS3Storage(suite.T())

	st.Role = os.Getenv(roleARNEnvKey)
	st.SessionName = "corso-repository-test"
	st.SessionDuration = roleDuration.String()

	r, err := Initialize(ctx, account.Account{}, st, control.Options{})
	require.NoError(suite.T(), err)

	defer func() {
		r.Close(ctx)
	}()
}

func (suite *RepositoryIntegrationSuite) TestConnect() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	repo, err := Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	// now re-connect
	_, err = Connect(ctx, account.Account{}, st, repo.GetID(), control.Defaults())
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *RepositoryIntegrationSuite) TestConnect_sameID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	oldID := r.GetID()

	err = r.Close(ctx)
	require.NoError(t, err, clues.ToCore(err))

	// now re-connect
	r, err = Connect(ctx, account.Account{}, st, oldID, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, oldID, r.GetID())
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	userID := tester.M365UserID(t)

	bo, err := r.NewBackup(ctx, selectors.Selector{DiscreteOwner: userID})
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(t)
	restoreCfg := testdata.DefaultRestoreConfig("")

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	ro, err := r.NewRestore(ctx, "backup-id", selectors.Selector{DiscreteOwner: "test"}, restoreCfg)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, ro)
}

func (suite *RepositoryIntegrationSuite) TestNewMaintenance() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := Initialize(ctx, acct, st, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	mo, err := r.NewMaintenance(ctx, ctrlRepo.Maintenance{})
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, mo)
}

func (suite *RepositoryIntegrationSuite) TestConnect_DisableMetrics() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	repo, err := Initialize(ctx, account.Account{}, st, control.Defaults())
	require.NoError(t, err)

	// now re-connect
	r, err := Connect(ctx, account.Account{}, st, repo.GetID(), control.Options{DisableMetrics: true})
	assert.NoError(t, err)

	// now we have repoID beforehand
	assert.Equal(t, r.GetID(), r.GetID())
}

// Test_Options tests that the options are passed through to the repository
// correctly
func (suite *RepositoryIntegrationSuite) Test_Options() {
	table := []struct {
		name        string
		opts        func() control.Options
		expectedLen int
	}{
		{
			name: "default options",
			opts: func() control.Options {
				return control.Defaults()
			},
			expectedLen: 0,
		},
		{
			name: "options with an extension factory",
			opts: func() control.Options {
				o := control.Defaults()
				o.ItemExtensionFactory = append(
					o.ItemExtensionFactory,
					&extensions.MockItemExtensionFactory{})

				return o
			},
			expectedLen: 1,
		},
		{
			name: "options with multiple extension factories",
			opts: func() control.Options {
				o := control.Defaults()
				f := []extensions.CreateItemExtensioner{
					&extensions.MockItemExtensionFactory{},
					&extensions.MockItemExtensionFactory{},
				}

				o.ItemExtensionFactory = f

				return o
			},
			expectedLen: 2,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			acct := tester.NewM365Account(t)
			st := tester.NewPrefixedS3Storage(t)

			ctx, flush := tester.NewContext(t)
			defer flush()

			repo, err := Initialize(ctx, acct, st, test.opts())
			require.NoError(t, err)

			r := repo.(*repository)
			assert.Equal(t, test.expectedLen, len(r.Opts.ItemExtensionFactory))

			repo, err = Connect(ctx, acct, st, repo.GetID(), test.opts())
			assert.NoError(t, err)

			r = repo.(*repository)
			assert.Equal(t, test.expectedLen, len(r.Opts.ItemExtensionFactory))
		})
	}
}

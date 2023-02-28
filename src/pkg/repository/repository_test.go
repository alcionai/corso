package repository_test

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/storage"
)

// ---------------
// unit tests
// ---------------

type RepositorySuite struct {
	tester.Suite
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, &RepositorySuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RepositorySuite) TestInitialize() {
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
			assert.NoError(t, err)
			_, err = repository.Initialize(ctx, test.account, st, control.Options{})
			test.errCheck(t, err, "")
		})
	}
}

// repository.Connect involves end-to-end communication with kopia, therefore this only
// tests expected error cases
func (suite *RepositorySuite) TestConnect() {
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
			assert.NoError(t, err)
			_, err = repository.Connect(ctx, test.account, st, control.Options{})
			test.errCheck(t, err)
		})
	}
}

// ---------------
// integration tests
// ---------------

type RepositoryIntegrationSuite struct {
	tester.Suite
	userID string
}

func TestRepositoryIntegrationSuite(t *testing.T) {
	suite.Run(t, &RepositoryIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs},
			tester.CorsoRepositoryTests,
		),
	})
}

func (suite *RepositoryIntegrationSuite) SetupSuite() {
	suite.userID = tester.M365UserID(suite.T())
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
			r, err := repository.Initialize(ctx, test.account, st, control.Options{})
			if err == nil {
				defer func() {
					assert.NoError(t, r.Close(ctx))
				}()
			}

			test.errCheck(t, err)
		})
	}
}

func (suite *RepositoryIntegrationSuite) TestConnect() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	_, err := repository.Initialize(ctx, account.Account{}, st, control.Options{})
	require.NoError(t, err)

	// now re-connect
	_, err = repository.Connect(ctx, account.Account{}, st, control.Options{})
	assert.NoError(t, err)
}

func (suite *RepositoryIntegrationSuite) TestConnect_sameID() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, account.Account{}, st, control.Options{})
	require.NoError(t, err)

	oldID := r.GetID()

	require.NoError(t, r.Close(ctx))

	// now re-connect
	r, err = repository.Connect(ctx, account.Account{}, st, control.Options{})
	require.NoError(t, err)
	assert.Equal(t, oldID, r.GetID())
}

func (suite *RepositoryIntegrationSuite) TestNewBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err)

	bo, err := r.NewBackup(ctx, selectors.Selector{DiscreteOwner: "test"})
	require.NoError(t, err)
	require.NotNil(t, bo)
}

func (suite *RepositoryIntegrationSuite) TestNewRestore() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	acct := tester.NewM365Account(t)
	dest := tester.DefaultTestRestoreDestination()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	r, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err)

	ro, err := r.NewRestore(ctx, "backup-id", selectors.Selector{DiscreteOwner: "test"}, dest)
	require.NoError(t, err)
	require.NotNil(t, ro)
}

func (suite *RepositoryIntegrationSuite) TestBackupDetails_regression() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		st   = tester.NewPrefixedS3Storage(t)
		dest = "Corso_Restore_empty_" + common.FormatNow(common.SimpleTimeTesting)
	)

	m365, err := acct.M365Config()
	require.NoError(t, err)

	adpt, err := graph.CreateAdapter(acct.ID(), m365.AzureClientID, m365.AzureClientSecret)
	require.NoError(t, err)

	srv := graph.NewService(adpt)

	pager, err := onedrive.PagerForSource(onedrive.OneDriveSource, srv, suite.userID, nil)
	require.NoError(t, err)

	drives, err := onedrive.Drives(ctx, pager, false)
	require.NoError(t, err)

	d0 := drives[0]
	body := models.DriveItem{}
	body.SetName(&dest)

	fld := models.Folder{}
	fld.SetChildCount(ptr.To[int32](0))
	body.SetFolder(&fld)

	_, err = srv.Client().
		UsersById(suite.userID).
		DrivesById(*d0.GetId()).
		Items().
		Post(ctx, &body, nil)
	require.NoErrorf(t, err, "%+v", graph.ErrData(err))

	r, err := repository.Initialize(ctx, acct, st, control.Options{})
	require.NoError(t, err)

	sel := selectors.NewOneDriveBackup([]string{suite.userID})
	sel.Include(sel.Folders([]string{dest}))

	op, err := r.NewBackup(ctx, sel.Selector)
	require.NoError(t, err)
	require.NoError(t, op.Run(ctx))
	require.NotZero(t, op.Results.ItemsWritten)

	// the actual test.  The backup details, having backed up an empty folder,
	// should not return the folder within the backup details.  That value
	// should get filtered out, along with .meta and .dirmeta files.
	deets, _, ferr := r.BackupDetails(ctx, string(op.Results.BackupID))
	require.NoError(t, ferr.Failure())

	for _, ent := range deets.Entries {
		assert.NotContains(t, ent.RepoRef, dest)
		assert.NotContains(t, ent.RepoRef, onedrive.MetaFileSuffix)
		assert.NotContains(t, ent.RepoRef, onedrive.DirMetaFileSuffix)
	}
}

package operations

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/graph"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
	"github.com/alcionai/corso/src/pkg/store"
)

func getKopiaHandles(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
) (*kopia.Wrapper, *kopia.ModelStore) {
	st := storeTD.NewPrefixedS3Storage(t)
	k := kopia.NewConn(st)
	err := k.Initialize(ctx, repository.Options{}, repository.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	kw, err := kopia.NewWrapper(k)
	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	k.Close(ctx)

	require.NoError(t, err, "getting kopia wrapper: %s", clues.ToCore(err))

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err, "getting model store: %s", clues.ToCore(err))

	return kw, ms
}

type MaintenanceOpIntegrationSuite struct {
	tester.Suite
}

func TestMaintenanceOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &MaintenanceOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MaintenanceOpIntegrationSuite) TestRepoMaintenance() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	kw, ms := getKopiaHandles(t, ctx)

	defer kw.Close(ctx)
	defer ms.Close(ctx)

	mo, err := NewMaintenanceOperation(
		ctx,
		control.DefaultOptions(),
		kw,
		store.NewWrapper(ms),
		repository.Maintenance{
			Type: repository.MetadataMaintenance,
		},
		evmock.NewBus())
	require.NoError(t, err, clues.ToCore(err))

	err = mo.Run(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}

func TestMaintenanceOpNightlyIntegrationSuite(t *testing.T) {
	suite.Run(t, &MaintenanceOpIntegrationSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MaintenanceOpIntegrationSuite) TestRepoMaintenance_GarbageCollection() {
	var (
		t        = suite.T()
		acct     = tconfig.NewM365Account(suite.T())
		tenantID = acct.Config[config.AzureTenantIDKey]
		opts     = control.DefaultOptions()
		osel     = selectors.NewOneDriveBackup([]string{userID})
		// Default policy used by SDK clients
		failurePolicy = control.FailAfterRecovery
		T1            = time.Now().Truncate(0)
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	pathElements := []string{odConsts.DrivesPathDir, "drive-id", odConsts.RootPathDir, folderID}

	tmp, err := path.Build(tenantID, userID, path.OneDriveService, path.FilesCategory, false, pathElements...)
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPath := path.Builder{}.Append(tmp.Folders()...)

	kw, ms := getKopiaHandles(t, ctx)
	storer := store.NewWrapper(ms)

	var bupIDs []model.StableID

	// Make two failed backups so the garbage collection code will try to delete
	// something.
	for i := 0; i < 2; i++ {
		suite.Run(fmt.Sprintf("Setup%d", i), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			cs := []data.BackupCollection{
				makeBackupCollection(
					tmp,
					locPath,
					[]dataMock.Item{
						makeMockItem("file1", nil, T1, false, nil),
						makeMockItem("file2", nil, T1, false, assert.AnError),
					}),
			}

			mc, err := graph.MakeMetadataCollection(
				tenantID,
				userID,
				path.OneDriveService,
				path.FilesCategory,
				makeMetadataCollectionEntries("url/1", driveID, folderID, tmp),
				func(*support.ControllerOperationStatus) {})
			require.NoError(t, err, clues.ToCore(err))

			cs = append(cs, mc)
			bp := &mockBackupProducer{
				colls: cs,
			}

			opts.FailureHandling = failurePolicy

			bo, err := NewBackupOperation(
				ctx,
				opts,
				kw,
				storer,
				bp,
				acct,
				osel.Selector,
				selectors.Selector{DiscreteOwner: userID},
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			err = bo.Run(ctx)
			assert.Error(t, err, clues.ToCore(err))

			require.NotEmpty(t, bo.Results.BackupID)

			bupIDs = append(bupIDs, bo.Results.BackupID)
		})
	}

	// Double check we have two backup models. This is not an exhaustive check but
	// will give us some comfort that things are working as expected.
	bups, err := storer.GetBackups(ctx)
	require.NoError(
		t,
		err,
		"checking backup model existence: %s",
		clues.ToCore(err))

	var gotBupIDs []model.StableID

	for _, bup := range bups {
		gotBupIDs = append(gotBupIDs, bup.ID)
	}

	require.ElementsMatch(t, bupIDs, gotBupIDs)

	// Run maintenance with garbage collection.

	suite.Run("RunMaintenance", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mo, err := NewMaintenanceOperation(
			ctx,
			control.DefaultOptions(),
			kw,
			store.NewWrapper(ms),
			repository.Maintenance{
				Type: repository.CompletePlusMaintenance,
				// Set buffer to 0 so things will actually be garbage collected.
				CleanupBuffer: ptr.To(time.Duration(0)),
			},
			evmock.NewBus())
		require.NoError(t, err, clues.ToCore(err))

		err = mo.Run(ctx)
		assert.NoError(t, err, clues.ToCore(err))

		// Check for backup models again. Only the second one should still be present.
		bups, err = storer.GetBackups(ctx)
		require.NoError(
			t,
			err,
			"checking backup model existence after maintenance: %s",
			clues.ToCore(err))

		gotBupIDs = nil

		for _, bup := range bups {
			gotBupIDs = append(gotBupIDs, bup.ID)
		}

		assert.ElementsMatch(t, bupIDs[1:], gotBupIDs)
	})
}

package kopia

import (
	"context"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/data"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

func BenchmarkHierarchyMerge(b *testing.B) {
	ctx, flush := tester.NewContext(b)
	defer flush()

	c, err := openKopiaRepo(b, ctx)
	require.NoError(b, err, clues.ToCore(err))

	w := &Wrapper{c}

	defer func() {
		err := w.Close(ctx)
		assert.NoError(b, err, clues.ToCore(err))
	}()

	var (
		cols                 []data.BackupCollection
		collectionLimit      = 1000
		collectionItemsLimit = 3
		itemData             = []byte("abcdefghijklmnopqrstuvwxyz")
	)

	baseStorePath, err := path.Build(
		"a-tenant",
		"a-user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"Inbox")
	require.NoError(b, err, clues.ToCore(err))

	for i := 0; i < collectionLimit; i++ {
		folderName := fmt.Sprintf("folder%d", i)

		storePath, err := baseStorePath.Append(false, folderName)
		require.NoError(b, err, clues.ToCore(err))

		col := exchMock.NewCollection(
			storePath,
			storePath,
			collectionItemsLimit)

		for j := 0; j < collectionItemsLimit; j++ {
			itemName := fmt.Sprintf("item%d", j)
			col.Names[j] = itemName
			col.Data[j] = itemData
		}

		cols = append(cols, col)
	}

	reasons := []identity.Reasoner{
		identity.NewReason(
			testTenant,
			baseStorePath.ProtectedResource(),
			baseStorePath.Service(),
			baseStorePath.Category()),
	}

	type testCase struct {
		name        string
		baseBackups func(base BackupBase) BackupBases
		collections []data.BackupCollection
	}

	// Initial backup. All files should be considered new by kopia.
	baseBackupCase := testCase{
		name: "Setup",
		baseBackups: func(BackupBase) BackupBases {
			return NewMockBackupBases()
		},
		collections: cols,
	}

	runAndTestBackup := func(
		t tester.TestT,
		ctx context.Context,
		test testCase,
		base BackupBase,
	) BackupBase {
		bbs := test.baseBackups(base)
		counter := count.New()

		stats, _, _, err := w.ConsumeBackupCollections(
			ctx,
			reasons,
			bbs,
			test.collections,
			nil,
			nil,
			true,
			counter,
			fault.New(true))
		require.NoError(t, err, clues.ToCore(err))

		assert.Zero(t, stats.IgnoredErrorCount)
		assert.Zero(t, stats.ErrorCount)
		assert.Zero(t, counter.Get(count.PersistenceIgnoredErrs))
		assert.Zero(t, counter.Get(count.PersistenceErrs))
		assert.False(t, stats.Incomplete)

		snap, err := snapshot.LoadSnapshot(
			ctx,
			w.c,
			manifest.ID(stats.SnapshotID))
		require.NoError(t, err, clues.ToCore(err))

		return BackupBase{
			ItemDataSnapshot: snap,
			Reasons:          reasons,
		}
	}

	b.Logf("setting up base backup\n")

	base := runAndTestBackup(b, ctx, baseBackupCase, BackupBase{})

	table := []testCase{
		{
			name: "Merge All",
			baseBackups: func(base BackupBase) BackupBases {
				return NewMockBackupBases().WithMergeBases(base)
			},
			collections: func() []data.BackupCollection {
				p, err := baseStorePath.Dir()
				require.NoError(b, err, clues.ToCore(err))

				col := exchMock.NewCollection(p, p, 0)
				col.ColState = data.NotMovedState
				col.PrevPath = p

				return []data.BackupCollection{col}
			}(),
		},
	}

	b.ResetTimer()

	for _, test := range table {
		b.Run(fmt.Sprintf("num_dirs_%d", collectionLimit), func(b *testing.B) {
			ctx, flush := tester.NewContext(b)
			defer flush()

			for i := 0; i < b.N; i++ {
				runAndTestBackup(b, ctx, test, base)
			}
		})
	}
}

package kopia

import (
	"bytes"
	"context"
	"io"
	stdpath "path"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	pmMock "github.com/alcionai/corso/src/internal/common/prefixmatcher/mock"
	"github.com/alcionai/corso/src/internal/data"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

func makePath(t *testing.T, elements []string, isItem bool) path.Path {
	p, err := path.FromDataLayerPath(stdpath.Join(elements...), isItem)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func newExpectedFile(name string, fileData []byte) *expectedNode {
	return &expectedNode{
		name:     name,
		data:     fileData,
		children: []*expectedNode{},
	}
}

// baseWithChildren returns an fs.Entry hierarchy where the first len(basic)
// levels are the encoded values of basic in order. All items in children are
// used as the direct descendents of the final entry in basic.
func baseWithChildren(
	basic []string,
	children []fs.Entry,
) fs.Entry {
	if len(basic) == 0 {
		return nil
	}

	if len(basic) == 1 {
		return virtualfs.NewStaticDirectory(
			encodeElements(basic[0])[0],
			children)
	}

	return virtualfs.NewStaticDirectory(
		encodeElements(basic[0])[0],
		[]fs.Entry{
			baseWithChildren(basic[1:], children),
		})
}

type expectedNode struct {
	name     string
	children []*expectedNode
	data     []byte
}

// expectedTreeWithChildren returns an expectedNode hierarchy where the first
// len(basic) levels are the values of basic in order. All items in children are
// made a direct descendent of the final entry in basic.
func expectedTreeWithChildren(
	basic []string,
	children []*expectedNode,
) *expectedNode {
	if len(basic) == 0 {
		return nil
	}

	if len(basic) == 1 {
		return &expectedNode{
			name:     basic[0],
			children: children,
		}
	}

	return &expectedNode{
		name: basic[0],
		children: []*expectedNode{
			expectedTreeWithChildren(basic[1:], children),
		},
	}
}

// Currently only works for files that Corso has serialized as it expects a
// version specifier at the start of the file.
func expectFileData(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	expected []byte,
	f fs.StreamingFile,
) {
	t.Helper()

	if len(expected) == 0 {
		return
	}

	name, err := decodeElement(f.Name())
	if err != nil {
		name = f.Name()
	}

	r, err := f.GetReader(ctx)
	if !assert.NoError(t, err, "getting reader for file:", name, clues.ToCore(err)) {
		return
	}

	got, err := io.ReadAll(r)
	if !assert.NoError(t, err, "reading data in file", name, clues.ToCore(err)) {
		return
	}

	assert.Equal(t, expected, got, "data in file", name, clues.ToCore(err))
}

func expectTree(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	expected *expectedNode,
	got fs.Entry,
) {
	t.Helper()

	if expected == nil {
		return
	}

	names := make([]string, 0, len(expected.children))
	mapped := make(map[string]*expectedNode, len(expected.children))

	for _, child := range expected.children {
		encoded := encodeElements(child.name)[0]

		names = append(names, encoded)
		mapped[encoded] = child
	}

	entries := getDirEntriesForEntry(t, ctx, got)
	expectDirs(t, entries, names, true)

	for _, e := range entries {
		expectedSubtree := mapped[e.Name()]
		if !assert.NotNil(t, expectedSubtree) {
			continue
		}

		if f, ok := e.(fs.StreamingFile); ok {
			expectFileData(t, ctx, expectedSubtree.data, f)
			continue
		}

		dir, ok := e.(fs.Directory)
		if !ok {
			continue
		}

		expectTree(t, ctx, expectedSubtree, dir)
	}
}

func expectDirs(
	t *testing.T,
	entries []fs.Entry,
	dirs []string,
	exactly bool,
) {
	t.Helper()

	ents := make([]string, 0, len(entries))
	for _, e := range entries {
		ents = append(ents, e.Name())
	}

	dd, err := decodeElements(dirs...)
	require.NoError(t, err, clues.ToCore(err))

	de, err := decodeElements(ents...)
	require.NoError(t, err, clues.ToCore(err))

	if exactly {
		require.Lenf(t, entries, len(dirs), "expected exactly %+v\ngot %+v", dd, de)
	}

	assert.Subsetf(t, dirs, ents, "expected at least %+v\ngot %+v", dd, de)
}

func getDirEntriesForEntry(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	entry fs.Entry,
) []fs.Entry {
	d, ok := entry.(fs.Directory)
	require.True(t, ok, "entry is not a directory")

	entries, err := fs.GetAllEntries(ctx, d)
	require.NoError(t, err, clues.ToCore(err))

	return entries
}

// ---------------
// unit tests
// ---------------
type CorsoProgressUnitSuite struct {
	tester.Suite
	targetFilePath path.Path
	targetFileLoc  *path.Builder
	targetFileName string
}

func TestCorsoProgressUnitSuite(t *testing.T) {
	suite.Run(t, &CorsoProgressUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CorsoProgressUnitSuite) SetupSuite() {
	p, err := path.Build(
		testTenant,
		testUser,
		path.ExchangeService,
		path.EmailCategory,
		true,
		testInboxDir, "testFile")
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.targetFilePath = p
	suite.targetFileLoc = path.Builder{}.Append(testInboxDir)
	suite.targetFileName = suite.targetFilePath.ToBuilder().Dir().String()
}

var _ data.ItemInfo = &mockExchangeMailInfoer{}

type mockExchangeMailInfoer struct{}

func (m mockExchangeMailInfoer) Info() (details.ItemInfo, error) {
	return details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			ItemType: details.ExchangeMail,
		},
	}, nil
}

type testInfo struct {
	info       *itemDetails
	err        error
	totalBytes int64
}

var finishedFileTable = []struct {
	name               string
	cachedItems        func(fname string, fpath path.Path) map[string]testInfo
	expectedBytes      int64
	expectedNumEntries int
	// Non-folder items.
	expectedNumItems int
	err              error
}{
	{
		name: "DetailsExist",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info: &itemDetails{
						infoer:       mockExchangeMailInfoer{},
						repoPath:     fpath,
						locationPath: path.Builder{}.Append(fpath.Folders()...),
					},
					err:        nil,
					totalBytes: 100,
				},
			}
		},
		expectedBytes: 100,
		// 1 file and 5 folders.
		expectedNumEntries: 2,
		expectedNumItems:   1,
	},
	{
		name: "PendingNoDetails",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info: nil,
					err:  nil,
				},
			}
		},
		expectedNumEntries: 0,
	},
	{
		name: "HadError",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info: &itemDetails{
						infoer:   mockExchangeMailInfoer{},
						repoPath: fpath,
					},
					err: assert.AnError,
				},
			}
		},
		expectedNumEntries: 0,
	},
	{
		name: "NotPending",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return nil
		},
		expectedNumEntries: 0,
	},
}

func (suite *CorsoProgressUnitSuite) TestFinishedFile() {
	table := []struct {
		name                 string
		cached               bool
		differentPrevPath    bool
		dropInfo             bool
		expectToMergeEntries bool
	}{
		{
			name:   "all updated",
			cached: false,
		},
		{
			name:                 "all cached from assist base",
			cached:               true,
			expectToMergeEntries: true,
		},
		{
			name:                 "all cached from merge base",
			cached:               true,
			differentPrevPath:    true,
			dropInfo:             true,
			expectToMergeEntries: true,
		},
		{
			name:                 "all not cached from merge base",
			cached:               false,
			differentPrevPath:    true,
			dropInfo:             true,
			expectToMergeEntries: true,
		},
	}

	for _, cachedTest := range table {
		suite.Run(cachedTest.name, func() {
			for _, test := range finishedFileTable {
				suite.Run(test.name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					bd := &details.Builder{}
					cp := corsoProgress{
						ctx:            ctx,
						UploadProgress: &snapshotfs.NullUploadProgress{},
						deets:          bd,
						toMerge:        newMergeDetails(),
						pending:        map[string]*itemDetails{},
						errs:           fault.New(true),
					}

					ci := test.cachedItems(suite.targetFileName, suite.targetFilePath)

					for k, v := range ci {
						if v.info != nil {
							v.info.prevPath = v.info.repoPath

							if cachedTest.differentPrevPath {
								// Doesn't really matter how we change the path as long as it's
								// different somehow.
								p, err := path.FromDataLayerPath(
									suite.targetFilePath.String()+"2",
									true)
								require.NoError(
									t,
									err,
									"making prevPath: %v",
									clues.ToCore(err))

								v.info.prevPath = p
							}

							if cachedTest.dropInfo {
								v.info.infoer = nil
							}
						}

						cp.put(k, v.info)
					}

					require.Len(t, cp.pending, len(ci))

					foundItems := map[string]bool{}

					for k, v := range ci {
						if cachedTest.cached {
							cp.CachedFile(k, v.totalBytes)
						}

						if v.info != nil && v.info.repoPath != nil {
							foundItems[v.info.repoPath.Item()] = false
						}

						cp.FinishedFile(k, v.err)
					}

					assert.Empty(t, cp.pending)

					entries := bd.Details().Entries

					if cachedTest.expectToMergeEntries {
						assert.Equal(
							t,
							test.expectedNumItems,
							cp.toMerge.ItemsToMerge(),
							"merge entries")

						return
					}

					assert.Len(t, entries, test.expectedNumEntries)

					for _, entry := range entries {
						foundItems[entry.ItemRef] = true
					}

					if test.expectedNumEntries > 0 {
						for item, found := range foundItems {
							assert.Truef(t, found, "details missing item: %s", item)
						}
					}
				})
			}
		})
	}
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileCachedNoPrevPathErrors() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	bd := &details.Builder{}
	cachedItems := map[string]testInfo{
		suite.targetFileName: {
			info:       &itemDetails{repoPath: suite.targetFilePath},
			err:        nil,
			totalBytes: 100,
		},
	}
	cp := corsoProgress{
		ctx:            ctx,
		UploadProgress: &snapshotfs.NullUploadProgress{},
		deets:          bd,
		pending:        map[string]*itemDetails{},
		errs:           fault.New(true),
	}

	for k, v := range cachedItems {
		cp.put(k, v.info)
	}

	require.Len(t, cp.pending, len(cachedItems))

	for k, v := range cachedItems {
		cp.CachedFile(k, v.totalBytes)
		cp.FinishedFile(k, v.err)
	}

	assert.Empty(t, cp.pending)
	assert.Empty(t, bd.Details().Entries)
	assert.Error(t, cp.errs.Failure(), clues.ToCore(cp.errs.Failure()))
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileBaseItemDoesntBuildHierarchy() {
	type expectedRef struct {
		oldRef *path.Builder
		newRef path.Path
	}

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	prevPath := makePath(
		suite.T(),
		[]string{testTenant, service, testUser, category, testInboxDir, testFileName2},
		true)

	// Location is sourced from collections now so we don't need to check it here.
	expectedToMerge := []expectedRef{
		{
			oldRef: prevPath.ToBuilder(),
			newRef: suite.targetFilePath,
		},
	}

	// Setup stuff.
	db := &details.Builder{}
	cp := corsoProgress{
		ctx:            ctx,
		UploadProgress: &snapshotfs.NullUploadProgress{},
		deets:          db,
		pending:        map[string]*itemDetails{},
		toMerge:        newMergeDetails(),
		errs:           fault.New(true),
	}

	deets := &itemDetails{
		repoPath:     suite.targetFilePath,
		prevPath:     prevPath,
		locationPath: suite.targetFileLoc,
	}

	cp.put(suite.targetFileName, deets)
	require.Len(t, cp.pending, 1)

	cp.FinishedFile(suite.targetFileName, nil)
	assert.Empty(t, cp.deets)

	for _, expected := range expectedToMerge {
		gotRef, _, _ := cp.toMerge.GetNewPathRefs(
			expected.oldRef,
			time.Now(),
			nil)
		if !assert.NotNil(t, gotRef) {
			continue
		}

		assert.Equal(t, expected.newRef.String(), gotRef.String())
	}
}

func (suite *CorsoProgressUnitSuite) TestFinishedHashingFile() {
	for _, test := range finishedFileTable {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			bd := &details.Builder{}
			cp := corsoProgress{
				ctx:            ctx,
				UploadProgress: &snapshotfs.NullUploadProgress{},
				deets:          bd,
				pending:        map[string]*itemDetails{},
				errs:           fault.New(true),
			}

			ci := test.cachedItems(suite.targetFileName, suite.targetFilePath)

			for k, v := range ci {
				cp.FinishedHashingFile(k, v.totalBytes)
			}

			assert.Empty(t, cp.pending)
			assert.Equal(t, test.expectedBytes, cp.totalBytes)
		})
	}
}

type HierarchyBuilderUnitSuite struct {
	tester.Suite
	testStoragePath  path.Path
	testLocationPath path.Path
}

func (suite *HierarchyBuilderUnitSuite) SetupSuite() {
	suite.testStoragePath = makePath(
		suite.T(),
		[]string{testTenant, service, testUser, category, testInboxID},
		false)
	suite.testLocationPath = makePath(
		suite.T(),
		[]string{testTenant, service, testUser, category, testInboxDir},
		false)
}

func TestHierarchyBuilderUnitSuite(t *testing.T) {
	suite.Run(t, &HierarchyBuilderUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree() {
	t := suite.T()
	tester.LogTimeOfTest(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		tenant       = "a-tenant"
		user1        = testUser
		user1Encoded = encodeAsPath(user1)
		user2        = "user2"
		user2Encoded = encodeAsPath(user2)
		storeP2      = makePath(t, []string{tenant, service, user2, category, testInboxID}, false)
		locP2        = makePath(t, []string{tenant, service, user2, category, testInboxDir}, false)
	)

	// Encode user names here so we don't have to decode things later.
	expectedFileCount := map[string]int{
		user1Encoded: 5,
		user2Encoded: 42,
	}

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		toMerge: newMergeDetails(),
		errs:    fault.New(true),
	}

	collections := []data.BackupCollection{
		exchMock.NewCollection(
			suite.testStoragePath,
			suite.testLocationPath,
			expectedFileCount[user1Encoded]),
		exchMock.NewCollection(
			storeP2,
			locP2,
			expectedFileCount[user2Encoded]),
	}

	// Returned directory structure should look like:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - emails
	//         - Inbox
	//           - 5 separate files
	//     - user2
	//       - emails
	//         - Inbox
	//           - 42 separate files
	dirTree, err := inflateDirTree(ctx, nil, nil, collections, pmMock.NewPrefixMap(nil), progress)
	require.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(t, err, clues.ToCore(err))

	expectDirs(t, entries, encodeElements(service), true)

	entries = getDirEntriesForEntry(t, ctx, entries[0])
	expectDirs(t, entries, encodeElements(user1, user2), true)

	for _, entry := range entries {
		userName := entry.Name()

		entries = getDirEntriesForEntry(t, ctx, entry)
		expectDirs(t, entries, encodeElements(category), true)

		entries = getDirEntriesForEntry(t, ctx, entries[0])
		expectDirs(t, entries, encodeElements(testInboxID), true)

		entries = getDirEntriesForEntry(t, ctx, entries[0])
		assert.Len(t, entries, expectedFileCount[userName])
	}

	totalFileCount := 0
	for _, c := range expectedFileCount {
		totalFileCount += c
	}

	assert.Len(t, progress.pending, totalFileCount)
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree_MixedDirectory() {
	var (
		subfldID  = "subfolder_ID"
		subfldDir = "subfolder"
		storeP2   = makePath(suite.T(), append(suite.testStoragePath.Elements(), subfldID), false)
		locP2     = makePath(suite.T(), append(suite.testLocationPath.Elements(), subfldDir), false)
	)

	// Test multiple orders of items because right now order can matter. Both
	// orders result in a directory structure like:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - emails
	//         - Inbox_ID
	//           - subfolder_ID
	//             - 5 separate files
	//           - 42 separate files
	table := []struct {
		name   string
		layout []data.BackupCollection
	}{
		{
			name: "SubdirFirst",
			layout: []data.BackupCollection{
				exchMock.NewCollection(
					storeP2,
					locP2,
					5),
				exchMock.NewCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					42),
			},
		},
		{
			name: "SubdirLast",
			layout: []data.BackupCollection{
				exchMock.NewCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					42),
				exchMock.NewCollection(
					storeP2,
					locP2,
					5),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				pending: map[string]*itemDetails{},
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}

			dirTree, err := inflateDirTree(ctx, nil, nil, test.layout, pmMock.NewPrefixMap(nil), progress)
			require.NoError(t, err, clues.ToCore(err))

			assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

			entries, err := fs.GetAllEntries(ctx, dirTree)
			require.NoError(t, err, clues.ToCore(err))

			expectDirs(t, entries, encodeElements(service), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(testUser), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(category), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			expectDirs(t, entries, encodeElements(testInboxID), true)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			// 42 files and 1 subdirectory.
			assert.Len(t, entries, 43)

			// One of these entries should be a subdirectory with items in it.
			subDirs := []fs.Directory(nil)
			for _, e := range entries {
				d, ok := e.(fs.Directory)
				if !ok {
					continue
				}

				subDirs = append(subDirs, d)
				assert.Equal(t, encodeAsPath(subfldID), d.Name())
			}

			require.Len(t, subDirs, 1)

			entries = getDirEntriesForEntry(t, ctx, entries[0])
			assert.Len(t, entries, 5)
		})
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree_Fails() {
	storeP2 := makePath(
		suite.T(),
		[]string{"tenant2", service, "user2", category, testInboxID},
		false)
	locP2 := makePath(
		suite.T(),
		[]string{"tenant2", service, "user2", category, testInboxDir},
		false)

	table := []struct {
		name   string
		layout []data.BackupCollection
	}{
		{
			"MultipleRoots",
			// Directory structure would look like:
			// - tenant1
			//   - exchange
			//     - user1
			//       - emails
			//         - Inbox
			//           - 5 separate files
			// - tenant2
			//   - exchange
			//     - user2
			//       - emails
			//         - Inbox
			//           - 42 separate files
			[]data.BackupCollection{
				exchMock.NewCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					5),
				exchMock.NewCollection(
					storeP2,
					locP2,
					42),
			},
		},
		{
			"NoCollectionPath",
			[]data.BackupCollection{
				exchMock.NewCollection(
					nil,
					nil,
					5),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}

			_, err := inflateDirTree(ctx, nil, nil, test.layout, pmMock.NewPrefixMap(nil), progress)
			assert.Error(t, err, clues.ToCore(err))
		})
	}
}

type mockSnapshotWalker struct {
	snapshotRoot fs.Entry
}

func (msw *mockSnapshotWalker) SnapshotRoot(*snapshot.Manifest) (fs.Entry, error) {
	return msw.snapshotRoot, nil
}

func makeManifestEntry(
	id, tenant, resourceOwner string,
	service path.ServiceType,
	categories ...path.CategoryType,
) ManifestEntry {
	var reasons []identity.Reasoner

	for _, c := range categories {
		reasons = append(reasons, NewReason(tenant, resourceOwner, service, c))
	}

	return ManifestEntry{
		Manifest: &snapshot.Manifest{
			ID: manifest.ID(id),
		},
		Reasons: reasons,
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeErrors() {
	var (
		storePath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxID},
			false)
		storePath2 = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveID},
			false)
		locPath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveDir},
			false)
	)

	table := []struct {
		name   string
		states []data.CollectionState
	}{
		{
			name: "DeletedAndNotMoved",
			states: []data.CollectionState{
				data.NotMovedState,
				data.DeletedState,
			},
		},
		{
			name: "NotMovedAndDeleted",
			states: []data.CollectionState{
				data.DeletedState,
				data.NotMovedState,
			},
		},
		{
			name: "DeletedAndMoved",
			states: []data.CollectionState{
				data.DeletedState,
				data.MovedState,
			},
		},
		{
			name: "NotMovedAndMoved",
			states: []data.CollectionState{
				data.NotMovedState,
				data.MovedState,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tester.LogTimeOfTest(t)

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				pending: map[string]*itemDetails{},
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}

			cols := []data.BackupCollection{}
			for _, s := range test.states {
				prevPath := storePath
				nowPath := storePath

				switch s {
				case data.DeletedState:
					nowPath = nil
				case data.MovedState:
					nowPath = storePath2
				}

				mc := exchMock.NewCollection(nowPath, locPath, 0)
				mc.ColState = s
				mc.PrevPath = prevPath

				cols = append(cols, mc)
			}

			_, err := inflateDirTree(ctx, nil, nil, cols, pmMock.NewPrefixMap(nil), progress)
			require.Error(t, err, clues.ToCore(err))
		})
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSingleSubtree() {
	var (
		storePath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxID},
			false)
		storePath2 = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveID},
			false)
		locPath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxDir},
			false)
		locPath2 = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveDir},
			false)

		prefixDirs = []string{
			testTenant,
			service,
			testUser,
			category,
		}
	)

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	getBaseSnapshot := func() fs.Entry {
		return baseWithChildren(
			prefixDirs,
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(testInboxID)[0],
					[]fs.Entry{
						virtualfs.StreamingFileWithModTimeFromReader(
							encodeElements(testFileName)[0],
							time.Time{},
							io.NopCloser(bytes.NewReader(testFileData))),
					}),
			})
	}

	table := []struct {
		name             string
		inputCollections func() []data.BackupCollection
		expected         *expectedNode
	}{
		{
			name: "SkipsDeletedItems",
			inputCollections: func() []data.BackupCollection {
				mc := exchMock.NewCollection(storePath, locPath, 1)
				mc.Names[0] = testFileName
				mc.DeletedItems[0] = true

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name:     testInboxID,
						children: []*expectedNode{},
					},
				}),
		},
		{
			name: "AddsNewItems",
			inputCollections: func() []data.BackupCollection {
				mc := exchMock.NewCollection(storePath, locPath, 1)
				mc.PrevPath = storePath
				mc.Names[0] = testFileName2
				mc.Data[0] = testFileData2
				mc.ColState = data.NotMovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(testFileName, nil),
							newExpectedFile(testFileName2, testFileData2),
						},
					},
				}),
		},
		{
			name: "SkipsUpdatedItems",
			inputCollections: func() []data.BackupCollection {
				mc := exchMock.NewCollection(storePath, locPath, 1)
				mc.PrevPath = storePath
				mc.Names[0] = testFileName
				mc.Data[0] = testFileData2
				mc.ColState = data.NotMovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(testFileName, testFileData2),
						},
					},
				}),
		},
		{
			name: "DeleteAndNew",
			inputCollections: func() []data.BackupCollection {
				mc1 := exchMock.NewCollection(storePath, locPath, 0)
				mc1.ColState = data.DeletedState
				mc1.PrevPath = storePath

				mc2 := exchMock.NewCollection(storePath, locPath, 1)
				mc2.ColState = data.NewState
				mc2.Names[0] = testFileName2
				mc2.Data[0] = testFileData2

				return []data.BackupCollection{mc1, mc2}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(testFileName2, testFileData2),
						},
					},
				}),
		},
		{
			name: "MovedAndNew",
			inputCollections: func() []data.BackupCollection {
				mc1 := exchMock.NewCollection(storePath2, locPath2, 0)
				mc1.ColState = data.MovedState
				mc1.PrevPath = storePath

				mc2 := exchMock.NewCollection(storePath, locPath, 1)
				mc2.ColState = data.NewState
				mc2.Names[0] = testFileName2
				mc2.Data[0] = testFileData2

				return []data.BackupCollection{mc1, mc2}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(testFileName2, testFileData2),
						},
					},
					{
						name: testArchiveID,
						children: []*expectedNode{
							newExpectedFile(testFileName, nil),
						},
					},
				}),
		},
		{
			name: "NewDoesntMerge",
			inputCollections: func() []data.BackupCollection {
				mc1 := exchMock.NewCollection(storePath, locPath, 1)
				mc1.ColState = data.NewState
				mc1.Names[0] = testFileName2
				mc1.Data[0] = testFileData2

				return []data.BackupCollection{mc1}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(testFileName2, testFileData2),
						},
					},
				}),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tester.LogTimeOfTest(t)

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				pending: map[string]*itemDetails{},
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}
			msw := &mockSnapshotWalker{
				snapshotRoot: getBaseSnapshot(),
			}

			dirTree, err := inflateDirTree(
				ctx,
				msw,
				[]ManifestEntry{
					makeManifestEntry("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
				},
				test.inputCollections(),
				pmMock.NewPrefixMap(nil),
				progress)
			require.NoError(t, err, clues.ToCore(err))

			expectTree(t, ctx, test.expected, dirTree)
		})
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeMultipleSubdirectories() {
	const (
		personalID  = "personal_ID"
		workID      = "work_ID"
		personalDir = "personal"
		workDir     = "work"
	)

	var (
		inboxStorePath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxID},
			false)
		inboxLocPath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxDir},
			false)
		inboxFileName1 = testFileName
		inboxFileData1 = testFileData4
		inboxFileName2 = testFileName5
		inboxFileData2 = testFileData5

		personalStorePath = makePath(
			suite.T(),
			append(inboxStorePath.Elements(), personalID),
			false)
		personalLocPath = makePath(
			suite.T(),
			append(inboxLocPath.Elements(), personalDir),
			false)
		personalFileName1 = inboxFileName1
		personalFileName2 = testFileName2

		workStorePath = makePath(
			suite.T(),
			append(inboxStorePath.Elements(), workID),
			false)
		workLocPath = makePath(
			suite.T(),
			append(inboxLocPath.Elements(), workDir),
			false)
		workFileName1 = testFileName3
		workFileName2 = testFileName4
		workFileData2 = testFileData

		prefixDirs = []string{
			testTenant,
			service,
			testUser,
			category,
		}
	)

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox_ID
	//           - file1
	//           - personal_ID
	//             - file1
	//             - file2
	//           - work_ID
	//             - file3
	getBaseSnapshot := func() fs.Entry {
		return baseWithChildren(
			prefixDirs,
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(testInboxID)[0],
					[]fs.Entry{
						virtualfs.StreamingFileWithModTimeFromReader(
							encodeElements(inboxFileName1)[0],
							time.Time{},
							io.NopCloser(bytes.NewReader(inboxFileData1))),
						virtualfs.NewStaticDirectory(
							encodeElements(personalID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(personalFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData))),
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(personalFileName2)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData2))),
							}),
						virtualfs.NewStaticDirectory(
							encodeElements(workID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(workFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData3))),
							}),
					}),
			})
	}

	table := []struct {
		name             string
		inputCollections func(t *testing.T) []data.BackupCollection
		inputExcludes    *pmMock.PrefixMap
		expected         *expectedNode
	}{
		{
			name: "GlobalExcludeSet",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				return nil
			},
			inputExcludes: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				"": {
					inboxFileName1: {},
				},
			}),
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName2, nil),
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
				}),
		},
		{
			name: "MovesSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, testInboxID + "2"},
					false)
				newLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, testInboxDir + "2"},
					false)

				mc := exchMock.NewCollection(newStorePath, newLocPath, 0)
				mc.PrevPath = inboxStorePath
				mc.ColState = data.MovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID + "2",
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
				}),
		},
		{
			name: "MovesChildAfterAncestorMove",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newInboxStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, testInboxID + "2"},
					false)
				newWorkStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workID},
					false)
				newInboxLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, testInboxDir + "2"},
					false)
				newWorkLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workID},
					false)

				inbox := exchMock.NewCollection(newInboxStorePath, newInboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.MovedState

				work := exchMock.NewCollection(newWorkStorePath, newWorkLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID + "2",
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
						},
					},
					{
						name: workID,
						children: []*expectedNode{
							newExpectedFile(workFileName1, nil),
						},
					},
				}),
		},
		{
			name: "MovesChildAfterAncestorDelete",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newWorkStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workID},
					false)
				newWorkLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workDir},
					false)

				inbox := exchMock.NewCollection(inboxStorePath, inboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.DeletedState

				work := exchMock.NewCollection(newWorkStorePath, newWorkLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: workID,
						children: []*expectedNode{
							newExpectedFile(workFileName1, nil),
						},
					},
				}),
		},
		{
			name: "ReplaceDeletedDirectory",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				personal := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.DeletedState

				work := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{personal, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
				}),
		},
		{
			name: "ReplaceDeletedDirectoryWithNew",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				personal := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.DeletedState

				newCol := exchMock.NewCollection(personalStorePath, personalLocPath, 1)
				newCol.ColState = data.NewState
				newCol.Names[0] = workFileName2
				newCol.Data[0] = workFileData2

				return []data.BackupCollection{personal, newCol}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(workFileName2, workFileData2),
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
				}),
		},
		{
			name: "ReplaceDeletedSubtreeWithNew",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				oldInbox := exchMock.NewCollection(inboxStorePath, inboxLocPath, 0)
				oldInbox.PrevPath = inboxStorePath
				oldInbox.ColState = data.DeletedState

				newCol := exchMock.NewCollection(inboxStorePath, inboxLocPath, 1)
				newCol.ColState = data.NewState
				newCol.Names[0] = workFileName2
				newCol.Data[0] = workFileData2

				return []data.BackupCollection{oldInbox, newCol}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(workFileName2, workFileData2),
						},
					},
				}),
		},
		{
			name: "ReplaceMovedDirectory",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newPersonalStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalID},
					false)
				newPersonalLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalDir},
					false)

				personal := exchMock.NewCollection(newPersonalStorePath, newPersonalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.MovedState

				work := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{personal, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
					{
						name: personalID,
						children: []*expectedNode{
							newExpectedFile(personalFileName1, nil),
							newExpectedFile(personalFileName2, nil),
						},
					},
				}),
		},
		{
			name: "MoveDirectoryAndMergeItems",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newPersonalStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workID},
					false)
				newPersonalLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, workDir},
					false)

				personal := exchMock.NewCollection(newPersonalStorePath, newPersonalLocPath, 2)
				personal.PrevPath = personalStorePath
				personal.ColState = data.MovedState
				personal.Names[0] = personalFileName2
				personal.Data[0] = testFileData5
				personal.Names[1] = testFileName4
				personal.Data[1] = testFileData4

				return []data.BackupCollection{personal}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, nil),
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
					{
						name: workID,
						children: []*expectedNode{
							newExpectedFile(personalFileName1, nil),
							newExpectedFile(personalFileName2, nil),
							newExpectedFile(testFileName4, testFileData4),
						},
					},
				}),
		},
		{
			name: "MoveParentDeleteFileNoMergeSubtreeMerge",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newInboxStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalID},
					false)
				newInboxLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalDir},
					false)

				// This path is implicitly updated because we update the inbox path. If
				// we didn't update it here then it would end up at the old location
				// still.
				newWorkStorePath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalID, workID},
					false)
				newWorkLocPath := makePath(
					t,
					[]string{testTenant, service, testUser, category, personalDir, workDir},
					false)

				inbox := exchMock.NewCollection(newInboxStorePath, newInboxLocPath, 1)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.MovedState
				inbox.DoNotMerge = true
				// First file in inbox is implicitly deleted as we're not merging items
				// and it's not listed.
				inbox.Names[0] = inboxFileName2
				inbox.Data[0] = inboxFileData2

				work := exchMock.NewCollection(newWorkStorePath, newWorkLocPath, 1)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState
				work.Names[0] = testFileName6
				work.Data[0] = testFileData6

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: personalID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName2, inboxFileData2),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
									newExpectedFile(testFileName6, testFileData6),
								},
							},
						},
					},
				}),
		},
		{
			name: "NoMoveParentDeleteFileNoMergeSubtreeMerge",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				inbox := exchMock.NewCollection(inboxStorePath, inboxLocPath, 1)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.NotMovedState
				inbox.DoNotMerge = true
				// First file in inbox is implicitly deleted as we're not merging items
				// and it's not listed.
				inbox.Names[0] = inboxFileName2
				inbox.Data[0] = inboxFileData2

				work := exchMock.NewCollection(workStorePath, workLocPath, 1)
				work.PrevPath = workStorePath
				work.ColState = data.NotMovedState
				work.Names[0] = testFileName6
				work.Data[0] = testFileData6

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName2, inboxFileData2),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
									newExpectedFile(testFileName6, testFileData6),
								},
							},
						},
					},
				}),
		},
		{
			// This could happen if a subfolder is moved out of the parent, the parent
			// is deleted, a new folder at the same location as the parent is created,
			// and then the subfolder is moved back to the same location.
			name: "Delete Parent But Child Marked Not Moved Explicit New Parent",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				inbox := exchMock.NewCollection(nil, inboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.DeletedState

				inbox2 := exchMock.NewCollection(inboxStorePath, inboxLocPath, 1)
				inbox2.PrevPath = nil
				inbox2.ColState = data.NewState
				inbox2.Names[0] = workFileName1

				personal := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.NotMovedState

				return []data.BackupCollection{inbox, inbox2, personal}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(workFileName1, nil),
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
						},
					},
				}),
		},
		{
			// This could happen if a subfolder is moved out of the parent, the parent
			// is deleted, a new folder at the same location as the parent is created,
			// and then the subfolder is moved back to the same location.
			name: "Delete Parent But Child Marked Not Moved Implicit New Parent",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				inbox := exchMock.NewCollection(nil, inboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.DeletedState

				// New folder not explicitly listed as it may not have had new items.
				personal := exchMock.NewCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.NotMovedState

				return []data.BackupCollection{inbox, personal}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(personalFileName1, nil),
									newExpectedFile(personalFileName2, nil),
								},
							},
						},
					},
				}),
		},
		{
			// This could happen if a subfolder is moved out of the parent, the parent
			// is deleted, a new folder at the same location as the parent is created,
			// and then the subfolder is moved back to the same location.
			name: "Delete Parent But Child Marked Not Moved Implicit New Parent Child Do Not Merge",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				inbox := exchMock.NewCollection(nil, inboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.DeletedState

				// New folder not explicitly listed as it may not have had new items.
				personal := exchMock.NewCollection(personalStorePath, personalLocPath, 1)
				personal.PrevPath = personalStorePath
				personal.ColState = data.NotMovedState
				personal.DoNotMerge = true
				personal.Names[0] = workFileName1

				return []data.BackupCollection{inbox, personal}
			},
			expected: expectedTreeWithChildren(
				prefixDirs,
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name: personalID,
								children: []*expectedNode{
									newExpectedFile(workFileName1, nil),
								},
							},
						},
					},
				}),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tester.LogTimeOfTest(t)

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				pending: map[string]*itemDetails{},
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}
			msw := &mockSnapshotWalker{
				snapshotRoot: getBaseSnapshot(),
			}

			ie := pmMock.NewPrefixMap(nil)
			if test.inputExcludes != nil {
				ie = test.inputExcludes
			}

			dirTree, err := inflateDirTree(
				ctx,
				msw,
				[]ManifestEntry{
					makeManifestEntry("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
				},
				test.inputCollections(t),
				ie,
				progress)
			require.NoError(t, err, clues.ToCore(err))

			expectTree(t, ctx, test.expected, dirTree)
		})
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSkipsDeletedSubtree() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		personalDir = "personal"
		workDir     = "work"

		prefixDirs = []string{
			testTenant,
			service,
			testUser,
			category,
		}
	)

	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox
	//           - personal
	//             - file1
	//           - work
	//             - file2
	//         - Archive
	//           - personal
	//             - file3
	//           - work
	//             - file4
	getBaseSnapshot := func() fs.Entry {
		return baseWithChildren(
			prefixDirs,
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(testInboxID)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(personalDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData))),
							}),
						virtualfs.NewStaticDirectory(
							encodeElements(workDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName2)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData2))),
							}),
					}),
				virtualfs.NewStaticDirectory(
					encodeElements(testArchiveID)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(personalDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName3)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData3))),
							}),
						virtualfs.NewStaticDirectory(
							encodeElements(workDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName4)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData4))),
							}),
					}),
			})
	}

	expected := expectedTreeWithChildren(
		prefixDirs,
		[]*expectedNode{
			{
				name: testArchiveID,
				children: []*expectedNode{
					{
						name: personalDir,
						children: []*expectedNode{
							newExpectedFile(testFileName3, nil),
						},
					},
					{
						name: workDir,
						children: []*expectedNode{
							newExpectedFile(testFileName4, nil),
						},
					},
				},
			},
		})

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		toMerge: newMergeDetails(),
		errs:    fault.New(true),
	}
	mc := exchMock.NewCollection(suite.testStoragePath, suite.testStoragePath, 1)
	mc.PrevPath = mc.FullPath()
	mc.ColState = data.DeletedState
	msw := &mockSnapshotWalker{
		snapshotRoot: getBaseSnapshot(),
	}

	collections := []data.BackupCollection{mc}

	// Returned directory structure should look like:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - emails
	//         - Archive
	//           - personal
	//             - file3
	//           - work
	//             - file4
	dirTree, err := inflateDirTree(
		ctx,
		msw,
		[]ManifestEntry{
			makeManifestEntry("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		pmMock.NewPrefixMap(nil),
		progress)
	require.NoError(t, err, clues.ToCore(err))

	expectTree(t, ctx, expected, dirTree)
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree_HandleEmptyBase() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		archiveStorePath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveID},
			false)
		archiveLocPath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testArchiveDir},
			false)
	)

	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchangeMetadata
	//     - user1
	//       - email
	//         - file1
	getBaseSnapshot := func() fs.Entry {
		return baseWithChildren(
			[]string{
				testTenant,
				path.ExchangeMetadataService.String(),
				testUser,
				category,
			},
			[]fs.Entry{
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(testFileName)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(testFileData))),
			})
	}

	// Metadata subtree doesn't appear because we don't select it as one of the
	// subpaths and we're not passing in a metadata collection.
	expected := expectedTreeWithChildren(
		[]string{
			testTenant,
			path.ExchangeService.String(),
			testUser,
			category,
		},
		[]*expectedNode{
			{
				name: testArchiveID,
				children: []*expectedNode{
					newExpectedFile(testFileName2, nil),
				},
			},
		})

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		toMerge: newMergeDetails(),
		errs:    fault.New(true),
	}
	mc := exchMock.NewCollection(archiveStorePath, archiveLocPath, 1)
	mc.ColState = data.NewState
	mc.Names[0] = testFileName2
	mc.Data[0] = testFileData2

	msw := &mockSnapshotWalker{
		snapshotRoot: getBaseSnapshot(),
	}

	collections := []data.BackupCollection{mc}

	// Returned directory structure should look like:
	// - a-tenant
	//   - exchangeMetadata
	//     - user1
	//       - emails
	//         - file1
	//   - exchange
	//     - user1
	//       - emails
	//         - Archive
	//           - file2
	dirTree, err := inflateDirTree(
		ctx,
		msw,
		[]ManifestEntry{
			makeManifestEntry("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		pmMock.NewPrefixMap(nil),
		progress)
	require.NoError(t, err, clues.ToCore(err))

	expectTree(t, ctx, expected, dirTree)
}

type mockMultiSnapshotWalker struct {
	snaps map[string]fs.Entry
}

func (msw *mockMultiSnapshotWalker) SnapshotRoot(man *snapshot.Manifest) (fs.Entry, error) {
	if snap := msw.snaps[string(man.ID)]; snap != nil {
		return snap, nil
	}

	return nil, clues.New("snapshot not found")
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSelectsCorrectSubtrees() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	const contactsDir = "contacts"

	var (
		inboxPath = makePath(
			suite.T(),
			[]string{testTenant, service, testUser, category, testInboxID},
			false)

		inboxFileName1 = testFileName
		inboxFileName2 = testFileName2

		inboxFileData1   = testFileData
		inboxFileData1v2 = testFileData5
		inboxFileData2   = testFileData2

		contactsFileName1 = testFileName3
		contactsFileData1 = testFileData3

		eventsFileName1 = testFileName5
		eventsFileData1 = testFileData

		prefixDirs = []string{
			testTenant,
			service,
			testUser,
		}
	)

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox
	//           - file1
	//       - contacts
	//         - contacts
	//           - file2
	getBaseSnapshot1 := func() fs.Entry {
		return baseWithChildren(
			prefixDirs,
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(category)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(testInboxID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(inboxFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(inboxFileData1))),
							}),
					}),
				virtualfs.NewStaticDirectory(
					encodeElements(path.ContactsCategory.String())[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(contactsDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(contactsFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(contactsFileData1))),
							}),
					}),
			})
	}

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox
	//           - file1 <- has different data version
	//       - events
	//         - events
	//           - file3
	getBaseSnapshot2 := func() fs.Entry {
		return baseWithChildren(
			prefixDirs,
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(category)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(testInboxID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(inboxFileName1)[0],
									time.Time{},
									// Wrap with a backup reader so it gets the version injected.
									io.NopCloser(bytes.NewReader(inboxFileData1v2))),
							}),
					}),
				virtualfs.NewStaticDirectory(
					encodeElements(path.EventsCategory.String())[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements("events")[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(eventsFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(eventsFileData1))),
							}),
					}),
			})
	}

	// Check the following:
	//   * contacts pulled from base1 unchanged even if no collections reference
	//     it
	//   * email pulled from base2
	//   * new email added
	//   * events not pulled from base2 as it's not listed as a Reason
	//
	// Expected output:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox
	//           - file1 <- version of data from second base
	//           - file2
	//       - contacts
	//         - contacts
	//           - file2
	expected := expectedTreeWithChildren(
		prefixDirs,
		[]*expectedNode{
			{
				name: category,
				children: []*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, inboxFileData1v2),
							newExpectedFile(inboxFileName2, inboxFileData2),
						},
					},
				},
			},
			{
				name: path.ContactsCategory.String(),
				children: []*expectedNode{
					{
						name: contactsDir,
						children: []*expectedNode{
							newExpectedFile(contactsFileName1, nil),
						},
					},
				},
			},
		})

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		toMerge: newMergeDetails(),
		errs:    fault.New(true),
	}

	mc := exchMock.NewCollection(inboxPath, inboxPath, 1)
	mc.PrevPath = mc.FullPath()
	mc.ColState = data.NotMovedState
	mc.Names[0] = inboxFileName2
	mc.Data[0] = inboxFileData2

	msw := &mockMultiSnapshotWalker{
		snaps: map[string]fs.Entry{
			"id1": getBaseSnapshot1(),
			"id2": getBaseSnapshot2(),
		},
	}

	collections := []data.BackupCollection{mc}

	dirTree, err := inflateDirTree(
		ctx,
		msw,
		[]ManifestEntry{
			makeManifestEntry("id1", testTenant, testUser, path.ExchangeService, path.ContactsCategory),
			makeManifestEntry("id2", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		pmMock.NewPrefixMap(nil),
		progress)
	require.NoError(t, err, clues.ToCore(err))

	expectTree(t, ctx, expected, dirTree)
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSelectsMigrateSubtrees() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	const (
		contactsDir  = "contacts"
		migratedUser = "user_migrate"
	)

	oldPrefixPathEmail, err := path.BuildPrefix(testTenant, testUser, path.ExchangeService, path.EmailCategory)
	require.NoError(t, err, clues.ToCore(err))

	newPrefixPathEmail, err := path.BuildPrefix(testTenant, migratedUser, path.ExchangeService, path.EmailCategory)
	require.NoError(t, err, clues.ToCore(err))

	oldPrefixPathCont, err := path.BuildPrefix(testTenant, testUser, path.ExchangeService, path.ContactsCategory)
	require.NoError(t, err, clues.ToCore(err))

	newPrefixPathCont, err := path.BuildPrefix(testTenant, migratedUser, path.ExchangeService, path.ContactsCategory)
	require.NoError(t, err, clues.ToCore(err))

	var (
		inboxFileName1 = testFileName

		inboxFileData1 = testFileData
		// inboxFileData1v2 = testFileData5

		contactsFileName1 = testFileName3
		contactsFileData1 = testFileData3
	)

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	// baseSnapshot with the following layout:
	// - a-tenant
	//   - exchange
	//     - user1
	//       - email
	//         - Inbox
	//           - file1
	//       - contacts
	//         - contacts
	//           - file2
	getBaseSnapshot1 := func() fs.Entry {
		return baseWithChildren(
			[]string{testTenant, service, testUser},
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(category)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(testInboxID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(inboxFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(inboxFileData1))),
							}),
					}),
				virtualfs.NewStaticDirectory(
					encodeElements(path.ContactsCategory.String())[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(contactsDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(contactsFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(contactsFileData1))),
							}),
					}),
			})
	}

	// Check the following:
	//   * contacts pulled from base1 unchanged even if no collections reference
	//     it
	//   * email pulled from base2
	//
	// Expected output:
	// - a-tenant
	//   - exchange
	//     - user1new
	//       - email
	//         - Inbox
	//           - file1
	//       - contacts
	//         - contacts
	//           - file1
	expected := expectedTreeWithChildren(
		[]string{testTenant, service, migratedUser},
		[]*expectedNode{
			{
				name: category,
				children: []*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							newExpectedFile(inboxFileName1, inboxFileData1),
						},
					},
				},
			},
			{
				name: path.ContactsCategory.String(),
				children: []*expectedNode{
					{
						name: contactsDir,
						children: []*expectedNode{
							newExpectedFile(contactsFileName1, nil),
						},
					},
				},
			},
		})

	progress := &corsoProgress{
		ctx:     ctx,
		pending: map[string]*itemDetails{},
		toMerge: newMergeDetails(),
		errs:    fault.New(true),
	}

	mce := exchMock.NewCollection(newPrefixPathEmail, nil, 0)
	mce.PrevPath = oldPrefixPathEmail
	mce.ColState = data.MovedState

	mcc := exchMock.NewCollection(newPrefixPathCont, nil, 0)
	mcc.PrevPath = oldPrefixPathCont
	mcc.ColState = data.MovedState

	msw := &mockMultiSnapshotWalker{
		snaps: map[string]fs.Entry{"id1": getBaseSnapshot1()},
	}

	dirTree, err := inflateDirTree(
		ctx,
		msw,
		[]ManifestEntry{
			makeManifestEntry("id1", testTenant, testUser, path.ExchangeService, path.EmailCategory, path.ContactsCategory),
		},
		[]data.BackupCollection{mce, mcc},
		pmMock.NewPrefixMap(nil),
		progress)
	require.NoError(t, err, clues.ToCore(err))

	expectTree(t, ctx, expected, dirTree)
}

func newMockStaticDirectory(
	name string,
	entries []fs.Entry,
) (fs.Directory, *int) {
	res := &mockStaticDirectory{
		Directory: virtualfs.NewStaticDirectory(name, entries),
	}

	return res, &res.iterateCount
}

type mockStaticDirectory struct {
	fs.Directory
	iterateCount int
}

func (msd *mockStaticDirectory) IterateEntries(
	ctx context.Context,
	callback func(context.Context, fs.Entry) error,
) error {
	msd.iterateCount++
	return msd.Directory.IterateEntries(ctx, callback)
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree_SelectiveSubtreePruning() {
	var (
		tenant   = "tenant-id"
		service  = path.OneDriveService.String()
		user     = "user-id"
		category = path.FilesCategory.String()

		// Not using drive/drive-id/root folders for brevity.
		folderID1 = "folder1-id"
		folderID2 = "folder2-id"
		folderID3 = "folder3-id"
		folderID4 = "folder4-id"
		folderID5 = "folder5-id"

		folderName1 = "folder1-name"
		folderName2 = "folder2-name"
		folderName3 = "folder3-name"
		folderName5 = "folder5-name"

		fileName1 = "file1"
		fileName2 = "file2"
		fileName3 = "file3"
		fileName4 = "file4"
		fileName5 = "file5"
		fileName6 = "file6"
		fileName7 = "file7"
		fileName8 = "file8"

		fileData1 = []byte("1")
		fileData2 = []byte("2")
		fileData3 = []byte("3")
		fileData4 = []byte("4")
		fileData5 = []byte("5")
		fileData6 = []byte("6")
		fileData7 = []byte("7")
		fileData8 = []byte("8")
	)

	var (
		folderPath1 = makePath(
			suite.T(),
			[]string{tenant, service, user, category, folderID1},
			false)
		folderLocPath1 = makePath(
			suite.T(),
			[]string{tenant, service, user, category, folderName1},
			false)

		folderPath2 = makePath(
			suite.T(),
			append(folderPath1.Elements(), folderID2),
			false)
		folderLocPath2 = makePath(
			suite.T(),
			append(folderLocPath1.Elements(), folderName2),
			false)

		folderPath3 = makePath(
			suite.T(),
			append(folderPath2.Elements(), folderID3),
			false)

		folderPath5 = makePath(
			suite.T(),
			[]string{tenant, service, user, category, folderID5},
			false)
		folderLocPath5 = makePath(
			suite.T(),
			[]string{tenant, service, user, category, folderName5},
			false)

		prefixFolders = []string{
			tenant,
			service,
			user,
			category,
		}
	)

	folder5Unchanged := exchMock.NewCollection(folderPath5, folderLocPath5, 0)
	folder5Unchanged.PrevPath = folderPath5
	folder5Unchanged.ColState = data.NotMovedState

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once. Each directory below the prefix directory
	// is also wrapped in a mock so we can count the number of times
	// IterateEntries was called on it.
	// baseSnapshot with the following layout:
	// - tenant-id
	//   - onedrive
	//     - user-id
	//       - files
	//         - folder1-id
	//           - file1
	//           - file2
	//           - folder2-id
	//             - file3
	//             - file4
	//             - folder3-id
	//               - file5
	//               - file6
	//           - folder4-id
	//         - folder5-id
	//           - file7
	//           - file8
	getBaseSnapshot := func() (fs.Entry, map[string]*int) {
		counters := map[string]*int{}

		folder, count := newMockStaticDirectory(
			encodeElements(folderID3)[0],
			[]fs.Entry{
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName5)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData5))),
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName6)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData6))),
			})
		counters[folderID3] = count

		folder, count = newMockStaticDirectory(
			encodeElements(folderID2)[0],
			[]fs.Entry{
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName3)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData3))),
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName4)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData4))),
				folder,
			})
		counters[folderID2] = count

		folder4, count := newMockStaticDirectory(
			encodeElements(folderID4)[0],
			[]fs.Entry{})
		counters[folderID4] = count

		folder, count = newMockStaticDirectory(
			encodeElements(folderID1)[0],
			[]fs.Entry{
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName1)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData1))),
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName2)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData2))),
				folder,
				folder4,
			})
		counters[folderID1] = count

		folder5, count := newMockStaticDirectory(
			encodeElements(folderID5)[0],
			[]fs.Entry{
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName7)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData7))),
				virtualfs.StreamingFileWithModTimeFromReader(
					encodeElements(fileName8)[0],
					time.Time{},
					io.NopCloser(bytes.NewReader(fileData8))),
			})
		counters[folderID5] = count

		return baseWithChildren(
				prefixFolders,
				[]fs.Entry{
					folder,
					folder5,
				}),
			counters
	}

	table := []struct {
		name                  string
		inputCollections      func(t *testing.T) []data.BackupCollection
		inputExcludes         *pmMock.PrefixMap
		expected              *expectedNode
		expectedIterateCounts map[string]int
	}{
		{
			// Test that even if files are excluded in the subtree selective subtree
			// pruning skips traversing, the file is properly excluded during upload.
			//
			// It's safe to prune the subtree during merging because the directory
			// layout hasn't changed. We still require traversal of all directories
			// during data upload which allows us to exclude the file properly.
			name: "NoDirectoryChanges ExcludedFile PrunesSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				return []data.BackupCollection{folder5Unchanged}
			},
			inputExcludes: pmMock.NewPrefixMap(map[string]map[string]struct{}{
				"": {
					fileName3: {},
					fileName5: {},
					fileName6: {},
				},
			}),
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName4, fileData4),
									{
										name: folderID3,
									},
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 0,
				folderID2: 0,
				folderID3: 0,
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that if a subtree is deleted in its entirety selective subtree
			// pruning skips traversing it during hierarchy merging.
			name: "SubtreeDelete PrunesSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				mc := exchMock.NewCollection(nil, nil, 0)
				mc.PrevPath = folderPath1
				mc.ColState = data.DeletedState

				return []data.BackupCollection{folder5Unchanged, mc}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				// Deleted collections aren't added to the in-memory tree.
				folderID1: 0,
				folderID2: 0,
				folderID3: 0,
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that if a directory is moved but the subtree rooted at the moved
			// directory is unchanged selective subtree pruning skips traversing all
			// directories under the moved directory during hierarchy merging even if
			// a new directory is created at the path of one of the unchanged (pruned)
			// subdirectories of the moved directory.
			name: "ParentMoved NewFolderAtOldPath PrunesSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newPath := makePath(
					suite.T(),
					[]string{tenant, service, user, category, "foo-id"},
					false)
				newLoc := makePath(
					suite.T(),
					[]string{tenant, service, user, category, "foo"},
					false)

				mc := exchMock.NewCollection(newPath, newLoc, 0)
				mc.PrevPath = folderPath1
				mc.ColState = data.MovedState

				newMC := exchMock.NewCollection(folderPath2, folderLocPath2, 0)

				return []data.BackupCollection{folder5Unchanged, mc, newMC}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: "foo-id",
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
									{
										name: folderID3,
										children: []*expectedNode{
											newExpectedFile(fileName5, fileData5),
											newExpectedFile(fileName6, fileData6),
										},
									},
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID1,
						children: []*expectedNode{
							{
								name: folderID2,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				folderID2: 0,
				folderID3: 0,
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that if a directory and its subtree is deleted in its entirety
			// selective subtree pruning skips traversing the subtree during hierarchy
			// merging even if a new directory is created at the path of one of the
			// deleted (pruned) directories.
			name: "SubtreeDelete NewFolderAtOldPath PrunesSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				mc := exchMock.NewCollection(nil, nil, 0)
				mc.PrevPath = folderPath2
				mc.ColState = data.DeletedState

				newMC := exchMock.NewCollection(folderPath2, folderLocPath2, 0)

				return []data.BackupCollection{folder5Unchanged, mc, newMC}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				// Deleted collections aren't added to the in-memory tree.
				folderID2: 0,
				folderID3: 0,
				// Skipped because it's unchanged.
				folderID4: 0,
				folderID5: 1,
			},
		},
		// These tests check that subtree pruning isn't triggered for some parts of
		// the hierarchy.
		{
			// Test that creating a new directory in an otherwise unchanged subtree
			// doesn't trigger selective subtree merging for any subtree of the full
			// hierarchy that includes the new directory but does trigger selective
			// subtree pruning for unchanged subtrees without the new directory.
			name: "NewDirectory DoesntPruneSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newP, err := folderPath2.Append(false, "foo-id")
				require.NoError(t, err, clues.ToCore(err))

				newL, err := folderLocPath2.Append(false, "foo")
				require.NoError(t, err, clues.ToCore(err))

				newMC := exchMock.NewCollection(newP, newL, 0)

				return []data.BackupCollection{folder5Unchanged, newMC}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
									{
										name: "foo-id",
									},
									{
										name: folderID3,
										children: []*expectedNode{
											newExpectedFile(fileName5, fileData5),
											newExpectedFile(fileName6, fileData6),
										},
									},
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				folderID2: 1,
				// Folder 3 triggers pruning because it has nothing changed under it.
				folderID3: 0,
				// Folder 4 triggers pruning because it doesn't include foo and hasn't
				// changed.
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that moving a directory within a subtree doesn't trigger selective
			// subtree merging for any subtree of the full hierarchy that includes the
			// moved directory.
			name: "MoveWithinSubtree DoesntPruneSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newP, err := folderPath1.Append(false, folderID3)
				require.NoError(t, err, clues.ToCore(err))

				newL, err := folderLocPath1.Append(false, folderName3)
				require.NoError(t, err, clues.ToCore(err))

				mc := exchMock.NewCollection(newP, newL, 0)
				mc.PrevPath = folderPath3
				mc.ColState = data.MovedState

				return []data.BackupCollection{folder5Unchanged, mc}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
								},
							},
							{
								name: folderID3,
								children: []*expectedNode{
									newExpectedFile(fileName5, fileData5),
									newExpectedFile(fileName6, fileData6),
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				// Folder 2 can't be pruned because there's subtree changes under it
				// (folder3 move).
				folderID2: 1,
				// Folder 3 can't be pruned because it has a collection associated with
				// it.
				folderID3: 1,
				// Folder 4 is pruned because it didn't and still doesn't include
				// Folder 3 and it hasn't changed.
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that moving a directory out of a subtree doesn't trigger selective
			// subtree merging for any subtree of the full hierarchy that includes the
			// moved directory in the previous hierarchy.
			name: "MoveOutOfSubtree DoesntPruneSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newP := makePath(
					suite.T(),
					[]string{tenant, service, user, category, folderID3},
					false)
				newL := makePath(
					suite.T(),
					[]string{tenant, service, user, category, folderName3},
					false)

				mc := exchMock.NewCollection(newP, newL, 0)
				mc.PrevPath = folderPath3
				mc.ColState = data.MovedState

				return []data.BackupCollection{folder5Unchanged, mc}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID3,
						children: []*expectedNode{
							newExpectedFile(fileName5, fileData5),
							newExpectedFile(fileName6, fileData6),
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				// Folder 2 can't be pruned because there's subtree changes under it
				// (folder3 move).
				folderID2: 1,
				// Folder 3 can't be pruned because it has a collection associated with
				// it.
				folderID3: 1,
				// Folder 4 is pruned because it didn't and still doesn't include
				// Folder 3 and it hasn't changed.
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that deleting a directory in a subtree doesn't trigger selective
			// subtree merging for any subtree of the full hierarchy that includes the
			// deleted directory in the previous hierarchy.
			name: "DeleteInSubtree DoesntPruneSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				mc := exchMock.NewCollection(nil, nil, 0)
				mc.PrevPath = folderPath3
				mc.ColState = data.DeletedState

				return []data.BackupCollection{folder5Unchanged, mc}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
								},
							},
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				folderID1: 1,
				// Folder 2 can't be pruned because there's subtree changes under it
				// (folder3 delete).
				folderID2: 1,
				// Folder3 is pruned because there's no changes under it.
				folderID3: 0,
				// Folder 4 is pruned because it didn't and still doesn't include
				// Folder 3 and it hasn't changed.
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that moving an existing directory into a subtree doesn't trigger
			// selective subtree merging for any subtree of the full hierarchy that
			// includes the moved directory in the current hierarchy.
			name: "MoveIntoSubtree DoesntPruneSubtree",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newP, err := folderPath1.Append(false, folderID5)
				require.NoError(t, err, clues.ToCore(err))

				newL, err := folderLocPath1.Append(false, folderName5)
				require.NoError(t, err, clues.ToCore(err))

				mc := exchMock.NewCollection(newP, newL, 0)
				mc.PrevPath = folderPath5
				mc.ColState = data.MovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
									{
										name: folderID3,
										children: []*expectedNode{
											newExpectedFile(fileName5, fileData5),
											newExpectedFile(fileName6, fileData6),
										},
									},
								},
							},
							{
								name: folderID4,
							},
							{
								name: folderID5,
								children: []*expectedNode{
									newExpectedFile(fileName7, fileData7),
									newExpectedFile(fileName8, fileData8),
								},
							},
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				// Folder 1 can't be pruned because there's subtree changes under it
				// (folder4 move).
				folderID1: 1,
				// Folder 2 is pruned because nothing changes below it and it has no
				// collection associated with it.
				folderID2: 0,
				folderID3: 0,
				// Folder 4 is pruned because nothing changes below it and it has no
				// collection associated with it.
				folderID4: 0,
				folderID5: 1,
			},
		},
		// These test check a mix of not pruning and pruning.
		{
			// Test that deleting a directory with subdirectories that are re-added to
			// the tree still results in a proper hierarchy and prunes where possible
			// for the re-added subdirectories.
			name: "DeleteSubtree ReaddChildren PrunesWherePossible",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				mc := exchMock.NewCollection(nil, nil, 0)
				mc.PrevPath = folderPath1
				mc.ColState = data.DeletedState

				mc2 := exchMock.NewCollection(folderPath2, folderLocPath2, 0)
				mc2.PrevPath = folderPath2
				mc2.ColState = data.NotMovedState

				return []data.BackupCollection{folder5Unchanged, mc, mc2}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: folderID1,
						children: []*expectedNode{
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
									{
										name: folderID3,
										children: []*expectedNode{
											newExpectedFile(fileName5, fileData5),
											newExpectedFile(fileName6, fileData6),
										},
									},
								},
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				// Traversed because it's children still exist.
				folderID1: 1,
				// Folder 2 can't be pruned because it has a collection associated with
				// it.
				folderID2: 1,
				// Folder3 is pruned because there's no changes under it.
				folderID3: 0,
				// Not traversed because it's deleted.
				folderID4: 0,
				folderID5: 1,
			},
		},
		{
			// Test that moving a directory with subdirectories that are themselves
			// moved back to their original location still results in a proper
			// hierarchy and prunes where possible for the subdirectories.
			name: "MoveSubtree ReaddChildren PrunesWherePossible",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				newPath := makePath(
					suite.T(),
					[]string{tenant, service, user, category, "foo-id"},
					false)
				newLoc := makePath(
					suite.T(),
					[]string{tenant, service, user, category, "foo"},
					false)

				mc := exchMock.NewCollection(newPath, newLoc, 0)
				mc.PrevPath = folderPath1
				mc.ColState = data.MovedState

				mc2 := exchMock.NewCollection(folderPath2, folderLocPath2, 0)
				mc2.PrevPath = folderPath2
				mc2.ColState = data.NotMovedState

				return []data.BackupCollection{folder5Unchanged, mc, mc2}
			},
			expected: expectedTreeWithChildren(
				prefixFolders,
				[]*expectedNode{
					{
						name: "foo-id",
						children: []*expectedNode{
							newExpectedFile(fileName1, fileData1),
							newExpectedFile(fileName2, fileData2),
							{
								name: folderID4,
							},
						},
					},
					{
						name: folderID1,
						children: []*expectedNode{
							{
								name: folderID2,
								children: []*expectedNode{
									newExpectedFile(fileName3, fileData3),
									newExpectedFile(fileName4, fileData4),
									{
										name: folderID3,
										children: []*expectedNode{
											newExpectedFile(fileName5, fileData5),
											newExpectedFile(fileName6, fileData6),
										},
									},
								},
							},
						},
					},
					{
						name: folderID5,
						children: []*expectedNode{
							newExpectedFile(fileName7, fileData7),
							newExpectedFile(fileName8, fileData8),
						},
					},
				}),
			expectedIterateCounts: map[string]int{
				// Traversed because it has a collection associated with it.
				folderID1: 1,
				// Traversed because it has a collection associated with it.
				folderID2: 1,
				// Folder3 is pruned because there's no changes under it.
				folderID3: 0,
				// Folder4 is pruned because there's no changes under it.
				folderID4: 0,
				folderID5: 1,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			progress := &corsoProgress{
				ctx:     ctx,
				pending: map[string]*itemDetails{},
				toMerge: newMergeDetails(),
				errs:    fault.New(true),
			}
			snapshotRoot, counters := getBaseSnapshot()
			msw := &mockSnapshotWalker{
				snapshotRoot: snapshotRoot,
			}

			ie := pmMock.NewPrefixMap(nil)
			if test.inputExcludes != nil {
				ie = test.inputExcludes
			}

			dirTree, err := inflateDirTree(
				ctx,
				msw,
				[]ManifestEntry{
					makeManifestEntry("", tenant, user, path.OneDriveService, path.FilesCategory),
				},
				test.inputCollections(t),
				ie,
				progress)
			require.NoError(t, err, clues.ToCore(err))

			// Check iterate counts before checking tree content as checking tree
			// content can disturb the counter values.
			for name, count := range test.expectedIterateCounts {
				c, ok := counters[name]
				assert.True(t, ok, "unexpected counter %q", name)
				assert.Equal(t, count, *c, "folder %q iterate count", name)
			}

			expectTree(t, ctx, test.expected, dirTree)
		})
	}
}

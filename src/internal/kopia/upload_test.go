package kopia

import (
	"bytes"
	"context"
	"io"
	stdpath "path"
	"testing"
	"time"

	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/fs/virtualfs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/kopia/kopia/snapshot/snapshotfs"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

func makePath(t *testing.T, elements []string, isItem bool) path.Path {
	p, err := path.FromDataLayerPath(stdpath.Join(elements...), isItem)
	require.NoError(t, err)

	return p
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
			children,
		)
	}

	return virtualfs.NewStaticDirectory(
		encodeElements(basic[0])[0],
		[]fs.Entry{
			baseWithChildren(basic[1:], children),
		},
	)
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
	if !assert.NoErrorf(t, err, "getting reader for file: %s", name) {
		return
	}

	// Need to wrap with a restore stream reader to remove the version.
	r = &restoreStreamReader{
		ReadCloser:      io.NopCloser(r),
		expectedVersion: serializationVersion,
	}

	got, err := io.ReadAll(r)
	if !assert.NoErrorf(t, err, "reading data in file: %s", name) {
		return
	}

	assert.Equalf(t, expected, got, "data in file: %s", name)
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

	if exactly {
		require.Len(t, entries, len(dirs))
	}

	names := make([]string, 0, len(entries))
	for _, e := range entries {
		names = append(names, e.Name())
	}

	assert.Subset(t, names, dirs)
}

func getDirEntriesForEntry(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	entry fs.Entry,
) []fs.Entry {
	d, ok := entry.(fs.Directory)
	require.True(t, ok, "entry is not a directory")

	entries, err := fs.GetAllEntries(ctx, d)
	require.NoError(t, err)

	return entries
}

// ---------------
// unit tests
// ---------------
type limitedRangeReader struct {
	readLen int
	io.ReadCloser
}

func (lrr *limitedRangeReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		// Not well specified behavior, defer to underlying reader.
		return lrr.ReadCloser.Read(p)
	}

	toRead := lrr.readLen
	if len(p) < toRead {
		toRead = len(p)
	}

	return lrr.ReadCloser.Read(p[:toRead])
}

type VersionReadersUnitSuite struct {
	tester.Suite
}

func TestVersionReadersUnitSuite(t *testing.T) {
	suite.Run(t, &VersionReadersUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *VersionReadersUnitSuite) TestWriteAndRead() {
	inputData := []byte("This is some data for the reader to test with")
	table := []struct {
		name         string
		readVersion  uint32
		writeVersion uint32
		check        assert.ErrorAssertionFunc
	}{
		{
			name:         "SameVersionSucceeds",
			readVersion:  42,
			writeVersion: 42,
			check:        assert.NoError,
		},
		{
			name:         "DifferentVersionsFail",
			readVersion:  7,
			writeVersion: 42,
			check:        assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			baseReader := bytes.NewReader(inputData)

			reversible := &restoreStreamReader{
				expectedVersion: test.readVersion,
				ReadCloser: newBackupStreamReader(
					test.writeVersion,
					io.NopCloser(baseReader),
				),
			}

			defer reversible.Close()

			allData, err := io.ReadAll(reversible)
			test.check(t, err)

			if err != nil {
				return
			}

			assert.Equal(t, inputData, allData)
		})
	}
}

func readAllInParts(
	t *testing.T,
	partLen int,
	reader io.ReadCloser,
) ([]byte, int) {
	res := []byte{}
	read := 0
	tmp := make([]byte, partLen)

	for {
		n, err := reader.Read(tmp)
		if errors.Is(err, io.EOF) {
			break
		}

		require.NoError(t, err)

		read += n
		res = append(res, tmp[:n]...)
	}

	return res, read
}

func (suite *VersionReadersUnitSuite) TestWriteHandlesShortReads() {
	t := suite.T()
	inputData := []byte("This is some data for the reader to test with")
	version := uint32(42)
	baseReader := bytes.NewReader(inputData)
	versioner := newBackupStreamReader(version, io.NopCloser(baseReader))
	expectedToWrite := len(inputData) + int(versionSize)

	// "Write" all the data.
	versionedData, writtenLen := readAllInParts(t, 1, versioner)
	assert.Equal(t, expectedToWrite, writtenLen)

	// Read all of the data back.
	baseReader = bytes.NewReader(versionedData)
	reader := &restoreStreamReader{
		expectedVersion: version,
		// Be adversarial and only allow reads of length 1 from the byte reader.
		ReadCloser: &limitedRangeReader{
			readLen:    1,
			ReadCloser: io.NopCloser(baseReader),
		},
	}
	readData, readLen := readAllInParts(t, 1, reader)
	// This reports the bytes read and returned to the user, excluding the version
	// that is stripped off at the start.
	assert.Equal(t, len(inputData), readLen)
	assert.Equal(t, inputData, readData)
}

type CorsoProgressUnitSuite struct {
	tester.Suite
	targetFilePath path.Path
	targetFileName string
}

func TestCorsoProgressUnitSuite(t *testing.T) {
	suite.Run(t, &CorsoProgressUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CorsoProgressUnitSuite) SetupSuite() {
	p, err := path.Builder{}.Append(
		testInboxDir,
		"testFile",
	).ToDataLayerExchangePathForCategory(
		testTenant,
		testUser,
		path.EmailCategory,
		true,
	)
	require.NoError(suite.T(), err)

	suite.targetFilePath = p
	suite.targetFileName = suite.targetFilePath.ToBuilder().Dir().String()
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
	err                error
}{
	{
		name: "DetailsExist",
		cachedItems: func(fname string, fpath path.Path) map[string]testInfo {
			return map[string]testInfo{
				fname: {
					info:       &itemDetails{info: &details.ItemInfo{}, repoPath: fpath},
					err:        nil,
					totalBytes: 100,
				},
			}
		},
		expectedBytes: 100,
		// 1 file and 5 folders.
		expectedNumEntries: 6,
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
					info: &itemDetails{info: &details.ItemInfo{}, repoPath: fpath},
					err:  assert.AnError,
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
		name   string
		cached bool
	}{
		{
			name:   "all updated",
			cached: false,
		},
		{
			name:   "all cached",
			cached: true,
		},
	}

	for _, cachedTest := range table {
		suite.Run(cachedTest.name, func() {
			for _, test := range finishedFileTable {
				suite.Run(test.name, func() {
					t := suite.T()

					bd := &details.Builder{}
					cp := corsoProgress{
						UploadProgress: &snapshotfs.NullUploadProgress{},
						deets:          bd,
						pending:        map[string]*itemDetails{},
						errs:           fault.New(true),
					}

					ci := test.cachedItems(suite.targetFileName, suite.targetFilePath)

					for k, v := range ci {
						cp.put(k, v.info)
					}

					require.Len(t, cp.pending, len(ci))

					for k, v := range ci {
						if cachedTest.cached {
							cp.CachedFile(k, v.totalBytes)
						}

						cp.FinishedFile(k, v.err)
					}

					assert.Empty(t, cp.pending)

					entries := bd.Details().Entries

					assert.Len(t, entries, test.expectedNumEntries)

					for _, entry := range entries {
						assert.Equal(t, !cachedTest.cached, entry.Updated)
					}
				})
			}
		})
	}
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileCachedNoPrevPathErrors() {
	t := suite.T()
	bd := &details.Builder{}
	cachedItems := map[string]testInfo{
		suite.targetFileName: {
			info:       &itemDetails{info: nil, repoPath: suite.targetFilePath},
			err:        nil,
			totalBytes: 100,
		},
	}
	cp := corsoProgress{
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
	assert.Error(t, cp.errs.Failure())
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileBuildsHierarchyNewItem() {
	t := suite.T()
	// Order of folders in hierarchy from root to leaf (excluding the item).
	expectedFolderOrder := suite.targetFilePath.ToBuilder().Dir().Elements()

	// Setup stuff.
	bd := &details.Builder{}
	cp := corsoProgress{
		UploadProgress: &snapshotfs.NullUploadProgress{},
		deets:          bd,
		pending:        map[string]*itemDetails{},
		toMerge:        map[string]PrevRefs{},
		errs:           fault.New(true),
	}

	deets := &itemDetails{info: &details.ItemInfo{}, repoPath: suite.targetFilePath}
	cp.put(suite.targetFileName, deets)
	require.Len(t, cp.pending, 1)

	cp.FinishedFile(suite.targetFileName, nil)

	assert.Empty(t, cp.toMerge)

	// Gather information about the current state.
	var (
		curRef     *details.DetailsEntry
		refToEntry = map[string]*details.DetailsEntry{}
	)

	entries := bd.Details().Entries

	for i := 0; i < len(entries); i++ {
		e := &entries[i]
		if e.Folder == nil {
			continue
		}

		refToEntry[e.ShortRef] = e

		if e.Folder.DisplayName == expectedFolderOrder[len(expectedFolderOrder)-1] {
			curRef = e
		}
	}

	// Actual tests start here.
	var rootRef *details.DetailsEntry

	// Traverse the details entries from leaf to root, following the ParentRef
	// fields. At the end rootRef should point to the root of the path.
	for i := len(expectedFolderOrder) - 1; i >= 0; i-- {
		name := expectedFolderOrder[i]

		require.NotNil(t, curRef)
		assert.Equal(t, name, curRef.Folder.DisplayName)

		rootRef = curRef
		curRef = refToEntry[curRef.ParentRef]
	}

	// Hierarchy root's ParentRef = "" and map will return nil.
	assert.Nil(t, curRef)
	require.NotNil(t, rootRef)
	assert.Empty(t, rootRef.ParentRef)
}

func (suite *CorsoProgressUnitSuite) TestFinishedFileBaseItemDoesntBuildHierarchy() {
	t := suite.T()

	prevPath := makePath(
		suite.T(),
		[]string{testTenant, service, testUser, category, testInboxDir, testFileName2},
		true,
	)

	expectedToMerge := map[string]PrevRefs{
		prevPath.ShortRef(): {
			Repo:     suite.targetFilePath,
			Location: suite.targetFilePath,
		},
	}

	// Setup stuff.
	db := &details.Builder{}
	cp := corsoProgress{
		UploadProgress: &snapshotfs.NullUploadProgress{},
		deets:          db,
		pending:        map[string]*itemDetails{},
		toMerge:        map[string]PrevRefs{},
		errs:           fault.New(true),
	}

	deets := &itemDetails{
		info:         nil,
		repoPath:     suite.targetFilePath,
		prevPath:     prevPath,
		locationPath: suite.targetFilePath,
	}

	cp.put(suite.targetFileName, deets)
	require.Len(t, cp.pending, 1)

	cp.FinishedFile(suite.targetFileName, nil)
	assert.Equal(t, expectedToMerge, cp.toMerge)
	assert.Empty(t, cp.deets)
}

func (suite *CorsoProgressUnitSuite) TestFinishedHashingFile() {
	for _, test := range finishedFileTable {
		suite.Run(test.name, func() {
			t := suite.T()

			bd := &details.Builder{}
			cp := corsoProgress{
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
	tester.LogTimeOfTest(suite.T())
	ctx, flush := tester.NewContext()

	defer flush()

	var (
		t            = suite.T()
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
		pending: map[string]*itemDetails{},
		errs:    fault.New(true),
	}

	collections := []data.BackupCollection{
		mockconnector.NewMockExchangeCollection(
			suite.testStoragePath,
			suite.testLocationPath,
			expectedFileCount[user1Encoded]),
		mockconnector.NewMockExchangeCollection(
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
	dirTree, err := inflateDirTree(ctx, nil, nil, collections, nil, progress)
	require.NoError(t, err)

	assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

	entries, err := fs.GetAllEntries(ctx, dirTree)
	require.NoError(t, err)

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
	ctx, flush := tester.NewContext()
	defer flush()

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
				mockconnector.NewMockExchangeCollection(
					storeP2,
					locP2,
					5),
				mockconnector.NewMockExchangeCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					42),
			},
		},
		{
			name: "SubdirLast",
			layout: []data.BackupCollection{
				mockconnector.NewMockExchangeCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					42),
				mockconnector.NewMockExchangeCollection(
					storeP2,
					locP2,
					5),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			progress := &corsoProgress{
				pending: map[string]*itemDetails{},
				errs:    fault.New(true),
			}

			dirTree, err := inflateDirTree(ctx, nil, nil, test.layout, nil, progress)
			require.NoError(t, err)

			assert.Equal(t, encodeAsPath(testTenant), dirTree.Name())

			entries, err := fs.GetAllEntries(ctx, dirTree)
			require.NoError(t, err)

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
				mockconnector.NewMockExchangeCollection(
					suite.testStoragePath,
					suite.testLocationPath,
					5),
				mockconnector.NewMockExchangeCollection(
					storeP2,
					locP2,
					42),
			},
		},
		{
			"NoCollectionPath",
			[]data.BackupCollection{
				mockconnector.NewMockExchangeCollection(
					nil,
					nil,
					5),
			},
		},
	}

	for _, test := range table {
		ctx, flush := tester.NewContext()
		defer flush()

		suite.Run(test.name, func() {
			t := suite.T()

			_, err := inflateDirTree(ctx, nil, nil, test.layout, nil, nil)
			assert.Error(t, err)
		})
	}
}

type mockSnapshotWalker struct {
	snapshotRoot fs.Entry
}

func (msw *mockSnapshotWalker) SnapshotRoot(*snapshot.Manifest) (fs.Entry, error) {
	return msw.snapshotRoot, nil
}

func mockIncrementalBase(
	id, tenant, resourceOwner string,
	service path.ServiceType,
	category path.CategoryType,
) IncrementalBase {
	return IncrementalBase{
		Manifest: &snapshot.Manifest{
			ID: manifest.ID(id),
		},
		SubtreePaths: []*path.Builder{
			path.Builder{}.Append(tenant, service.String(), resourceOwner, category.String()),
		},
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

			ctx, flush := tester.NewContext()
			defer flush()

			progress := &corsoProgress{
				pending: map[string]*itemDetails{},
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

				mc := mockconnector.NewMockExchangeCollection(nowPath, locPath, 0)
				mc.ColState = s
				mc.PrevPath = prevPath

				cols = append(cols, mc)
			}

			_, err := inflateDirTree(ctx, nil, nil, cols, nil, progress)
			require.Error(t, err)
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
	)

	// Must be a function that returns a new instance each time as StreamingFile
	// can only return its Reader once.
	getBaseSnapshot := func() fs.Entry {
		return baseWithChildren(
			[]string{
				testTenant,
				service,
				testUser,
				category,
			},
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(testInboxID)[0],
					[]fs.Entry{
						virtualfs.StreamingFileWithModTimeFromReader(
							encodeElements(testFileName)[0],
							time.Time{},
							io.NopCloser(bytes.NewReader(testFileData)),
						),
					},
				),
			},
		)
	}

	table := []struct {
		name             string
		inputCollections func() []data.BackupCollection
		expected         *expectedNode
	}{
		{
			name: "SkipsDeletedItems",
			inputCollections: func() []data.BackupCollection {
				mc := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc.Names[0] = testFileName
				mc.DeletedItems[0] = true

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name:     testInboxID,
						children: []*expectedNode{},
					},
				},
			),
		},
		{
			name: "AddsNewItems",
			inputCollections: func() []data.BackupCollection {
				mc := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc.Names[0] = testFileName2
				mc.Data[0] = testFileData2
				mc.ColState = data.NotMovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     testFileName,
								children: []*expectedNode{},
							},
							{
								name:     testFileName2,
								children: []*expectedNode{},
								data:     testFileData2,
							},
						},
					},
				},
			),
		},
		{
			name: "SkipsUpdatedItems",
			inputCollections: func() []data.BackupCollection {
				mc := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc.Names[0] = testFileName
				mc.Data[0] = testFileData2
				mc.ColState = data.NotMovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     testFileName,
								children: []*expectedNode{},
								data:     testFileData2,
							},
						},
					},
				},
			),
		},
		{
			name: "DeleteAndNew",
			inputCollections: func() []data.BackupCollection {
				mc1 := mockconnector.NewMockExchangeCollection(storePath, locPath, 0)
				mc1.ColState = data.DeletedState
				mc1.PrevPath = storePath

				mc2 := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc2.ColState = data.NewState
				mc2.Names[0] = testFileName2
				mc2.Data[0] = testFileData2

				return []data.BackupCollection{mc1, mc2}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     testFileName2,
								children: []*expectedNode{},
								data:     testFileData2,
							},
						},
					},
				},
			),
		},
		{
			name: "MovedAndNew",
			inputCollections: func() []data.BackupCollection {
				mc1 := mockconnector.NewMockExchangeCollection(storePath2, locPath2, 0)
				mc1.ColState = data.MovedState
				mc1.PrevPath = storePath

				mc2 := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc2.ColState = data.NewState
				mc2.Names[0] = testFileName2
				mc2.Data[0] = testFileData2

				return []data.BackupCollection{mc1, mc2}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     testFileName2,
								children: []*expectedNode{},
								data:     testFileData2,
							},
						},
					},
					{
						name: testArchiveID,
						children: []*expectedNode{
							{
								name:     testFileName,
								children: []*expectedNode{},
							},
						},
					},
				},
			),
		},
		{
			name: "NewDoesntMerge",
			inputCollections: func() []data.BackupCollection {
				mc1 := mockconnector.NewMockExchangeCollection(storePath, locPath, 1)
				mc1.ColState = data.NewState
				mc1.Names[0] = testFileName2
				mc1.Data[0] = testFileData2

				return []data.BackupCollection{mc1}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     testFileName2,
								children: []*expectedNode{},
								data:     testFileData2,
							},
						},
					},
				},
			),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tester.LogTimeOfTest(t)

			ctx, flush := tester.NewContext()
			defer flush()

			progress := &corsoProgress{
				pending: map[string]*itemDetails{},
				errs:    fault.New(true),
			}
			msw := &mockSnapshotWalker{
				snapshotRoot: getBaseSnapshot(),
			}

			dirTree, err := inflateDirTree(
				ctx,
				msw,
				[]IncrementalBase{
					mockIncrementalBase("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
				},
				test.inputCollections(),
				nil,
				progress,
			)
			require.NoError(t, err)

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
			[]string{
				testTenant,
				service,
				testUser,
				category,
			},
			[]fs.Entry{
				virtualfs.NewStaticDirectory(
					encodeElements(testInboxID)[0],
					[]fs.Entry{
						virtualfs.StreamingFileWithModTimeFromReader(
							encodeElements(inboxFileName1)[0],
							time.Time{},
							io.NopCloser(bytes.NewReader(inboxFileData1)),
						),
						virtualfs.NewStaticDirectory(
							encodeElements(personalID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(personalFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData)),
								),
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(personalFileName2)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData2)),
								),
							},
						),
						virtualfs.NewStaticDirectory(
							encodeElements(workID)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(workFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData3)),
								),
							},
						),
					},
				),
			},
		)
	}

	table := []struct {
		name             string
		inputCollections func(t *testing.T) []data.BackupCollection
		inputExcludes    map[string]struct{}
		expected         *expectedNode
	}{
		{
			name: "GlobalExcludeSet",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				return nil
			},
			inputExcludes: map[string]struct{}{
				inboxFileName1: {},
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name: personalID,
								children: []*expectedNode{
									{
										name:     personalFileName2,
										children: []*expectedNode{},
									},
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name:     workFileName1,
										children: []*expectedNode{},
									},
								},
							},
						},
					},
				},
			),
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

				mc := mockconnector.NewMockExchangeCollection(newStorePath, newLocPath, 0)
				mc.PrevPath = inboxStorePath
				mc.ColState = data.MovedState

				return []data.BackupCollection{mc}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID + "2",
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name:     personalFileName1,
										children: []*expectedNode{},
									},
									{
										name:     personalFileName2,
										children: []*expectedNode{},
									},
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name:     workFileName1,
										children: []*expectedNode{},
									},
								},
							},
						},
					},
				},
			),
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

				inbox := mockconnector.NewMockExchangeCollection(newInboxStorePath, newInboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.MovedState

				work := mockconnector.NewMockExchangeCollection(newWorkStorePath, newWorkLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID + "2",
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name:     personalFileName1,
										children: []*expectedNode{},
									},
									{
										name:     personalFileName2,
										children: []*expectedNode{},
									},
								},
							},
						},
					},
					{
						name: workID,
						children: []*expectedNode{
							{
								name:     workFileName1,
								children: []*expectedNode{},
							},
						},
					},
				},
			),
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

				inbox := mockconnector.NewMockExchangeCollection(inboxStorePath, inboxLocPath, 0)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.DeletedState

				work := mockconnector.NewMockExchangeCollection(newWorkStorePath, newWorkLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: workID,
						children: []*expectedNode{
							{
								name:     workFileName1,
								children: []*expectedNode{},
							},
						},
					},
				},
			),
		},
		{
			name: "ReplaceDeletedDirectory",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				personal := mockconnector.NewMockExchangeCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.DeletedState

				work := mockconnector.NewMockExchangeCollection(personalStorePath, personalLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{personal, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name: workFileName1,
									},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "ReplaceDeletedDirectoryWithNew",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				personal := mockconnector.NewMockExchangeCollection(personalStorePath, personalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.DeletedState

				newCol := mockconnector.NewMockExchangeCollection(personalStorePath, personalLocPath, 1)
				newCol.ColState = data.NewState
				newCol.Names[0] = workFileName2
				newCol.Data[0] = workFileData2

				return []data.BackupCollection{personal, newCol}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name: workFileName2,
										data: workFileData2,
									},
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name: workFileName1,
									},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "ReplaceDeletedSubtreeWithNew",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				oldInbox := mockconnector.NewMockExchangeCollection(inboxStorePath, inboxLocPath, 0)
				oldInbox.PrevPath = inboxStorePath
				oldInbox.ColState = data.DeletedState

				newCol := mockconnector.NewMockExchangeCollection(inboxStorePath, inboxLocPath, 1)
				newCol.ColState = data.NewState
				newCol.Names[0] = workFileName2
				newCol.Data[0] = workFileData2

				return []data.BackupCollection{oldInbox, newCol}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name: workFileName2,
								data: workFileData2,
							},
						},
					},
				},
			),
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

				personal := mockconnector.NewMockExchangeCollection(newPersonalStorePath, newPersonalLocPath, 0)
				personal.PrevPath = personalStorePath
				personal.ColState = data.MovedState

				work := mockconnector.NewMockExchangeCollection(personalStorePath, personalLocPath, 0)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState

				return []data.BackupCollection{personal, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name: workFileName1,
									},
								},
							},
						},
					},
					{
						name: personalID,
						children: []*expectedNode{
							{
								name: personalFileName1,
							},
							{
								name: personalFileName2,
							},
						},
					},
				},
			),
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

				personal := mockconnector.NewMockExchangeCollection(newPersonalStorePath, newPersonalLocPath, 2)
				personal.PrevPath = personalStorePath
				personal.ColState = data.MovedState
				personal.Names[0] = personalFileName2
				personal.Data[0] = testFileData5
				personal.Names[1] = testFileName4
				personal.Data[1] = testFileData4

				return []data.BackupCollection{personal}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name:     workFileName1,
										children: []*expectedNode{},
									},
								},
							},
						},
					},
					{
						name: workID,
						children: []*expectedNode{
							{
								name: personalFileName1,
							},
							{
								name: personalFileName2,
								data: testFileData5,
							},
							{
								name: testFileName4,
								data: testFileData4,
							},
						},
					},
				},
			),
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

				inbox := mockconnector.NewMockExchangeCollection(newInboxStorePath, newInboxLocPath, 1)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.MovedState
				inbox.DoNotMerge = true
				// First file in inbox is implicitly deleted as we're not merging items
				// and it's not listed.
				inbox.Names[0] = inboxFileName2
				inbox.Data[0] = inboxFileData2

				work := mockconnector.NewMockExchangeCollection(newWorkStorePath, newWorkLocPath, 1)
				work.PrevPath = workStorePath
				work.ColState = data.MovedState
				work.Names[0] = testFileName6
				work.Data[0] = testFileData6

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: personalID,
						children: []*expectedNode{
							{
								name:     inboxFileName2,
								children: []*expectedNode{},
								data:     inboxFileData2,
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name:     personalFileName1,
										children: []*expectedNode{},
									},
									{
										name:     personalFileName2,
										children: []*expectedNode{},
									},
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name:     workFileName1,
										children: []*expectedNode{},
									},
									{
										name:     testFileName6,
										children: []*expectedNode{},
										data:     testFileData6,
									},
								},
							},
						},
					},
				},
			),
		},
		{
			name: "NoMoveParentDeleteFileNoMergeSubtreeMerge",
			inputCollections: func(t *testing.T) []data.BackupCollection {
				inbox := mockconnector.NewMockExchangeCollection(inboxStorePath, inboxLocPath, 1)
				inbox.PrevPath = inboxStorePath
				inbox.ColState = data.NotMovedState
				inbox.DoNotMerge = true
				// First file in inbox is implicitly deleted as we're not merging items
				// and it's not listed.
				inbox.Names[0] = inboxFileName2
				inbox.Data[0] = inboxFileData2

				work := mockconnector.NewMockExchangeCollection(workStorePath, workLocPath, 1)
				work.PrevPath = workStorePath
				work.ColState = data.NotMovedState
				work.Names[0] = testFileName6
				work.Data[0] = testFileData6

				return []data.BackupCollection{inbox, work}
			},
			expected: expectedTreeWithChildren(
				[]string{
					testTenant,
					service,
					testUser,
					category,
				},
				[]*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName2,
								children: []*expectedNode{},
								data:     inboxFileData2,
							},
							{
								name: personalID,
								children: []*expectedNode{
									{
										name:     personalFileName1,
										children: []*expectedNode{},
									},
									{
										name:     personalFileName2,
										children: []*expectedNode{},
									},
								},
							},
							{
								name: workID,
								children: []*expectedNode{
									{
										name:     workFileName1,
										children: []*expectedNode{},
									},
									{
										name:     testFileName6,
										children: []*expectedNode{},
										data:     testFileData6,
									},
								},
							},
						},
					},
				},
			),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			tester.LogTimeOfTest(t)

			ctx, flush := tester.NewContext()
			defer flush()

			progress := &corsoProgress{
				pending: map[string]*itemDetails{},
				errs:    fault.New(true),
			}
			msw := &mockSnapshotWalker{
				snapshotRoot: getBaseSnapshot(),
			}

			dirTree, err := inflateDirTree(
				ctx,
				msw,
				[]IncrementalBase{
					mockIncrementalBase("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
				},
				test.inputCollections(t),
				test.inputExcludes,
				progress)
			require.NoError(t, err)

			expectTree(t, ctx, test.expected, dirTree)
		})
	}
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSkipsDeletedSubtree() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	const (
		personalDir = "personal"
		workDir     = "work"
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
			[]string{
				testTenant,
				service,
				testUser,
				category,
			},
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
									io.NopCloser(bytes.NewReader(testFileData)),
								),
							},
						),
						virtualfs.NewStaticDirectory(
							encodeElements(workDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName2)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData2)),
								),
							},
						),
					},
				),
				virtualfs.NewStaticDirectory(
					encodeElements(testArchiveID)[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(personalDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName3)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData3)),
								),
							},
						),
						virtualfs.NewStaticDirectory(
							encodeElements(workDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(testFileName4)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(testFileData4)),
								),
							},
						),
					},
				),
			},
		)
	}

	expected := expectedTreeWithChildren(
		[]string{
			testTenant,
			service,
			testUser,
			category,
		},
		[]*expectedNode{
			{
				name: testArchiveID,
				children: []*expectedNode{
					{
						name: personalDir,
						children: []*expectedNode{
							{
								name:     testFileName3,
								children: []*expectedNode{},
							},
						},
					},
					{
						name: workDir,
						children: []*expectedNode{
							{
								name:     testFileName4,
								children: []*expectedNode{},
							},
						},
					},
				},
			},
		},
	)

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		errs:    fault.New(true),
	}
	mc := mockconnector.NewMockExchangeCollection(suite.testStoragePath, suite.testStoragePath, 1)
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
		[]IncrementalBase{
			mockIncrementalBase("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		nil,
		progress)
	require.NoError(t, err)

	expectTree(t, ctx, expected, dirTree)
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTree_HandleEmptyBase() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext()
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
					io.NopCloser(bytes.NewReader(testFileData)),
				),
			},
		)
	}

	// Metadata subtree doesn't appear because we don't select it as one of the
	// subpaths and we're not passing in a metadata collection.
	expected := expectedTreeWithChildren(
		[]string{
			testTenant,
			service,
			testUser,
			category,
		},
		[]*expectedNode{
			{
				name: testArchiveID,
				children: []*expectedNode{
					{
						name:     testFileName2,
						children: []*expectedNode{},
					},
				},
			},
		},
	)

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		errs:    fault.New(true),
	}
	mc := mockconnector.NewMockExchangeCollection(archiveStorePath, archiveLocPath, 1)
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
		[]IncrementalBase{
			mockIncrementalBase("", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		nil,
		progress)
	require.NoError(t, err)

	expectTree(t, ctx, expected, dirTree)
}

type mockMultiSnapshotWalker struct {
	snaps map[string]fs.Entry
}

func (msw *mockMultiSnapshotWalker) SnapshotRoot(man *snapshot.Manifest) (fs.Entry, error) {
	if snap := msw.snaps[string(man.ID)]; snap != nil {
		return snap, nil
	}

	return nil, errors.New("snapshot not found")
}

func (suite *HierarchyBuilderUnitSuite) TestBuildDirectoryTreeSelectsCorrectSubtrees() {
	tester.LogTimeOfTest(suite.T())
	t := suite.T()

	ctx, flush := tester.NewContext()
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
			[]string{
				testTenant,
				service,
				testUser,
			},
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
									io.NopCloser(bytes.NewReader(inboxFileData1)),
								),
							},
						),
					},
				),
				virtualfs.NewStaticDirectory(
					encodeElements(path.ContactsCategory.String())[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements(contactsDir)[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(contactsFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(contactsFileData1)),
								),
							},
						),
					},
				),
			},
		)
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
			[]string{
				testTenant,
				service,
				testUser,
			},
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
									newBackupStreamReader(
										serializationVersion,
										io.NopCloser(bytes.NewReader(inboxFileData1v2)),
									),
								),
							},
						),
					},
				),
				virtualfs.NewStaticDirectory(
					encodeElements(path.EventsCategory.String())[0],
					[]fs.Entry{
						virtualfs.NewStaticDirectory(
							encodeElements("events")[0],
							[]fs.Entry{
								virtualfs.StreamingFileWithModTimeFromReader(
									encodeElements(eventsFileName1)[0],
									time.Time{},
									io.NopCloser(bytes.NewReader(eventsFileData1)),
								),
							},
						),
					},
				),
			},
		)
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
		[]string{
			testTenant,
			service,
			testUser,
		},
		[]*expectedNode{
			{
				name: category,
				children: []*expectedNode{
					{
						name: testInboxID,
						children: []*expectedNode{
							{
								name:     inboxFileName1,
								children: []*expectedNode{},
								data:     inboxFileData1v2,
							},
							{
								name:     inboxFileName2,
								children: []*expectedNode{},
								data:     inboxFileData2,
							},
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
							{
								name:     contactsFileName1,
								children: []*expectedNode{},
							},
						},
					},
				},
			},
		},
	)

	progress := &corsoProgress{
		pending: map[string]*itemDetails{},
		errs:    fault.New(true),
	}

	mc := mockconnector.NewMockExchangeCollection(inboxPath, inboxPath, 1)
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
		[]IncrementalBase{
			mockIncrementalBase("id1", testTenant, testUser, path.ExchangeService, path.ContactsCategory),
			mockIncrementalBase("id2", testTenant, testUser, path.ExchangeService, path.EmailCategory),
		},
		collections,
		nil,
		progress,
	)
	require.NoError(t, err)

	expectTree(t, ctx, expected, dirTree)
}

package selectors

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type OneDriveSelectorSuite struct {
	tester.Suite
}

func TestOneDriveSelectorSuite(t *testing.T) {
	suite.Run(t, &OneDriveSelectorSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup(Any())
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup(Any())
	s := ob.Selector
	ob, err := s.ToOneDriveBackup()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_AllData() {
	t := suite.T()

	var (
		users     = []string{"u1", "u2"}
		sel       = NewOneDriveBackup(users)
		allScopes = sel.AllData()
	)

	assert.ElementsMatch(t, users, sel.DiscreteResourceOwners())

	// Initialize the selector Include, Exclude, Filter
	sel.Exclude(allScopes)
	sel.Include(allScopes)
	sel.Filter(allScopes)

	table := []struct {
		name          string
		scopesToCheck []scope
	}{
		{"Include Scopes", sel.Includes},
		{"Exclude Scopes", sel.Excludes},
		{"info scopes", sel.Filters},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			require.Len(t, test.scopesToCheck, 1)
			for _, scope := range test.scopesToCheck {
				scopeMustHave(
					t,
					OneDriveScope(scope),
					map[categorizer][]string{
						OneDriveItem:   Any(),
						OneDriveFolder: Any(),
					},
				)
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Include_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	var (
		users     = []string{u1, u2}
		sel       = NewOneDriveBackup(users)
		allScopes = sel.AllData()
	)

	sel.Include(allScopes)
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer][]string{
				OneDriveItem:   Any(),
				OneDriveFolder: Any(),
			},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Exclude_AllData() {
	t := suite.T()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	var (
		users     = []string{u1, u2}
		sel       = NewOneDriveBackup(users)
		allScopes = sel.AllData()
	)

	sel.Exclude(allScopes)
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer][]string{
				OneDriveItem:   Any(),
				OneDriveFolder: Any(),
			},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveRestore() {
	t := suite.T()
	or := NewOneDriveRestore(Any())
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveRestore() {
	t := suite.T()
	eb := NewOneDriveRestore(Any())
	s := eb.Selector
	or, err := s.ToOneDriveRestore()
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}

func (suite *OneDriveSelectorSuite) TestOneDriveRestore_Reduce() {
	var (
		file = stubRepoRef(
			path.OneDriveService,
			path.FilesCategory,
			"uid",
			"drive/driveID/root:/folderA.d/folderB.d",
			"file")
		fileParent = "folderA/folderB"
		file2      = stubRepoRef(
			path.OneDriveService,
			path.FilesCategory,
			"uid",
			"drive/driveID/root:/folderA.d/folderC.d",
			"file2")
		fileParent2 = "folderA/folderC"
		file3       = stubRepoRef(
			path.OneDriveService,
			path.FilesCategory,
			"uid",
			"drive/driveID/root:/folderD.d/folderE.d",
			"file3")
		fileParent3 = "folderD/folderE"
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.Entry{
				{
					RepoRef: file,
					ItemRef: "file",
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType:   details.OneDriveItem,
							ItemName:   "fileName",
							ParentPath: fileParent,
						},
					},
				},
				{
					RepoRef: file2,
					ItemRef: "file2",
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType:   details.OneDriveItem,
							ItemName:   "fileName2",
							ParentPath: fileParent2,
						},
					},
				},
				{
					RepoRef: file3,
					// item ref intentionally blank to assert fallback case
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType:   details.OneDriveItem,
							ItemName:   "fileName3",
							ParentPath: fileParent3,
						},
					},
				},
			},
		},
	}

	arr := func(s ...string) []string {
		return s
	}

	table := []struct {
		name         string
		makeSelector func() *OneDriveRestore
		expect       []string
		cfg          Config
	}{
		{
			name: "all",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.AllData())
				return odr
			},
			expect: arr(file, file2, file3),
		},
		{
			name: "only match file",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.Items(Any(), []string{"file2"}))
				return odr
			},
			expect: arr(file2),
		},
		{
			name: "id doesn't match name",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.Items(Any(), []string{"file2"}))
				return odr
			},
			expect: []string{},
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "only match file name",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.Items(Any(), []string{"fileName2"}))
				return odr
			},
			expect: arr(file2),
			cfg:    Config{OnlyMatchItemNames: true},
		},
		{
			name: "name doesn't match id",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.Items(Any(), []string{"fileName2"}))
				return odr
			},
			expect: []string{},
		},
		{
			name: "only match folder",
			makeSelector: func() *OneDriveRestore {
				odr := NewOneDriveRestore([]string{"uid"})
				odr.Include(odr.Folders([]string{"folderA/folderB", "folderA/folderC"}))
				return odr
			},
			expect: arr(file, file2),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			sel := test.makeSelector()
			sel.Configure(test.cfg)
			results := sel.Reduce(ctx, deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveCategory_PathValues() {
	t := suite.T()

	fileName := "file"
	fileID := fileName + "-id"
	shortRef := "short"
	elems := []string{"drives", "driveID", "root:", "dir1.d", "dir2.d", fileID}

	filePath, err := path.Build("tenant", "user", path.OneDriveService, path.FilesCategory, true, elems...)
	require.NoError(t, err, clues.ToCore(err))

	fileLoc := path.Builder{}.Append("dir1", "dir2")

	table := []struct {
		name      string
		pathElems []string
		expected  map[categorizer][]string
		cfg       Config
	}{
		{
			name:      "items",
			pathElems: elems,
			expected: map[categorizer][]string{
				OneDriveFolder: {"dir1/dir2"},
				OneDriveItem:   {fileID, shortRef},
			},
			cfg: Config{},
		},
		{
			name:      "items w/ name",
			pathElems: elems,
			expected: map[categorizer][]string{
				OneDriveFolder: {"dir1/dir2"},
				OneDriveItem:   {fileName, shortRef},
			},
			cfg: Config{OnlyMatchItemNames: true},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			itemPath, err := path.Build(
				"tenant",
				"site",
				path.OneDriveService,
				path.FilesCategory,
				true,
				test.pathElems...)
			require.NoError(t, err, clues.ToCore(err))

			ent := details.Entry{
				RepoRef:  filePath.String(),
				ShortRef: shortRef,
				ItemRef:  fileID,
				ItemInfo: details.ItemInfo{
					OneDrive: &details.OneDriveInfo{
						ItemName:   fileName,
						ParentPath: fileLoc.String(),
					},
				},
			}

			pv, err := OneDriveItem.pathValues(itemPath, ent, test.cfg)
			require.NoError(t, err)
			assert.Equal(t, test.expected, pv)
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveScope_MatchesInfo() {
	ods := NewOneDriveRestore(Any())

	var (
		epoch  = time.Time{}
		now    = time.Now()
		future = now.Add(1 * time.Minute)
	)

	itemInfo := details.ItemInfo{
		OneDrive: &details.OneDriveInfo{
			ItemType:   details.OneDriveItem,
			ParentPath: "folder1/folder2",
			ItemName:   "file1",
			Size:       10,
			Owner:      "user@email.com",
			Created:    now,
			Modified:   now,
		},
	}

	table := []struct {
		name   string
		scope  []OneDriveScope
		expect assert.BoolAssertionFunc
	}{
		{"file create after the epoch", ods.CreatedAfter(dttm.Format(epoch)), assert.True},
		{"file create after now", ods.CreatedAfter(dttm.Format(now)), assert.False},
		{"file create after later", ods.CreatedAfter(dttm.Format(future)), assert.False},
		{"file create before future", ods.CreatedBefore(dttm.Format(future)), assert.True},
		{"file create before now", ods.CreatedBefore(dttm.Format(now)), assert.False},
		{"file create before epoch", ods.CreatedBefore(dttm.Format(now)), assert.False},
		{"file modified after the epoch", ods.ModifiedAfter(dttm.Format(epoch)), assert.True},
		{"file modified after now", ods.ModifiedAfter(dttm.Format(now)), assert.False},
		{"file modified after later", ods.ModifiedAfter(dttm.Format(future)), assert.False},
		{"file modified before future", ods.ModifiedBefore(dttm.Format(future)), assert.True},
		{"file modified before now", ods.ModifiedBefore(dttm.Format(now)), assert.False},
		{"file modified before epoch", ods.ModifiedBefore(dttm.Format(now)), assert.False},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			scopes := setScopesToDefault(test.scope)
			for _, scope := range scopes {
				test.expect(t, scope.matchesInfo(itemInfo))
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestCategory_PathType() {
	table := []struct {
		cat      oneDriveCategory
		pathType path.CategoryType
	}{
		{OneDriveCategoryUnknown, path.UnknownCategory},
		{OneDriveUser, path.UnknownCategory},
		{OneDriveItem, path.FilesCategory},
		{OneDriveFolder, path.FilesCategory},
		{FileInfoCreatedAfter, path.FilesCategory},
		{FileInfoCreatedBefore, path.FilesCategory},
		{FileInfoModifiedAfter, path.FilesCategory},
		{FileInfoModifiedBefore, path.FilesCategory},
	}
	for _, test := range table {
		suite.Run(test.cat.String(), func() {
			assert.Equal(suite.T(), test.pathType, test.cat.PathType())
		})
	}
}

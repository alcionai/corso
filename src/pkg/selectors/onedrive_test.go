package selectors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
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
	require.NoError(t, err)
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
		{"Filter Scopes", sel.Filters},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			require.Len(t, test.scopesToCheck, 1)
			for _, scope := range test.scopesToCheck {
				scopeMustHave(
					t,
					OneDriveScope(scope),
					map[categorizer]string{
						OneDriveItem:   AnyTgt,
						OneDriveFolder: AnyTgt,
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
			map[categorizer]string{
				OneDriveItem:   AnyTgt,
				OneDriveFolder: AnyTgt,
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
			map[categorizer]string{
				OneDriveItem:   AnyTgt,
				OneDriveFolder: AnyTgt,
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
	require.NoError(t, err)
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}

func (suite *OneDriveSelectorSuite) TestOneDriveRestore_Reduce() {
	var (
		file  = stubRepoRef(path.OneDriveService, path.FilesCategory, "uid", "drive/driveID/root:/folderA/folderB", "file")
		file2 = stubRepoRef(path.OneDriveService, path.FilesCategory, "uid", "drive/driveID/root:/folderA/folderC", "file2")
		file3 = stubRepoRef(path.OneDriveService, path.FilesCategory, "uid", "drive/driveID/root:/folderD/folderE", "file3")
	)

	deets := &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.DetailsEntry{
				{
					RepoRef: file,
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType: details.OneDriveItem,
						},
					},
				},
				{
					RepoRef: file2,
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType: details.OneDriveItem,
						},
					},
				},
				{
					RepoRef: file3,
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType: details.OneDriveItem,
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
		deets        *details.Details
		makeSelector func() *OneDriveRestore
		expect       []string
	}{
		{
			"all",
			deets,
			func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.AllData())
				return odr
			},
			arr(file, file2, file3),
		},
		{
			"only match file",
			deets,
			func() *OneDriveRestore {
				odr := NewOneDriveRestore(Any())
				odr.Include(odr.Items(Any(), []string{"file2"}))
				return odr
			},
			arr(file2),
		},
		{
			"only match folder",
			deets,
			func() *OneDriveRestore {
				odr := NewOneDriveRestore([]string{"uid"})
				odr.Include(odr.Folders([]string{"folderA/folderB", "folderA/folderC"}))
				return odr
			},
			arr(file, file2),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets, fault.New(true))
			paths := results.Paths()
			assert.Equal(t, test.expect, paths)
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveCategory_PathValues() {
	t := suite.T()

	pathBuilder := path.Builder{}.Append("drive", "driveID", "root:", "dir1", "dir2", "file")
	filePath, err := pathBuilder.ToDataLayerOneDrivePath("tenant", "user", true)
	require.NoError(t, err)

	expected := map[categorizer][]string{
		OneDriveFolder: {"dir1/dir2"},
		OneDriveItem:   {"file", "short"},
	}

	ent := details.DetailsEntry{
		RepoRef:  filePath.String(),
		ShortRef: "short",
	}

	r := OneDriveItem.pathValues(filePath, ent)
	assert.Equal(t, expected, r)
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
		{"file create after the epoch", ods.CreatedAfter(common.FormatTime(epoch)), assert.True},
		{"file create after now", ods.CreatedAfter(common.FormatTime(now)), assert.False},
		{"file create after later", ods.CreatedAfter(common.FormatTime(future)), assert.False},
		{"file create before future", ods.CreatedBefore(common.FormatTime(future)), assert.True},
		{"file create before now", ods.CreatedBefore(common.FormatTime(now)), assert.False},
		{"file create before epoch", ods.CreatedBefore(common.FormatTime(now)), assert.False},
		{"file modified after the epoch", ods.ModifiedAfter(common.FormatTime(epoch)), assert.True},
		{"file modified after now", ods.ModifiedAfter(common.FormatTime(now)), assert.False},
		{"file modified after later", ods.ModifiedAfter(common.FormatTime(future)), assert.False},
		{"file modified before future", ods.ModifiedBefore(common.FormatTime(future)), assert.True},
		{"file modified before now", ods.ModifiedBefore(common.FormatTime(now)), assert.False},
		{"file modified before epoch", ods.ModifiedBefore(common.FormatTime(now)), assert.False},
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
		{FileFilterCreatedAfter, path.FilesCategory},
		{FileFilterCreatedBefore, path.FilesCategory},
		{FileFilterModifiedAfter, path.FilesCategory},
		{FileFilterModifiedBefore, path.FilesCategory},
	}
	for _, test := range table {
		suite.T().Run(test.cat.String(), func(t *testing.T) {
			assert.Equal(t, test.pathType, test.cat.PathType())
		})
	}
}

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
	"github.com/alcionai/corso/src/pkg/path"
)

type OneDriveSelectorSuite struct {
	suite.Suite
}

func TestOneDriveSelectorSuite(t *testing.T) {
	suite.Run(t, new(OneDriveSelectorSuite))
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveBackup() {
	t := suite.T()
	ob := NewOneDriveBackup()
	s := ob.Selector
	ob, err := s.ToOneDriveBackup()
	require.NoError(t, err)
	assert.Equal(t, ob.Service, ServiceOneDrive)
	assert.NotZero(t, ob.Scopes())
}

func (suite *OneDriveSelectorSuite) TestOneDriveBackup_DiscreteScopes() {
	usrs := []string{"u1", "u2"}
	table := []struct {
		name     string
		include  []string
		discrete []string
		expect   []string
	}{
		{
			name:     "any user",
			include:  Any(),
			discrete: usrs,
			expect:   usrs,
		},
		{
			name:     "discrete user",
			include:  []string{"u3"},
			discrete: usrs,
			expect:   []string{"u3"},
		},
		{
			name:     "nil discrete slice",
			include:  Any(),
			discrete: nil,
			expect:   Any(),
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			eb := NewOneDriveBackup()
			eb.Include(eb.Users(test.include))

			scopes := eb.DiscreteScopes(test.discrete)
			for _, sc := range scopes {
				users := sc.Get(OneDriveUser)
				assert.Equal(t, test.expect, users)
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	userScopes := sel.Users([]string{u1, u2})
	for _, scope := range userScopes {
		// Scope value is either u1 or u2
		assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()].Target)
	}

	// Initialize the selector Include, Exclude, Filter
	sel.Exclude(userScopes)
	sel.Include(userScopes)
	sel.Filter(userScopes)

	table := []struct {
		name          string
		scopesToCheck []scope
	}{
		{"Include Scopes", sel.Includes},
		{"Exclude Scopes", sel.Excludes},
		{"Filter Scopes", sel.Filters},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			require.Len(t, test.scopesToCheck, 1)
			for _, scope := range test.scopesToCheck {
				// Scope value is u1,u2
				assert.Contains(t, join(u1, u2), scope[OneDriveUser.String()].Target)
			}
		})
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Include_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Include(sel.Users([]string{u1, u2}))
	scopes := sel.Includes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer]string{OneDriveUser: join(u1, u2)},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestOneDriveSelector_Exclude_Users() {
	t := suite.T()
	sel := NewOneDriveBackup()

	const (
		u1 = "u1"
		u2 = "u2"
	)

	sel.Exclude(sel.Users([]string{u1, u2}))
	scopes := sel.Excludes
	require.Len(t, scopes, 1)

	for _, sc := range scopes {
		scopeMustHave(
			t,
			OneDriveScope(sc),
			map[categorizer]string{OneDriveUser: join(u1, u2)},
		)
	}
}

func (suite *OneDriveSelectorSuite) TestNewOneDriveRestore() {
	t := suite.T()
	or := NewOneDriveRestore()
	assert.Equal(t, or.Service, ServiceOneDrive)
	assert.NotZero(t, or.Scopes())
}

func (suite *OneDriveSelectorSuite) TestToOneDriveRestore() {
	t := suite.T()
	eb := NewOneDriveRestore()
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
				odr := NewOneDriveRestore()
				odr.Include(odr.Users(Any()))
				return odr
			},
			arr(file, file2, file3),
		},
		{
			"only match file",
			deets,
			func() *OneDriveRestore {
				odr := NewOneDriveRestore()
				odr.Include(odr.Items(Any(), Any(), []string{"file2"}))
				return odr
			},
			arr(file2),
		},
		{
			"only match folder",
			deets,
			func() *OneDriveRestore {
				odr := NewOneDriveRestore()
				odr.Include(odr.Folders([]string{"uid"}, []string{"folderA/folderB", "folderA/folderC"}))
				return odr
			},
			arr(file, file2),
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			sel := test.makeSelector()
			results := sel.Reduce(ctx, test.deets)
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

	expected := map[categorizer]string{
		OneDriveUser:   "user",
		OneDriveFolder: "dir1/dir2",
		OneDriveItem:   "file",
	}

	assert.Equal(t, expected, OneDriveItem.pathValues(filePath))
}

func (suite *OneDriveSelectorSuite) TestOneDriveScope_MatchesInfo() {
	ods := NewOneDriveRestore()

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
		suite.T().Run(test.name, func(t *testing.T) {
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

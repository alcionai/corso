package kopia

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type mockLocationIDer struct {
	pb *path.Builder
}

func (ul mockLocationIDer) ID() *path.Builder {
	return ul.pb
}

func (ul mockLocationIDer) InDetails() *path.Builder {
	return ul.pb
}

type DetailsMergeInfoerUnitSuite struct {
	tester.Suite
}

func TestDetailsMergeInfoerUnitSuite(t *testing.T) {
	suite.Run(t, &DetailsMergeInfoerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DetailsMergeInfoerUnitSuite) TestAddRepoRef_DuplicateFails() {
	t := suite.T()
	oldRef1 := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false)

	dm := newMergeDetails()

	err := dm.addRepoRef(oldRef1.ToBuilder(), oldRef1, nil)
	require.NoError(t, err, clues.ToCore(err))

	err = dm.addRepoRef(oldRef1.ToBuilder(), oldRef1, nil)
	require.Error(t, err, clues.ToCore(err))
}

// TestRepoRefs is a basic sanity test to ensure lookups are working properly
// for stored RepoRefs.
func (suite *DetailsMergeInfoerUnitSuite) TestGetNewPathRefs() {
	t := suite.T()
	oldRef1 := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false)
	oldRef2 := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
			"folder2",
		},
		false)
	newRef1 := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder3",
		},
		false)
	newRef2 := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder3",
			"folder2",
		},
		false)
	newLoc1 := path.Builder{}.Append(newRef1.Folders()...)
	newLoc2 := path.Builder{}.Append(newRef2.Folders()...)
	oldLoc1 := path.Builder{}.Append(oldRef1.Folders()...)
	oldLoc2 := path.Builder{}.Append(oldRef2.Folders()...)

	searchLoc1 := mockLocationIDer{oldLoc1}
	searchLoc2 := mockLocationIDer{oldLoc2}

	dm := newMergeDetails()

	err := dm.addRepoRef(oldRef1.ToBuilder(), newRef1, newLoc1)
	require.NoError(t, err, clues.ToCore(err))

	err = dm.addRepoRef(oldRef2.ToBuilder(), newRef2, nil)
	require.NoError(t, err, clues.ToCore(err))

	// Add prefix matcher entry.
	err = dm.addLocation(searchLoc1, newLoc1)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name        string
		searchRef   *path.Builder
		searchLoc   mockLocationIDer
		errCheck    require.ErrorAssertionFunc
		expectedRef path.Path
		expectedLoc *path.Builder
	}{
		{
			name:        "Exact Match With Loc",
			searchRef:   oldRef1.ToBuilder(),
			searchLoc:   searchLoc1,
			errCheck:    require.NoError,
			expectedRef: newRef1,
			expectedLoc: newLoc1,
		},
		{
			name:        "Exact Match Without Loc",
			searchRef:   oldRef1.ToBuilder(),
			errCheck:    require.NoError,
			expectedRef: newRef1,
			expectedLoc: newLoc1,
		},
		{
			name:        "Prefix Match",
			searchRef:   oldRef2.ToBuilder(),
			searchLoc:   searchLoc2,
			errCheck:    require.NoError,
			expectedRef: newRef2,
			expectedLoc: newLoc2,
		},
		{
			name:      "Would Be Prefix Match Without Old Loc Errors",
			searchRef: oldRef2.ToBuilder(),
			errCheck:  require.Error,
		},
		{
			name:      "Not Found With Old Loc",
			searchRef: newRef1.ToBuilder(),
			searchLoc: searchLoc2,
			errCheck:  require.NoError,
		},
		{
			name:      "Not Found Without Old Loc",
			searchRef: newRef1.ToBuilder(),
			errCheck:  require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			newRef, newLoc, err := dm.GetNewPathRefs(test.searchRef, test.searchLoc)
			test.errCheck(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectedRef, newRef, "RepoRef")
			assert.Equal(t, test.expectedLoc, newLoc, "LocationRef")
		})
	}
}

type LocationPrefixMatcherUnitSuite struct {
	tester.Suite
}

func TestLocationPrefixMatcherUnitSuite(t *testing.T) {
	suite.Run(t, &LocationPrefixMatcherUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *LocationPrefixMatcherUnitSuite) TestAdd_Twice_Fails() {
	t := suite.T()
	p := mockLocationIDer{makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false).ToBuilder()}
	loc1 := path.Builder{}.Append("folder1")
	loc2 := path.Builder{}.Append("folder2")

	lpm := newLocationPrefixMatcher()

	err := lpm.add(p, loc1)
	require.NoError(t, err, clues.ToCore(err))

	err = lpm.add(p, loc2)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *LocationPrefixMatcherUnitSuite) TestAdd_And_Match() {
	loc1 := mockLocationIDer{path.Builder{}.Append("folder1")}
	loc2 := mockLocationIDer{loc1.InDetails().Append("folder2")}
	loc3 := mockLocationIDer{path.Builder{}.Append("foo")}

	res1 := mockLocationIDer{path.Builder{}.Append("1")}

	lpm := newLocationPrefixMatcher()

	err := lpm.add(loc1, res1.InDetails())
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name      string
		searchKey mockLocationIDer
		found     bool
	}{
		{
			name:      "Exact Match",
			searchKey: loc1,
			found:     true,
		},
		{
			name:      "No Match",
			searchKey: loc3,
		},
		{
			name:      "Prefix Match",
			searchKey: loc2,
			found:     true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			prefixes := lpm.longestPrefix(test.searchKey)

			if !test.found {
				assert.Nil(t, prefixes.oldLoc)
				assert.Nil(t, prefixes.newLoc)

				return
			}

			assert.Equal(t, loc1.InDetails(), prefixes.oldLoc, "old prefix")
			assert.Equal(t, res1.InDetails(), prefixes.newLoc, "new prefix")
		})
	}
}

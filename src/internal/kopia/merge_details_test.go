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

type fakeUniqueLocation struct {
	pb *path.Builder
}

func (ul fakeUniqueLocation) ID() *path.Builder {
	return ul.pb
}

func (ul fakeUniqueLocation) InDetails() *path.Builder {
	return ul.pb
}

type DetailsMergeInfoerUnitSuite struct {
	tester.Suite
}

func TestDetailsMergeInfoerUnitSuite(t *testing.T) {
	suite.Run(t, &DetailsMergeInfoerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DetailsMergeInfoerUnitSuite) TestAdd_Twice_Fails() {
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
			"folder4",
		},
		false)
	newLoc1 := path.Builder{}.Append(newRef1.Folders()...)
	oldLoc1 := path.Builder{}.Append(oldRef1.Folders()...)
	oldLoc2 := path.Builder{}.Append(oldRef2.Folders()...)

	searchLoc1 := fakeUniqueLocation{oldLoc1}
	searchLoc2 := fakeUniqueLocation{oldLoc2}

	dm := newMergeDetails()

	err := dm.addRepoRef(oldRef1.ToBuilder(), newRef1, newLoc1)
	require.NoError(t, err, clues.ToCore(err))

	err = dm.addRepoRef(oldRef2.ToBuilder(), newRef2, nil)
	require.NoError(t, err, clues.ToCore(err))

	// Add prefix matcher entry.
	err = dm.addLocation(searchLoc1, newLoc1)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name              string
		searchRef         *path.Builder
		searchLoc         fakeUniqueLocation
		expectedRef       path.Path
		prefixFound       bool
		expectedOldPrefix *path.Builder
	}{
		{
			name:              "Exact Match With Loc",
			searchRef:         oldRef1.ToBuilder(),
			searchLoc:         searchLoc1,
			expectedRef:       newRef1,
			prefixFound:       true,
			expectedOldPrefix: oldLoc1,
		},
		{
			name:              "Exact Match Without Loc",
			searchRef:         oldRef1.ToBuilder(),
			expectedRef:       newRef1,
			prefixFound:       true,
			expectedOldPrefix: nil,
		},
		{
			name:              "Prefix Match",
			searchRef:         oldRef2.ToBuilder(),
			searchLoc:         searchLoc2,
			expectedRef:       newRef2,
			prefixFound:       true,
			expectedOldPrefix: oldLoc1,
		},
		{
			name:        "Not Found",
			searchRef:   newRef1.ToBuilder(),
			expectedRef: nil,
		},
		{
			name:        "Not Found With Loc",
			searchRef:   newRef1.ToBuilder(),
			searchLoc:   searchLoc1,
			expectedRef: nil,
		},
		{
			name:        "Ref Found Loc Not",
			searchRef:   oldRef2.ToBuilder(),
			searchLoc:   fakeUniqueLocation{path.Builder{}.Append("foo")},
			expectedRef: newRef2,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			newRef, oldPrefix, newPrefix := dm.GetNewPathRefs(test.searchRef, test.searchLoc)
			assert.Equal(t, test.expectedRef, newRef, "RepoRef")

			if !test.prefixFound {
				assert.Nil(t, oldPrefix)
				assert.Nil(t, newPrefix)
				return
			}

			assert.Equal(t, test.expectedOldPrefix, oldPrefix, "old prefix")
			assert.Equal(t, newLoc1, newPrefix, "new prefix")
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
	p := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false).ToBuilder()
	loc1 := path.Builder{}.Append("folder1")
	loc2 := path.Builder{}.Append("folder2")

	lpm := newLocationPrefixMatcher()

	err := lpm.add(p, loc1)
	require.NoError(t, err, clues.ToCore(err))

	err = lpm.add(p, loc2)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *LocationPrefixMatcherUnitSuite) TestAdd_And_Match() {
	loc1 := path.Builder{}.Append("folder1")
	loc2 := loc1.Append("folder2")
	loc3 := path.Builder{}.Append("foo")

	res1 := path.Builder{}.Append("1")

	lpm := newLocationPrefixMatcher()

	err := lpm.add(loc1, res1)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name      string
		searchKey *path.Builder
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

			oldPrefix, newPrefix := lpm.longestPrefix(test.searchKey)

			if !test.found {
				assert.Nil(t, oldPrefix)
				assert.Nil(t, newPrefix)

				return
			}

			assert.Equal(t, loc1, oldPrefix, "old prefix")
			assert.Equal(t, res1, newPrefix, "new prefix")
		})
	}
}

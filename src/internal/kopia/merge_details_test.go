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

type DetailsMergeInfoerUnitSuite struct {
	tester.Suite
}

func TestDetailsMergeInfoerUnitSuite(t *testing.T) {
	suite.Run(t, &DetailsMergeInfoerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// TestRepoRefs is a basic sanity test to ensure lookups are working properly
// for stored RepoRefs.
func (suite *DetailsMergeInfoerUnitSuite) TestRepoRefs() {
	t := suite.T()
	oldRef := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false).ToBuilder()
	newRef := makePath(
		t,
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder2",
		},
		false)

	dm := newMergeDetails()

	err := dm.addRepoRef(oldRef, newRef)
	require.NoError(t, err, clues.ToCore(err))

	got := dm.GetNewRepoRef(oldRef)
	require.NotNil(t, got)
	assert.Equal(t, newRef.String(), got.String())

	got = dm.GetNewRepoRef(newRef.ToBuilder())
	assert.Nil(t, got)
}

type LocationPrefixMatcherUnitSuite struct {
	tester.Suite
}

func TestLocationPrefixMatcherUnitSuite(t *testing.T) {
	suite.Run(t, &LocationPrefixMatcherUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type inputData struct {
	repoRef path.Path
	locRef  *path.Builder
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
	p1 := makePath(
		suite.T(),
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder1",
		},
		false)

	loc1 := path.Builder{}.Append("folder1")
	p1Parent, err := p1.Dir()
	require.NoError(suite.T(), err, clues.ToCore(err))

	p2 := makePath(
		suite.T(),
		[]string{
			testTenant,
			service,
			testUser,
			category,
			"folder2",
		},
		false)

	table := []struct {
		name      string
		inputs    []inputData
		searchKey string
		check     require.ValueAssertionFunc
		expected  *path.Builder
	}{
		{
			name: "Exact Match",
			inputs: []inputData{
				{
					repoRef: p1,
					locRef:  loc1,
				},
			},
			searchKey: p1.String(),
			check:     require.NotNil,
			expected:  loc1,
		},
		{
			name: "No Match",
			inputs: []inputData{
				{
					repoRef: p1,
					locRef:  loc1,
				},
			},
			searchKey: p2.String(),
			check:     require.Nil,
		},
		{
			name: "No Prefix Match",
			inputs: []inputData{
				{
					repoRef: p1Parent,
					locRef:  loc1,
				},
			},
			searchKey: p1.String(),
			check:     require.Nil,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			lpm := newLocationPrefixMatcher()

			for _, input := range test.inputs {
				err := lpm.add(input.repoRef.ToBuilder(), input.locRef)
				require.NoError(t, err, clues.ToCore(err))
			}

			loc := lpm.longestPrefix(test.searchKey)
			test.check(t, loc)

			if loc == nil {
				return
			}

			assert.Equal(t, test.expected.String(), loc.String())
		})
	}
}

package kopia_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	testTenant = "a-tenant"
	testUser   = "a-user"
	service    = path.ExchangeService
	category   = path.EmailCategory
)

type LocationPrefixMatcherUnitSuite struct {
	tester.Suite
}

func makePath(
	t *testing.T,
	service path.ServiceType,
	category path.CategoryType,
	tenant, user string,
	folders []string,
) path.Path {
	p, err := path.Build(tenant, user, service, category, false, folders...)
	require.NoError(t, err, clues.ToCore(err))

	return p
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
		service,
		category,
		testTenant,
		testUser,
		[]string{"folder1"})
	loc1 := path.Builder{}.Append("folder1")
	loc2 := path.Builder{}.Append("folder2")

	lpm := kopia.NewLocationPrefixMatcher()

	err := lpm.Add(p, loc1)
	require.NoError(t, err, clues.ToCore(err))

	err = lpm.Add(p, loc2)
	assert.Error(t, err, clues.ToCore(err))
}

func (suite *LocationPrefixMatcherUnitSuite) TestAdd_And_Match() {
	p1 := makePath(
		suite.T(),
		service,
		category,
		testTenant,
		testUser,
		[]string{"folder1"})

	p1Parent, err := p1.Dir()
	require.NoError(suite.T(), err, clues.ToCore(err))

	p2 := makePath(
		suite.T(),
		service,
		category,
		testTenant,
		testUser,
		[]string{"folder2"})
	loc1 := path.Builder{}.Append("folder1")

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
			lpm := kopia.NewLocationPrefixMatcher()

			for _, input := range test.inputs {
				err := lpm.Add(input.repoRef, input.locRef)
				require.NoError(t, err, clues.ToCore(err))
			}

			loc := lpm.LongestPrefix(test.searchKey)
			test.check(t, loc)

			if loc == nil {
				return
			}

			assert.Equal(t, test.expected.String(), loc.String())
		})
	}
}

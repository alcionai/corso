package kopia

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type MigrationsUnitSuite struct {
	tester.Suite
}

func TestMigrationsUnitSuite(t *testing.T) {
	suite.Run(t, &MigrationsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MigrationsUnitSuite) TestSubtreeOwnerMigration_GetNewSubtree() {
	table := []struct {
		name   string
		input  *path.Builder
		expect *path.Builder
	}{
		{
			name: "nil builder",
		},
		{
			name:   "builder too small",
			input:  path.Builder{}.Append("foo"),
			expect: path.Builder{}.Append("foo"),
		},
		{
			name:   "non-matching owner",
			input:  path.Builder{}.Append("foo", "bar", "ownerronwo", "baz"),
			expect: path.Builder{}.Append("foo", "bar", "ownerronwo", "baz"),
		},
		{
			name:   "migrated",
			input:  path.Builder{}.Append("foo", "bar", "owner", "baz"),
			expect: path.Builder{}.Append("foo", "bar", "migrated", "baz"),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t      = suite.T()
				stm    = NewSubtreeOwnerMigration("migrated", "owner")
				result = stm.GetNewSubtree(test.input)
			)

			if result == nil {
				assert.Nil(t, test.expect)
				return
			}

			assert.Equal(t, test.expect.String(), result.String())
		})
	}
}

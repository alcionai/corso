package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type GroupBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestGroupBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &GroupBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupBackupHandlerUnitSuite) TestCanonicalPath() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "group",
			expect:    "tenant/groups/resourceOwner/libraries/sites/site-id/prefix",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := NewGroupBackupHandler(resourceOwner, "site-id", api.Drives{}, nil)
			p := path.Builder{}.Append("prefix")

			result, err := h.CanonicalPath(p, tenantID)
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *GroupBackupHandlerUnitSuite) TestServiceCat() {
	t := suite.T()

	s, c := groupBackupHandler{
		libraryBackupHandler: libraryBackupHandler{service: path.GroupsService},
	}.ServiceCat()
	assert.Equal(t, path.GroupsService, s)
	assert.Equal(t, path.LibrariesCategory, c)
}

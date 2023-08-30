package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type LibraryBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestLibraryBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &LibraryBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *LibraryBackupHandlerUnitSuite) TestCanonicalPath() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "sharepoint",
			expect:    "tenant/sharepoint/resourceOwner/libraries/prefix",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := libraryBackupHandler{service: path.SharePointService, siteID: resourceOwner}
			p := path.Builder{}.Append("prefix")

			result, err := h.CanonicalPath(p, tenantID)
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *LibraryBackupHandlerUnitSuite) TestServiceCat() {
	t := suite.T()

	s, c := libraryBackupHandler{service: path.SharePointService}.ServiceCat()
	assert.Equal(t, path.SharePointService, s)
	assert.Equal(t, path.LibrariesCategory, c)
}

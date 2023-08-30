package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type ItemBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestItemBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ItemBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemBackupHandlerUnitSuite) TestCanonicalPath() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "onedrive",
			expect:    "tenant/onedrive/resourceOwner/files/prefix",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := itemBackupHandler{userID: resourceOwner}
			p := path.Builder{}.Append("prefix")

			result, err := h.CanonicalPath(p, tenantID)
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *ItemBackupHandlerUnitSuite) TestServiceCat() {
	t := suite.T()

	s, c := itemBackupHandler{}.ServiceCat()
	assert.Equal(t, path.OneDriveService, s)
	assert.Equal(t, path.FilesCategory, c)
}

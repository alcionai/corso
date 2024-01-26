package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type ItemBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestItemBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &ItemBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemBackupHandlerUnitSuite) TestPathPrefix() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "onedrive",
			expect:    "tenant/onedrive/resourceOwner/files/drives/driveID/root:",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := userDriveBackupHandler{
				baseUserDriveHandler: baseUserDriveHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
			}

			result, err := h.PathPrefix("driveID")
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *ItemBackupHandlerUnitSuite) TestMetadataPathPrefix() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "onedrive",
			expect:    "tenant/onedriveMetadata/resourceOwner/files",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := userDriveBackupHandler{
				baseUserDriveHandler: baseUserDriveHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
			}

			result, err := h.MetadataPathPrefix()
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
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
			h := userDriveBackupHandler{
				baseUserDriveHandler: baseUserDriveHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
			}
			p := path.Builder{}.Append("prefix")

			result, err := h.CanonicalPath(p)
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *ItemBackupHandlerUnitSuite) TestServiceCat() {
	t := suite.T()

	s, c := userDriveBackupHandler{}.ServiceCat()
	assert.Equal(t, path.OneDriveService, s)
	assert.Equal(t, path.FilesCategory, c)
}

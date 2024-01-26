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

type LibraryBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestLibraryBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &LibraryBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *LibraryBackupHandlerUnitSuite) TestPathPrefix() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "sharepoint",
			expect:    "tenant/sharepoint/resourceOwner/libraries/drives/driveID/root:",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := siteBackupHandler{
				baseSiteHandler: baseSiteHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
				service: path.SharePointService,
			}

			result, err := h.PathPrefix("driveID")
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
}

func (suite *LibraryBackupHandlerUnitSuite) TestMetadataPathPrefix() {
	tenantID, resourceOwner := "tenant", "resourceOwner"

	table := []struct {
		name      string
		expect    string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "sharepoint",
			expect:    "tenant/sharepointMetadata/resourceOwner/libraries",
			expectErr: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			h := siteBackupHandler{
				baseSiteHandler: baseSiteHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
				service: path.SharePointService,
			}

			result, err := h.MetadataPathPrefix()
			test.expectErr(t, err, clues.ToCore(err))

			if result != nil {
				assert.Equal(t, test.expect, result.String())
			}
		})
	}
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
			h := siteBackupHandler{
				baseSiteHandler: baseSiteHandler{
					qp: graph.QueryParams{
						ProtectedResource: idname.NewProvider(resourceOwner, resourceOwner),
						TenantID:          tenantID,
					},
				},
				service: path.SharePointService,
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

func (suite *LibraryBackupHandlerUnitSuite) TestServiceCat() {
	t := suite.T()

	s, c := siteBackupHandler{service: path.SharePointService}.ServiceCat()
	assert.Equal(t, path.SharePointService, s)
	assert.Equal(t, path.LibrariesCategory, c)
}

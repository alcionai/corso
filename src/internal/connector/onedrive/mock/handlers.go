package mock

import (
	"context"
	"net/http"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Backup Handler
// ---------------------------------------------------------------------------

type BackupHandler struct {
	ItemInfo details.ItemInfo

	GI  GetsItem
	GIP GetsItemPermission

	PathPrefixV   path.Path
	PathPrefixErr error

	CanonPath    path.Path
	CanonPathErr error

	Service  path.ServiceType
	Category path.CategoryType

	DrivePagerV api.DrivePager
	ItemPagerV  api.DriveItemEnumerator

	DisplayPath string

	LocationIDer details.LocationIDer

	getCall  int
	GetResps []*http.Response
	GetErrs  []error
}

func (h BackupHandler) PathPrefix(string, string, string) (path.Path, error) {
	return h.PathPrefixV, h.PathPrefixErr
}

func (h BackupHandler) CanonicalPath(*path.Builder, string, string) (path.Path, error) {
	return h.CanonPath, h.CanonPathErr
}

func (h BackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return h.Service, h.Category
}

func (h BackupHandler) DrivePager(string, []string) api.DrivePager {
	return h.DrivePagerV
}

func (h BackupHandler) ItemPager(string, string, []string) api.DriveItemEnumerator {
	return h.ItemPagerV
}

func (h BackupHandler) FormatDisplayPath(string, *path.Builder) string {
	return h.DisplayPath
}

func (h BackupHandler) NewLocationIDer(string, ...string) details.LocationIDer {
	return h.LocationIDer
}

func (h BackupHandler) AugmentItemInfo(details.ItemInfo, models.DriveItemable, int64, *path.Builder) details.ItemInfo {
	return h.ItemInfo
}

func (h *BackupHandler) Get(context.Context, string, map[string]string) (*http.Response, error) {
	c := h.getCall
	h.getCall++

	return h.GetResps[c], h.GetErrs[c]
}

func (h BackupHandler) GetItem(context.Context, string, string) (models.DriveItemable, error) {
	return h.GI.GetItem(nil, "", "")
}

func (h BackupHandler) GetItemPermission(
	context.Context,
	string, string,
) (models.PermissionCollectionResponseable, error) {
	return h.GIP.GetItemPermission(nil, "", "")
}

// ---------------------------------------------------------------------------
// Get Itemer
// ---------------------------------------------------------------------------

type GetsItem struct {
	Item models.DriveItemable
	Err  error
}

func (m GetsItem) GetItem(
	_ context.Context,
	_, _ string,
) (models.DriveItemable, error) {
	return m.Item, m.Err
}

// ---------------------------------------------------------------------------
// Get Item Permissioner
// ---------------------------------------------------------------------------

type GetsItemPermission struct {
	Perm models.PermissionCollectionResponseable
	Err  error
}

func (m GetsItemPermission) GetItemPermission(
	_ context.Context,
	_, _ string,
) (models.PermissionCollectionResponseable, error) {
	return m.Perm, m.Err
}

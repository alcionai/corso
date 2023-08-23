package drive

import (
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ BackupHandler = &groupBackupHandler{}

type groupBackupHandler struct {
	libraryBackupHandler
	groupID string
	scope   selectors.GroupsScope
}

func NewGroupBackupHandler(groupID string, ac api.Drives, scope selectors.GroupsScope) groupBackupHandler {
	return groupBackupHandler{
		libraryBackupHandler{
			ac: ac,
			// Not adding scope here. Anything that needs scope has to
			// be from group handler
			service: path.GroupsService,
		},
		groupID,
		scope,
	}
}

func (h groupBackupHandler) CanonicalPath(
	folders *path.Builder,
	tenantID, resourceOwner string,
) (path.Path, error) {
	// TODO(meain): path fixes
	return folders.ToDataLayerPath(tenantID, h.groupID, h.service, path.LibrariesCategory, false)
}

func (h groupBackupHandler) ServiceCat() (path.ServiceType, path.CategoryType) {
	return path.GroupsService, path.LibrariesCategory
}

func (h groupBackupHandler) IsAllPass() bool {
	return h.scope.IsAny(selectors.GroupsLibraryFolder)
}

func (h groupBackupHandler) IncludesDir(dir string) bool {
	return h.scope.Matches(selectors.GroupsLibraryFolder, dir)
}

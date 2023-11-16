package exchange

import (
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ backupHandler = &eventBackupHandler{}

type eventBackupHandler struct {
	ac api.Events
}

func newEventBackupHandler(
	ac api.Client,
) eventBackupHandler {
	ace := ac.Events()

	return eventBackupHandler{
		ac: ace,
	}
}

func (h eventBackupHandler) itemEnumerator() addedAndRemovedItemGetter {
	return h.ac
}

func (h eventBackupHandler) itemHandler() itemGetterSerializer {
	return h.ac
}

func (h eventBackupHandler) folderGetter() containerGetter {
	return h.ac
}

func (h eventBackupHandler) previewIncludeContainers() []string {
	return []string{
		"calendar",
	}
}

func (h eventBackupHandler) previewExcludeContainers() []string {
	return nil
}

func (h eventBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return api.DefaultCalendar, &eventContainerCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

package exchange

import (
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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

func (h eventBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return DefaultCalendar, &eventCalendarCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

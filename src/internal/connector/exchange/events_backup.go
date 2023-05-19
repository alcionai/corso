package exchange

import (
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &eventBackupHandler{}

type eventBackupHandler struct {
	ac api.Mail
}

func newEventBackupHandler(
	ac api.Client,
) eventBackupHandler {
	acm := ac.Mail()

	return eventBackupHandler{
		ac: acm,
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

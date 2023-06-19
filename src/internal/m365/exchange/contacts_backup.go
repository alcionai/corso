package exchange

import (
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &contactBackupHandler{}

type contactBackupHandler struct {
	ac api.Contacts
}

func newContactBackupHandler(
	ac api.Client,
) contactBackupHandler {
	acc := ac.Contacts()

	return contactBackupHandler{
		ac: acc,
	}
}

func (h contactBackupHandler) itemEnumerator() addedAndRemovedItemGetter {
	return h.ac
}

func (h contactBackupHandler) itemHandler() itemGetterSerializer {
	return h.ac
}

func (h contactBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return DefaultContactFolder, &contactContainerCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

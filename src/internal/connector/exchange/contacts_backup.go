package exchange

import (
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &contactBackupHandler{}

type contactBackupHandler struct {
	ac api.Mail
}

func newContactBackupHandler(
	ac api.Client,
) contactBackupHandler {
	acm := ac.Mail()

	return contactBackupHandler{
		ac: acm,
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
	return DefaultContactFolder, &contactFolderCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

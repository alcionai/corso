package exchange

import (
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &mailBackupHandler{}

type mailBackupHandler struct {
	ac api.Mail
}

func newMailBackupHandler(
	ac api.Client,
) mailBackupHandler {
	acm := ac.Mail()

	return mailBackupHandler{
		ac: acm,
	}
}

func (h mailBackupHandler) itemEnumerator() addedAndRemovedItemGetter {
	return h.ac
}

func (h mailBackupHandler) itemHandler() api.GetAndSerializeItemer[details.ExchangeInfo] {
	return h.ac
}

func (h mailBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return api.MsgFolderRoot, &mailContainerCache{
		userID: userID,
		enumer: h.ac,
		getter: h.ac,
	}
}

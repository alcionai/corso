package exchange

import (
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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

func (h mailBackupHandler) itemHandler() itemGetterSerializer {
	return h.ac
}

func (h mailBackupHandler) folderGetter() containerGetter {
	return h.ac
}

func (h mailBackupHandler) previewIncludeContainers() []string {
	return []string{
		"inbox",
	}
}

func (h mailBackupHandler) previewExcludeContainers() []string {
	return []string{
		"drafts",
		"outbox",
		"recoverableitemsdeletions",
		"junkemail",
	}
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

func (h mailBackupHandler) CanSkipItemFailure(
	err error,
	resourceID string,
	opts control.Options,
) (fault.SkipCause, bool) {
	return "", false
}

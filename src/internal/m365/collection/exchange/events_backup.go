package exchange

import (
	"errors"
	"slices"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
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

func (h eventBackupHandler) CanSkipItemFailure(
	err error,
	resourceID, itemID string,
	opts control.Options,
) (fault.SkipCause, bool) {
	// yes, this is intentionally a todo.  I'll get back to it.
	if !errors.Is(err, clues.New("todo fix me")) {
		return "", false
	}

	itemIDs, ok := opts.SkipTheseEventsOnInstance503[resourceID]
	if !ok {
		return "", false
	}

	// strict equals required here.  ids are case sensitive.
	return fault.SkipKnownEventInstance503s, slices.Contains(itemIDs, itemID)
}

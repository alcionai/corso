package m365

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewServiceHandler returns an instance of a struct capable of running various
// operations for a given service.
func (ctrl *Controller) NewServiceHandler(
	opts control.Options,
	service path.ServiceType,
) (inject.ServiceHandler, error) {
	ctrl.setResourceHandler(service)

	switch service {
	case path.OneDriveService:
		return onedrive.NewOneDriveHandler(opts), nil

	case path.SharePointService:
		return sharepoint.NewSharePointHandler(opts), nil

	case path.GroupsService:
		return groups.NewGroupsHandler(opts), nil
	}

	return nil, clues.New("unrecognized service").
		With("service_type", service.String())
}

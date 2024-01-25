package m365

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/groups"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/m365/service/teamschats"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewServiceHandler returns an instance of a struct capable of running various
// operations for a given service.
func (ctrl *Controller) NewServiceHandler(
	service path.ServiceType,
) (inject.ServiceHandler, error) {
	ctrl.setResourceHandler(service)

	switch service {
	case path.OneDriveService:
		return onedrive.NewOneDriveHandler(ctrl.AC, ctrl.resourceHandler), nil

	case path.SharePointService:
		return sharepoint.NewSharePointHandler(ctrl.AC, ctrl.resourceHandler), nil

	case path.GroupsService:
		return groups.NewGroupsHandler(ctrl.AC, ctrl.resourceHandler), nil

	case path.ExchangeService:
		return exchange.NewExchangeHandler(ctrl.AC, ctrl.resourceHandler), nil

	case path.TeamsChatsService:
		return teamschats.NewTeamsChatsHandler(ctrl.AC, ctrl.resourceHandler), nil
	}

	return nil, clues.New("unrecognized service").
		With("service_type", service.String())
}

package path

import (
	"strings"

	"github.com/alcionai/clues"
)

var ErrorUnknownService = clues.New("unknown service string")

// ServiceType denotes what service the path corresponds to. Metadata services
// are also included though they are only used for paths that house metadata for
// Corso backups.
//
// Metadata services are not considered valid service types for resource paths
// though they can be used for metadata paths.
//
// The string representaton of each enum _must remain the same_.  In case of
// changes to those values, we'll need migration code to handle transitions
// across states else we'll get marshalling/unmarshalling errors.
type ServiceType int

//go:generate stringer -type=ServiceType -linecomment
const (
	UnknownService            ServiceType = 0
	ExchangeService           ServiceType = 1 // exchange
	OneDriveService           ServiceType = 2 // onedrive
	SharePointService         ServiceType = 3 // sharepoint
	ExchangeMetadataService   ServiceType = 4 // exchangeMetadata
	OneDriveMetadataService   ServiceType = 5 // onedriveMetadata
	SharePointMetadataService ServiceType = 6 // sharepointMetadata
	GroupsService             ServiceType = 7 // groups
	GroupsMetadataService     ServiceType = 8 // groupsMetadata
)

func ToServiceType(service string) ServiceType {
	s := strings.ToLower(service)

	switch s {
	case strings.ToLower(ExchangeService.String()):
		return ExchangeService
	case strings.ToLower(OneDriveService.String()):
		return OneDriveService
	case strings.ToLower(SharePointService.String()):
		return SharePointService
	case strings.ToLower(GroupsService.String()):
		return GroupsService
	case strings.ToLower(ExchangeMetadataService.String()):
		return ExchangeMetadataService
	case strings.ToLower(OneDriveMetadataService.String()):
		return OneDriveMetadataService
	case strings.ToLower(SharePointMetadataService.String()):
		return SharePointMetadataService
	case strings.ToLower(GroupsMetadataService.String()):
		return GroupsMetadataService
	default:
		return UnknownService
	}
}

var serviceToHuman = map[ServiceType]string{
	ExchangeService:   "Exchange",
	OneDriveService:   "OneDrive",
	SharePointService: "SharePoint",
	GroupsService:     "Groups",
}

// HumanString produces a more human-readable string version of the service.
func (svc ServiceType) HumanString() string {
	hs, ok := serviceToHuman[svc]
	if ok {
		return hs
	}

	return "Unknown Service"
}

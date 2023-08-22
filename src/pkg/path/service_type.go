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
// The order of the enums below can be changed, but the string representation of
// each enum must remain the same or migration code needs to be added to handle
// changes to the string format.
type ServiceType int

//go:generate stringer -type=ServiceType -linecomment
const (
	UnknownService            ServiceType = 0
	ExchangeService           ServiceType = 1  // exchange
	OneDriveService           ServiceType = 2  // onedrive
	SharePointService         ServiceType = 3  // sharepoint
	ExchangeMetadataService   ServiceType = 4  // exchangeMetadata
	OneDriveMetadataService   ServiceType = 5  // onedriveMetadata
	SharePointMetadataService ServiceType = 6  // sharepointMetadata
	GroupsService             ServiceType = 7  // groups
	GroupsMetadataService     ServiceType = 8  // groupsMetadata
	TeamsService              ServiceType = 9  // teams
	TeamsMetadataService      ServiceType = 10 // teamsMetadata
)

func toServiceType(service string) ServiceType {
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
	case strings.ToLower(TeamsService.String()):
		return TeamsService
	case strings.ToLower(ExchangeMetadataService.String()):
		return ExchangeMetadataService
	case strings.ToLower(OneDriveMetadataService.String()):
		return OneDriveMetadataService
	case strings.ToLower(SharePointMetadataService.String()):
		return SharePointMetadataService
	case strings.ToLower(GroupsMetadataService.String()):
		return GroupsMetadataService
	case strings.ToLower(TeamsMetadataService.String()):
		return TeamsMetadataService
	default:
		return UnknownService
	}
}

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
	UnknownService            ServiceType = iota
	ExchangeService                       // exchange
	OneDriveService                       // onedrive
	SharePointService                     // sharepoint
	ExchangeMetadataService               // exchangeMetadata
	OneDriveMetadataService               // onedriveMetadata
	SharePointMetadataService             // sharepointMetadata
	GroupsService                         // groups
	GroupsMetadataService                 // groupsMetadata
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
	case strings.ToLower(ExchangeMetadataService.String()):
		return ExchangeMetadataService
	case strings.ToLower(OneDriveMetadataService.String()):
		return OneDriveMetadataService
	case strings.ToLower(SharePointMetadataService.String()):
		return SharePointMetadataService
	default:
		return UnknownService
	}
}

// subServices is a mapping of all valid service/subService pairs.
// a subService pair occurs when one service contains a reference
// to a protected resource of another service type, and the resource
// for that second service is the identifier which is used to discover
// data.  A subService relationship may imply that the subservice data
// is wholly replicated/owned by the primary service, or it may not,
// each case differs.
//
// Ex:
//   - groups/<gID>/sharepoint/<siteID> => each team in groups contains a
//     complete sharepoint site.
//   - groups/<gID>/member/<userID> => each user in a team can own one or
//     more Chats.  But the group does not contain the complete user data.
var subServices = map[ServiceType]map[ServiceType]struct{}{
	GroupsService: {
		SharePointService: {},
	},
}

func ValidateServiceAndSubService(service, subService ServiceType) error {
	subs, ok := subServices[service]
	if !ok {
		return clues.New("unsupported service").With("service", service)
	}

	if _, ok := subs[subService]; !ok {
		return clues.New("unknown service/subService combination").
			With("service", service, "subService", subService)
	}

	return nil
}

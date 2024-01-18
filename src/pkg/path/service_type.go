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
	ExchangeService           ServiceType = 1  // exchange
	OneDriveService           ServiceType = 2  // onedrive
	SharePointService         ServiceType = 3  // sharepoint
	ExchangeMetadataService   ServiceType = 4  // exchangeMetadata
	OneDriveMetadataService   ServiceType = 5  // onedriveMetadata
	SharePointMetadataService ServiceType = 6  // sharepointMetadata
	GroupsService             ServiceType = 7  // groups
	GroupsMetadataService     ServiceType = 8  // groupsMetadata
	TeamsChatsService         ServiceType = 9  // teamsChats
	TeamsChatsMetadataService ServiceType = 10 // teamsChatsMetadata
)

var strToSvc = map[string]ServiceType{
	ExchangeService.String():           ExchangeService,
	ExchangeMetadataService.String():   ExchangeMetadataService,
	OneDriveService.String():           OneDriveService,
	OneDriveMetadataService.String():   OneDriveMetadataService,
	SharePointService.String():         SharePointService,
	SharePointMetadataService.String(): SharePointMetadataService,
	GroupsService.String():             GroupsService,
	GroupsMetadataService.String():     GroupsMetadataService,
	TeamsChatsService.String():         TeamsChatsService,
	TeamsChatsMetadataService.String(): TeamsChatsMetadataService,
}

func ToServiceType(service string) ServiceType {
	st, ok := strToSvc[strings.ToLower(service)]
	if !ok {
		st = UnknownService
	}

	return st
}

var serviceToHuman = map[ServiceType]string{
	ExchangeService:   "Exchange",
	OneDriveService:   "OneDrive",
	SharePointService: "SharePoint",
	GroupsService:     "Groups",
	TeamsChatsService: "Chats",
}

// HumanString produces a more human-readable string version of the service.
func (svc ServiceType) HumanString() string {
	hs, ok := serviceToHuman[svc]
	if ok {
		return hs
	}

	return "Unknown Service"
}

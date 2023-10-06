package inject

import (
	"context"
	"strconv"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// RestoreConsumerConfig is a container-of-things for holding options and
// configurations from various packages, all of which are widely used by
// restore consumers independent of service or data category.
type RestoreConsumerConfig struct {
	BackupVersion     int
	Options           control.Options
	ProtectedResource idname.Provider
	RestoreConfig     control.RestoreConfig
	Selector          selectors.Selector
}

// BackupProducerConfig is a container-of-things for holding options and
// configurations from various packages, all of which are widely used by
// backup producers independent of service or data category.
type BackupProducerConfig struct {
	LastBackupVersion   int
	MetadataCollections []data.RestoreCollection
	Options             control.Options
	ProtectedResource   idname.Provider
	Selector            selectors.Selector
}

type BackupProducerResults struct {
	Collections          []data.BackupCollection
	Excludes             prefixmatcher.StringSetReader
	CanUsePreviousBackup bool
	DiscoveredItems      Stats
}

// Stats is a oneOf that contains service specific
// information
type Stats struct {
	Exchange   *ExchangeStats   `json:"exchange,omitempty"`
	SharePoint *SharePointStats `json:"sharePoint,omitempty"`
	OneDrive   *OneDriveStats   `json:"oneDrive,omitempty"`
	Groups     *GroupsStats     `json:"groups,omitempty"`
}

type ExchangeStats struct {
	ContactsAdded   int `json:"contactsAdded,omitempty"`
	ContactsDeleted int `json:"contactsDeleted,omitempty"`
	ContactFolders  int `json:"contactFolders,omitempty"`

	EventsAdded   int `json:"eventsAdded,omitempty"`
	EventsDeleted int `json:"eventsDeleted,omitempty"`
	EventFolders  int `json:"eventFolders,omitempty"`

	EmailsAdded   int `json:"emailsAdded,omitempty"`
	EmailsDeleted int `json:"emailsDeleted,omitempty"`
	EmailFolders  int `json:"emailFolders,omitempty"`
}

type SharePointStats struct {
}
type OneDriveStats struct {
	Folders int `json:"folders,omitempty"`
	Items   int `json:"items,omitempty"`
}
type GroupsStats struct {
}

// interface compliance checks
var _ print.Printable = &Stats{}

// Print writes the Backup to StdOut, in the format requested by the caller.
func (s Stats) Print(ctx context.Context) {
	print.Item(ctx, s)
}

// MinimumPrintable reduces the Backup to its minimally printable details.
func (s Stats) MinimumPrintable() any {
	return s
}

// Headers returns the human-readable names of properties in a Backup
// for printing out to a terminal in a columnar display.
func (s Stats) Headers() []string {
	switch {
	case s.OneDrive != nil:
		return []string{
			"Folders",
			"Items",
		}

	default:
		return []string{}
	}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (s Stats) Values() []string {
	switch {
	case s.OneDrive != nil:
		return []string{
			strconv.Itoa(s.OneDrive.Folders),
			strconv.Itoa(s.OneDrive.Items),
		}

	default:
		return []string{}

	}
}

package details

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/alcionai/corso/cli/print"
	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/model"
)

// --------------------------------------------------------------------------------
// Model
// --------------------------------------------------------------------------------

// DetailsModel describes what was stored in a Backup
type DetailsModel struct {
	model.BaseModel
	Entries []DetailsEntry `json:"entries"`
}

// Print writes the DetailModel Entries to StdOut, in the format
// requested by the caller.
func (dm DetailsModel) PrintEntries(ctx context.Context) {
	perType := map[string][]print.Printable{}

	for _, de := range dm.Entries {
		it := de.infoType()
		ps, ok := perType[it]

		if !ok {
			ps = []print.Printable{}
		}

		perType[it] = append(ps, print.Printable(de))
	}

	for _, ps := range perType {
		print.All(ctx, ps...)
	}
}

// Paths returns the list of Paths extracted from the Entries slice.
func (dm DetailsModel) Paths() []string {
	ents := dm.Entries
	r := make([]string, len(ents))

	for i := range ents {
		r[i] = ents[i].RepoRef
	}

	return r
}

// --------------------------------------------------------------------------------
// Details
// --------------------------------------------------------------------------------

// Details augments the core with a mutex for processing.
// Should be sliced back to d.DetailsModel for storage and
// printing.
type Details struct {
	DetailsModel

	// internal
	mu sync.Mutex `json:"-"`
}

func (d *Details) Add(repoRef string, info ItemInfo) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Entries = append(d.Entries, DetailsEntry{RepoRef: repoRef, ItemInfo: info})
}

// --------------------------------------------------------------------------------
// Entry
// --------------------------------------------------------------------------------

// DetailsEntry describes a single item stored in a Backup
type DetailsEntry struct {
	// TODO: `RepoRef` is currently the full path to the item in Kopia
	// This can be optimized.
	RepoRef string `json:"repoRef"`
	ItemInfo
}

// --------------------------------------------------------------------------------
// CLI Output
// --------------------------------------------------------------------------------

// interface compliance checks
var _ print.Printable = &DetailsEntry{}

// MinimumPrintable DetailsEntries is a passthrough func, because no
// reduction is needed for the json output.
func (de DetailsEntry) MinimumPrintable() any {
	return de
}

// Headers returns the human-readable names of properties in a DetailsEntry
// for printing out to a terminal in a columnar display.
func (de DetailsEntry) Headers() []string {
	hs := []string{"Repo Ref"}

	if de.ItemInfo.Exchange != nil {
		hs = append(hs, de.ItemInfo.Exchange.Headers()...)
	}

	if de.ItemInfo.Sharepoint != nil {
		hs = append(hs, de.ItemInfo.Sharepoint.Headers()...)
	}

	if de.ItemInfo.OneDrive != nil {
		hs = append(hs, de.ItemInfo.OneDrive.Headers()...)
	}

	return hs
}

// Values returns the values matching the Headers list.
func (de DetailsEntry) Values() []string {
	vs := []string{de.RepoRef}

	if de.ItemInfo.Exchange != nil {
		vs = append(vs, de.ItemInfo.Exchange.Values()...)
	}

	if de.ItemInfo.Sharepoint != nil {
		vs = append(vs, de.ItemInfo.Sharepoint.Values()...)
	}

	if de.ItemInfo.OneDrive != nil {
		vs = append(vs, de.ItemInfo.OneDrive.Values()...)
	}

	return vs
}

// ItemInfo is a oneOf that contains service specific
// information about the item it tracks
type ItemInfo struct {
	Exchange   *ExchangeInfo   `json:"exchange,omitempty"`
	Sharepoint *SharepointInfo `json:"sharepoint,omitempty"`
	OneDrive   *OneDriveInfo   `json:"oneDrive,omitempty"`
}

// infoType provides internal categorization for collecting like-typed ItemInfos.
// It should return the most granular value type (ex: "event" for an exchange
// calendar event).
func (i ItemInfo) infoType() string {
	switch {
	case i.Exchange != nil:
		return i.Exchange.infoType()

	case i.Sharepoint != nil:
		return i.Sharepoint.infoType()

	case i.OneDrive != nil:
		return i.OneDrive.infoType()
	}

	return ""
}

// ExchangeInfo describes an exchange item
type ExchangeInfo struct {
	Sender      string    `json:"sender,omitempty"`
	Subject     string    `json:"subject,omitempty"`
	Received    time.Time `json:"received,omitempty"`
	EventStart  time.Time `json:"eventStart,omitempty"`
	Organizer   string    `json:"organizer,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
	EventRecurs bool      `json:"eventRecurs,omitempty"`
}

func (i ExchangeInfo) infoType() string {
	hasContactName := len(i.ContactName) > 0
	hasReceived := !(&i.Received).Equal(time.Time{})
	hasEventStart := !(&i.EventStart).Equal(time.Time{})

	switch {
	case !hasContactName && !hasReceived:
		return "event"

	case !hasContactName && !hasEventStart:
		return "mail"

	case !hasEventStart && !hasReceived:
		return "contact"
	}

	return ""
}

// Headers returns the human-readable names of properties in an ExchangeInfo
// for printing out to a terminal in a columnar display.
func (i ExchangeInfo) Headers() []string {
	switch i.infoType() {
	case "event":
		return []string{"Organizer", "Subject", "Starts", "Recurring"}

	case "contact":
		return []string{"Contact Name"}

	case "mail":
		return []string{"Sender", "Subject", "Received"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i ExchangeInfo) Values() []string {
	switch i.infoType() {
	case "event":
		return []string{i.Organizer, i.Subject, common.FormatTime(i.EventStart), strconv.FormatBool(i.EventRecurs)}

	case "contact":
		return []string{i.ContactName}

	case "mail":
		return []string{i.Sender, i.Subject, common.FormatTime(i.Received)}
	}

	return []string{}
}

// SharepointInfo describes a sharepoint item
// TODO: Implement this. This is currently here
// just to illustrate usage
type SharepointInfo struct{}

func (i SharepointInfo) infoType() string {
	return "sharepoint"
}

// Headers returns the human-readable names of properties in a SharepointInfo
// for printing out to a terminal in a columnar display.
func (i SharepointInfo) Headers() []string {
	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i SharepointInfo) Values() []string {
	return []string{}
}

// OneDriveInfo describes a oneDrive item
type OneDriveInfo struct {
	ParentPath string `json:"parentPath"`
	ItemName   string `json:"itemName"`
}

// Headers returns the human-readable names of properties in a OneDriveInfo
// for printing out to a terminal in a columnar display.
func (i OneDriveInfo) Headers() []string {
	return []string{"ItemName", "ParentPath"}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i OneDriveInfo) Values() []string {
	return []string{i.ItemName, i.ParentPath}
}

func (i OneDriveInfo) infoType() string {
	return "item"
}

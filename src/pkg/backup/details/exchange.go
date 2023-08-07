package details

import (
	"strconv"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewExchangeLocationIDer builds a LocationIDer for the given category and
// folder path. The path denoted by the folders should be unique within the
// category.
func NewExchangeLocationIDer(
	category path.CategoryType,
	escapedFolders ...string,
) (uniqueLoc, error) {
	if err := path.ValidateServiceAndCategory(path.ExchangeService, category); err != nil {
		return uniqueLoc{}, clues.Wrap(err, "making exchange LocationIDer")
	}

	pb := path.Builder{}.Append(category.String()).Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: 1,
	}, nil
}

// ExchangeInfo describes an exchange item
type ExchangeInfo struct {
	ItemType    ItemType  `json:"itemType,omitempty"`
	Sender      string    `json:"sender,omitempty"`
	Subject     string    `json:"subject,omitempty"`
	Recipient   []string  `json:"recipient,omitempty"`
	ParentPath  string    `json:"parentPath,omitempty"`
	Received    time.Time `json:"received,omitempty"`
	EventStart  time.Time `json:"eventStart,omitempty"`
	EventEnd    time.Time `json:"eventEnd,omitempty"`
	Organizer   string    `json:"organizer,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
	EventRecurs bool      `json:"eventRecurs,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Modified    time.Time `json:"modified,omitempty"`
	Size        int64     `json:"size,omitempty"`
}

// Headers returns the human-readable names of properties in an ExchangeInfo
// for printing out to a terminal in a columnar display.
func (i ExchangeInfo) Headers() []string {
	switch i.ItemType {
	case ExchangeEvent:
		return []string{"Organizer", "Subject", "Starts", "Ends", "Recurring"}

	case ExchangeContact:
		return []string{"Contact Name"}

	case ExchangeMail:
		return []string{"Sender", "Folder", "Subject", "Received"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i ExchangeInfo) Values() []string {
	switch i.ItemType {
	case ExchangeEvent:
		return []string{
			i.Organizer,
			i.Subject,
			dttm.FormatToTabularDisplay(i.EventStart),
			dttm.FormatToTabularDisplay(i.EventEnd),
			strconv.FormatBool(i.EventRecurs),
		}

	case ExchangeContact:
		return []string{i.ContactName}

	case ExchangeMail:
		return []string{
			i.Sender, i.ParentPath, i.Subject,
			dttm.FormatToTabularDisplay(i.Received),
		}
	}

	return []string{}
}

func (i *ExchangeInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.String()
}

func (i *ExchangeInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	var category path.CategoryType

	switch i.ItemType {
	case ExchangeEvent:
		category = path.EventsCategory
	case ExchangeContact:
		category = path.ContactsCategory
	case ExchangeMail:
		category = path.EmailCategory
	}

	loc, err := NewExchangeLocationIDer(category, baseLoc.Elements()...)

	return &loc, err
}

func (i *ExchangeInfo) updateFolder(f *FolderInfo) error {
	// Use a switch instead of a rather large if-statement. Just make sure it's an
	// Exchange type. If it's not return an error.
	switch i.ItemType {
	case ExchangeContact, ExchangeEvent, ExchangeMail:
	default:
		return clues.New("unsupported non-Exchange ItemType").
			With("item_type", i.ItemType)
	}

	f.DataType = i.ItemType

	return nil
}

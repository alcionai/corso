package path

import (
	"fmt"
	"strings"

	"github.com/alcionai/clues"
)

var ErrorUnknownCategory = clues.New("unknown category string")

// CategoryType denotes what category of data the path corresponds to. The order
// of the enums below can be changed, but the string representation of each enum
// must remain the same or migration code needs to be added to handle changes to
// the string format.
type CategoryType int

//go:generate stringer -type=CategoryType -linecomment
const (
	UnknownCategory         CategoryType = 0
	EmailCategory           CategoryType = 1 // email
	ContactsCategory        CategoryType = 2 // contacts
	EventsCategory          CategoryType = 3 // events
	FilesCategory           CategoryType = 4 // files
	ListsCategory           CategoryType = 5 // lists
	LibrariesCategory       CategoryType = 6 // libraries
	PagesCategory           CategoryType = 7 // pages
	DetailsCategory         CategoryType = 8 // details
	ChannelMessagesCategory CategoryType = 9 // channelMessages
)

var strToCat = map[string]CategoryType{
	strings.ToLower(EmailCategory.String()):           EmailCategory,
	strings.ToLower(ContactsCategory.String()):        ContactsCategory,
	strings.ToLower(EventsCategory.String()):          EventsCategory,
	strings.ToLower(FilesCategory.String()):           FilesCategory,
	strings.ToLower(LibrariesCategory.String()):       LibrariesCategory,
	strings.ToLower(ListsCategory.String()):           ListsCategory,
	strings.ToLower(PagesCategory.String()):           PagesCategory,
	strings.ToLower(DetailsCategory.String()):         DetailsCategory,
	strings.ToLower(ChannelMessagesCategory.String()): ChannelMessagesCategory,
}

func ToCategoryType(s string) CategoryType {
	cat, ok := strToCat[strings.ToLower(s)]
	if ok {
		return cat
	}

	return UnknownCategory
}

var catToHuman = map[CategoryType]string{
	EmailCategory:           "Emails",
	ContactsCategory:        "Contacts",
	EventsCategory:          "Events",
	FilesCategory:           "Files",
	LibrariesCategory:       "Libraries",
	ListsCategory:           "Lists",
	PagesCategory:           "Pages",
	DetailsCategory:         "Details",
	ChannelMessagesCategory: "Messages",
}

// HumanString produces a more human-readable string version of the category.
func (cat CategoryType) HumanString() string {
	hs, ok := catToHuman[cat]
	if ok {
		return hs
	}

	return "Unknown Category"
}

// ---------------------------------------------------------------------------
// Service-Category pairings
// ---------------------------------------------------------------------------

// serviceCategories is a mapping of all valid service/category pairs for
// non-metadata paths.
var serviceCategories = map[ServiceType]map[CategoryType]struct{}{
	ExchangeService: {
		EmailCategory:    {},
		ContactsCategory: {},
		EventsCategory:   {},
	},
	OneDriveService: {
		FilesCategory: {},
	},
	SharePointService: {
		LibrariesCategory: {},
		ListsCategory:     {},
		PagesCategory:     {},
	},
	GroupsService: {
		ChannelMessagesCategory: {},
		LibrariesCategory:       {},
	},
}

func validateServiceAndCategoryStrings(s, c string) (ServiceType, CategoryType, error) {
	service := ToServiceType(s)
	if service == UnknownService {
		return UnknownService, UnknownCategory, clues.Stack(ErrorUnknownService).With("service", fmt.Sprintf("%q", s))
	}

	category := ToCategoryType(c)
	if category == UnknownCategory {
		return UnknownService, UnknownCategory, clues.Stack(ErrorUnknownService).With("category", fmt.Sprintf("%q", c))
	}

	if err := ValidateServiceAndCategory(service, category); err != nil {
		return UnknownService, UnknownCategory, err
	}

	return service, category, nil
}

func ValidateServiceAndCategory(service ServiceType, category CategoryType) error {
	cats, ok := serviceCategories[service]
	if !ok {
		return clues.New("unsupported service").With("service", fmt.Sprintf("%q", service))
	}

	if _, ok := cats[category]; !ok {
		return clues.New("unknown service/category combination").
			With("service", fmt.Sprintf("%q", service), "category", fmt.Sprintf("%q", category))
	}

	return nil
}

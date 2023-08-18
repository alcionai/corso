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
	UnknownCategory   CategoryType = 0
	EmailCategory     CategoryType = 1 // email
	ContactsCategory  CategoryType = 2 // contacts
	EventsCategory    CategoryType = 3 // events
	FilesCategory     CategoryType = 4 // files
	ListsCategory     CategoryType = 5 // lists
	LibrariesCategory CategoryType = 6 // libraries
	PagesCategory     CategoryType = 7 // pages
	DetailsCategory   CategoryType = 8 // details
)

func ToCategoryType(category string) CategoryType {
	cat := strings.ToLower(category)

	switch cat {
	case strings.ToLower(EmailCategory.String()):
		return EmailCategory
	case strings.ToLower(ContactsCategory.String()):
		return ContactsCategory
	case strings.ToLower(EventsCategory.String()):
		return EventsCategory
	case strings.ToLower(FilesCategory.String()):
		return FilesCategory
	case strings.ToLower(LibrariesCategory.String()):
		return LibrariesCategory
	case strings.ToLower(ListsCategory.String()):
		return ListsCategory
	case strings.ToLower(PagesCategory.String()):
		return PagesCategory
	case strings.ToLower(DetailsCategory.String()):
		return DetailsCategory
	default:
		return UnknownCategory
	}
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
		LibrariesCategory: {}, // TODO(meain)
	},
}

func validateServiceAndCategoryStrings(s, c string) (ServiceType, CategoryType, error) {
	service := toServiceType(s)
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

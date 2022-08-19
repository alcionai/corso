package selectors

// ---------------------------------------------------------------------------
// Selectors
// ---------------------------------------------------------------------------

type (
	// onedrive provides an api for selecting
	// data scopes applicable to the OneDrive service.
	onedrive struct {
		Selector
	}

	// OneDriveBackup provides an api for selecting
	// data scopes applicable to the OneDrive service,
	// plus backup-specific methods.
	OneDriveBackup struct {
		onedrive
	}
)

// NewOneDriveBackup produces a new Selector with the service set to ServiceOneDrive.
func NewOneDriveBackup() *OneDriveBackup {
	src := OneDriveBackup{
		onedrive{
			newSelector(ServiceOneDrive),
		},
	}
	return &src
}

// ToOneDriveBackup transforms the generic selector into an OneDriveBackup.
// Errors if the service defined by the selector is not ServiceOneDrive.
func (s Selector) ToOneDriveBackup() (*OneDriveBackup, error) {
	if s.Service != ServiceOneDrive {
		return nil, badCastErr(ServiceOneDrive, s.Service)
	}
	src := OneDriveBackup{onedrive{s}}
	return &src, nil
}

// ---------------------------------------------------------------------------
// Scopes
// ---------------------------------------------------------------------------

type (
	// OneDriveScope specifies the data available
	// when interfacing with the OneDrive service.
	OneDriveScope scope
	// onedriveCategory enumerates the type of the lowest level
	// of data () in a scope.
	onedriveCategory int
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=onedriveCategory
const (
	OneDriveCategoryUnknown onedriveCategory = iota
	// types of data identified by OneDrive
	OneDriveUser
)

// Scopes retrieves the list of onedriveScopes in the selector.
func (s *onedrive) Scopes() []OneDriveScope {
	scopes := []OneDriveScope{}
	for _, v := range s.Includes {
		scopes = append(scopes, OneDriveScope(v))
	}
	return scopes
}

// Get returns the data category in the scope.  If the scope
// contains all data types for a user, it'll return the
// OneDriveUser category.
func (s OneDriveScope) Get(cat onedriveCategory) []string {
	v, ok := s[cat.String()]
	if !ok {
		return None()
	}
	return split(v)
}

// Produces one or more onedrive user scopes.
// One scope is created per user entry.
// If any slice contains selectors.Any, that slice is reduced to [selectors.Any]
// If any slice contains selectors.None, that slice is reduced to [selectors.None]
// If any slice is empty, it defaults to [selectors.None]
func (s *onedrive) Users(users []string) []OneDriveScope {
	users = normalize(users)
	scopes := []OneDriveScope{}
	for _, u := range users {
		userScope := OneDriveScope{
			OneDriveUser.String(): u,
		}
		scopes = append(scopes, userScope)
	}
	return scopes
}

// nop-transform method
func nopTransform(sl []OneDriveScope) []OneDriveScope { return sl }

func (s *onedrive) Include(scopes ...[]OneDriveScope) {
	// appendIncludes(&s.Selector, nopTransform, scopes...)
}

func (s *onedrive) Exclude(scopes ...[]OneDriveScope) {
	// appendExcludes(&s.Selector, nopTransform, scopes...)
}

func (s *onedrive) Filter(scopes ...[]OneDriveScope) {
	// appendFilters(&s.Selector, nopTransform, scopes...)
}

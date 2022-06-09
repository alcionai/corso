package source

type applicationType int

//go:generate stringer -type=storageProvider -linecomment
const (
	AppUnknown  applicationType = iota // Unknown Application
	AppExchange                        // Exchange
)

type Source struct {
	App     applicationType
	userIDs []string
}

// Creates a new source for the given application.
func NewSource(app applicationType) *Source {
	return &Source{
		App: app,
	}
}

// AddUsers adds the provide user IDs to the source.
// Data retrieval will be scoped to include the unioned
// set of specified users.
func (s *Source) AddUsers(uids ...string) error {
	// future todo: not all applications support users.
	// this should error when users are added in that context.
	if s.userIDs == nil {
		s.userIDs = []string{}
	}
	s.userIDs = append(s.userIDs, uids...)
	return nil
}

// Users retrieves the userIDs specified by the caller.
func (s *Source) Users() []string {
	if s.userIDs == nil {
		return []string{}
	}
	return s.userIDs
}

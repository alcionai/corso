package source

type service int

//go:generate stringer -type=service -linecomment
const (
	ServiceUnknown  service = iota // Unknown Service
	ServiceExchange                // Exchange
)

type Source struct {
	Service service
	userIDs []string
}

// Creates a new source for the given application.
func NewSource(s service) *Source {
	return &Source{
		Service: s,
	}
}

// AddUsers adds the provide user IDs to the source.
// Data retrieval will be scoped to include the unioned
// set of specified users.
func (s *Source) AddUsers(uids ...string) error {
	// future todo: not all services identify users.
	// this should error when users are added in that context.
	if s.userIDs == nil {
		s.userIDs = []string{}
	}
	for _, uid := range uids {
		if len(uid) > 0 {
			s.userIDs = append(s.userIDs, uid)
		}
	}
	return nil
}

// Users retrieves the userIDs specified by the caller.
func (s *Source) Users() []string {
	if s.userIDs == nil {
		return []string{}
	}
	return s.userIDs
}

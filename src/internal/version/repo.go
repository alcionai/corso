package version

const Repo = 1

// Capability constants denote the type of capability supported by a repo version.
type Capability int

//go:generate go run golang.org/x/tools/cmd/stringer -type=Capability
const (
	UnknownCapability = Capability(iota)
	ImmutableIDCapability
)

func RepoCapabilities(v int) map[string]struct{} {
	if v > 1 {
		return map[string]struct{}{
			ImmutableIDCapability.String(): {},
		}
	}
	return map[string]struct{}{}
}

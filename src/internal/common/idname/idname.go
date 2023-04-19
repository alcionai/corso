package idname

import (
	"strings"

	"golang.org/x/exp/maps"
)

// Provider is a tuple containing an ID and a Name.  Names are
// assumed to be human-displayable versions of system IDs.
// Providers should always be populated, while a nil values is
// likely an error.  Compliant structs should provide both a name
// and an ID, never just one.  Values are not validated, so both
// values being empty is an allowed conditions, but the assumption
// is that downstream consumers will have problems as a result.
type Provider interface {
	// ID returns the canonical id of the thing, generated and
	// usable  by whichever system has ownership of it.
	ID() string
	// the human-readable name of the thing.
	Name() string
}

var _ Provider = &Is{}

// Is provides an id-name tuple.
type Is struct {
	id   string
	name string
}

func (is Is) ID() string   { return is.id }
func (is Is) Name() string { return is.name }

type Cacher interface {
	IDOf(name string) (string, bool)
	NameOf(id string) (string, bool)
	IDs() []string
	Names() []string
	ProviderForID(id string) Provider
	ProviderForName(id string) Provider
}

var _ Cacher = &Cache{}

// Cache holds a cache of id-name mappings.
type Cache struct {
	idToName map[string]string
	nameToID map[string]string
}

func NewCache(idToName map[string]string) Cache {
	nti := make(map[string]string, len(idToName))

	for id, name := range idToName {
		nti[name] = id
	}

	return Cache{
		idToName: idToName,
		nameToID: nti,
	}
}

// IDOf returns the id associated with the given name.
func (c Cache) IDOf(name string) (string, bool) {
	id, ok := c.nameToID[strings.ToLower(name)]
	return id, ok
}

// NameOf returns the name associated with the given id.
func (c Cache) NameOf(id string) (string, bool) {
	name, ok := c.idToName[strings.ToLower(id)]
	return name, ok
}

// IDs returns all known ids.
func (c Cache) IDs() []string {
	return maps.Keys(c.idToName)
}

// Names returns all known names.
func (c Cache) Names() []string {
	return maps.Keys(c.nameToID)
}

func (c Cache) ProviderForID(id string) Provider {
	n, ok := c.NameOf(id)
	if !ok {
		return &Is{}
	}

	return &Is{
		id:   id,
		name: n,
	}
}

func (c Cache) ProviderForName(name string) Provider {
	i, ok := c.IDOf(name)
	if !ok {
		return &Is{}
	}

	return &Is{
		id:   i,
		name: name,
	}
}

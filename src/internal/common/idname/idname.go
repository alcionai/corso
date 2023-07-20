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

var _ Provider = &is{}

type is struct {
	id   string
	name string
}

func NewProvider(id, name string) *is {
	return &is{id, name}
}

func (is is) ID() string   { return is.id }
func (is is) Name() string { return is.name }

type Cacher interface {
	IDOf(name string) (string, bool)
	NameOf(id string) (string, bool)
	IDs() []string
	Names() []string
	ProviderForID(id string) Provider
	ProviderForName(id string) Provider
}

type CacheBuilder interface {
	Add(id, name string)
	Cacher
}

var _ Cacher = &cache{}

type cache struct {
	idToName map[string]string
	nameToID map[string]string
}

func NewCache(idToName map[string]string) *cache {
	c := cache{
		idToName: map[string]string{},
		nameToID: map[string]string{},
	}

	if len(idToName) > 0 {
		nti := make(map[string]string, len(idToName))

		for id, name := range idToName {
			nti[name] = id
		}

		c.idToName = idToName
		c.nameToID = nti
	}

	return &c
}

func (c *cache) Add(id, name string) {
	c.idToName[strings.ToLower(id)] = name
	c.nameToID[strings.ToLower(name)] = id
}

// IDOf returns the id associated with the given name.
func (c cache) IDOf(name string) (string, bool) {
	id, ok := c.nameToID[strings.ToLower(name)]
	return id, ok
}

// NameOf returns the name associated with the given id.
func (c cache) NameOf(id string) (string, bool) {
	name, ok := c.idToName[strings.ToLower(id)]
	return name, ok
}

// IDs returns all known ids.
func (c cache) IDs() []string {
	return maps.Keys(c.idToName)
}

// Names returns all known names.
func (c cache) Names() []string {
	return maps.Keys(c.nameToID)
}

func (c cache) ProviderForID(id string) Provider {
	n, ok := c.NameOf(id)
	if !ok {
		return &is{}
	}

	return &is{
		id:   id,
		name: n,
	}
}

func (c cache) ProviderForName(name string) Provider {
	i, ok := c.IDOf(name)
	if !ok {
		return &is{}
	}

	return &is{
		id:   i,
		name: name,
	}
}

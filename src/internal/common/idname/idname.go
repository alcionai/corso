package idname

import (
	"strings"

	"golang.org/x/exp/maps"
)

type Provider interface {
	// the canonical id of the thing, generated and usable
	//  by whichever system has ownership of it.
	ID() string
	// the human-readable name of the thing.
	Name() string
}

var _ Provider = &Is{}

// Is provides an id-name tuple.
type Is struct {
	IDV   string
	NameV string
}

func (is Is) ID() string   { return is.IDV }
func (is Is) Name() string { return is.NameV }

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
	IDToName map[string]string
	NameToID map[string]string
}

// IDOf returns the id associated with the given name.
func (c Cache) IDOf(name string) (string, bool) {
	id, ok := c.NameToID[strings.ToLower(name)]
	return id, ok
}

// NameOf returns the name associated with the given id.
func (c Cache) NameOf(id string) (string, bool) {
	name, ok := c.IDToName[strings.ToLower(id)]
	return name, ok
}

// IDs returns all known ids.
func (c Cache) IDs() []string {
	return maps.Keys(c.IDToName)
}

// Names returns all known names.
func (c Cache) Names() []string {
	return maps.Keys(c.NameToID)
}

func (c Cache) ProviderForID(id string) Provider {
	n, ok := c.NameOf(id)
	if !ok {
		return &Is{}
	}

	return &Is{
		IDV:   id,
		NameV: n,
	}
}

func (c Cache) ProviderForName(name string) Provider {
	i, ok := c.IDOf(name)
	if !ok {
		return &Is{}
	}

	return &Is{
		IDV:   i,
		NameV: name,
	}
}

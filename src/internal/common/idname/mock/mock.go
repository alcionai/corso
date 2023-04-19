package mock

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/idname"
	"golang.org/x/exp/maps"
)

var _ idname.Provider = &in{}

func NewProvider(id, name string) *in {
	return &in{
		id:   id,
		name: name,
	}
}

type in struct {
	id   string
	name string
}

func (i in) ID() string   { return i.id }
func (i in) Name() string { return i.name }

type Cache struct {
	IDToName map[string]string
	NameToID map[string]string
}

func NewCache(itn, nti map[string]string) Cache {
	return Cache{
		IDToName: itn,
		NameToID: nti,
	}
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

func (c Cache) ProviderForID(id string) idname.Provider {
	n, ok := c.NameOf(id)
	if !ok {
		return nil
	}

	return &in{
		id:   id,
		name: n,
	}
}

func (c Cache) ProviderForName(name string) idname.Provider {
	i, ok := c.IDOf(name)
	if !ok {
		return nil
	}

	return &in{
		id:   i,
		name: name,
	}
}

package common

import (
	"strings"

	"golang.org/x/exp/maps"
)

type IDNamer interface {
	// the canonical id of the thing, generated and usable
	//  by whichever system has ownership of it.
	ID() string
	// the human-readable name of the thing.
	Name() string
}

type IDNameSwapper interface {
	IDOf(name string) (string, bool)
	NameOf(id string) (string, bool)
	IDs() []string
	Names() []string
}

var _ IDNameSwapper = &IDsNames{}

type IDsNames struct {
	IDToName map[string]string
	NameToID map[string]string
}

// IDOf returns the id associated with the given name.
func (in IDsNames) IDOf(name string) (string, bool) {
	id, ok := in.NameToID[strings.ToLower(name)]
	return id, ok
}

// NameOf returns the name associated with the given id.
func (in IDsNames) NameOf(id string) (string, bool) {
	name, ok := in.IDToName[strings.ToLower(id)]
	return name, ok
}

// IDs returns all known ids.
func (in IDsNames) IDs() []string {
	return maps.Keys(in.IDToName)
}

// Names returns all known names.
func (in IDsNames) Names() []string {
	return maps.Keys(in.NameToID)
}

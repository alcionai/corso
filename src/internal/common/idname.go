package common

type IDNamer interface {
	// the canonical id of the thing, generated and usable
	//  by whichever system has ownership of it.
	ID() string
	// the human-readable name of the thing.
	Name() string
}

package resource

type Category int

const (
	UnknownResource Category = iota
	AllResources             // deprecated, kept for iota marker
	Users
	Sites
)

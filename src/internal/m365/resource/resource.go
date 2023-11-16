package resource

import "github.com/alcionai/clues"

var ErrNoResourceLookup = clues.New("missing resource lookup client")

type Category string

const (
	UnknownResource Category = ""
	Users           Category = "users"
	Sites           Category = "sites"
	Groups          Category = "groups"
)

package resource

type Category string

const (
	UnknownResource Category = ""
	Users           Category = "users"
	Sites           Category = "sites"
	Groups          Category = "groups"
)

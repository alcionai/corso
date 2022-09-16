package control

// CollisionPolicy describes how the datalayer behaves in case of a collision.
type CollisionPolicy int

//go:generate stringer -type=CollisionPolicy
const (
	Unknown CollisionPolicy = iota
	Copy
	Skip
	Replace
)

// Options holds the optional configurations for a process
type Options struct {
	Collision      CollisionPolicy `json:"-"`
	DisableMetrics bool            `json:"disableMetrics"`
	FailFast       bool            `json:"failFast"`
}

// Defaults provides an Options with the default values set.
func Defaults() Options {
	return Options{
		FailFast: true,
	}
}

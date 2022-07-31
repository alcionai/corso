package control

// CollisionPolicy describes how the datalayer behaves in case of a collision.
type CollisionPolicy int

const (
	Unknown CollisionPolicy = iota
	Copy
	Skip
	Replace
)

// Options holds the optional configurations for a process
type Options struct {
	FailFast  bool            `json:"failFast"`
	Collision CollisionPolicy `json:"-"`
}

func NewOptions(failFast bool) Options {
	return Options{
		FailFast: failFast,
	}
}

package common

// Policy is a type that defines the actions taken
type RestorePolicy int

//go:generate stringer -type=RestorePolicy
const (
	Unknown RestorePolicy = iota
	Copy
	Drop
	Replace
)
3
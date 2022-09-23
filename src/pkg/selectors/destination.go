package selectors

// RestoreDestination is a POD that contains the resource owner to restore data
// under and the name of the root of the restored directory hierarchy.
type RestoreDestination struct {
	ResourceOwner string
	ContainerName string
}

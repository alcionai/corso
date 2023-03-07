package fault

type itemKind string

const (
	ItemKindFile          itemKind = "file"
	ItemKindContainer     itemKind = "container"
	ItemKindResourceOwner itemKind = "resource_owner"
)

// Item contains a concrete reference to a thing that failed
// during processing.  The categorization of the item is determined
// by its Kind: file, container, or reourceOwner.
//
// Item is compliant with the error interface so that it can be
// aggregated with the fault bus, and deserialized using the
// errors.As() func.  The idea is that fault,Items, during
// processing, will get packed into bus.AddRecoverable (or failure)
// as part of standard error handling, and later deserialized
// by the end user (cli or sdk) for surfacing human-readable and
// identifiable points of failure.
type Item struct {
	// deduplication identifier; the ID of the observed item.
	ID string `json:"id"`

	// a human-readable reference: file/container name, email, etc
	Name string `json:"name"`

	// the name and id of the container holding this item, if the
	// item is normally stored in a folder.  Ex: a oneDrive file
	// would be contained in a folder.
	ContainerName string `json:"containerName"`
	ContainerID   string `json:"containerID"`

	// tracks the type of item represented by this entry.
	Kind itemKind `json:"kind"`

	// Error() of the causal error, or a sentinel if this is the
	// source of the error.  In case of ID collisions, the first
	// item takes priority.
	Cause string `json:"cause"`
}

// Error complies with the error interface.
func (i *Item) Error() string {
	if i == nil {
		return "<nil>"
	}

	if len(i.Kind) == 0 {
		return "processing item of unknown kind"
	}

	return string("processing " + i.Kind)
}

// ContainerErr constructs a Container-kind Item.
func ContainerErr(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Kind:          ItemKindContainer,
		Cause:         cause.Error(),
	}
}

// FileErr constructs a File-kind Item.
func FileErr(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Kind:          ItemKindFile,
		Cause:         cause.Error(),
	}
}

// OnwerErr constructs a ResourceOwner-kind Item.
func OwnerErr(cause error, id, name string) *Item {
	return &Item{
		ID:    id,
		Name:  name,
		Kind:  ItemKindResourceOwner,
		Cause: cause.Error(),
	}
}

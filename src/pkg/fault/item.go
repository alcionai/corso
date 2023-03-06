package fault

type itemKind string

const (
	ItemKindFile          itemKind = "file"
	ItemKindContainer     itemKind = "container"
	ItemKindResourceOwner itemKind = "resource_owner"
)

// in pkg fault.  eg: fault.Item{}
type Item struct {
	// deduplication identifier; the ID of the item under scrutiny.
	ID string `json:"id"`

	// a human-readable reference: file/container name, email, etc
	Name string `json:"name"`

	// the name and id of the container inside which this item is
	// stored.  May be empty if the item is not stored within a
	// container (ex: resourceOwner)
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

// ContainerItem produces a Container-kind Item.
func ContainerItem(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Kind:          ItemKindContainer,
		Cause:         cause.Error(),
	}
}

// FileItem produces a File-kind Item.
func FileItem(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Kind:          ItemKindFile,
		Cause:         cause.Error(),
	}
}

// OnwerItem produces a ResourceOwner-kind Item.
func OwnerItem(cause error, id, name string) *Item {
	return &Item{
		ID:    id,
		Name:  name,
		Kind:  ItemKindResourceOwner,
		Cause: cause.Error(),
	}
}

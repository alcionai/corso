package fault

type itemType string

const (
	FileType          itemType = "file"
	ContainerType     itemType = "container"
	ResourceOwnerType itemType = "resource_owner"
)

var _ error = &Item{}

// in pkg fault.  eg: fault.Item{}
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
	Type itemType `json:"type"`

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

	if len(i.Type) == 0 {
		return "processing item of unknown type"
	}

	return string("processing " + i.Type)
}

// ---------------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------------

// ContainerErr produces a Container-kind Item.
func ContainerErr(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Type:          ContainerType,
		Cause:         cause.Error(),
	}
}

// FileErr constructs a File-type Item.
func FileErr(cause error, id, name, containerID, containerName string) *Item {
	return &Item{
		ID:            id,
		Name:          name,
		ContainerID:   containerID,
		ContainerName: containerName,
		Type:          FileType,
		Cause:         cause.Error(),
	}
}

// OnwerErr constructs a ResourceOwner-type Item.
func OwnerErr(cause error, id, name string) *Item {
	return &Item{
		ID:    id,
		Name:  name,
		Type:  ResourceOwnerType,
		Cause: cause.Error(),
	}
}

package fault

type itemType string

const (
	FileType          itemType = "file"
	ContainerType     itemType = "container"
	ResourceOwnerType itemType = "resource_owner"
)

var _ error = &Item{}

// Item contains a concrete reference to a thing that failed
// during processing.  The categorization of the item is determined
// by its Type: file, container, or reourceOwner.
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

// ContainerErr produces a Container-type Item for tracking erronous items
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

// FileErr produces a File-type Item for tracking erronous items.
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

// OnwerErr produces a ResourceOwner-type Item for tracking erronous items.
func OwnerErr(cause error, id, name string) *Item {
	return &Item{
		ID:    id,
		Name:  name,
		Type:  ResourceOwnerType,
		Cause: cause.Error(),
	}
}

// ---------------------------------------------------------------------------
// Skipped Items
// ---------------------------------------------------------------------------

// skipCause identifies the well-known conditions to Skip an item.  It is
// important that skip cause enumerations do not overlap with general error
// handling.  Skips must be well known, well documented, and consistent.
// Transient failures, undocumented or unknown conditions, and arbitrary
// handling should never produce a skipped item. Those cases should get
// handled as normal errors.
type skipCause string

// SkipMalware identifies a malware detection case.  Files that graph api
// identifies as malware cannot be downloaded or uploaded, and will permanently
// fail any attempts to backup or restore.
const SkipMalware skipCause = "malware_detected"

// Skipped items are permanently unprocessable due to well-known conditions.
// In order to skip an item, the following conditions should be met:
// 1. The conditions for skipping the item are well-known and
// well-documented.  End users need to be able to understand
// both the conditions and identifications of skips.
// 2. Skipping avoids a permanent and consistent failure.  If
// the underlying reason is transient or otherwise recoverable,
// the item should not be skipped.
//
// Skipped wraps Item primarily to minimze confusion when sharing the
// fault interface.  Skipped items are not errors, and Item{} errors are
// not the basis for a Skip.
type Skipped struct {
	item Item
}

// String complies with the stringer interface. func (s *Skipped) String() string {
func (s *Skipped) String() string {
	if s == nil {
		return "<nil>"
	}

	return "skipped " + s.item.Error() + ": " + s.item.Cause
}

// ContainerSkip produces a Container-kind Item for tracking skipped items.
func ContainerSkip(cause skipCause, id, name, containerID, containerName string) *Skipped {
	return &Skipped{
		item: Item{
			ID:            id,
			Name:          name,
			ContainerID:   containerID,
			ContainerName: containerName,
			Type:          ContainerType,
			Cause:         string(cause),
		},
	}
}

// FileSkip produces a File-kind Item for tracking skipped items.
func FileSkip(cause skipCause, id, name, containerID, containerName string) *Skipped {
	return &Skipped{
		item: Item{
			ID:            id,
			Name:          name,
			ContainerID:   containerID,
			ContainerName: containerName,
			Type:          FileType,
			Cause:         string(cause),
		},
	}
}

// OnwerSkip produces a ResourceOwner-kind Item for tracking skipped items.
func OwnerSkip(cause skipCause, id, name string) *Skipped {
	return &Skipped{
		item: Item{
			ID:    id,
			Name:  name,
			Type:  ResourceOwnerType,
			Cause: string(cause),
		},
	}
}

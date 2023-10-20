package fault

import "github.com/alcionai/corso/src/cli/print"

const (
	AddtlCreatedBy     = "created_by"
	AddtlLastModBy     = "last_modified_by"
	AddtlContainerID   = "container_id"
	AddtlContainerName = "container_name"
	AddtlContainerPath = "container_path"
	AddtlMalwareDesc   = "malware_description"
)

type ItemType string

const (
	FileType          ItemType = "file"
	ContainerType     ItemType = "container"
	ResourceOwnerType ItemType = "resource_owner"
)

func (it ItemType) Printable() string {
	switch it {
	case FileType:
		return "File"
	case ContainerType:
		return "Container"
	case ResourceOwnerType:
		return "Resource Owner"
	}

	return "Unknown"
}

var (
	_ error           = &Item{}
	_ print.Printable = &Item{}
)

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
	// deduplication namespace; the maximally-unique boundary of the
	// item ID.  The scope of this boundary depends on the service.
	// ex: exchange items are unique within their category, drive items
	// are only unique within a given drive.
	Namespace string `json:"namespace"`

	// deduplication identifier; the ID of the observed item.
	ID string `json:"id"`

	// a human-readable reference: file/container name, email, etc
	Name string `json:"name"`

	// tracks the type of item represented by this entry.
	Type ItemType `json:"type"`

	// Error() of the causal error, or a sentinel if this is the
	// source of the error.  In case of ID collisions, the first
	// item takes priority.
	Cause string `json:"cause"`

	// Additional is a catch-all map for storing data that might
	// be relevant to particular types or contexts of items without
	// being globally relevant.  Ex: parent container references,
	// created-by ids, last modified, etc.  Should be used sparingly,
	// only for information that might be immediately relevant to the
	// end user.
	Additional map[string]any `json:"additional"`
}

// dedupeID is the id used to deduplicate items when aggreagating
// errors in fault.Errors().
func (i *Item) dedupeID() string {
	return i.Namespace + i.ID
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

func (i Item) MinimumPrintable() any {
	return i
}

// Headers returns the human-readable names of properties of an Item
// for printing out to a terminal.
func (i Item) Headers(bool) []string {
	// NOTE: skipID does not make sense in this context
	return []string{"Action", "Type", "Name", "Container", "Cause"}
}

// Values populates the printable values matching the Headers list.
func (i Item) Values(bool) []string {
	var cn string

	acn, ok := i.Additional[AddtlContainerName]
	if ok {
		str, ok := acn.(string)
		if ok {
			cn = str
		}
	}

	return []string{"Error", i.Type.Printable(), i.Name, cn, i.Cause}
}

// ContainerErr produces a Container-type Item for tracking erroneous items
func ContainerErr(cause error, namespace, id, name string, addtl map[string]any) *Item {
	return itemErr(ContainerType, cause, namespace, id, name, addtl)
}

// FileErr produces a File-type Item for tracking erroneous items.
func FileErr(cause error, namespace, id, name string, addtl map[string]any) *Item {
	return itemErr(FileType, cause, namespace, id, name, addtl)
}

// OnwerErr produces a ResourceOwner-type Item for tracking erroneous items.
func OwnerErr(cause error, namespace, id, name string, addtl map[string]any) *Item {
	return itemErr(ResourceOwnerType, cause, namespace, id, name, addtl)
}

// itemErr produces a Item of the provided type for tracking erroneous items.
func itemErr(t ItemType, cause error, namespace, id, name string, addtl map[string]any) *Item {
	return &Item{
		Namespace:  namespace,
		ID:         id,
		Name:       name,
		Type:       t,
		Cause:      cause.Error(),
		Additional: addtl,
	}
}

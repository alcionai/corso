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

type itemType string

const (
	FileType          itemType = "file"
	AttachmentType    itemType = "attachment"
	ContainerType     itemType = "container"
	ResourceOwnerType itemType = "resource_owner"
)

func (it itemType) Printable() string {
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
	Type itemType `json:"type"`

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
func (i Item) Headers() []string {
	return []string{"Action", "Type", "Name", "Container", "Cause"}
}

// Values populates the printable values matching the Headers list.
func (i Item) Values() []string {
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
func itemErr(t itemType, cause error, namespace, id, name string, addtl map[string]any) *Item {
	return &Item{
		Namespace:  namespace,
		ID:         id,
		Name:       name,
		Type:       t,
		Cause:      cause.Error(),
		Additional: addtl,
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

const (
	// SkipMalware identifies a malware detection case.  Files that graph
	// api identifies as malware cannot be downloaded or uploaded, and will
	// permanently fail any attempts to backup or restore.
	SkipMalware skipCause = "malware_detected"

	// SkipNotFound identifies that a file was skipped because we could
	// not find it when trying to download contents
	SkipNotFound skipCause = "file_not_found"

	// SkipBigOneNote identifies that a file was skipped because it
	// was big OneNote file and we can only download OneNote files which
	// are less that 2GB in size.
	//nolint:lll
	// https://support.microsoft.com/en-us/office/restrictions-and-limitations-in-onedrive-and-sharepoint-64883a5d-228e-48f5-b3d2-eb39e07630fa#onenotenotebooks
	SkipBigOneNote skipCause = "big_one_note_file"
)

var _ print.Printable = &Skipped{}

// Skipped items are permanently unprocessable due to well-known conditions.
// In order to skip an item, the following conditions should be met:
// 1. The conditions for skipping the item are well-known and
// well-documented.  End users need to be able to understand
// both the conditions and identifications of skips.
// 2. Skipping avoids a permanent and consistent failure.  If
// the underlying reason is transient or otherwise recoverable,
// the item should not be skipped.
//
// Skipped wraps Item primarily to minimize confusion when sharing the
// fault interface.  Skipped items are not errors, and Item{} errors are
// not the basis for a Skip.
type Skipped struct {
	Item Item `json:"item"`
}

// String complies with the stringer interface.
func (s *Skipped) String() string {
	if s == nil {
		return "<nil>"
	}

	return "skipped " + s.Item.Error() + ": " + s.Item.Cause
}

// HasCause compares the underlying cause against the parameter.
func (s *Skipped) HasCause(c skipCause) bool {
	if s == nil {
		return false
	}

	return s.Item.Cause == string(c)
}

func (s Skipped) MinimumPrintable() any {
	return s
}

// Headers returns the human-readable names of properties of a skipped Item
// for printing out to a terminal.
func (s Skipped) Headers() []string {
	return []string{"Action", "Type", "Name", "Container", "Cause"}
}

// Values populates the printable values matching the Headers list.
func (s Skipped) Values() []string {
	var cn string

	acn, ok := s.Item.Additional[AddtlContainerName]
	if ok {
		str, ok := acn.(string)
		if ok {
			cn = str
		}
	}

	return []string{"Skip", s.Item.Type.Printable(), s.Item.Name, cn, s.Item.Cause}
}

// ContainerSkip produces a Container-kind Item for tracking skipped items.
func ContainerSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(ContainerType, cause, namespace, id, name, addtl)
}

// FileSkip produces a File-kind Item for tracking skipped items.
func FileSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(FileType, cause, namespace, id, name, addtl)
}

// AttachmentSkip produces a Attachment-kind Item for tracking skipped items.
func AttachmentSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(AttachmentType, cause, namespace, id, name, addtl)
}

// OnwerSkip produces a ResourceOwner-kind Item for tracking skipped items.
func OwnerSkip(cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return itemSkip(ResourceOwnerType, cause, namespace, id, name, addtl)
}

// itemSkip produces a Item of the provided type for tracking skipped items.
func itemSkip(t itemType, cause skipCause, namespace, id, name string, addtl map[string]any) *Skipped {
	return &Skipped{
		Item: Item{
			Namespace:  namespace,
			ID:         id,
			Name:       name,
			Type:       t,
			Cause:      string(cause),
			Additional: addtl,
		},
	}
}

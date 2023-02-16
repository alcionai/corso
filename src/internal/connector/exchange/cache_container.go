package exchange

import (
	"github.com/pkg/errors"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
)

// checkIDAndName is a helper function to ensure that
// the ID and name pointers are set prior to being called.
func checkIDAndName(c graph.Container) error {
	id, ok := ptr.ValOK(c.GetId())
	if !ok {
		return errors.New("container missing ID")
	}

	if _, ok := ptr.ValOK(c.GetDisplayName()); !ok {
		return clues.New("container missing display name").With("container_id", id)
	}

	return nil
}

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func checkRequiredValues(c graph.Container) error {
	if err := checkIDAndName(c); err != nil {
		return err
	}

	if _, ok := ptr.ValOK(c.GetParentFolderId()); !ok {
		return clues.New("container missing parent ID").With("container_id", ptr.Val(c.GetId()))
	}

	return nil
}

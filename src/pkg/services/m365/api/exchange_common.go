package api

import (
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
)

// checkIDAndName is a helper function to ensure that
// the ID and name pointers are set prior to being called.
func checkIDAndName(c graph.Container) error {
	id := ptr.Val(c.GetId())
	if len(id) == 0 {
		return clues.New("container missing ID")
	}

	dn := ptr.Val(c.GetDisplayName())
	if len(dn) == 0 {
		return clues.New("container missing display name").With("container_id", id)
	}

	return nil
}

func HasAttachments(body models.ItemBodyable) bool {
	if body == nil {
		return false
	}

	if ct, ok := ptr.ValOK(body.GetContentType()); !ok || ct == models.TEXT_BODYTYPE {
		return false
	}

	if body, ok := ptr.ValOK(body.GetContent()); !ok || len(body) == 0 {
		return false
	}

	return strings.Contains(ptr.Val(body.GetContent()), "src=\"cid:")
}

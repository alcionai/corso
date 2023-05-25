package api

import (
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

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

func makeSessionAttachment(name string, size int64) *models.AttachmentItem {
	attItem := models.NewAttachmentItem()
	attType := models.FILE_ATTACHMENTTYPE
	attItem.SetAttachmentType(&attType)
	attItem.SetName(&name)
	attItem.SetSize(&size)

	return attItem
}

func GetAttachmentContent(attachment models.Attachmentable) ([]byte, error) {
	ibs, err := attachment.GetBackingStore().Get("contentBytes")
	if err != nil {
		return nil, err
	}

	bs, ok := ibs.([]byte)
	if !ok {
		return nil, clues.New(fmt.Sprintf("unexpected type for attachment content: %T", ibs))
	}

	return bs, nil
}

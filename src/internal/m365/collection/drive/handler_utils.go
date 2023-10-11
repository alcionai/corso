package drive

import (
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

func getItemCreator(item models.DriveItemable) string {
	if item.GetCreatedBy() == nil || item.GetCreatedBy().GetUser() == nil {
		return ""
	}

	// User is sometimes not available when created via some
	// external applications (like backup/restore solutions)
	additionalData := item.GetCreatedBy().GetUser().GetAdditionalData()

	ed, ok := additionalData["email"]
	if !ok {
		ed = additionalData["displayName"]
	}

	if ed == nil {
		return ""
	}

	return *ed.(*string)
}

func getItemDriveInfo(item models.DriveItemable) (string, string) {
	if item.GetParentReference() == nil {
		return "", ""
	}

	driveName := strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	driveID := ptr.Val(item.GetParentReference().GetDriveId())

	return driveName, driveID
}

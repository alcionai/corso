package drive

import (
	"strings"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

func getItemCreator(item *custom.DriveItem) string {
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

	// TODO(ashmrtn): Replace with str package with fallbacks.
	return *ed.(*string)
}

func getItemDriveInfo(item *custom.DriveItem) (string, string) {
	if item.GetParentReference() == nil {
		return "", ""
	}

	driveName := strings.TrimSpace(ptr.Val(item.GetParentReference().GetName()))
	driveID := ptr.Val(item.GetParentReference().GetDriveId())

	return driveName, driveID
}

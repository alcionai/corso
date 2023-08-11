package testdata

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func NewStubDriveItem(
	id, name string,
	size int64,
	created, modified time.Time,
	isFile, isShared bool,
) models.DriveItemable {
	stubItem := models.NewDriveItem()
	stubItem.SetId(&id)
	stubItem.SetName(&name)
	stubItem.SetSize(&size)
	stubItem.SetCreatedDateTime(&created)
	stubItem.SetLastModifiedDateTime(&modified)
	stubItem.SetAdditionalData(map[string]any{"@microsoft.graph.downloadUrl": "https://corsobackup.io"})

	if isFile {
		stubItem.SetFile(models.NewFile())
	}

	if isShared {
		stubItem.SetShared(&models.Shared{})
	}

	return stubItem
}

package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// downloadUrlKeys is used to find the download URL in a DriveItem response.
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

func downloadItem(
	ctx context.Context,
	ag api.Getter,
	item models.DriveItemable,
) (io.ReadCloser, error) {
	var (
		rc     io.ReadCloser
		isFile = item.GetFile() != nil
	)

	if isFile {
		var (
			url string
			ad  = item.GetAdditionalData()
		)

		for _, key := range downloadURLKeys {
			if v, err := str.AnyValueToString(key, ad); err == nil {
				url = v
				break
			}
		}

		if len(url) == 0 {
			return nil, clues.New("extracting file url")
		}

		resp, err := ag.Get(ctx, url, nil)
		if err != nil {
			return nil, clues.Wrap(err, "getting item")
		}

		if graph.IsMalwareResp(ctx, resp) {
			return nil, clues.New("malware detected").Label(graph.LabelsMalware)
		}

		if (resp.StatusCode / 100) != 2 {
			// upstream error checks can compare the status with
			// clues.HasLabel(err, graph.LabelStatus(http.KnownStatusCode))
			return nil, clues.
				Wrap(clues.New(resp.Status), "non-2xx http response").
				Label(graph.LabelStatus(resp.StatusCode))
		}

		rc = resp.Body
	}

	return rc, nil
}

func downloadItemMeta(
	ctx context.Context,
	gip GetItemPermissioner,
	driveID string,
	item models.DriveItemable,
) (io.ReadCloser, int, error) {
	meta := metadata.Metadata{FileName: ptr.Val(item.GetName())}

	if item.GetShared() == nil {
		meta.SharingMode = metadata.SharingModeInherited
	} else {
		meta.SharingMode = metadata.SharingModeCustom
	}

	if meta.SharingMode == metadata.SharingModeCustom {
		perm, err := gip.GetItemPermission(ctx, driveID, ptr.Val(item.GetId()))
		if err != nil {
			return nil, 0, err
		}

		meta.Permissions = metadata.FilterPermissions(ctx, perm.GetValue())
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, 0, clues.Wrap(err, "serializing item metadata").WithClues(ctx)
	}

	return io.NopCloser(bytes.NewReader(metaJSON)), len(metaJSON), nil
}

// driveItemWriter is used to initialize and return an io.Writer to upload data for the specified item
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
// TODO: @vkamra verify if var session is the desired input
func driveItemWriter(
	ctx context.Context,
	nicu NewItemContentUploader,
	driveID, itemID string,
	itemSize int64,
) (io.Writer, string, error) {
	ctx = clues.Add(ctx, "upload_item_id", itemID)

	icu, err := nicu.NewItemContentUpload(ctx, driveID, itemID)
	if err != nil {
		return nil, "", clues.Stack(err)
	}

	iw := graph.NewLargeItemWriter(itemID, ptr.Val(icu.GetUploadUrl()), itemSize)

	return iw, ptr.Val(icu.GetUploadUrl()), nil
}

func setName(orig models.ItemReferenceable, driveName string) models.ItemReferenceable {
	if orig == nil {
		return nil
	}

	orig.SetName(&driveName)

	return orig
}

package drive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	acceptHeaderKey   = "Accept"
	acceptHeaderValue = "*/*"
)

// downloadUrlKeys is used to find the download URL in a DriveItem response.
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

func downloadItem(
	ctx context.Context,
	ag api.Getter,
	item CorsoDriveItemable,
) (io.ReadCloser, error) {
	if item == nil {
		return nil, clues.New("nil item")
	}

	var (
		rc     io.ReadCloser
		isFile = item.GetFile() != nil
		err    error
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

		rc, err = downloadFile(ctx, ag, url)
		if err != nil {
			return nil, clues.Stack(err)
		}
	}

	return rc, nil
}

type downloadWithRetries struct {
	getter api.Getter
	url    string
}

func (dg *downloadWithRetries) SupportsRange() bool {
	return true
}

func (dg *downloadWithRetries) Get(
	ctx context.Context,
	additionalHeaders map[string]string,
) (io.ReadCloser, error) {
	headers := maps.Clone(additionalHeaders)
	// Set the accept header like curl does. Local testing showed range headers
	// wouldn't work without it (get 416 responses instead of 206).
	headers[acceptHeaderKey] = acceptHeaderValue

	resp, err := dg.getter.Get(ctx, dg.url, headers)
	if err != nil {
		return nil, clues.Wrap(err, "getting file")
	}

	if graph.IsMalwareResp(ctx, resp) {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}

		return nil, clues.New("malware detected").Label(graph.LabelsMalware)
	}

	if resp != nil && (resp.StatusCode/100) != 2 {
		if resp.Body != nil {
			resp.Body.Close()
		}

		// upstream error checks can compare the status with
		// clues.HasLabel(err, graph.LabelStatus(http.KnownStatusCode))
		return nil, clues.
			Wrap(clues.New(resp.Status), "non-2xx http response").
			Label(graph.LabelStatus(resp.StatusCode))
	}

	return resp.Body, nil
}

func downloadFile(
	ctx context.Context,
	ag api.Getter,
	url string,
) (io.ReadCloser, error) {
	if len(url) == 0 {
		return nil, clues.New("empty file url").WithClues(ctx)
	}

	rc, err := readers.NewResetRetryHandler(
		ctx,
		&downloadWithRetries{
			getter: ag,
			url:    url,
		})

	return rc, clues.Stack(err).OrNil()
}

func downloadItemMeta(
	ctx context.Context,
	gip GetItemPermissioner,
	driveID string,
	item CorsoDriveItemable,
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
		meta.LinkShares = metadata.FilterLinkShares(ctx, perm.GetValue())
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

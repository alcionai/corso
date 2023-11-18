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
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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
		return nil, clues.NewWC(ctx, "empty file url")
	}

	// Precheck for url expiry before we make a call to graph to download the
	// file. If the url is expiredErr, we can return early and save a call to graph.
	//
	// Ignore all errors encountered during the check. We can rely on graph to
	// return errors on malformed urls. Ignoring errors also future proofs against
	// any sudden graph changes, for e.g. if graph decides to embed the token in a
	// new query param.
	expiredErr, err := graph.IsURLExpired(ctx, url)
	if expiredErr != nil {
		logger.CtxErr(ctx, expiredErr).Debug("expired item download url")
		return nil, clues.Stack(expiredErr)
	} else if err != nil {
		logger.CtxErr(ctx, err).Info("checking item download url for expiration")
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
	getter GetItemPermissioner,
	driveID string,
	item CorsoDriveItemable,
) (io.ReadCloser, int, error) {
	meta := metadata.Metadata{
		FileName:    ptr.Val(item.GetName()),
		SharingMode: metadata.SharingModeInherited,
	}

	if item.GetShared() != nil {
		meta.SharingMode = metadata.SharingModeCustom

		perm, err := getter.GetItemPermission(ctx, driveID, ptr.Val(item.GetId()))
		if err != nil {
			return nil, 0, err
		}

		meta.Permissions = metadata.FilterPermissions(ctx, perm.GetValue())
		meta.LinkShares = metadata.FilterLinkShares(ctx, perm.GetValue())
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, 0, clues.WrapWC(ctx, err, "serializing item metadata")
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
	counter *count.Bus,
) (io.Writer, string, error) {
	ctx = clues.Add(ctx, "upload_item_id", itemID)

	icu, err := nicu.NewItemContentUpload(ctx, driveID, itemID)
	if err != nil {
		return nil, "", clues.Stack(err)
	}

	iw := graph.NewLargeItemWriter(
		itemID,
		ptr.Val(icu.GetUploadUrl()),
		itemSize,
		counter)

	return iw, ptr.Val(icu.GetUploadUrl()), nil
}

func setName(orig models.ItemReferenceable, driveName string) models.ItemReferenceable {
	if orig == nil {
		return nil
	}

	orig.SetName(&driveName)

	return orig
}

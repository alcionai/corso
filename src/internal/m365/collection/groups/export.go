package groups

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func NewExportCollection(
	baseDir string,
	backingCollections []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	stats *data.ExportStats,
) export.Collectioner {
	return export.BaseCollection{
		BaseDir:           baseDir,
		BackingCollection: backingCollections,
		BackupVersion:     backupVersion,
		Cfg:               cec,
		Stream:            streamItems,
		Stats:             stats,
	}
}

// streamItems streams the items in the backingCollection into the export stream chan
func streamItems(
	ctx context.Context,
	drc []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	ch chan<- export.Item,
	stats *data.ExportStats,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		for item := range rc.Items(ctx, errs) {
			body, err := formatChannelMessage(cec, item.ToReader())
			if err != nil {
				ch <- export.Item{
					ID:    item.ID(),
					Error: err,
				}
			} else {
				stats.UpdateResourceCount(details.GroupsChannelMessage)
				body = data.ReaderWithStats(body, details.GroupsChannelMessage, stats)

				ch <- export.Item{
					ID: item.ID(),
					// channel message items have no name
					Name: item.ID(),
					Body: body,
				}
			}
		}

		items, recovered := errs.ItemsAndRecovered()

		// Return all the items that we failed to source from the persistence layer
		for _, item := range items {
			ch <- export.Item{
				ID:    item.ID,
				Error: &item,
			}
		}

		for _, err := range recovered {
			ch <- export.Item{
				Error: err,
			}
		}
	}
}

type (
	minimumChannelMessage struct {
		Content              string    `json:"content"`
		CreatedDateTime      time.Time `json:"createdDateTime"`
		From                 string    `json:"from"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	}

	minimumChannelMessageAndReplies struct {
		minimumChannelMessage
		Replies []minimumChannelMessage `json:"replies,omitempty"`
	}
)

func formatChannelMessage(
	cec control.ExportConfig,
	rc io.ReadCloser,
) (io.ReadCloser, error) {
	if cec.Format == control.JSONFormat {
		return rc, nil
	}

	bs, err := io.ReadAll(rc)
	if err != nil {
		return nil, clues.Wrap(err, "reading item bytes")
	}

	defer rc.Close()

	cfb, err := api.CreateFromBytes(bs, models.CreateChatMessageFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to message")
	}

	msg, ok := cfb.(models.ChatMessageable)
	if !ok {
		return nil, clues.New("expected deserialized item to implement models.ChatMessageable")
	}

	mItem := makeMinimumChannelMesasge(msg)
	replies := msg.GetReplies()

	mcmar := minimumChannelMessageAndReplies{
		minimumChannelMessage: mItem,
		Replies:               make([]minimumChannelMessage, 0, len(replies)),
	}

	for _, r := range replies {
		mcmar.Replies = append(mcmar.Replies, makeMinimumChannelMesasge(r))
	}

	bs, err = json.Marshal(mcmar)
	if err != nil {
		return nil, clues.Wrap(err, "serializing minimized channel message")
	}

	return io.NopCloser(bytes.NewReader(bs)), nil
}

func makeMinimumChannelMesasge(item models.ChatMessageable) minimumChannelMessage {
	var content string

	if item.GetBody() != nil {
		content = ptr.Val(item.GetBody().GetContent())
	}

	return minimumChannelMessage{
		Content:              content,
		CreatedDateTime:      ptr.Val(item.GetCreatedDateTime()),
		From:                 api.GetChatMessageFrom(item),
		LastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
	}
}

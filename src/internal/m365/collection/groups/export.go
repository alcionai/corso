package groups

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/converters/eml"
	"github.com/alcionai/canario/src/internal/data"
	groupMeta "github.com/alcionai/canario/src/internal/m365/collection/groups/metadata"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/export"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/logger"
	"github.com/alcionai/canario/src/pkg/metrics"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph/metadata"
)

func NewExportCollection(
	baseDir string,
	backingCollections []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	stats *metrics.ExportStats,
	cat path.CategoryType,
) export.Collectioner {
	var streamItems export.ItemStreamer

	switch cat {
	case path.ChannelMessagesCategory:
		streamItems = streamChannelMessages
	case path.ConversationPostsCategory:
		streamItems = streamConversationPosts
	default:
		return nil
	}

	return export.BaseCollection{
		BaseDir:           baseDir,
		BackingCollection: backingCollections,
		BackupVersion:     backupVersion,
		Cfg:               cec,
		Stream:            streamItems,
		Stats:             stats,
	}
}

//-------------------------------------------------------------
// Channel Messages
//-------------------------------------------------------------

// streamChannelMessages streams the items in the backingCollection into the export stream chan
func streamChannelMessages(
	ctx context.Context,
	drc []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	ch chan<- export.Item,
	stats *metrics.ExportStats,
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
				stats.UpdateResourceCount(path.ChannelMessagesCategory)
				body = metrics.ReaderWithStats(body, path.ChannelMessagesCategory, stats)

				// messages are exported as json and should be named as such
				name := item.ID() + ".json"

				ch <- export.Item{
					ID:   item.ID(),
					Name: name,
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
		Attachments          []minimumAttachment `json:"attachments"`
		Content              string              `json:"content"`
		CreatedDateTime      time.Time           `json:"createdDateTime"`
		From                 string              `json:"from"`
		LastModifiedDateTime time.Time           `json:"lastModifiedDateTime"`
		Subject              string              `json:"subject"`
	}

	minimumChannelMessageAndReplies struct {
		minimumChannelMessage
		Replies []minimumChannelMessage `json:"replies,omitempty"`
	}

	minimumAttachment struct {
		ID   string `json:"id"`
		Name string `json:"name"`
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

	bs, err = marshalJSONContainingHTML(mcmar)
	if err != nil {
		return nil, clues.Wrap(err, "serializing minimized channel message")
	}

	return io.NopCloser(bytes.NewReader(bs)), nil
}

// json.Marshal will replace many markup tags (ex: "<" and ">") with their unicode
// equivalent.  In order to maintain parity with original content that contains html,
// we have to use this alternative encoding behavior.
// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
func marshalJSONContainingHTML(a any) ([]byte, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(a)

	return buffer.Bytes(), clues.Stack(err).OrNil()
}

func makeMinimumChannelMesasge(item models.ChatMessageable) minimumChannelMessage {
	var content string

	if item.GetBody() != nil {
		content = ptr.Val(item.GetBody().GetContent())
	}

	attachments := item.GetAttachments()
	minAttachments := make([]minimumAttachment, 0, len(attachments))

	for _, a := range attachments {
		minAttachments = append(minAttachments, minimumAttachment{
			ID:   ptr.Val(a.GetId()),
			Name: ptr.Val(a.GetName()),
		})
	}

	return minimumChannelMessage{
		Attachments:          minAttachments,
		Content:              content,
		CreatedDateTime:      ptr.Val(item.GetCreatedDateTime()),
		From:                 api.GetChatMessageFrom(item),
		LastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
		Subject:              ptr.Val(item.GetSubject()),
	}
}

//-------------------------------------------------------------
// Conversation Posts
//-------------------------------------------------------------

// streamConversationPosts adds the post items into the export stream channel.
func streamConversationPosts(
	ctx context.Context,
	drc []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	ch chan<- export.Item,
	stats *metrics.ExportStats,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		for item := range rc.Items(ctx, errs) {
			ictx := clues.Add(
				ctx,
				"path_short_ref", rc.FullPath().ShortRef(),
				"stream_item_id", item.ID())

			// Trim .data suffix from itemID. Also, we don't expect .meta files
			// here since details are not persisted for metadata files.
			trimmedID := strings.TrimSuffix(item.ID(), metadata.DataFileSuffix)
			exportName := trimmedID + ".eml"

			postMetadata, err := fetchAndReadMetadata(ictx, trimmedID, rc)
			if err != nil {
				ch <- export.Item{
					ID:    item.ID(),
					Error: err,
				}

				continue
			}

			reader := item.ToReader()
			content, err := io.ReadAll(reader)

			reader.Close()

			if err != nil {
				ch <- export.Item{
					ID:    item.ID(),
					Error: err,
				}

				continue
			}

			// Convert JSON to eml.
			email, err := eml.FromJSONPostToEML(ictx, content, postMetadata)
			if err != nil {
				err = clues.Wrap(err, "converting JSON to eml")

				logger.CtxErr(ictx, err).Info("processing collection item")

				ch <- export.Item{
					ID:    item.ID(),
					Error: err,
				}

				continue
			}

			emlReader := io.NopCloser(bytes.NewReader([]byte(email)))

			stats.UpdateResourceCount(path.ConversationPostsCategory)
			body := metrics.ReaderWithStats(emlReader, path.ConversationPostsCategory, stats)

			ch <- export.Item{
				ID:   item.ID(),
				Name: exportName,
				Body: body,
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

func fetchAndReadMetadata(
	ctx context.Context,
	itemID string,
	fin data.FetchItemByNamer,
) (groupMeta.ConversationPostMetadata, error) {
	metaName := itemID + metadata.MetaFileSuffix

	ctx = clues.Add(ctx, "meta_file_name", metaName)

	meta, err := fin.FetchItemByName(ctx, metaName)
	if err != nil {
		return groupMeta.ConversationPostMetadata{},
			clues.WrapWC(ctx, err, "fetching metadata")
	}

	metaReader := meta.ToReader()
	defer metaReader.Close()

	metaFormatted, err := readMetadata(metaReader)
	if err != nil {
		return groupMeta.ConversationPostMetadata{},
			clues.WrapWC(ctx, err, "deserializing metadata")
	}

	return metaFormatted, nil
}

// getMetadata reads and parses the metadata info for an item
func readMetadata(metaRC io.ReadCloser) (groupMeta.ConversationPostMetadata, error) {
	var meta groupMeta.ConversationPostMetadata

	metaraw, err := io.ReadAll(metaRC)
	if err != nil {
		return groupMeta.ConversationPostMetadata{}, err
	}

	err = json.Unmarshal(metaraw, &meta)
	if err != nil {
		return groupMeta.ConversationPostMetadata{}, err
	}

	return meta, nil
}

package teamschats

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
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func NewExportCollection(
	baseDir string,
	backingCollections []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	stats *metrics.ExportStats,
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
	stats *metrics.ExportStats,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		for item := range rc.Items(ctx, errs) {
			body, err := formatChat(cec, item.ToReader())
			if err != nil {
				ch <- export.Item{
					ID:    item.ID(),
					Error: err,
				}
			} else {
				stats.UpdateResourceCount(path.ChatsCategory)
				body = metrics.ReaderWithStats(body, path.ChatsCategory, stats)

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
	minimumChat struct {
		CreatedDateTime     time.Time            `json:"createdDateTime"`
		LastUpdatedDateTime time.Time            `json:"lastUpdatedDateTime"`
		Topic               string               `json:"topic"`
		Messages            []minimumChatMessage `json:"replies,omitempty"`
		Members             []minimumChatMember  `json:"members"`
	}

	minimumChatMember struct {
		Name                    string    `json:"name"`
		VisibleHistoryStartedAt time.Time `json:"visibleHistoryStartedAt"`
	}

	minimumChatMessage struct {
		Attachments          []minimumAttachment `json:"attachments"`
		Content              string              `json:"content"`
		CreatedDateTime      time.Time           `json:"createdDateTime"`
		From                 string              `json:"from"`
		LastModifiedDateTime time.Time           `json:"lastModifiedDateTime"`
		IsDeleted            bool                `json:"isDeleted"`
	}

	minimumAttachment struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func formatChat(
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

	cfb, err := api.CreateFromBytes(bs, models.CreateChatFromDiscriminatorValue)
	if err != nil {
		return nil, clues.Wrap(err, "deserializing bytes to message")
	}

	chat, ok := cfb.(models.Chatable)
	if !ok {
		return nil, clues.New("expected deserialized item to implement models.Chatable")
	}

	var (
		members  = chat.GetMembers()
		messages = chat.GetMessages()
	)

	result := minimumChat{
		CreatedDateTime:     ptr.Val(chat.GetCreatedDateTime()),
		LastUpdatedDateTime: ptr.Val(chat.GetLastUpdatedDateTime()),
		Topic:               ptr.Val(chat.GetTopic()),
		Members:             make([]minimumChatMember, 0, len(members)),
		Messages:            make([]minimumChatMessage, 0, len(messages)),
	}

	for _, r := range messages {
		result.Messages = append(result.Messages, makeMinimumChatMessage(r))
	}

	for _, r := range members {
		result.Members = append(result.Members, makeMinimumChatMember(r))
	}

	bs, err = marshalJSONContainingHTML(result)
	if err != nil {
		return nil, clues.Wrap(err, "serializing minimized chat")
	}

	return io.NopCloser(bytes.NewReader(bs)), nil
}

func makeMinimumChatMessage(item models.ChatMessageable) minimumChatMessage {
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

	var isDeleted bool

	deletedAt, ok := ptr.ValOK(item.GetDeletedDateTime())
	isDeleted = ok && deletedAt.After(time.Time{})

	return minimumChatMessage{
		Attachments:          minAttachments,
		Content:              content,
		CreatedDateTime:      ptr.Val(item.GetCreatedDateTime()),
		From:                 api.GetChatMessageFrom(item),
		LastModifiedDateTime: ptr.Val(item.GetLastModifiedDateTime()),
		IsDeleted:            isDeleted,
	}
}

func makeMinimumChatMember(item models.ConversationMemberable) minimumChatMember {
	return minimumChatMember{
		Name:                    ptr.Val(item.GetDisplayName()),
		VisibleHistoryStartedAt: ptr.Val(item.GetVisibleHistoryStartDateTime()),
	}
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

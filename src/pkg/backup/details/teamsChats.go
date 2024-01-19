package details

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/path"
)

// NewChatsLocationIDer builds a LocationIDer for the chats.
func NewChatsLocationIDer(
	category path.CategoryType,
	escapedFolders ...string,
) (uniqueLoc, error) {
	if err := path.ValidateServiceAndCategory(path.TeamsChatsService, category); err != nil {
		return uniqueLoc{}, clues.Wrap(err, "making chats LocationIDer")
	}

	pb := path.Builder{}.Append(category.String()).Append(escapedFolders...)

	return uniqueLoc{
		pb:          pb,
		prefixElems: 1,
	}, nil
}

// TeamsChatsInfo describes a chat within teams chats.
type TeamsChatsInfo struct {
	ItemType   ItemType  `json:"itemType,omitempty"`
	Modified   time.Time `json:"modified,omitempty"`
	ParentPath string    `json:"parentPath,omitempty"`

	Chat ChatInfo `json:"chat,omitempty"`
}

type ChatInfo struct {
	CreatedAt          time.Time `json:"createdAt,omitempty"`
	HasExternalMembers bool      `json:"hasExternalMemebers,omitempty"`
	LastMessageAt      time.Time `json:"lastMessageAt,omitempty"`
	LastMessagePreview string    `json:"preview,omitempty"`
	Members            []string  `json:"members,omitempty"`
	MessageCount       int       `json:"size,omitempty"`
	Name               string    `json:"name,omitempty"`
}

// Headers returns the human-readable names of properties in a ChatsInfo
// for printing out to a terminal in a columnar display.
func (i TeamsChatsInfo) Headers() []string {
	switch i.ItemType {
	case TeamsChat:
		return []string{"Name", "Last message", "Last message at", "Message count", "Created", "Members"}
	}

	return []string{}
}

// Values returns the values matching the Headers list for printing
// out to a terminal in a columnar display.
func (i TeamsChatsInfo) Values() []string {
	switch i.ItemType {
	case TeamsChat:
		members := ""
		icmLen := len(i.Chat.Members)

		if icmLen > 0 {
			members = i.Chat.Members[0]
		}

		if icmLen > 1 {
			members = fmt.Sprintf("%s, and %d more", members, icmLen-1)
		}

		return []string{
			i.Chat.Name,
			i.Chat.LastMessagePreview,
			dttm.FormatToTabularDisplay(i.Chat.LastMessageAt),
			strconv.Itoa(i.Chat.MessageCount),
			dttm.FormatToTabularDisplay(i.Chat.CreatedAt),
			members,
		}
	}

	return []string{}
}

func (i *TeamsChatsInfo) UpdateParentPath(newLocPath *path.Builder) {
	i.ParentPath = newLocPath.String()
}

func (i *TeamsChatsInfo) uniqueLocation(baseLoc *path.Builder) (*uniqueLoc, error) {
	var category path.CategoryType

	switch i.ItemType {
	case TeamsChat:
		category = path.ChatsCategory
	}

	loc, err := NewChatsLocationIDer(category, baseLoc.Elements()...)

	return &loc, err
}

func (i *TeamsChatsInfo) updateFolder(f *FolderInfo) error {
	// Use a switch instead of a rather large if-statement. Just make sure it's an
	// Exchange type. If it's not return an error.
	switch i.ItemType {
	case TeamsChat:
	default:
		return clues.New("unsupported non-Chats ItemType").
			With("item_type", i.ItemType)
	}

	f.DataType = i.ItemType

	return nil
}

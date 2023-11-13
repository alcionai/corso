package graph

import (
	"time"
)

type GetIDer interface {
	GetId() *string
}

type GetLastModifiedDateTimer interface {
	GetLastModifiedDateTime() *time.Time
}

type GetAdditionalDataer interface {
	GetAdditionalData() map[string]any
}

type GetDeletedDateTimer interface {
	GetDeletedDateTime() *time.Time
}

type GetDisplayNamer interface {
	GetDisplayName() *string
}

type GetParentFolderIDer interface {
	GetParentFolderId() *string
}

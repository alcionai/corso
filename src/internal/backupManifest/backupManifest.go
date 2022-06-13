package backupManifest

import (
	_ "modernc.org/sqlite"
)

type BackEntityMetadata interface {
}

//implements BackEntityMetadata
type MessageMetadata struct {
	attachment []AttachmentMetadata
}

//implements BackEntityMetadata
type EventsMetadata struct {
	attachment []AttachmentMetadata
}

//implements BackEntityMetadata
type ContactMetadata struct {
}

type AttachmentMetadata struct {
}

type BackManifestHandler interface {
	Open(path string) error
	Insert(bem BackEntityMetadata) error
	Search(callbck func(bem BackEntityMetadata, err error) error, filters ...string) error
	Close() error
	Delete() error
}

//implements BackManifestHandler
type ExchangeBackManifestHandler struct {
}

package backupManifest

import (
	_ "modernc.org/sqlite"
)

type BackEntityMetadata interface {
}

//implements BackEntityMeta
type MessageMetadata struct {
	attachment []AttachmentMetadata
}

//implements BackEntityMeta
type EventsMetadata struct {
	attachment []AttachmentMetadata
}

//implements BackEntityMeta
type ContactMetadata struct {
}

type AttachmentMetadata struct {
}

type BackManifestHandler interface {
	Open(path string) error
	Insert(bem BackEntityMeta) error
	Search(calbck func(bem BackEntityMeta, err error) error, filters ...string) error
	Close() error
	Delete() error
}

//implements BackManifestHandler
type ExchangeBackManifestHandler struct {
}

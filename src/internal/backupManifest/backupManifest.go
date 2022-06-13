package backupManifest

import (
	_ "modernc.org/sqlite"
)

type BackEntityMeta interface {
}

//implements BackEntityMeta
type MessageMeta struct {
}

//implements BackEntityMeta
type AttachmentMeta struct {
}

//implements BackEntityMeta
type EventsMeta struct {
}

//implements BackEntityMeta
type ContactMeta struct {
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

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
	//Open the database
	Open(path string) error
	//Insert into the database
	Insert(bem BackEntityMetadata) error
	//Search in the database with filters
	Search(callbck func(bem BackEntityMetadata, err error) error, filters ...string) error
	//Close the database
	Close() error
	//Destory the database
	Destroy() error
}

//implements BackManifestHandler
type ExchangeBackManifestHandler struct {
}

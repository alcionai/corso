package backupManifest

import (
	"database/sql"
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

/* For Future reference
//implements BackEntityMetadata
type FileMetadata struct {
}

//implements BackEntityMetadata
type ObjectMetadata struct {
}
*/

type AttachmentMetadata struct {
}

type searchCallBack func(bem BackEntityMetadata, err error)

type BackManifestHandler interface {
	//Open the database
	Open(name string) error
	//Insert into the database
	Insert(bem BackEntityMetadata) error
	//Search in the database with filters
	Search(callbck searchCallBack, filters ...string) error
	//Close the database
	Close() error
	//Destory the database
	Destroy() error
}

//implements BackManifestHandler
type ExchangeBackupManifestHandler struct {
	db sql.DB
}

func NewExchangeBackupManifestHandler(
	path string,
	callbck searchCallBack) (ExchangeBackupManifestHandler, error) {

	return ExchangeBackupManifestHandler{}, nil
}

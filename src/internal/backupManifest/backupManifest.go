package backupManifest

import (
	"database/sql"
)

type BackEntityMetadata interface {
	GetID() string
	//Will fill out other common getter functions here.
}

//implements BackEntityMetadata
type MessageMetadata struct {
	//Message specific fields 
	attachment []AttachmentMetadata
}

//implements BackEntityMetadata
type EventsMetadata struct {
	//Event specific fields
	attachment []AttachmentMetadata
}

//implements BackEntityMetadata
type ContactMetadata struct {
	//Contact specific fields
}

/* For Future reference
//implements BackEntityMetadata
type FileMetadata struct {
	//File specific fields
}

//implements BackEntityMetadata
type ObjectMetadata struct {
	//Object specific fields
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
	Search(callback searchCallBack, filters ...string) error
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

package backupmanifest

import (
	"database/sql"
)

type BackupEntityMetadata interface {
	GetID() string
	Insert(bemh *BackupManifestHandler) error
}

/*******************************************************************

							MessageMetadata

********************************************************************/

//implements BackEntityMetadata
type MessageMetadata struct {
	//Message specific fields
	attachment []AttachmentMetadata
}

func (mm *MessageMetadata) GetID() string {
	//TODO
	return ""
}

func (mm *MessageMetadata) Insert(bemh *BackupManifestHandler) error {
	var ebmh *ExchangeBackupManifestHandler
	ebmh = bemh
	return
}

/*******************************************************************

							EventsMetadata

********************************************************************/

//implements BackEntityMetadata
type EventsMetadata struct {
	//Event specific fields
	attachment []AttachmentMetadata
}

func (em *EventsMetadata) GetID() string {
	//TODO
	return ""
}

/*******************************************************************

							ContactMetadata

********************************************************************/

//implements BackEntityMetadata
type ContactMetadata struct {
	//Event specific fields
	attachment []AttachmentMetadata
}

func (em *ContactMetadata) GetID() string {
	//TODO
	return ""
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

/*******************************************************************

						BackManifestHandler

********************************************************************/

type FilterMap map[string]string

type BackupManifestHandler interface {
	//Open the database
	Open(name string) error
	//Close the database
	Close() error
	//Destory the database
	Destroy() error
	//Abstarct Insert
	Insert(bem *BackupEntityMetadata) error
}

/*******************************************************************

						ExchangeBackupManifestHandler

********************************************************************/

//implements BackManifestHandler
type ExchangeBackupManifestHandler struct {
	exchangeSearchQueryMap map[int]string
	db                     sql.DB
}

type ExchangeSearchIterator struct {
	rows *sql.Rows
}

//Insert MessageMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertMessageMetadata(mm MessageMetadata) error {
	return nil
}

//Insert EventMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertEventMetadata(mm EventMetadata) error {
	return nil
}

//Insert EventMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertContactMetadata(mm ContactMetadata) error {
	return nil
}

//Search
func (ebmh *ExchangeBackupManifestHandler) SearchMessageMetadata() error {

}

//Etc..more to be added according to List/Restore requirement

func NewExchangeBackupManifestHandler(path string) (ExchangeBackupManifestHandler, error) {
	return ExchangeBackupManifestHandler{}, nil
}

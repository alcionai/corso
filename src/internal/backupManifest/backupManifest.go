package backupmanifest

import (
	"database/sql"
)

type BackEntityMetadata interface {
	GetID() string
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
	//Validate Search Query Index And Filters
	validateSearchQueryIndexAndFilters(searchQueryPrefixIndex []int, XXXfilters []string) error
	//Get Search Query Prefix using index
	getSearchQueryPrefixString(searchQueryPrefixIndex int) string
	//Open the database
	Open(name string) error
	//Close the database
	Close() error
	//Destory the database
	Destroy() error
}

/*******************************************************************

						ExchangeBackupManifestHandler

********************************************************************/

//implements BackManifestHandler
type ExchangeBackupManifestHandler struct {
	exchangeSearchQueryMap map[int]string
	db                     sql.DB
}

//Insert into the database
func (ebmh *ExchangeBackupManifestHandler) InsertMessageMetadata(mm MessageMetadata) error {
	return nil
}

//Etc..more to be added according to List/Restore requirement

func NewExchangeBackupManifestHandler(
	path string,
	callbck searchCallBack) (ExchangeBackupManifestHandler, error) {
	return ExchangeBackupManifestHandler{}, nil
}

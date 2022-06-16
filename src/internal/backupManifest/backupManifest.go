package backupManifest

import (
	"database/sql"
)

type BackEntityMetadata interface {
	GetID() string
	Insert(db *sql.DB) error
	readCallBack(callback searchCallBack, row sql.Row) error // call the user call back internally
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

func (mm *MessageMetadata) Insert(db *sql.DB) error {
	//TODO
	return nil
}

func (mm *MessageMetadata) readCallBack(callback searchCallBack, row sql.Row) error {
	//TODO
	return nil
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

func (em *EventsMetadata) Insert(db *sql.DB) error {
	//TODO
	return nil
}

func (em *EventsMetadata) readCallBack(callback searchCallBack, row sql.Row) error {
	//TODO
	return nil
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

func (em *ContactMetadata) Insert(db *sql.DB) error {
	//TODO
	return nil
}

func (em *ContactMetadata) readCallBack(callback searchCallBack, row sql.Row) error {
	//TODO
	return nil
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

/*******************************************************************

						BackManifestHandler

********************************************************************/

type BackManifestHandler interface {
	//Validate Search Query Index And Filters
	validateSearchQueryIndexAndFilters(searchQueryPrefixIndex []int, XXXfilters []string) error
	//Get Search Query Prefix using index
	getSearchQueryPrefixString(searchQueryPrefixIndex int) string
	//Open the database
	Open(name string) error
	//Insert into the database
	Insert(bem BackEntityMetadata) error
	//Search in the database with filters
	/*

		THIS COMMENT IS ONLY FOR DESIGN UNDERSTANDING!


		The user calls Search() with a call back function.
		//Pseudo code for a XXXBackupManifestHandler which handles metadata of
		// XXXMetadataType1, XXXMetadataType2, XXXMetadataType3 etc
		func (XBMH *XXXBackupManifestHandler) Search (callback searchCallBack, searchQueryPrefixIndex []int, XXXfilters []string) error {

			//Validation of searchQueryPrefixIndex and XXXfilters for the appropriate XXXBackupManifestHandler
			err := validateXXXSearchQueryIndexAndFilters (searchQueryPrefixIndex, XXXfilters)
			if err != nil {
				return err
			}

			for sqIndex in searchQueryPrefixIndex {
				var actualSearchQuery []string
				actualSearchQuery = getSearchQueryPrefixString(sqIndex)
				append(actualSearchQuery, " where ")
				for filter in XXXfilters {
					//Only append filter of appropriate to the search query/XXXMetadata type
					if (isFilterOfType(filter, sqIndex)) {
						append(actualSearchQuery, filter)
					}
				}

				rows = db.Query (actualSearchQuery)
				for row in rows {
					switch (sqIndex) {
						case type1:
							XXXMetadataType1 := NewXXXMetadataType1()
							XXXMetadataType1.readCBK(callback, row) // This call the users callback function for each resultant row
						case type2:
							XXXMetadataType1 := NewXXXMetadataType2(row, callback)
							XXXMetadataType1.readCBK(callback) // This call the users callback function for each resultant row
						case type3:
							....
					}
				}
			}

		}

	*/
	Search(callback searchCallBack, searchQueryIndex []int, filters []string) error
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

// Exchange Search query constant for exchange
const exchangeSearchQueryAllMessagesPrefixIndex int = 1 // Messages type
const exchangeSearchQueryAllEventsPrefixIndex int = 2   // Event type
const exchangeSearchQueryAllContactPrefixIndex = 3      //Contact type

const exchangeSearchQueryAllMessagesPrefixString string = "SELECT * FROM message MESSAGE"
const exchangeSearchQueryStringAllEventsPrefixString string = "SELECT * FROM event EVENTS"
const exchangeSearchQueryStringAllContactsPrefixString string = "SELECT * FROM contact CONTACTS"

// Exchange Search query filter constants examples:
const exchangeSearchQueryFilterMessageSender string = "message.sender"
const exchangeSearchQueryFilterEventSender string = "event.sender"

//Etc..more to be added according to List/Restore requirement

func NewExchangeBackupManifestHandler(
	path string,
	callbck searchCallBack) (ExchangeBackupManifestHandler, error) {
	return ExchangeBackupManifestHandler{}, nil
}

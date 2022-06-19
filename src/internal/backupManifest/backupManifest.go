package backupmanifest

import (
	"database/sql"
)

type BackupEntityMetadata interface {
	GetID() string
	//Abstract Insert
	Insert(bemh BackupManifestHandler) error
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

func (mm *MessageMetadata) Insert(bemh BackupManifestHandler) error {
	var ebmh *ExchangeBackupManifestHandler
	ebmh = bemh.(*ExchangeBackupManifestHandler)
	return ebmh.InsertMessageMetadata(mm)
}

/*******************************************************************

							EventsMetadata

********************************************************************/

//implements BackEntityMetadata
type EventMetadata struct {
	//Event specific fields
	attachment []AttachmentMetadata
}

func (em *EventMetadata) GetID() string {
	//TODO
	return ""
}

func (em *EventMetadata) Insert(bemh BackupManifestHandler) error {
	var ebmh *ExchangeBackupManifestHandler
	ebmh = bemh.(*ExchangeBackupManifestHandler)
	return ebmh.InsertEventMetadata(em)
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

func (cm *ContactMetadata) Insert(bemh BackupManifestHandler) error {
	var ebmh *ExchangeBackupManifestHandler
	ebmh = bemh.(*ExchangeBackupManifestHandler)
	return ebmh.InsertContactMetadata(cm)
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
	Open() error
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
	dbPath string
	db     sql.DB
}

//Abstarct Insert
func (ebmh *ExchangeBackupManifestHandler) Insert(bem *BackupEntityMetadata) error {
	var _bem BackupEntityMetadata
	_bem = *bem
	return _bem.Insert(ebmh)
}

//Open the database
func (ebmh *ExchangeBackupManifestHandler) Open() error {
	return nil
}

//Close the database
func (ebmh *ExchangeBackupManifestHandler) Close() error {
	return nil
}

//Destory the database
func (ebmh *ExchangeBackupManifestHandler) Destroy() error {
	return nil
}

//Insert MessageMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertMessageMetadata(mm *MessageMetadata) error {
	return nil
}

//Insert EventMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertEventMetadata(em *EventMetadata) error {
	return nil
}

//Insert EventMetadata into the database
func (ebmh *ExchangeBackupManifestHandler) InsertContactMetadata(cm *ContactMetadata) error {
	return nil
}

//Search Message Metadata
func (ebmh *ExchangeBackupManifestHandler) SearchMessageMetadata(filterMap FilterMap) (MessageSearchIterator, error) {
	return MessageSearchIterator{}, nil
}

//Search Event Metadata
func (ebmh *ExchangeBackupManifestHandler) SearchEventMetadata(filterMap FilterMap) (EventSearchIterator, error) {
	return EventSearchIterator{}, nil
}

//Search Contact Metadata
func (ebmh *ExchangeBackupManifestHandler) SearchContactMetadata(filterMap FilterMap) (ContactSearchIterator, error) {
	return ContactSearchIterator{}, nil
}

func NewExchangeBackupManifestHandler(path string) (ExchangeBackupManifestHandler, error) {
	return ExchangeBackupManifestHandler{}, nil
}

/*******************************************************************

                ExchangeSearchIterators

********************************************************************/

type ExchangeSearchIterator struct {
	rows *sql.Rows
}

type MessageSearchIterator struct {
	ExchangeSearchIterator
}

type EventSearchIterator struct {
	ExchangeSearchIterator
}

type ContactSearchIterator struct {
	ExchangeSearchIterator
}

func (msi *MessageSearchIterator) NextMessage() (MessageMetadata, error) {
	return MessageMetadata{}, nil
}

func (esi *EventSearchIterator) NextEvent() (EventMetadata, error) {
	return EventMetadata{}, nil
}

func (csi *ContactSearchIterator) NextContact() (ContactMetadata, error) {
	return ContactMetadata{}, nil
}

/*
Pseudo Code for List Messages

	//ExchangeBackupManifestJSONHandler not get defined above, just can example

	ebmjh ExchangeBackupManifestJSONHandler := NewExchangeBackupManifestJSONHandler ("list_snapshotID.json")
	ebmh ExchangeBackupManifestHandler := NewExchangeBackupManifestHandler ("snapshotID.db")

	err := ebmh.Open()
	defer ebmh.Close()
	if err != nil {
		//go crazy
	}


	err := ebmjh.Open()
	defer ebmjh.Close()
	if err != nil {
		//go crazy
	}


	filterMap FilterMap
	filterMap["ReceiveTimeStart"] = "XXX:XXX:XXXX:XXX"
	filterMap["ReceiveTimeEnd"] = "XXX:XXXX:XXXX:XXXX"
	filterMap["From"] = "mogambo@dongrila.com"


	msi, err = ebmh.SearchMessageMetadata (filterMap)
	if err != nil {
		//go crazy
	}

	for {
		mm,err := msi.NextMessage()
		if err != nil {
			if err == io.EOF {
				return
			} else {
				//go crazy
			}
		}

		//Add to JSON file
		err := ebmjh.AppendMessageMetadata (mm)
		if err != nil {
			//go crazy
		}

	}

	return nil

*/

package connector

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

//Local filesytem source
var CORSO_USER string = "corso"
var CORSO_PATH string = "./corso_sample"
var CORSO_SAMPLE_DATA_FILE string = "./sample.data"
var MAX_FILE_ID int = 10

// A DataCollection represents a collection of data of the
// same type (e.g. mail)
type DataCollection interface {

	//Tells the current position in Data Collection
	CurPosition() int

	//Reset the current position pointer in collection
	RestItemCur()

	//Tells if Data Collection is empty
	isEmpty() bool

	// Returns either the next item in the collection or an error if one occurred.
	// If not more items are available in the collection, returns (nil, nil).
	NextItem() (DataStream, error)

	//Initializs DataCollection
	InitDataCollection(string) error

	//Close DataCollection
	CloseDataCollection() error
}

// DataStream represents a single item within a DataCollection
// that can be consumed as a stream (it embeds io.Reader)
type DataStream interface {
	// Provides a unique identifier for this data
	UUID() string
	// Reads from data reader
	Read(bytes []byte) (int, error)
}

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type ExchangeDataCollection struct {

	//Data Collection is initialized
	isInit bool

	//curent iteration number
	curIndex int

	user string
	// TODO: We would want to replace this with a channel so that we
	// don't need to wait for all data to be retrieved before reading it out
	data []ExchangeData
}

// ExchangeData represents a single item retrieved from exchange
type ExchangeData struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	//message []

	//Ondisk file where emails will be saved
	file os.File

	//ondisk file data reader
	io io.Reader
}

// NextItem returns either the next item in the collection or an error if one occurred.
// If not more items are available in the collection, returns (nil, io.EOF).
func (eDataCol *ExchangeDataCollection) NextItem() (DataStream, error) {

	if eDataCol.data == nil {
		return nil, io.EOF
	}

	if eDataCol.curIndex < len(eDataCol.data) {
		ed := &eDataCol.data[eDataCol.curIndex]
		eDataCol.curIndex++
		return ed, nil
	}
	return nil, io.EOF
}

//Reset the DataCollection current pointer to first
func (eDataCol *ExchangeDataCollection) RestItemCur() {
	eDataCol.curIndex = 0
}

// Tells if ExchangeDataCollection is empty
func (eDatacol *ExchangeDataCollection) isEmpty() bool {
	return eDatacol.data == nil
}

// Tells the current position in Data Collection
func (eDatacol *ExchangeDataCollection) CurPosition() int {
	return eDatacol.curIndex
}

//Initialize the Data Collection
func (eDataCol *ExchangeDataCollection) InitDataCollection(user string) error {
	switch {
	//Sample path
	case user == CORSO_USER:
		eDataCol.user = CORSO_USER
		return eDataCol.InitDataCollectionSample(CORSO_PATH)
	default:
		return errors.New("invalid user")
	}
}

//Close Data Collection
func (eDataCol *ExchangeDataCollection) CloseDataCollection() error {

	if !eDataCol.isInit {
		return errors.New("DataCollection not initialized ")
	}

	switch {
	//Sample path
	case eDataCol.user == CORSO_USER:
		return eDataCol.ClearDataCollectionSample(CORSO_PATH)
	default:
		return errors.New("invalid user")
	}

}

func (ed *ExchangeData) UUID() string {
	return ed.id
}

func (ed *ExchangeData) Read(bytes []byte) (int, error) {
	return ed.io.Read(bytes)
}

/******************************************************************************

						Exchange Data Sample Code

*******************************************************************************/

func fileCopy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// populate local source fs
func (eDataCol *ExchangeDataCollection) populateCollectionSample(path string) error {

	// Create sample directory
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create files
	for i := 0; i < MAX_FILE_ID; i++ {

		//Create file
		file_name := fmt.Sprintf("%v/%d", path, i)
		// Print name.
		fmt.Println("Creating file : ", file_name)
		//Create file content from CORSO_SAMPLE_DATA_FILE
		if size, err := fileCopy(CORSO_SAMPLE_DATA_FILE, file_name); err != nil || size == 0 {
			eDataCol.depopulateCollectionSample(path)
			return err
		}

	}

	return nil
}

// depopulate local source fs
func (eDataCol *ExchangeDataCollection) depopulateCollectionSample(path string) error {

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	// Print name.
	fmt.Println("Deleting folder: ", path)
	err = os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func (eDataCol *ExchangeDataCollection) ClearDataCollectionSample(path string) error {

	//Close all open files
	for index := range eDataCol.data {
		eData := eDataCol.data[index]
		eData.file.Close()
	}

	//Clear the slice
	eDataCol.data = nil

	//Delete the source path
	return eDataCol.depopulateCollectionSample(path)
}

// Initiate Data Collection for Sample data
func (eDataCol *ExchangeDataCollection) InitDataCollectionSample(path string) error {

	//Clear any data that is present
	eDataCol.depopulateCollectionSample(path)

	//populate the sample path
	err := eDataCol.populateCollectionSample(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Open the directory.
	outputDirRead, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Call ReadDir to get all files.
	outputDirFiles, err := outputDirRead.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Loop over files.
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Get name of file.
		outputNameHere := outputFileHere.Name()

		//file path
		filePath := fmt.Sprintf("%v/%v", path, outputNameHere)

		// Print name.
		fmt.Println("Opening file : ", outputNameHere)

		file, err := os.Open(filePath)
		if err != nil {
			eDataCol.ClearDataCollectionSample(path)
			return err
		}

		io := bufio.NewReader(file)

		eData := new(ExchangeData)
		eData.file = *file
		eData.io = io
		eData.id = filePath

		eDataCol.data = append(eDataCol.data, *eData)

	}

	eDataCol.isInit = true
	eDataCol.curIndex = 0

	return nil
}

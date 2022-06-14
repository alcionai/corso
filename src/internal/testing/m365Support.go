package testing

import (
	"bufio"
	"os"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	js "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CreateFromBytes helper function to initialize m365 object form bytes.
// @param bytes -> source, createFunc -> abstract function for initialization
func CreateFromBytes(bytes []byte, createFunc absser.ParsableFactory) (absser.Parsable, error) {
	parseNode, err := js.NewJsonParseNodeFactory().GetRootParseNode("application/json", bytes)
	if err != nil {
		return nil, err
	}

	anObject, err := parseNode.GetObjectValue(createFunc)
	if err != nil {
		return nil, err
	}
	return anObject, nil
}

// CreateMessageFromBytes function to transform bytes into Messageable object
func CreateMessageFromBytes(bytes []byte) (models.Messageable, error) {

	aMessage, err := CreateFromBytes(bytes, models.CreateMessageFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	message := aMessage.(models.Messageable)
	return message, nil
}

func LoadAFile(aFile string) ([]byte, error) {
	// Preserves '\n' of original file. Uses incremental version when file too large
	bytes, err := os.ReadFile(aFile)
	if err != nil {
		f, err := os.Open(aFile)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		buffer := make([]byte, 0)
		reader := bufio.NewScanner(f)
		for reader.Scan() {
			temp := reader.Bytes()
			buffer = append(buffer, temp...)
		}
		aErr := reader.Err()
		if aErr != nil {
			return nil, aErr
		}
		return buffer, nil
	}
	return bytes, nil
}

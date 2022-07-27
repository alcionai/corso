package support

import (
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	js "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// TaskList is a a generic map of a list of items with a string index
type TaskList map[string][]string

// NewTaskList constructor for TaskList
func NewTaskList() TaskList {
	return make(map[string][]string, 0)
}

// AddTask helper method to ensure that keys and items are created properly
func (tl *TaskList) AddTask(key, value string) {
	aMap := *tl
	_, isCreated := aMap[key]
	if isCreated {
		aMap[key] = append(aMap[key], value)
	} else {
		aMap[key] = []string{value}
	}
}

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

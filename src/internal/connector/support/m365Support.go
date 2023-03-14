package support

import (
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	js "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	bmodels "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CreateFromBytes helper function to initialize m365 object form bytes.
// @param bytes -> source, createFunc -> abstract function for initialization
func CreateFromBytes(bytes []byte, createFunc absser.ParsableFactory) (absser.Parsable, error) {
	parseNode, err := js.NewJsonParseNodeFactory().GetRootParseNode("application/json", bytes)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes into base m365 object")
	}

	anObject, err := parseNode.GetObjectValue(createFunc)
	if err != nil {
		return nil, errors.Wrap(err, "parsing m365 object factory")
	}

	return anObject, nil
}

// CreateMessageFromBytes function to transform bytes into Messageable object
func CreateMessageFromBytes(bytes []byte) (models.Messageable, error) {
	aMessage, err := CreateFromBytes(bytes, models.CreateMessageFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes to exchange message")
	}

	message := aMessage.(models.Messageable)

	return message, nil
}

// CreateContactFromBytes function to transform bytes into Contactable object
// Error returned if ParsableFactory function does not accept given bytes
func CreateContactFromBytes(bytes []byte) (models.Contactable, error) {
	parsable, err := CreateFromBytes(bytes, models.CreateContactFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes to exchange contact")
	}

	contact := parsable.(models.Contactable)

	return contact, nil
}

// CreateEventFromBytes transforms given bytes into models.Eventable object
func CreateEventFromBytes(bytes []byte) (models.Eventable, error) {
	parsable, err := CreateFromBytes(bytes, models.CreateEventFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes to exchange event")
	}

	event := parsable.(models.Eventable)

	return event, nil
}

// CreateListFromBytes transforms given bytes into models.Listable object
func CreateListFromBytes(bytes []byte) (models.Listable, error) {
	parsable, err := CreateFromBytes(bytes, models.CreateListFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes to sharepoint list")
	}

	list := parsable.(models.Listable)

	return list, nil
}

// CreatePageFromBytes transforms given bytes in models.SitePageable object
func CreatePageFromBytes(bytes []byte) (bmodels.SitePageable, error) {
	parsable, err := CreateFromBytes(bytes, bmodels.CreateSitePageFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, "deserializing bytes to sharepoint page")
	}

	page := parsable.(bmodels.SitePageable)

	return page, nil
}

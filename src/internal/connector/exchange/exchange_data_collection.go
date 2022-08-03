// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"io"

	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/logger"
)

var _ data.Collection = &Collection{}
var _ data.Stream = &Stream{}
var _ data.StreamInfo = &Stream{}

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
	// RestorePropertyTag defined: https://docs.microsoft.com/en-us/office/client-developer/outlook/mapi/pidtagmessageflags-canonical-property
	RestorePropertyTag          = "Integer 0x0E07"
	RestoreCanonicalEnableValue = "4"
)

// Collection implements the interface from data.Collection
// Structure holds data for an Exchange application for a single user
type Collection struct {
	// M365 user
	user string // M365 user
	data chan data.Stream
	// jobs represents items from the inventory of M365 objectIds whose information
	// is desired to be sent through the data channel for eventual storage
	jobs []string
	// service - client/adapter pair used to access M365 back store
	service graph.Service
	// populate - Utility function to populate collection based on the M365 application type and granularity
	populate populater
	statusCh chan<- *support.ConnectorOperationStatus
	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// Populater are a class of functions that can be used to fill exchange.Collections with
// the corresponding information
type populater func(context.Context, graph.Service, *Collection, chan<- *support.ConnectorOperationStatus)

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewCollection(
	user string,
	fullPath []string,
	service graph.Service,
	aUser string,
	pathRepresentation []string,
	collectionType optionIdentifier,
	aService graph.Service,
	statusCh chan<- *support.ConnectorOperationStatus,
) Collection {
	collection := Collection{
		user:     user,
		data:     make(chan data.Stream, collectionChannelBufferSize),
		jobs:     make([]string, 0),
		service:  service,
		statusCh: statusCh,
		fullPath: fullPath,
		populate: getPopulateFunction(collectionType),
	}
	return collection
}

// getPopulateFunction is a function to set populate function field
// with exchange-application specific functions
func getPopulateFunction(optId optionIdentifier) populater {
	switch optId {
	case messages:
		return PopulateFromCollection
	default:
		return nil
	}
}

// AddJob appends additional objectID to structure's jobs field
func (eoc *Collection) AddJob(objID string) {
	eoc.jobs = append(eoc.jobs, objID)
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (eoc *Collection) Items() <-chan data.Stream {
	if eoc.populate != nil {
		go eoc.populate(context.TODO(), eoc.service, eoc, eoc.statusCh)
	}
	return eoc.data
}

// FullPath returns the Collection's fullPath []string
func (edc *Collection) FullPath() []string {
	return append([]string{}, edc.fullPath...)
}

// populateFromTaskList async call to fill DataCollection via channel implementation
func PopulateFromCollection(
	ctx context.Context,
	service graph.Service,
	edc *Collection,
	statusChannel chan<- *support.ConnectorOperationStatus,
) {
	var errs error
	var attemptedItems, success int
	objectWriter := kw.NewJsonSerializationWriter()

	for _, task := range edc.jobs {
		response, err := service.Client().UsersById(edc.user).MessagesById(task).Get()
		if err != nil {
			details := support.ConnectorStackErrorTrace(err)
			errs = support.WrapAndAppend(edc.user, errors.Wrapf(err, "unable to retrieve item %s; details %s", task, details), errs)
			continue
		}
		err = messageToDataCollection(service.Client(), ctx, objectWriter, edc.data, response, edc.user)
		success++
		if err != nil {
			errs = support.WrapAndAppendf(edc.user, err, errs)
			success--
		}
		if errs != nil && service.ErrPolicy() {
			break
		}
	}
	close(edc.data)
	attemptedItems += len(edc.jobs)

	status := support.CreateStatus(ctx, support.Backup, attemptedItems, success, 1, errs)
	logger.Ctx(ctx).Debug(status.String())
	statusChannel <- status
}

func messageToDataCollection(
	client *msgraphsdk.GraphServiceClient,
	ctx context.Context,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	message models.Messageable,
	user string,
) error {
	var err error
	aMessage := message
	adtl := message.GetAdditionalData()
	if len(adtl) > 2 {
		aMessage, err = support.ConvertFromMessageable(adtl, message)
		if err != nil {
			return err
		}
	}
	if *aMessage.GetHasAttachments() {
		// getting all the attachments might take a couple attempts due to filesize
		var retriesErr error
		for count := 0; count < numberOfRetries; count++ {
			attached, err := client.
				UsersById(user).
				MessagesById(*aMessage.GetId()).
				Attachments().
				Get()
			retriesErr = err
			if err == nil && attached != nil {
				aMessage.SetAttachments(attached.GetValue())
				break
			}
		}
		if retriesErr != nil {
			logger.Ctx(ctx).Debug("exceeded maximum retries")
			return support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(retriesErr, "attachment failed"), nil)
		}
	}
	err = objectWriter.WriteObjectValue("", aMessage)
	if err != nil {
		return support.SetNonRecoverableError(errors.Wrapf(err, "%s", *aMessage.GetId()))
	}

	byteArray, err := objectWriter.GetSerializedContent()
	objectWriter.Close()
	if err != nil {
		return support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(err, "serializing mail content"), nil)
	}
	if byteArray != nil {
		dataChannel <- &Stream{id: *aMessage.GetId(), message: byteArray, info: MessageInfo(aMessage)}
	}
	return nil
}

// Stream represents a single item retrieved from exchange
type Stream struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *details.ExchangeInfo //temporary change to bring populate function into directory
}

func (od *Stream) UUID() string {
	return od.id

}

func (od *Stream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.message))
}

func (od *Stream) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: od.info}
}

// NewStream constructor for exchange.Stream object
func NewStream(identifier string, bytes []byte, detail details.ExchangeInfo) Stream {
	return Stream{
		id:      identifier,
		message: bytes,
		info:    &detail,
	}

}

// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"fmt"
	"io"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	kw "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.Collection = &Collection{}
	_ data.Stream     = &Stream{}
	_ data.StreamInfo = &Stream{}
)

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4
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

	collectionType optionIdentifier
	statusUpdater  support.StatusUpdater
	// FullPath is the slice representation of the action context passed down through the hierarchy.
	// The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath path.Path
}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewCollection(
	user string,
	fullPath path.Path,
	collectionType optionIdentifier,
	service graph.Service,
	statusUpdater support.StatusUpdater,
) Collection {
	collection := Collection{
		user:           user,
		data:           make(chan data.Stream, collectionChannelBufferSize),
		jobs:           make([]string, 0),
		service:        service,
		statusUpdater:  statusUpdater,
		fullPath:       fullPath,
		collectionType: collectionType,
	}

	return collection
}

// AddJob appends additional objectID to structure's jobs field
func (col *Collection) AddJob(objID string) {
	col.jobs = append(col.jobs, objID)
}

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *Collection) Items() <-chan data.Stream {
	go col.populateByOptionIdentifier(context.TODO())
	return col.data
}

// GetQueryAndSerializeFunc helper function that returns the two functions functions
// required to convert M365 identifier into a byte array filled with the serialized data
func GetQueryAndSerializeFunc(optID optionIdentifier) (GraphRetrievalFunc, GraphSerializeFunc) {
	switch optID {
	case contacts:
		return RetrieveContactDataForUser, contactToDataCollection
	case events:
		return RetrieveEventDataForUser, eventToDataCollection
	case messages:
		return RetrieveMessageDataForUser, messageToDataCollection
	// Unsupported options returns nil, nil
	default:
		return nil, nil
	}
}

// FullPath returns the Collection's fullPath []string
func (col *Collection) FullPath() path.Path {
	return col.fullPath
}

// populateByOptionIdentifier is a utility function that uses col.collectionType to be able to serialize
// all the M365IDs defined in the jobs field. data channel is closed by this function
func (col *Collection) populateByOptionIdentifier(
	ctx context.Context,
) {
	var (
		errs       error
		success    int
		totalBytes int64

		user         = col.user
		objectWriter = kw.NewJsonSerializationWriter()
	)

	colProgress, closer := observe.CollectionProgress(user, col.fullPath.Category().String(), col.fullPath.Folder())
	go closer()

	defer func() {
		close(colProgress)
		col.finishPopulation(ctx, success, totalBytes, errs)
	}()

	// get QueryBasedonIdentifier
	// verify that it is the correct type in called function
	// serializationFunction
	query, serializeFunc := GetQueryAndSerializeFunc(col.collectionType)
	if query == nil {
		errs = fmt.Errorf("unrecognized collection type: %s", col.collectionType.String())
		return
	}

	for _, identifier := range col.jobs {
		response, err := query(ctx, col.service, user, identifier)
		if err != nil {
			errs = support.WrapAndAppendf(user, err, errs)

			if col.service.ErrPolicy() {
				break
			}

			continue
		}

		byteCount, err := serializeFunc(ctx, col.service.Client(), objectWriter, col.data, response, user)
		if err != nil {
			errs = support.WrapAndAppendf(user, err, errs)

			if col.service.ErrPolicy() {
				break
			}

			continue
		}

		success++

		totalBytes += int64(byteCount)
		colProgress <- struct{}{}
	}
}

// terminatePopulateSequence is a utility function used to close a Collection's data channel
// and to send the status update through the channel.
func (col *Collection) finishPopulation(ctx context.Context, success int, totalBytes int64, errs error) {
	close(col.data)
	attempted := len(col.jobs)
	status := support.CreateStatus(ctx,
		support.Backup,
		1,
		support.CollectionMetrics{
			Objects:    attempted,
			Successes:  success,
			TotalBytes: totalBytes,
		},
		errs,
		col.fullPath.Folder())
	logger.Ctx(ctx).Debug(status.String())
	col.statusUpdater(status)
}

// GraphSerializeFunc are class of functions that are used by Collections to transform GraphRetrievalFunc
// responses into data.Stream items contained within the Collection
type GraphSerializeFunc func(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error)

// eventToDataCollection is a GraphSerializeFunc used to serialize models.Eventable objects into
// data.Stream objects. Returns an error the process finishes unsuccessfully.
func eventToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error) {
	var err error

	defer objectWriter.Close()

	event, ok := parsable.(models.Eventable)
	if !ok {
		return 0, fmt.Errorf("expected Eventable, got %T", parsable)
	}

	if *event.GetHasAttachments() {
		var retriesErr error

		for count := 0; count < numberOfRetries; count++ {
			attached, err := client.
				UsersById(user).
				EventsById(*event.GetId()).
				Attachments().
				Get(ctx, nil)
			retriesErr = err

			if err == nil && attached != nil {
				event.SetAttachments(attached.GetValue())
				break
			}
		}

		if retriesErr != nil {
			logger.Ctx(ctx).Debug("exceeded maximum retries")

			return 0, support.WrapAndAppend(
				*event.GetId(),
				errors.Wrap(retriesErr, "attachment failed"),
				nil)
		}
	}

	err = objectWriter.WriteObjectValue("", event)
	if err != nil {
		return 0, support.SetNonRecoverableError(errors.Wrap(err, *event.GetId()))
	}

	byteArray, err := objectWriter.GetSerializedContent()
	if err != nil {
		return 0, support.WrapAndAppend(*event.GetId(), errors.Wrap(err, "serializing content"), nil)
	}

	if len(byteArray) > 0 {
		dataChannel <- &Stream{id: *event.GetId(), message: byteArray, info: EventInfo(event, int64(len(byteArray)))}
	}

	return len(byteArray), nil
}

// contactToDataCollection is a GraphSerializeFunc for models.Contactable
func contactToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error) {
	defer objectWriter.Close()

	contact, ok := parsable.(models.Contactable)
	if !ok {
		return 0, fmt.Errorf("expected Contactable, got %T", parsable)
	}

	err := objectWriter.WriteObjectValue("", contact)
	if err != nil {
		return 0, support.SetNonRecoverableError(errors.Wrap(err, *contact.GetId()))
	}

	byteArray, err := objectWriter.GetSerializedContent()
	if err != nil {
		return 0, support.WrapAndAppend(*contact.GetId(), err, nil)
	}

	if len(byteArray) > 0 {
		dataChannel <- &Stream{id: *contact.GetId(), message: byteArray, info: ContactInfo(contact, int64(len(byteArray)))}
	}

	return len(byteArray), nil
}

// messageToDataCollection is the GraphSerializeFunc for models.Messageable
func messageToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kw.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error) {
	var err error

	defer objectWriter.Close()

	aMessage, ok := parsable.(models.Messageable)
	if !ok {
		return 0, fmt.Errorf("expected Messageable, got %T", parsable)
	}

	adtl := aMessage.GetAdditionalData()
	if len(adtl) > 2 {
		aMessage, err = support.ConvertFromMessageable(adtl, aMessage)
		if err != nil {
			return 0, err
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
				Get(ctx, nil)
			retriesErr = err

			if err == nil {
				aMessage.SetAttachments(attached.GetValue())
				break
			}
		}

		if retriesErr != nil {
			logger.Ctx(ctx).Debug("exceeded maximum retries")
			return 0, support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(retriesErr, "attachment failed"), nil)
		}
	}

	err = objectWriter.WriteObjectValue("", aMessage)
	if err != nil {
		return 0, support.SetNonRecoverableError(errors.Wrapf(err, "%s", *aMessage.GetId()))
	}

	byteArray, err := objectWriter.GetSerializedContent()
	if err != nil {
		err = support.WrapAndAppend(*aMessage.GetId(), errors.Wrap(err, "serializing mail content"), nil)
		return 0, support.SetNonRecoverableError(err)
	}

	dataChannel <- &Stream{id: *aMessage.GetId(), message: byteArray, info: MessageInfo(aMessage, int64(len(byteArray)))}

	return len(byteArray), nil
}

// Stream represents a single item retrieved from exchange
type Stream struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *details.ExchangeInfo // temporary change to bring populate function into directory
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
func NewStream(identifier string, dataBytes []byte, detail details.ExchangeInfo) Stream {
	return Stream{
		id:      identifier,
		message: dataBytes,
		info:    &detail,
	}
}

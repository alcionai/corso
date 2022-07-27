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

	"github.com/alcionai/corso/internal/connector/service"
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
)

// PopulateFunc are a class of functions that can be used to fill exchange.Collections with
// the corresponding information
type PopulateFunc func(context.Context, service.Service, Collection, chan *support.ConnectorOperationStatus)

// ExchangeDataCollection represents exchange mailbox
// data for a single user.
//
// It implements the DataCollection interface
type Collection struct {
	// M365 user
	User string // M365 user
	Data chan data.Stream

	// FullPath is the slice representation of the action context passed down through the hierarchy.
	//The original request can be gleaned from the slice. (e.g. {<tenant ID>, <user ID>, "emails"})
	fullPath []string
}

// NewExchangeDataCollection creates an ExchangeDataCollection with fullPath is annotated
func NewCollection(aUser string, pathRepresentation []string) Collection {
	collection := Collection{
		User:     aUser,
		Data:     make(chan data.Stream, collectionChannelBufferSize),
		fullPath: pathRepresentation,
	}
	return collection
}

func (eoc *Collection) PopulateCollection(newData *Stream) {
	eoc.Data <- newData
}

// FinishPopulation is used to indicate data population of the collection is complete
// TODO: This should be an internal method once we move the message retrieval logic into `ExchangeDataCollection`
func (eoc *Collection) FinishPopulation() {
	if eoc.Data != nil {
		close(eoc.Data)
	}
}

// NOTE: Refactor has not happened moving into folders
// populateFromTaskList async call to fill DataCollection via channel implementation
func PopulateFromTaskList(
	ctx context.Context,
	tasklist support.TaskList,
	service service.Service,
	collections map[string]*Collection,
	statusChannel chan<- *support.ConnectorOperationStatus,
) {
	var errs error
	var attemptedItems, success int
	objectWriter := kw.NewJsonSerializationWriter()

	//Todo this has to return all the errors in the status
	for aFolder, tasks := range tasklist {
		// Get the same folder
		edc := collections[aFolder]
		if edc == nil {
			for _, task := range tasks {
				errs = support.WrapAndAppend(task, errors.New("unable to query: collection not found during populateFromTaskList"), errs)
			}
			continue
		}

		for _, task := range tasks {
			response, err := service.Client().UsersById(edc.User).MessagesById(task).Get()
			if err != nil {
				details := support.ConnectorStackErrorTrace(err)
				errs = support.WrapAndAppend(edc.User, errors.Wrapf(err, "unable to retrieve %s, %s", task, details), errs)
				continue
			}
			err = messageToDataCollection(service.Client(), ctx, objectWriter, edc.Data, response, edc.User)
			success++
			if err != nil {
				errs = support.WrapAndAppendf(edc.User, err, errs)
				success--
			}
			if errs != nil && service.Policy() {
				break
			}
		}

		edc.FinishPopulation()
		attemptedItems += len(tasks)
	}

	status := support.CreateStatus(ctx, support.Backup, attemptedItems, success, len(tasklist), errs)
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
		for count := 0; count < 5; count++ {
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

func (eoc *Collection) Items() <-chan data.Stream {
	return eoc.Data
}

func (edc *Collection) FullPath() []string {
	return append([]string{}, edc.fullPath...)
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
func NewStream(identifier string, bytes []byte, detail details.ExchangeInfo) Stream {
	return Stream{
		id:      identifier,
		message: bytes,
		info:    &detail,
	}

}

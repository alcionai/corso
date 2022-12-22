// Package exchange provides support for retrieving M365 Exchange objects
// from M365 servers using the Graph API. M365 object support centers
// on the applications: Mail, Contacts, and Calendar.
package exchange

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.Collection    = &Collection{}
	_ data.Stream        = &Stream{}
	_ data.StreamInfo    = &Stream{}
	_ data.StreamModTime = &Stream{}
)

const (
	collectionChannelBufferSize = 1000
	numberOfRetries             = 4

	// Outlooks expects max 4 concurrent requests
	// https://learn.microsoft.com/en-us/graph/throttling-limits#outlook-service-limits
	urlPrefetchChannelBufferSize = 4
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
	service graph.Servicer

	collectionType optionIdentifier
	statusUpdater  support.StatusUpdater
	ctrl           control.Options

	// FullPath is the current hierarchical path used by this collection.
	fullPath path.Path

	// PrevPath is the previous hierarchical path used by this collection.
	// It may be the same as fullPath, if the folder was not renamed or
	// moved.  It will be empty on its first retrieval.
	prevPath path.Path

	state data.CollectionState

	// doNotMergeItems should only be true if the old delta token expired.
	doNotMergeItems bool
}

// NewExchangeDataCollection creates an ExchangeDataCollection.
// State of the collection is set as an observation of the current
// and previous paths.  If the curr path is nil, the state is assumed
// to be deleted.  If the prev path is nil, it is assumed newly created.
// If both are populated, then state is either moved (if they differ),
// or notMoved (if they match).
func NewCollection(
	user string,
	curr, prev path.Path,
	collectionType optionIdentifier,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) Collection {
	collection := Collection{
		collectionType:  collectionType,
		ctrl:            ctrlOpts,
		data:            make(chan data.Stream, collectionChannelBufferSize),
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		jobs:            make([]string, 0),
		prevPath:        prev,
		service:         service,
		state:           stateOf(prev, curr),
		statusUpdater:   statusUpdater,
		user:            user,
	}

	return collection
}

func stateOf(prev, curr path.Path) data.CollectionState {
	if curr == nil || len(curr.String()) == 0 {
		return data.DeletedState
	}

	if prev == nil || len(prev.String()) == 0 {
		return data.NewState
	}

	if curr.Folder() != prev.Folder() {
		return data.MovedState
	}

	return data.NotMovedState
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

// TODO(ashmrtn): Fill in with previous path once GraphConnector compares old
// and new folder hierarchies.
func (col Collection) PreviousPath() path.Path {
	return nil
}

func (col Collection) State() data.CollectionState {
	return col.state
}

func (col Collection) DoNotMergeItems() bool {
	return col.doNotMergeItems
}

// populateByOptionIdentifier is a utility function that uses col.collectionType to be able to serialize
// all the M365IDs defined in the jobs field. data channel is closed by this function
func (col *Collection) populateByOptionIdentifier(ctx context.Context) {
	var (
		errs       error
		success    int64
		totalBytes int64
		wg         sync.WaitGroup

		user = col.user
	)

	colProgress, closer := observe.CollectionProgress(user, col.fullPath.Category().String(), col.fullPath.Folder())
	go closer()

	defer func() {
		close(colProgress)
		col.finishPopulation(ctx, int(success), totalBytes, errs)
	}()

	// get QueryBasedonIdentifier
	// verify that it is the correct type in called function
	// serializationFunction
	query, serializeFunc := GetQueryAndSerializeFunc(col.collectionType)
	if query == nil {
		errs = fmt.Errorf("unrecognized collection type: %s", col.collectionType.String())
		return
	}

	// Limit the max number of active requests to GC
	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	errUpdater := func(user string, err error) {
		errs = support.WrapAndAppend(user, err, errs)
	}

	for _, identifier := range col.jobs {
		if col.ctrl.FailFast && errs != nil {
			break
		}
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(identifier string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			var (
				response absser.Parsable
				err      error
			)

			for i := 1; i <= numberOfRetries; i++ {
				response, err = query(ctx, col.service, user, identifier)
				if err == nil {
					break
				}
				// TODO: Tweak sleep times
				if i < numberOfRetries {
					time.Sleep(time.Duration(3*(i+1)) * time.Second)
				}
			}

			if err != nil {
				errUpdater(user, err)
				return
			}

			byteCount, err := serializeFunc(
				ctx,
				col.service.Client(),
				kioser.NewJsonSerializationWriter(),
				col.data,
				response,
				user)
			if err != nil {
				errUpdater(user, err)
				return
			}

			atomic.AddInt64(&success, 1)
			atomic.AddInt64(&totalBytes, int64(byteCount))

			colProgress <- struct{}{}
		}(identifier)
	}

	wg.Wait()
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

type modTimer interface {
	GetLastModifiedDateTime() *time.Time
}

func getModTime(mt modTimer) time.Time {
	res := time.Now()

	if t := mt.GetLastModifiedDateTime(); t != nil {
		res = *t
	}

	return res
}

// GraphSerializeFunc are class of functions that are used by Collections to transform GraphRetrievalFunc
// responses into data.Stream items contained within the Collection
type GraphSerializeFunc func(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kioser.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error)

// eventToDataCollection is a GraphSerializeFunc used to serialize models.Eventable objects into
// data.Stream objects. Returns an error the process finishes unsuccessfully.
func eventToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kioser.JsonSerializationWriter,
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
		dataChannel <- &Stream{
			id:      *event.GetId(),
			message: byteArray,
			info:    EventInfo(event, int64(len(byteArray))),
			modTime: getModTime(event),
		}
	}

	return len(byteArray), nil
}

// contactToDataCollection is a GraphSerializeFunc for models.Contactable
func contactToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kioser.JsonSerializationWriter,
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

	bs, err := objectWriter.GetSerializedContent()
	if err != nil {
		return 0, support.WrapAndAppend(*contact.GetId(), err, nil)
	}

	addtl := contact.GetAdditionalData()
	_, removed := addtl[graph.AddtlDataRemoved]

	if len(bs) > 0 || removed {
		dataChannel <- &Stream{
			id:      *contact.GetId(),
			message: bs,
			info:    ContactInfo(contact, int64(len(bs))),
			modTime: getModTime(contact),
			deleted: removed,
		}
	}

	return len(bs), nil
}

// messageToDataCollection is the GraphSerializeFunc for models.Messageable
func messageToDataCollection(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	objectWriter *kioser.JsonSerializationWriter,
	dataChannel chan<- data.Stream,
	parsable absser.Parsable,
	user string,
) (int, error) {
	var err error

	defer objectWriter.Close()

	msg, ok := parsable.(models.Messageable)
	if !ok {
		return 0, fmt.Errorf("expected Messageable, got %T", parsable)
	}

	if *msg.GetHasAttachments() {
		// getting all the attachments might take a couple attempts due to filesize
		var retriesErr error

		for count := 0; count < numberOfRetries; count++ {
			attached, err := client.
				UsersById(user).
				MessagesById(*msg.GetId()).
				Attachments().
				Get(ctx, nil)
			retriesErr = err

			if err == nil {
				msg.SetAttachments(attached.GetValue())
				break
			}
		}

		if retriesErr != nil {
			logger.Ctx(ctx).Debug("exceeded maximum retries")
			return 0, support.WrapAndAppend(*msg.GetId(), errors.Wrap(retriesErr, "attachment failed"), nil)
		}
	}

	err = objectWriter.WriteObjectValue("", msg)
	if err != nil {
		return 0, support.SetNonRecoverableError(errors.Wrapf(err, "%s", *msg.GetId()))
	}

	bs, err := objectWriter.GetSerializedContent()
	if err != nil {
		err = support.WrapAndAppend(*msg.GetId(), errors.Wrap(err, "serializing mail content"), nil)
		return 0, support.SetNonRecoverableError(err)
	}

	addtl := msg.GetAdditionalData()
	_, removed := addtl[graph.AddtlDataRemoved]

	dataChannel <- &Stream{
		id:      *msg.GetId(),
		message: bs,
		info:    MessageInfo(msg, int64(len(bs))),
		modTime: getModTime(msg),
		deleted: removed,
	}

	return len(bs), nil
}

// Stream represents a single item retrieved from exchange
type Stream struct {
	id string
	// TODO: We may need this to be a "oneOf" of `message`, `contact`, etc.
	// going forward. Using []byte for now but I assume we'll have
	// some structured type in here (serialization to []byte can be done in `Read`)
	message []byte
	info    *details.ExchangeInfo // temporary change to bring populate function into directory
	// TODO(ashmrtn): Can probably eventually be sourced from info as there's a
	// request to provide modtime in ItemInfo structs.
	modTime time.Time

	// true if the item was marked by graph as deleted.
	deleted bool
}

func (od *Stream) UUID() string {
	return od.id
}

func (od *Stream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.message))
}

func (od Stream) Deleted() bool {
	return od.deleted
}

func (od *Stream) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: od.info}
}

func (od *Stream) ModTime() time.Time {
	return od.modTime
}

// NewStream constructor for exchange.Stream object
func NewStream(identifier string, dataBytes []byte, detail details.ExchangeInfo, modTime time.Time) Stream {
	return Stream{
		id:      identifier,
		message: dataBytes,
		info:    &detail,
		modTime: modTime,
	}
}

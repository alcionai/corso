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

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
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

	// added is a list of existing item IDs that were added to a container
	added []string
	// removed is a list of item IDs that were deleted from, or moved out, of a container
	removed []string

	// service - client/adapter pair used to access M365 back store
	service graph.Servicer

	category      path.CategoryType
	statusUpdater support.StatusUpdater
	ctrl          control.Options

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
	category path.CategoryType,
	service graph.Servicer,
	statusUpdater support.StatusUpdater,
	ctrlOpts control.Options,
	doNotMergeItems bool,
) Collection {
	collection := Collection{
		category:        category,
		ctrl:            ctrlOpts,
		data:            make(chan data.Stream, collectionChannelBufferSize),
		doNotMergeItems: doNotMergeItems,
		fullPath:        curr,
		added:           make([]string, 0),
		removed:         make([]string, 0),
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

// Items utility function to asynchronously execute process to fill data channel with
// M365 exchange objects and returns the data channel
func (col *Collection) Items() <-chan data.Stream {
	go col.streamItems(context.TODO())
	return col.data
}

// GetQueryAndSerializeFunc helper function that returns the two functions functions
// required to convert M365 identifier into a byte array filled with the serialized data
func GetQueryAndSerializeFunc(category path.CategoryType) (api.GraphRetrievalFunc, GraphSerializeFunc) {
	switch category {
	case path.ContactsCategory:
		return api.RetrieveContactDataForUser, serializeAndStreamContact
	case path.EventsCategory:
		return api.RetrieveEventDataForUser, serializeAndStreamEvent
	case path.EmailCategory:
		return api.RetrieveMessageDataForUser, serializeAndStreamMessage
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
	return col.prevPath
}

func (col Collection) State() data.CollectionState {
	return col.state
}

func (col Collection) DoNotMergeItems() bool {
	return col.doNotMergeItems
}

// ---------------------------------------------------------------------------
// Items() channel controller
// ---------------------------------------------------------------------------

// streamItems is a utility function that uses col.collectionType to be able to serialize
// all the M365IDs defined in the added field. data channel is closed by this function
func (col *Collection) streamItems(ctx context.Context) {
	var (
		errs        error
		success     int64
		totalBytes  int64
		wg          sync.WaitGroup
		colProgress chan<- struct{}

		user = col.user
	)

	defer func() {
		col.finishPopulation(ctx, int(success), totalBytes, errs)
	}()

	if len(col.added)+len(col.removed) > 0 {
		var closer func()
		colProgress, closer = observe.CollectionProgress(user, col.fullPath.Category().String(), col.fullPath.Folder())

		go closer()

		defer func() {
			close(colProgress)
		}()
	}

	// get QueryBasedonIdentifier
	// verify that it is the correct type in called function
	// serializationFunction
	query, serializeFunc := GetQueryAndSerializeFunc(col.category)
	if query == nil {
		errs = fmt.Errorf("unrecognized collection type: %s", col.category)
		return
	}

	// Limit the max number of active requests to GC
	semaphoreCh := make(chan struct{}, urlPrefetchChannelBufferSize)
	defer close(semaphoreCh)

	errUpdater := func(user string, err error) {
		errs = support.WrapAndAppend(user, err, errs)
	}

	// delete all removed items
	for _, id := range col.removed {
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			col.data <- &Stream{
				id:      id,
				modTime: time.Now().UTC(), // removed items have no modTime entry.
				deleted: true,
			}

			atomic.AddInt64(&success, 1)
			atomic.AddInt64(&totalBytes, 0)

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	// add any new items
	for _, id := range col.added {
		if col.ctrl.FailFast && errs != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			var (
				response absser.Parsable
				err      error
			)

			for i := 1; i <= numberOfRetries; i++ {
				response, err = query(ctx, col.service, user, id)
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

			if colProgress != nil {
				colProgress <- struct{}{}
			}
		}(id)
	}

	wg.Wait()
}

// terminatePopulateSequence is a utility function used to close a Collection's data channel
// and to send the status update through the channel.
func (col *Collection) finishPopulation(ctx context.Context, success int, totalBytes int64, errs error) {
	close(col.data)
	attempted := len(col.added) + len(col.removed)
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
	res := time.Now().UTC()

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

// serializeAndStreamEvent is a GraphSerializeFunc used to serialize models.Eventable objects into
// data.Stream objects. Returns an error the process finishes unsuccessfully.
func serializeAndStreamEvent(
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

// serializeAndStreamContact is a GraphSerializeFunc for models.Contactable
func serializeAndStreamContact(
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

	if len(bs) > 0 {
		dataChannel <- &Stream{
			id:      *contact.GetId(),
			message: bs,
			info:    ContactInfo(contact, int64(len(bs))),
			modTime: getModTime(contact),
		}
	}

	return len(bs), nil
}

// serializeAndStreamMessage is the GraphSerializeFunc for models.Messageable
func serializeAndStreamMessage(
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

	dataChannel <- &Stream{
		id:      *msg.GetId(),
		message: bs,
		info:    MessageInfo(msg, int64(len(bs))),
		modTime: getModTime(msg),
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

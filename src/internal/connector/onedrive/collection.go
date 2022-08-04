// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"context"
	"io"
	"net/http"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/logger"
)

const (
	collectionChannelBufferSize = 1000
)

var _ data.Collection = &Collection{}
var _ data.Stream = &Item{}
var _ data.StreamInfo = &Item{}

// Collection represents a set of OneDrive objects retreived from M365
type Collection struct {
	data       chan data.Stream
	folderPath string

	// file items in this collection
	driveItems []models.DriveItemable
	driveID    string
	service    graph.Service
}

// NewCollection creates a Collection
func NewCollection(folderPath, driveID string, service graph.Service) *Collection {
	return &Collection{
		folderPath: folderPath,
		driveItems: []models.DriveItemable{},
		driveID:    driveID,
		service:    service,
		data:       make(chan data.Stream, collectionChannelBufferSize),
	}
}

// Items() returns the channel containing M365 Exchange objects
func (oc *Collection) Items() <-chan data.Stream {
	for _, item := range oc.driveItems {
		d, err := oc.readItem(context.TODO(), item)
		if err != nil {
			panic(err)
		}
		oc.data <- &Item{id: *item.GetId(), data: d}
	}
	close(oc.data)

	return oc.data
}

func (oc *Collection) FullPath() []string {
	return []string{}
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id   string
	data io.ReadCloser
	info *details.OnedriveInfo //temporary change to bring populate function into directory
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return od.data
}

func (od *Item) Info() details.ItemInfo {
	return details.ItemInfo{Onedrive: od.info}
}

func (oc *Collection) readItem(ctx context.Context, item models.DriveItemable) (io.ReadCloser, error) {
	logger.Ctx(ctx).Debugf("Reading Item %s", *item.GetId())

	r, err := oc.service.Client().DrivesById(oc.driveID).ItemsById(*item.GetId()).Get()
	if err != nil {
		return nil, errors.Errorf("failed to get item %s", *item.GetId())
	}

	// Get the download URL - https://docs.microsoft.com/en-us/graph/api/driveitem-get-content
	// These URLs are pre-authenticated and can be used to download the data using the standard
	// http client

	if _, found := r.GetAdditionalData()[downloadUrlKey]; !found {
		return nil, errors.Errorf("file does not have a download URL. ID: %s, %#v", *item.GetId(), item.GetAdditionalData())
	}
	downloadUrl := r.GetAdditionalData()[downloadUrlKey].(*string)
	logger.Ctx(ctx).Debugf("Found download URL. Item %s, URL: %s", *item.GetId(), *downloadUrl)

	resp, err := http.Get(*downloadUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to download file from %s", *downloadUrl)
	}

	return resp.Body, nil
}

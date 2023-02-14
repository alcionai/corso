package graph

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// Idable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of an ID.
type Idable interface {
	GetId() *string
}

// Descendable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a "parent folder".
type Descendable interface {
	Idable
	GetParentFolderId() *string
}

// Displayable represents objects that implement msgraph-sdk-go/models.entityable
// and have the concept of a display name.
type Displayable interface {
	Idable
	GetDisplayName() *string
}

type Container interface {
	Descendable
	Displayable
}

// CachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type CachedContainer interface {
	Container
	// Location contains either the display names for the dirs (if this is a calendar)
	// or nil
	Location() *path.Builder
	SetLocation(*path.Builder)
	// Path contains either the ids for the dirs (if this is a calendar)
	// or the display names for the dirs
	Path() *path.Builder
	SetPath(*path.Builder)
}

// ContainerResolver houses functions for getting information about containers
// from remote APIs (i.e. resolve folder paths with Graph API). Resolvers may
// cache information about containers.
type ContainerResolver interface {
	// IDToPath takes an m365 container ID and converts it to a hierarchical path
	// to that container. The path has a similar format to paths on the local
	// file system.
	IDToPath(ctx context.Context, m365ID string, useIDInPath bool) (*path.Builder, *path.Builder, error)

	// Populate performs initialization steps for the resolver
	// @param ctx is necessary param for Graph API tracing
	// @param baseFolderID represents the M365ID base that the resolver will
	// conclude its search. Default input is "".
	Populate(ctx context.Context, errs *fault.Errors, baseFolderID string, baseContainerPather ...string) error

	// PathInCache performs a look up of a path reprensentation
	// and returns the m365ID of directory iff the pathString
	// matches the path of a container within the cache.
	// @returns bool represents if m365ID was found.
	PathInCache(pathString string) (string, bool)

	AddToCache(ctx context.Context, m365Container Container, useIDInPath bool) error

	// DestinationNameToID returns the ID of the destination container.  Dest is
	// assumed to be a display name.  The ID is only populated if the destination
	// was added using `AddToCache()`.  Returns an empty string if not found.
	DestinationNameToID(dest string) string

	// Items returns the containers in the cache.
	Items() []CachedContainer
}

// ======================================
// cachedContainer Implementations
// ======================================

var _ CachedContainer = &CacheFolder{}

type CacheFolder struct {
	Container
	l *path.Builder
	p *path.Builder
}

// NewCacheFolder public constructor for struct
func NewCacheFolder(c Container, pb, lpb *path.Builder) CacheFolder {
	cf := CacheFolder{
		Container: c,
		l:         lpb,
		p:         pb,
	}

	return cf
}

// =========================================
// Required Functions to satisfy interfaces
// =========================================

func (cf CacheFolder) Location() *path.Builder {
	return cf.l
}

func (cf *CacheFolder) SetLocation(newLocation *path.Builder) {
	cf.l = newLocation
}

func (cf CacheFolder) Path() *path.Builder {
	return cf.p
}

func (cf *CacheFolder) SetPath(newPath *path.Builder) {
	cf.p = newPath
}

// CalendarDisplayable is a transformative struct that aligns
// models.Calendarable interface with the container interface.
// Calendars do not have the 2 of the
type CalendarDisplayable struct {
	models.Calendarable
	parentID string
}

// GetDisplayName returns the *string of the calendar name
func (c CalendarDisplayable) GetDisplayName() *string {
	return c.GetName()
}

// GetParentFolderId returns the default calendar name address
// EventCalendars have a flat hierarchy and Calendars are rooted
// at the default
//
//nolint:revive
func (c CalendarDisplayable) GetParentFolderId() *string {
	return &c.parentID
}

// CreateCalendarDisplayable helper function to create the
// calendarDisplayable during msgraph-sdk-go iterative process
// @param entry is the input supplied by pageIterator.Iterate()
// @param parentID of Calendar sets. Only populate when used with
// EventCalendarCache
func CreateCalendarDisplayable(entry any, parentID string) *CalendarDisplayable {
	calendar, ok := entry.(models.Calendarable)
	if !ok {
		return nil
	}

	return &CalendarDisplayable{
		Calendarable: calendar,
		parentID:     parentID,
	}
}

// =========================================
// helper funcs
// =========================================

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func CheckRequiredValues(c Container) error {
	id := ptr.Val(c.GetId())
	if len(id) == 0 {
		return errors.New("container missing ID")
	}

	dn := ptr.Val(c.GetDisplayName())
	if len(dn) == 0 {
		return clues.New("container missing display name").With("container_id", id)
	}

	parentID := ptr.Val(c.GetParentFolderId())
	if len(parentID) == 0 {
		return clues.New("container missing parent ID").With("container_id", id)
	}

	return nil
}

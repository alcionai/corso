package onedrive

import (
	"context"
	"sync"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/puzpuzpuz/xsync/v2"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/metadata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type driveInfo struct {
	id           string
	name         string
	rootFolderID string
}

type restoreCaches struct {
	BackupDriveIDName     idname.Cacher
	collisionKeyToItemID  map[string]api.DriveItemIDType
	DriveIDToDriveInfo    *xsync.MapOf[string, driveInfo]
	DriveNameToDriveInfo  *xsync.MapOf[string, driveInfo]
	Folders               *folderCache
	OldLinkShareIDToNewID *xsync.MapOf[string, string]
	OldPermIDToNewID      *xsync.MapOf[string, string]
	ParentDirToMeta       *xsync.MapOf[string, metadata.Metadata]

	pool sync.Pool
}

func (rc *restoreCaches) AddDrive(
	ctx context.Context,
	md models.Driveable,
	grf GetRootFolderer,
) error {
	di := driveInfo{
		id:   ptr.Val(md.GetId()),
		name: ptr.Val(md.GetName()),
	}

	ctx = clues.Add(ctx, "drive_info", di)

	root, err := grf.GetRootFolder(ctx, di.id)
	if err != nil {
		return clues.Wrap(err, "getting drive root id")
	}

	di.rootFolderID = ptr.Val(root.GetId())

	rc.DriveIDToDriveInfo.Store(di.id, di)
	rc.DriveNameToDriveInfo.Store(di.name, di)

	return nil
}

// Populate looks up drive items available to the protectedResource
// and adds their info to the caches.
func (rc *restoreCaches) Populate(
	ctx context.Context,
	gdparf GetDrivePagerAndRootFolderer,
	protectedResourceID string,
) error {
	drives, err := api.GetAllDrives(
		ctx,
		gdparf.NewDrivePager(protectedResourceID, nil),
		true,
		maxDrivesRetries)
	if err != nil {
		return clues.Wrap(err, "getting drives")
	}

	for _, md := range drives {
		if err := rc.AddDrive(ctx, md, gdparf); err != nil {
			return clues.Wrap(err, "caching drive")
		}
	}

	return nil
}

type GetDrivePagerAndRootFolderer interface {
	GetRootFolderer
	NewDrivePagerer
}

func NewRestoreCaches(
	backupDriveIDNames idname.Cacher,
) *restoreCaches {
	// avoid nil panics
	if backupDriveIDNames == nil {
		backupDriveIDNames = idname.NewCache(nil)
	}

	return &restoreCaches{
		BackupDriveIDName:     backupDriveIDNames,
		collisionKeyToItemID:  map[string]api.DriveItemIDType{},
		DriveIDToDriveInfo:    xsync.NewMapOf[driveInfo](),
		DriveNameToDriveInfo:  xsync.NewMapOf[driveInfo](),
		Folders:               NewFolderCache(),
		OldLinkShareIDToNewID: xsync.NewMapOf[string](),
		OldPermIDToNewID:      xsync.NewMapOf[string](),
		ParentDirToMeta:       xsync.NewMapOf[metadata.Metadata](),
		// Buffer pool for uploads
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, graph.CopyBufferSize)
				return &b
			},
		},
	}
}

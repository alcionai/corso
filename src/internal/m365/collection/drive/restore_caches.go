package drive

import (
	"context"
	"sync"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/common/syncd"
	"github.com/alcionai/canario/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

type driveInfo struct {
	id           string
	name         string
	rootFolderID string
}

type ResourceIDNames struct {
	Users  idname.Cacher
	Groups idname.Cacher
}

type restoreCaches struct {
	BackupDriveIDName     idname.Cacher
	collisionKeyToItemID  map[string]api.DriveItemIDType
	DriveIDToDriveInfo    syncd.MapTo[driveInfo]
	DriveNameToDriveInfo  syncd.MapTo[driveInfo]
	Folders               *folderCache
	OldLinkShareIDToNewID syncd.MapTo[string]
	OldPermIDToNewID      syncd.MapTo[string]
	ParentDirToMeta       syncd.MapTo[metadata.Metadata]
	AvailableEntities     ResourceIDNames

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

type GetAllIDsAndNameser interface {
	GetAllIDsAndNames(ctx context.Context, errs *fault.Bus) (idname.Cacher, error)
}

// Populate looks up drive items available to the protectedResource
// and adds their info to the caches.
func (rc *restoreCaches) Populate(
	ctx context.Context,
	usersGetter, groupsGetter GetAllIDsAndNameser,
	gdparf GetDrivePagerAndRootFolderer,
	protectedResourceID string,
	errs *fault.Bus,
) error {
	drives, err := api.GetAllDrives(
		ctx,
		gdparf.NewDrivePager(protectedResourceID, nil))
	if err != nil {
		return clues.Wrap(err, "getting drives")
	}

	for _, md := range drives {
		if err := rc.AddDrive(ctx, md, gdparf); err != nil {
			return clues.Wrap(err, "caching drive")
		}
	}

	users, err := usersGetter.GetAllIDsAndNames(ctx, errs)
	if err != nil {
		return clues.Wrap(err, "getting users")
	}

	groups, err := groupsGetter.GetAllIDsAndNames(ctx, errs)
	if err != nil {
		return clues.Wrap(err, "getting groups")
	}

	rc.AvailableEntities.Users = users
	rc.AvailableEntities.Groups = groups

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
		DriveIDToDriveInfo:    syncd.NewMapTo[driveInfo](),
		DriveNameToDriveInfo:  syncd.NewMapTo[driveInfo](),
		Folders:               NewFolderCache(),
		OldLinkShareIDToNewID: syncd.NewMapTo[string](),
		OldPermIDToNewID:      syncd.NewMapTo[string](),
		ParentDirToMeta:       syncd.NewMapTo[metadata.Metadata](),
		// Buffer pool for uploads
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, graph.CopyBufferSize)
				return &b
			},
		},
	}
}

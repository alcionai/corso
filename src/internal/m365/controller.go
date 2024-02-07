package m365

import (
	"context"
	"runtime/trace"
	"sync"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/internal/m365/resource"
	"github.com/alcionai/canario/src/internal/m365/support"
	"github.com/alcionai/canario/src/internal/operations/inject"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

var ErrNoResourceLookup = clues.New("missing resource lookup client")

// must comply with BackupProducer and RestoreConsumer
var (
	_ inject.BackupProducer   = &Controller{}
	_ inject.ToServiceHandler = &Controller{}
)

// Controller is a struct used to wrap the GraphServiceClient and
// GraphRequestAdapter from the msgraph-sdk-go. Additional fields are for
// bookkeeping and interfacing with other component.
type Controller struct {
	AC api.Client

	tenant      string
	credentials account.M365Config

	resourceHandler idname.GetResourceIDAndNamer
	// maps of resource owner ids to names, and names to ids.
	// not guaranteed to be populated, only here as a post-population
	// reference for processes that choose to populate the values.
	IDNameLookup idname.Cacher

	// wg is used to track completion of tasks
	wg     *sync.WaitGroup
	region *trace.Region

	// mutex used to synchronize updates to `status`
	mu     sync.Mutex
	status support.ControllerOperationStatus // contains the status of the last run status

	// backupDriveIDNames is populated on restore and export.  It maps
	// the backup's drive names to their id. Primarily for use when
	// creating or looking up a new drive.
	backupDriveIDNames idname.CacheBuilder

	// backupSiteIDWebURL is populated on restore and export. It maps
	// the backup's site names to their id. Primarily for use in
	// exports for groups.
	backupSiteIDWebURL idname.CacheBuilder
}

func NewController(
	ctx context.Context,
	acct account.Account,
	pst path.ServiceType,
	co control.Options,
	counter *count.Bus,
) (*Controller, error) {
	graph.InitializeConcurrencyLimiter(ctx, pst == path.ExchangeService, co.Parallelism.ItemFetch)

	creds, err := acct.M365Config()
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "retrieving m365 account configuration")
	}

	ac, err := api.NewClient(creds, co, counter)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "creating api client")
	}

	ctrl := Controller{
		AC:           ac,
		IDNameLookup: idname.NewCache(nil),

		credentials:        creds,
		tenant:             acct.ID(),
		wg:                 &sync.WaitGroup{},
		backupDriveIDNames: idname.NewCache(nil),
		backupSiteIDWebURL: idname.NewCache(nil),
	}

	// TODO: remove in favor of calling setResourceHandler in the newServiceHandler once
	// all operations types are populated.  When we do that, we can also remove pst from
	// the func params in this call.
	ctrl.setResourceHandler(pst)

	return &ctrl, nil
}

func (ctrl *Controller) VerifyAccess(ctx context.Context) error {
	return ctrl.AC.Access().GetToken(ctx)
}

func (ctrl *Controller) setResourceHandler(
	serviceInOperation path.ServiceType,
) {
	var rh *resourceGetter

	switch serviceInOperation {
	case path.ExchangeService, path.OneDriveService, path.TeamsChatsService:
		rh = &resourceGetter{
			enum:   resource.Users,
			getter: ctrl.AC.Users(),
		}
	case path.GroupsService:
		rh = &resourceGetter{
			enum:   resource.Groups,
			getter: ctrl.AC.Groups(),
		}
	case path.SharePointService:
		rh = &resourceGetter{
			enum:   resource.Sites,
			getter: ctrl.AC.Sites(),
		}
	}

	ctrl.resourceHandler = rh
}

// ---------------------------------------------------------------------------
// Processing Status
// ---------------------------------------------------------------------------

// AwaitStatus waits for all tasks to complete and then returns status
func (ctrl *Controller) Wait() *data.CollectionStats {
	defer func() {
		if ctrl.region != nil {
			ctrl.region.End()
			ctrl.region = nil
		}
	}()
	ctrl.wg.Wait()

	// clean up and reset statefulness
	dcs := data.CollectionStats{
		Folders:   ctrl.status.Folders,
		Objects:   ctrl.status.Metrics.Objects,
		Successes: ctrl.status.Metrics.Successes,
		Bytes:     ctrl.status.Metrics.Bytes,
		Details:   ctrl.status.String(),
	}

	ctrl.wg = &sync.WaitGroup{}
	ctrl.status = support.ControllerOperationStatus{}

	return &dcs
}

// UpdateStatus is used by initiated tasks to indicate completion
func (ctrl *Controller) UpdateStatus(status *support.ControllerOperationStatus) {
	defer ctrl.wg.Done()

	if status == nil {
		return
	}

	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()
	ctrl.status = support.MergeStatus(ctrl.status, *status)
}

// Status returns the current status of the controller process.
func (ctrl *Controller) Status() support.ControllerOperationStatus {
	return ctrl.status
}

// PrintableStatus returns a string formatted version of the status.
func (ctrl *Controller) PrintableStatus() string {
	return ctrl.status.String()
}

func (ctrl *Controller) incrementAwaitingMessages() {
	ctrl.wg.Add(1)
}

func (ctrl *Controller) CacheItemInfo(dii details.ItemInfo) {
	if dii.Groups != nil {
		ctrl.backupDriveIDNames.Add(dii.Groups.DriveID, dii.Groups.DriveName)
		ctrl.backupSiteIDWebURL.Add(dii.Groups.SiteID, dii.Groups.WebURL)
	}

	if dii.SharePoint != nil {
		ctrl.backupDriveIDNames.Add(dii.SharePoint.DriveID, dii.SharePoint.DriveName)
		ctrl.backupSiteIDWebURL.Add(dii.SharePoint.SiteID, dii.SharePoint.WebURL)
	}

	if dii.OneDrive != nil {
		ctrl.backupDriveIDNames.Add(dii.OneDrive.DriveID, dii.OneDrive.DriveName)
	}
}

// ---------------------------------------------------------------------------
// Resource Lookup Handling
// ---------------------------------------------------------------------------

var _ idname.GetResourceIDAndNamer = &resourceGetter{}

type resourceGetter struct {
	enum   resource.Category
	getter getIDAndNamer
}

type getIDAndNamer interface {
	GetIDAndName(
		ctx context.Context,
		owner string,
		cc api.CallConfig,
	) (
		ownerID string,
		ownerName string,
		err error,
	)
}

// GetResourceIDAndNameFrom looks up the resource's canonical id and display name.
// If the resource is present in the idNameSwapper, then that interface's id and
// name values are returned.  As a fallback, the resource calls the discovery
// api to fetch the user or site using the resource value. This fallback assumes
// that the resource is a well formed ID or display name of appropriate design
// (PrincipalName for users, WebURL for sites).
func (r resourceGetter) GetResourceIDAndNameFrom(
	ctx context.Context,
	owner string,
	ins idname.Cacher,
) (idname.Provider, error) {
	if ins != nil {
		if n, ok := ins.NameOf(owner); ok {
			return idname.NewProvider(owner, n), nil
		} else if i, ok := ins.IDOf(owner); ok {
			return idname.NewProvider(i, owner), nil
		}
	}

	ctx = clues.Add(ctx, "owner_identifier", owner)

	var (
		id, name string
		err      error
	)

	id, name, err = r.getter.GetIDAndName(ctx, owner, api.CallConfig{})
	if err != nil {
		return nil, clues.Stack(err)
	}

	if len(id) == 0 || len(name) == 0 {
		return nil, clues.Stack(core.ErrNotFound)
	}

	return idname.NewProvider(id, name), nil
}

// PopulateProtectedResourceIDAndName takes the provided owner identifier and produces
// the owner's name and ID from that value.  Returns an error if the owner is
// not recognized by the current tenant.
//
// The id-name cacher is optional.  Some processes will look up all owners in
// the tenant before reaching this step.  In that case, the data gets handed
// down for this func to consume instead of performing further queries.  The
// data gets stored inside the controller instance for later re-use.
func (ctrl *Controller) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	resourceID string, // input value, can be either id or name
	ins idname.Cacher,
) (idname.Provider, error) {
	if ctrl.resourceHandler == nil {
		return nil, clues.StackWC(ctx, ErrNoResourceLookup)
	}

	pr, err := ctrl.resourceHandler.GetResourceIDAndNameFrom(ctx, resourceID, ins)
	if err != nil {
		return nil, clues.Wrap(err, "identifying resource owner")
	}

	ctrl.IDNameLookup = idname.NewCache(map[string]string{pr.ID(): pr.Name()})

	return pr, nil
}

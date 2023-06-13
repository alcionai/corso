package impl

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	Count         int
	Destination   string
	Site          string
	Tenant        string
	User          string
	SecondaryUser string
)

// TODO: ErrGenerating       = clues.New("not all items were successfully generated")

var ErrNotYetImplemented = clues.New("not yet implemented")

// ------------------------------------------------------------------------------------------
// Restoration
// ------------------------------------------------------------------------------------------

type dataBuilderFunc func(id, now, subject, body string) []byte

func generateAndRestoreItems(
	ctx context.Context,
	ctrl *m365.Controller,
	service path.ServiceType,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, userID, destFldr string,
	howMany int,
	dbf dataBuilderFunc,
	opts control.Options,
	errs *fault.Bus,
) (*details.Details, error) {
	items := make([]item, 0, howMany)

	for i := 0; i < howMany; i++ {
		var (
			now       = dttm.Now()
			nowLegacy = dttm.FormatToLegacy(time.Now())
			id        = uuid.NewString()
			subject   = "automated " + now[:16] + " - " + id[:8]
			body      = "automated " + cat.String() + " generation for " + userID + " at " + now + " - " + id
		)

		items = append(items, item{
			name: id,
			data: dbf(id, nowLegacy, subject, body),
		})
	}

	collections := []collection{{
		PathElements: []string{destFldr},
		category:     cat,
		items:        items,
	}}

	dest := control.DefaultRestoreConfig(dttm.SafeForTesting)
	dest.Location = destFldr
	print.Infof(ctx, "Restoring to folder %s", dest.Location)

	dataColls, err := buildCollections(
		service,
		tenantID, userID,
		dest,
		collections)
	if err != nil {
		return nil, err
	}

	print.Infof(ctx, "Generating %d %s items in %s\n", howMany, cat, Destination)

	return ctrl.ConsumeRestoreCollections(ctx, version.Backup, sel, dest, opts, dataColls, errs)
}

// ------------------------------------------------------------------------------------------
// Common Helpers
// ------------------------------------------------------------------------------------------

func getControllerAndVerifyResourceOwner(
	ctx context.Context,
	rc resource.Category,
	resourceOwner string,
) (
	*m365.Controller,
	account.Account,
	idname.Provider,
	error,
) {
	tid := str.First(Tenant, os.Getenv(account.AzureTenantID))

	if len(Tenant) == 0 {
		Tenant = tid
	}

	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: tid,
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, account.Account{}, nil, clues.Wrap(err, "finding m365 account details")
	}

	ctrl, err := m365.NewController(ctx, acct, rc)
	if err != nil {
		return nil, account.Account{}, nil, clues.Wrap(err, "connecting to graph api")
	}

	id, _, err := ctrl.PopulateOwnerIDAndNamesFrom(ctx, resourceOwner, nil)
	if err != nil {
		return nil, account.Account{}, nil, clues.Wrap(err, "verifying user")
	}

	return ctrl, acct, ctrl.IDNameLookup.ProviderForID(id), nil
}

type item struct {
	name string
	data []byte
}

type collection struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	PathElements []string
	category     path.CategoryType
	items        []item
}

func buildCollections(
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreConfig,
	colls []collection,
) ([]data.RestoreCollection, error) {
	collections := make([]data.RestoreCollection, 0, len(colls))

	for _, c := range colls {
		pth, err := path.Build(
			tenant,
			user,
			service,
			c.category,
			false,
			c.PathElements...)
		if err != nil {
			return nil, err
		}

		mc := exchMock.NewCollection(pth, pth, len(c.items))

		for i := 0; i < len(c.items); i++ {
			mc.Names[i] = c.items[i].name
			mc.Data[i] = c.items[i].data
		}

		collections = append(collections, data.NoFetchRestoreCollection{Collection: mc})
	}

	return collections, nil
}

var (
	folderAName = "folder-a"
	folderBName = "b"
	folderCName = "folder-c"

	fileAData = []byte(strings.Repeat("a", 33))
	fileBData = []byte(strings.Repeat("b", 65))
	fileEData = []byte(strings.Repeat("e", 257))

	// Cannot restore owner or empty permissions and so not testing them
	writePerm = []string{"write"}
	readPerm  = []string{"read"}
)

func generateAndRestoreDriveItems(
	ctrl *m365.Controller,
	resourceOwner, secondaryUserID, secondaryUserName string,
	acct account.Account,
	service path.ServiceType,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, destFldr string,
	count int,
	errs *fault.Bus,
) (
	*details.Details,
	error,
) {
	ctx, flush := tester.NewContext(nil)
	defer flush()

	dest := control.DefaultRestoreConfig(dttm.SafeForTesting)
	dest.Location = destFldr
	print.Infof(ctx, "Restoring to folder %s", dest.Location)

	var driveID string

	switch service {
	case path.SharePointService:
		d, err := ctrl.AC.Stable.Client().Sites().BySiteId(resourceOwner).Drive().Get(ctx, nil)
		if err != nil {
			return nil, clues.Wrap(err, "getting site's default drive")
		}

		driveID = ptr.Val(d.GetId())
	default:
		d, err := ctrl.AC.Stable.Client().Users().ByUserId(resourceOwner).Drive().Get(ctx, nil)
		if err != nil {
			return nil, clues.Wrap(err, "getting user's default drive")
		}

		driveID = ptr.Val(d.GetId())
	}

	var (
		cols []m365.OnedriveColInfo

		rootPath    = []string{"drives", driveID, "root:"}
		folderAPath = []string{"drives", driveID, "root:", folderAName}
		folderBPath = []string{"drives", driveID, "root:", folderBName}
		folderCPath = []string{"drives", driveID, "root:", folderCName}

		now              = time.Now()
		year, mnth, date = now.Date()
		hour, min, sec   = now.Clock()
		currentTime      = fmt.Sprintf("%d-%v-%d-%d-%d-%d", year, mnth, date, hour, min, sec)
	)

	for i := 0; i < count; i++ {
		col := []m365.OnedriveColInfo{
			// basic folder and file creation
			{
				PathElements: rootPath,
				Files: []m365.ItemData{
					{
						Name: fmt.Sprintf("file-1st-count-%d-at-%s", i, currentTime),
						Data: fileAData,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
					{
						Name: fmt.Sprintf("file-2nd-count-%d-at-%s", i, currentTime),
						Data: fileBData,
					},
				},
				Folders: []m365.ItemData{
					{
						Name: folderBName,
					},
					{
						Name: folderAName,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
					{
						Name: folderCName,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
				},
			},
			{
				// a folder that has permissions with an item in the folder with
				// the different permissions.
				PathElements: folderAPath,
				Files: []m365.ItemData{
					{
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileEData,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
				Perms: m365.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
			{
				// a folder that has permissions with an item in the folder with
				// no permissions.
				PathElements: folderCPath,
				Files: []m365.ItemData{
					{
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileAData,
					},
				},
				Perms: m365.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
			{
				PathElements: folderBPath,
				Files: []m365.ItemData{
					{
						// restoring a file in a non-root folder that doesn't inherit
						// permissions.
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileBData,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
				Folders: []m365.ItemData{
					{
						Name: folderAName,
						Perms: m365.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
				},
			},
		}

		cols = append(cols, col...)
	}

	input, err := m365.DataForInfo(service, cols, version.Backup)
	if err != nil {
		return nil, err
	}

	// collections := getCollections(
	// 	service,
	// 	tenantID,
	// 	[]string{resourceOwner},
	// 	input,
	// 	version.Backup)

	opts := control.Options{
		RestorePermissions: true,
		ToggleFeatures:     control.Toggles{},
	}

	config := m365.ConfigInfo{
		Opts:           opts,
		Resource:       resource.Users,
		Service:        service,
		Tenant:         tenantID,
		ResourceOwners: []string{resourceOwner},
		RestoreCfg:     ctrlTD.DefaultRestoreConfig(""),
	}

	_, _, collections, _, err := m365.GetCollectionsAndExpected(
		config,
		input,
		version.Backup)
	if err != nil {
		return nil, err
	}

	return ctrl.ConsumeRestoreCollections(ctx, version.Backup, sel, dest, opts, collections, errs)
}

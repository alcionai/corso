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
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var (
	Count         int
	Destination   string
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
	gc *connector.GraphConnector,
	acct account.Account,
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

	dest := control.DefaultRestoreDestination(dttm.SafeForTesting)
	dest.ContainerName = destFldr
	print.Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	dataColls, err := buildCollections(
		service,
		tenantID, userID,
		dest,
		collections)
	if err != nil {
		return nil, err
	}

	print.Infof(ctx, "Generating %d %s items in %s\n", howMany, cat, Destination)

	return gc.ConsumeRestoreCollections(ctx, version.Backup, acct, sel, dest, opts, dataColls, errs)
}

// ------------------------------------------------------------------------------------------
// Common Helpers
// ------------------------------------------------------------------------------------------

func getGCAndVerifyUser(
	ctx context.Context,
	userID string,
) (
	*connector.GraphConnector,
	account.Account,
	idname.Provider,
	error,
) {
	tid := common.First(Tenant, os.Getenv(account.AzureTenantID))

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

	gc, err := connector.NewGraphConnector(
		ctx,
		acct,
		connector.Users)
	if err != nil {
		return nil, account.Account{}, nil, clues.Wrap(err, "connecting to graph api")
	}

	id, _, err := gc.PopulateOwnerIDAndNamesFrom(ctx, userID, nil)
	if err != nil {
		return nil, account.Account{}, nil, clues.Wrap(err, "verifying user")
	}

	return gc, acct, gc.IDNameLookup.ProviderForID(id), nil
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
	dest control.RestoreDestination,
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

		collections = append(collections, data.NotFoundRestoreCollection{Collection: mc})
	}

	return collections, nil
}

// type connector.PermData struct {
// 	user        string // user is only for older versions
// 	entityID    string
// 	roles       []string
// 	sharingMode metadata.SharingMode
// }

// type connector.ItemData struct {
// 	name  string
// 	data  []byte
// 	perms connector.PermData
// }

// type itemInfo struct {
// 	// lookupKey is a string that can be used to find this data from a set of
// 	// other data in the same collection. This key should be something that will
// 	// be the same before and after restoring the item in M365 and may not be
// 	// the M365 ID. When restoring items out of place, the item is assigned a
// 	// new ID making it unsuitable for a lookup key.
// 	lookupKey string
// 	name      string
// 	data      []byte
// }

// type onedriveCollection struct {
// 	service       path.ServiceType
// 	PathElements  []string
// 	items         []itemInfo
// 	aux           []itemInfo
// 	backupVersion int
// }

// type connector.OnedriveColInfo struct {
// 	PathElements []string
// 	perms        connector.PermData
// 	Files        []connector.ItemData
// 	folders      []connector.ItemData
// }

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

func generateAndRestoreOnedriveItems(
	gc *connector.GraphConnector,
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
	ctx, flush := tester.NewContext()
	defer flush()

	dest := control.DefaultRestoreDestination(dttm.SafeForTesting)
	dest.ContainerName = destFldr
	print.Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	d, _ := gc.Service.Client().UsersById(resourceOwner).Drive().Get(ctx, nil)
	driveID := ptr.Val(d.GetId())

	var (
		cols []connector.OnedriveColInfo

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
		col := []connector.OnedriveColInfo{
			// basic folder and file creation
			{
				PathElements: rootPath,
				Files: []connector.ItemData{
					{
						Name: fmt.Sprintf("file-1st-count-%d-at-%s", i, currentTime),
						Data: fileAData,
						Perms: connector.PermData{
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
				Folders: []connector.ItemData{
					{
						Name: folderBName,
					},
					{
						Name: folderAName,
						Perms: connector.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    readPerm,
						},
					},
					{
						Name: folderCName,
						Perms: connector.PermData{
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
				Files: []connector.ItemData{
					{
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileEData,
						Perms: connector.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
				Perms: connector.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
			{
				// a folder that has permissions with an item in the folder with
				// no permissions.
				PathElements: folderCPath,
				Files: []connector.ItemData{
					{
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileAData,
					},
				},
				Perms: connector.PermData{
					User:     secondaryUserName,
					EntityID: secondaryUserID,
					Roles:    readPerm,
				},
			},
			{
				PathElements: folderBPath,
				Files: []connector.ItemData{
					{
						// restoring a file in a non-root folder that doesn't inherit
						// permissions.
						Name: fmt.Sprintf("file-count-%d-at-%s", i, currentTime),
						Data: fileBData,
						Perms: connector.PermData{
							User:     secondaryUserName,
							EntityID: secondaryUserID,
							Roles:    writePerm,
						},
					},
				},
				Folders: []connector.ItemData{
					{
						Name: folderAName,
						Perms: connector.PermData{
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

	input, err := connector.DataForInfo(service, cols, version.Backup)
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

	config := connector.ConfigInfo{
		Acct:           acct,
		Opts:           opts,
		Resource:       connector.Users,
		Service:        service,
		Tenant:         tenantID,
		ResourceOwners: []string{resourceOwner},
		Dest:           tester.DefaultTestRestoreDestination(""),
	}

	_, _, collections, _, err := connector.GetCollectionsAndExpected(
		// &t,
		config,
		// service,
		// tenantID,
		// []string{resourceOwner},
		input,
		version.Backup)

	if err != nil {
		return nil, err
	}

	return gc.ConsumeRestoreCollections(ctx, version.Backup, acct, sel, dest, opts, collections, errs)
}

// func getCollections(
// 	service path.ServiceType,
// 	tenant string,
// 	resourceOwners []string,
// 	testCollections []colInfo,
// 	backupVersion int,
// ) []data.RestoreCollection {
// 	var collections []data.RestoreCollection

// 	for _, owner := range resourceOwners {
// 		ownerCollections := collectionsForInfo(
// 			service,
// 			tenant,
// 			owner,
// 			testCollections,
// 			backupVersion,
// 		)

// 		collections = append(collections, ownerCollections...)
// 	}

// 	return collections
// }

// type mockRestoreCollection struct {
// 	data.Collection
// 	auxItems map[string]data.Stream
// }

// func (rc mockRestoreCollection) Fetch(
// 	ctx context.Context,
// 	name string,
// ) (data.Stream, error) {
// 	res := rc.auxItems[name]
// 	if res == nil {
// 		return nil, data.ErrNotFound
// 	}

// 	return res, nil
// }

// func collectionsForInfo(
// 	service path.ServiceType,
// 	tenant, user string,
// 	allInfo []colInfo,
// 	backupVersion int,
// ) []data.RestoreCollection {
// 	collections := make([]data.RestoreCollection, 0, len(allInfo))

// 	for _, info := range allInfo {
// 		pth := mustToDataLayerPath(
// 			service,
// 			tenant,
// 			user,
// 			info.category,
// 			info.PathElements,
// 			false)

// 		mc := exchMock.NewCollection(pth, pth, len(info.items))

// 		for i := 0; i < len(info.items); i++ {
// 			mc.Names[i] = info.items[i].name
// 			mc.Data[i] = info.items[i].data

// 			// We do not count metadata Files against item count
// 			if backupVersion > 0 && metadata.HasMetaSuffix(info.items[i].name) &&
// 				(service == path.OneDriveService || service == path.SharePointService) {
// 				continue
// 			}
// 		}

// 		c := mockRestoreCollection{Collection: mc, auxItems: map[string]data.Stream{}}

// 		for _, aux := range info.auxItems {
// 			c.auxItems[aux.name] = &exchMock.Data{
// 				ID:     aux.name,
// 				Reader: io.NopCloser(bytes.NewReader(aux.data)),
// 			}
// 		}

// 		collections = append(collections, c)
// 	}

// 	return collections
// }

// func mustToDataLayerPath(
// 	service path.ServiceType,
// 	tenant, resourceOwner string,
// 	category path.CategoryType,
// 	elements []string,
// 	isItem bool,
// ) path.Path {
// 	res, err := path.Build(tenant, resourceOwner, service, category, isItem, elements...)
// 	if err != nil {
// 		fmt.Println("building path", clues.ToCore(err))
// 	}

// 	return res
// }

// type colInfo struct {
// 	// Elements (in order) for the path representing this collection. Should
// 	// only contain elements after the prefix that corso uses for the path. For
// 	// example, a collection for the Inbox folder in exchange mail would just be
// 	// "Inbox".
// 	PathElements []string
// 	category     path.CategoryType
// 	items        []itemInfo
// 	// auxItems are items that can be retrieved with Fetch but won't be returned
// 	// by Items().
// 	auxItems []itemInfo
// }

// func newOneDriveCollection(
// 	service path.ServiceType,
// 	PathElements []string,
// 	backupVersion int,
// ) *onedriveCollection {
// 	return &onedriveCollection{
// 		service:       service,
// 		PathElements:  PathElements,
// 		backupVersion: backupVersion,
// 	}
// }

// func dataForInfo(
// 	service path.ServiceType,
// 	cols []connector.OnedriveColInfo,
// 	backupVersion int,
// ) []colInfo {
// 	var res []colInfo

// 	for _, c := range cols {
// 		onedriveCol := newOneDriveCollection(service, c.PathElements, backupVersion)

// 		for _, f := range c.Files {
// 			onedriveCol.withFile(f.Name, f.Data, f.Perms)
// 		}

// 		onedriveCol.withPermissions(c.Perms)

// 		res = append(res, onedriveCol.collection())
// 	}

// 	return res
// }

// func (c onedriveCollection) collection() colInfo {
// 	cat := path.FilesCategory
// 	if c.service == path.SharePointService {
// 		cat = path.LibrariesCategory
// 	}

// 	return colInfo{
// 		PathElements: c.PathElements,
// 		category:     cat,
// 		items:        c.items,
// 		auxItems:     c.aux,
// 	}
// }

// func (c *onedriveCollection) withFile(name string, fileData []byte, perm connector.PermData) *onedriveCollection {
// 	c.items = append(c.items, onedriveItemWithData(
// 		name+metadata.DataFileSuffix,
// 		name+metadata.DataFileSuffix,
// 		fileData))

// 	md := onedriveMetadata(
// 		name,
// 		name+metadata.MetaFileSuffix,
// 		name,
// 		perm,
// 		true)
// 	c.items = append(c.items, md)
// 	c.aux = append(c.aux, md)

// 	return c
// }

// // withPermissions adds permissions to the folder represented by this
// // onedriveCollection.
// func (c *onedriveCollection) withPermissions(perm connector.PermData) *onedriveCollection {
// 	if c.backupVersion < version.OneDrive4DirIncludesPermissions {
// 		return c
// 	}

// 	name := c.PathElements[len(c.PathElements)-1]
// 	metaName := name

// 	if c.backupVersion >= version.OneDrive5DirMetaNoName {
// 		// We switched to just .dirmeta for metadata file names.
// 		metaName = ""
// 	}

// 	if name == "root:" {
// 		return c
// 	}

// 	md := onedriveMetadata(
// 		name,
// 		metaName+metadata.DirMetaFileSuffix,
// 		metaName+metadata.DirMetaFileSuffix,
// 		perm,
// 		true)

// 	c.items = append(c.items, md)
// 	c.aux = append(c.aux, md)

// 	return c
// }

// type oneDriveData struct {
// 	FileName string `json:"fileName,omitempty"`
// 	Data     []byte `json:"data,omitempty"`
// }

// func onedriveItemWithData(
// 	name, lookupKey string,
// 	fileData []byte,
// ) itemInfo {
// 	content := oneDriveData{
// 		FileName: lookupKey,
// 		Data:     fileData,
// 	}

// 	serialized, _ := json.Marshal(content)

// 	return itemInfo{
// 		name:      name,
// 		data:      serialized,
// 		lookupKey: lookupKey,
// 	}
// }

// func onedriveMetadata(
// 	fileName, itemID, lookupKey string,
// 	perm connector.PermData,
// 	permUseID bool,
// ) itemInfo {
// 	meta := getMetadata(fileName, perm, permUseID)

// 	metaJSON, err := json.Marshal(meta)
// 	if err != nil {
// 		fmt.Println("marshalling metadata", clues.ToCore(err))
// 	}

// 	return itemInfo{
// 		name:      itemID,
// 		data:      metaJSON,
// 		lookupKey: lookupKey,
// 	}
// }

// func getMetadata(fileName string, perm connector.PermData, permUseID bool) metadata.Metadata {
// 	if len(perm.User) == 0 || len(perm.Roles) == 0 ||
// 		perm.SharingMode != metadata.SharingModeCustom {
// 		return metadata.Metadata{
// 			FileName:    fileName,
// 			SharingMode: perm.SharingMode,
// 		}
// 	}

// 	// In case of permissions, the id will usually be same for same
// 	// user/role combo unless deleted and readded, but we have to do
// 	// this as we only have two users of which one is already taken.
// 	id := uuid.NewString()
// 	uperm := metadata.Permission{ID: id, Roles: perm.Roles}

// 	if permUseID {
// 		uperm.EntityID = perm.EntityID
// 	} else {
// 		uperm.Email = perm.User
// 	}

// 	meta := metadata.Metadata{
// 		FileName:    fileName,
// 		Permissions: []metadata.Permission{uperm},
// 	}

// 	return meta
// }

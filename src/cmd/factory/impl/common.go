package impl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"github.com/google/uuid"

	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
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
	"github.com/alcionai/corso/src/pkg/services/m365"
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
			now       = common.Now()
			nowLegacy = common.FormatLegacyTime(time.Now())
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
		pathElements: []string{destFldr},
		category:     cat,
		items:        items,
	}}

	// TODO: fit the destination to the containers
	dest := control.DefaultRestoreDestination(common.SimpleTimeTesting)
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

func getGCAndVerifyUser(ctx context.Context, userID string) (*connector.GraphConnector, account.Account, common.IDsNames, error) {
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
		return nil, account.Account{}, common.IDsNames{}, clues.Wrap(err, "finding m365 account details")
	}

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	ins, err := m365.UsersMap(ctx, acct, errs)
	if err != nil {
		return nil, account.Account{}, common.IDsNames{}, clues.Wrap(err, "getting tenant users")
	}

	_, idOK := ins.NameOf(strings.ToLower(userID))
	_, nameOK := ins.IDOf(strings.ToLower(userID))

	if !idOK && !nameOK {
		return nil, account.Account{}, common.IDsNames{}, clues.New("user not found within tenant")
	}

	gc, err := connector.NewGraphConnector(
		ctx,
		acct,
		connector.Users,
		errs)
	if err != nil {
		return nil, account.Account{}, common.IDsNames{}, clues.Wrap(err, "connecting to graph api")
	}

	return gc, acct, ins, nil
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
	pathElements []string
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
			c.pathElements...)
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

type permData struct {
	user        string // user is only for older versions
	entityID    string
	roles       []string
	sharingMode onedrive.SharingMode
}

type itemData struct {
	name  string
	data  []byte
	perms permData
}

type itemInfo struct {
	// lookupKey is a string that can be used to find this data from a set of
	// other data in the same collection. This key should be something that will
	// be the same before and after restoring the item in M365 and may not be
	// the M365 ID. When restoring items out of place, the item is assigned a
	// new ID making it unsuitable for a lookup key.
	lookupKey string
	name      string
	data      []byte
}

type onedriveCollection struct {
	service       path.ServiceType
	pathElements  []string
	items         []itemInfo
	aux           []itemInfo
	backupVersion int
}

type onedriveColInfo struct {
	pathElements []string
	perms        permData
	files        []itemData
	folders      []itemData
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

type restoreBackupInfoMultiVersion struct {
	service       path.ServiceType
	collections   []colInfo
	backupVersion int
}

func generateAndRestoreOnedriveItems(
	gc *connector.GraphConnector,
	resourceOwner, secondaryUserID, secondaryUserName string,
	acct account.Account,
	service path.ServiceType,
	cat path.CategoryType,
	sel selectors.Selector,
	tenantID, destFldr string,
	count int,
	errs *fault.Bus) (*details.Details, error) {
	ctx, flush := tester.NewContext()
	defer flush()

	// TODO: fit the destination to the containers
	dest := control.DefaultRestoreDestination(common.SimpleTimeTesting)
	dest.ContainerName = destFldr
	print.Infof(ctx, "Restoring to folder %s", dest.ContainerName)

	d, _ := gc.Service.Client().UsersById(resourceOwner).Drive().Get(ctx, nil)
	driveID := ptr.Val(d.GetId())

	var cols []onedriveColInfo

	rootPath := []string{
		"drives",
		driveID,
		"root:",
	}
	folderAPath := []string{
		"drives",
		driveID,
		"root:",
		folderAName,
	}
	folderBPath := []string{
		"drives",
		driveID,
		"root:",
		folderBName,
	}

	folderCPath := []string{
		"drives",
		driveID,
		"root:",
		folderCName,
	}

	currentTime := fmt.Sprintf("%d-%d", time.Now().Hour(), time.Now().Minute())

	for i := 0; i < count; i++ {
		var col onedriveColInfo

		// basic folder and file creation
		if i == 0 {
			col = onedriveColInfo{
				pathElements: rootPath,
				files: []itemData{
					{
						// Test restoring a file that doesn't inherit permissions.
						name: fmt.Sprintf("testFile-%s-1-%d", currentTime, i),
						data: fileAData,
						perms: permData{
							user:     secondaryUserName,
							entityID: secondaryUserID,
							roles:    writePerm,
						},
					},
					{
						// Test restoring a file that doesn't inherit permissions and has
						// no permissions.
						name: fmt.Sprintf("testFile-%s-2-%d", currentTime, i),
						data: fileBData,
					},
				},
				folders: []itemData{
					{
						name: folderBName,
					},
					{
						name: folderAName,
						perms: permData{
							user:     secondaryUserName,
							entityID: secondaryUserID,
							roles:    readPerm,
						},
					},
					{
						name: folderCName,
						perms: permData{
							user:     secondaryUserName,
							entityID: secondaryUserID,
							roles:    readPerm,
						},
					},
				},
			}

			cols = append(cols, col)
			continue
		}

		if i%2 == 0 {
			col = onedriveColInfo{
				// Tests a folder that has permissions with an item in the folder with
				// the different permissions.
				pathElements: folderAPath,
				files: []itemData{
					{
						name: fmt.Sprintf("testFile-%s-1-%d", currentTime, i),
						data: fileEData,
						perms: permData{
							user:     secondaryUserName,
							entityID: secondaryUserID,
							roles:    writePerm,
						},
					},
				},
				perms: permData{
					user:     secondaryUserName,
					entityID: secondaryUserID,
					roles:    readPerm,
				},
			}

			cols = append(cols, col)
			continue
		}

		if i%3 == 0 {
			col = onedriveColInfo{
				// Tests a folder that has permissions with an item in the folder with
				// no permissions.
				pathElements: folderCPath,
				files: []itemData{
					{
						name: fmt.Sprintf("testFile-%s-1-%d", currentTime, i),
						data: fileAData,
					},
				},
				perms: permData{
					user:     secondaryUserName,
					entityID: secondaryUserID,
					roles:    readPerm,
				},
			}

			cols = append(cols, col)
			continue
		}

		col = onedriveColInfo{
			pathElements: folderBPath,
			files: []itemData{
				{
					// Test restoring a file in a non-root folder that doesn't inherit
					// permissions.
					name: fmt.Sprintf("testFile-%s-1-%d", currentTime, i),
					data: fileBData,
					perms: permData{
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    writePerm,
					},
				},
			},
			folders: []itemData{
				{
					name: folderAName,
					perms: permData{
						user:     secondaryUserName,
						entityID: secondaryUserID,
						roles:    readPerm,
					},
				},
			},
		}

		cols = append(cols, col)
	}

	input := testDataForInfo(service, cols, version.Backup)

	testData := restoreBackupInfoMultiVersion{
		service:       service,
		backupVersion: version.Backup,
		collections:   input,
	}

	collections := runRestoreBackupTestVersions(
		acct,
		testData,
		tenantID,
		[]string{resourceOwner},
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{},
		},
	)

	opts := control.Options{
		RestorePermissions: true,
		ToggleFeatures:     control.Toggles{},
	}
	return gc.ConsumeRestoreCollections(ctx, version.Backup, acct, sel, dest, opts, collections, errs)

}

type configInfo struct {
	acct           account.Account
	opts           control.Options
	service        path.ServiceType
	tenant         string
	resourceOwners []string
	dest           control.RestoreDestination
}

func runRestoreBackupTestVersions(
	acct account.Account,
	test restoreBackupInfoMultiVersion,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) []data.RestoreCollection {
	config := configInfo{
		acct:           acct,
		opts:           opts,
		service:        test.service,
		tenant:         tenant,
		resourceOwners: resourceOwners,
		dest:           tester.DefaultTestRestoreDestination(),
	}

	collections := getCollections(
		config,
		test.collections,
		test.backupVersion)

	return collections
}

func getCollections(
	config configInfo,
	testCollections []colInfo,
	backupVersion int,
) []data.RestoreCollection {
	var (
		collections []data.RestoreCollection
	)

	for _, owner := range config.resourceOwners {
		ownerCollections := collectionsForInfo(
			config.service,
			config.tenant,
			owner,
			config.dest,
			testCollections,
			backupVersion,
		)

		collections = append(collections, ownerCollections...)
	}

	return collections
}

type mockRestoreCollection struct {
	data.Collection
	auxItems map[string]data.Stream
}

func (rc mockRestoreCollection) Fetch(
	ctx context.Context,
	name string,
) (data.Stream, error) {
	res := rc.auxItems[name]
	if res == nil {
		return nil, data.ErrNotFound
	}

	return res, nil
}

func collectionsForInfo(
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreDestination,
	allInfo []colInfo,
	backupVersion int,
) []data.RestoreCollection {
	var (
		collections = make([]data.RestoreCollection, 0, len(allInfo))
	)

	for _, info := range allInfo {
		pth := mustToDataLayerPath(
			service,
			tenant,
			user,
			info.category,
			info.pathElements,
			false)

		mc := exchMock.NewCollection(pth, pth, len(info.items))

		for i := 0; i < len(info.items); i++ {
			mc.Names[i] = info.items[i].name
			mc.Data[i] = info.items[i].data

			// We do not count metadata files against item count
			if backupVersion > 0 && metadata.HasMetaSuffix(info.items[i].name) &&
				(service == path.OneDriveService || service == path.SharePointService) {
				continue
			}

		}

		c := mockRestoreCollection{Collection: mc, auxItems: map[string]data.Stream{}}

		for _, aux := range info.auxItems {
			c.auxItems[aux.name] = &exchMock.Data{
				ID:     aux.name,
				Reader: io.NopCloser(bytes.NewReader(aux.data)),
			}
		}

		collections = append(collections, c)
	}

	return collections
}

func mustToDataLayerPath(
	service path.ServiceType,
	tenant, resourceOwner string,
	category path.CategoryType,
	elements []string,
	isItem bool,
) path.Path {
	res, err := path.Build(tenant, resourceOwner, service, category, isItem, elements...)
	if err != nil {
		fmt.Println("building path", clues.ToCore(err))
	}

	return res
}

type colInfo struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	pathElements []string
	category     path.CategoryType
	items        []itemInfo
	// auxItems are items that can be retrieved with Fetch but won't be returned
	// by Items(). These files do not directly participate in comparisosn at the
	// end of a test.
	auxItems []itemInfo
}

func newOneDriveCollection(
	service path.ServiceType,
	pathElements []string,
	backupVersion int,
) *onedriveCollection {
	return &onedriveCollection{
		service:       service,
		pathElements:  pathElements,
		backupVersion: backupVersion,
	}
}

func testDataForInfo(
	service path.ServiceType,
	cols []onedriveColInfo,
	backupVersion int,
) []colInfo {
	var res []colInfo

	for _, c := range cols {
		onedriveCol := newOneDriveCollection(service, c.pathElements, backupVersion)

		for _, f := range c.files {
			onedriveCol.withFile(f.name, f.data, f.perms)
		}

		for _, d := range c.folders {
			onedriveCol.withFolder(d.name, d.perms)
		}

		onedriveCol.withPermissions(c.perms)

		res = append(res, onedriveCol.collection())
	}

	return res
}

func (c onedriveCollection) collection() colInfo {
	cat := path.FilesCategory
	if c.service == path.SharePointService {
		cat = path.LibrariesCategory
	}

	return colInfo{
		pathElements: c.pathElements,
		category:     cat,
		items:        c.items,
		auxItems:     c.aux,
	}
}

func (c *onedriveCollection) withFolder(name string, perm permData) *onedriveCollection {
	switch c.backupVersion {
	case 0, version.OneDrive4DirIncludesPermissions, version.OneDrive5DirMetaNoName,
		version.OneDrive6NameInMeta, version.OneDrive7LocationRef:
		return c
	}

	return c
}

func (c *onedriveCollection) withFile(name string, fileData []byte, perm permData) *onedriveCollection {

	c.items = append(c.items, onedriveItemWithData(
		name+metadata.DataFileSuffix,
		name+metadata.DataFileSuffix,
		fileData))

	md := onedriveMetadata(
		name,
		name+metadata.MetaFileSuffix,
		name,
		perm,
		true)
	c.items = append(c.items, md)
	c.aux = append(c.aux, md)

	return c
}

// withPermissions adds permissions to the folder represented by this
// onedriveCollection.
func (c *onedriveCollection) withPermissions(perm permData) *onedriveCollection {
	if c.backupVersion < version.OneDrive4DirIncludesPermissions {
		return c
	}

	name := c.pathElements[len(c.pathElements)-1]
	metaName := name

	if c.backupVersion >= version.OneDrive5DirMetaNoName {
		// We switched to just .dirmeta for metadata file names.
		metaName = ""
	}

	if name == "root:" {
		return c
	}

	md := onedriveMetadata(
		name,
		metaName+metadata.DirMetaFileSuffix,
		metaName+metadata.DirMetaFileSuffix,
		perm,
		true)

	c.items = append(c.items, md)
	c.aux = append(c.aux, md)

	return c
}

type testOneDriveData struct {
	FileName string `json:"fileName,omitempty"`
	Data     []byte `json:"data,omitempty"`
}

func onedriveItemWithData(
	name, lookupKey string,
	fileData []byte,
) itemInfo {

	content := testOneDriveData{
		FileName: lookupKey,
		Data:     fileData,
	}

	serialized, _ := json.Marshal(content)

	return itemInfo{
		name:      name,
		data:      serialized,
		lookupKey: lookupKey,
	}
}

func onedriveMetadata(
	fileName, itemID, lookupKey string,
	perm permData,
	permUseID bool,
) itemInfo {
	testMeta := getMetadata(fileName, perm, permUseID)

	testMetaJSON, err := json.Marshal(testMeta)
	if err != nil {
		fmt.Println("marshalling metadata", clues.ToCore(err))
	}
	return itemInfo{
		name:      itemID,
		data:      testMetaJSON,
		lookupKey: lookupKey,
	}
}

func getMetadata(fileName string, perm permData, permUseID bool) onedrive.Metadata {
	if len(perm.user) == 0 || len(perm.roles) == 0 ||
		perm.sharingMode != onedrive.SharingModeCustom {
		return onedrive.Metadata{
			FileName:    fileName,
			SharingMode: perm.sharingMode,
		}
	}

	// In case of permissions, the id will usually be same for same
	// user/role combo unless deleted and readded, but we have to do
	// this as we only have two users of which one is already taken.
	id := uuid.NewString()
	uperm := onedrive.UserPermission{ID: id, Roles: perm.roles}

	if permUseID {
		uperm.EntityID = perm.entityID
	} else {
		uperm.Email = perm.user
	}

	testMeta := onedrive.Metadata{
		FileName:    fileName,
		Permissions: []onedrive.UserPermission{uperm},
	}

	return testMeta
}

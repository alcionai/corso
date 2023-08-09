package stub

import (
	"bytes"
	"io"

	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/mock"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type ColInfo struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	PathElements []string
	Category     path.CategoryType
	Items        []ItemInfo
	// auxItems are items that can be retrieved with Fetch but won't be returned
	// by Items(). These files do not directly participate in comparisosn at the
	// end of a test.
	AuxItems []ItemInfo
}

type ItemInfo struct {
	// lookupKey is a string that can be used to find this data from a set of
	// other data in the same collection. This key should be something that will
	// be the same before and after restoring the item in M365 and may not be
	// the M365 ID. When restoring items out of place, the item is assigned a
	// new ID making it unsuitable for a lookup key.
	LookupKey string
	Name      string
	Data      []byte
}

type ConfigInfo struct {
	Opts           control.Options
	Resource       resource.Category
	Service        path.ServiceType
	Tenant         string
	ResourceOwners []string
	RestoreCfg     control.RestoreConfig
}

func GetCollectionsAndExpected(
	config ConfigInfo,
	testCollections []ColInfo,
	backupVersion int,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte, error) {
	var (
		collections     []data.RestoreCollection
		expectedData    = map[string]map[string][]byte{}
		totalItems      = 0
		totalKopiaItems = 0
	)

	for _, owner := range config.ResourceOwners {
		numItems, kopiaItems, ownerCollections, userExpectedData, err := CollectionsForInfo(
			config.Service,
			config.Tenant,
			owner,
			config.RestoreCfg,
			testCollections,
			backupVersion)
		if err != nil {
			return totalItems, totalKopiaItems, collections, expectedData, err
		}

		collections = append(collections, ownerCollections...)
		totalItems += numItems
		totalKopiaItems += kopiaItems

		maps.Copy(expectedData, userExpectedData)
	}

	return totalItems, totalKopiaItems, collections, expectedData, nil
}

func CollectionsForInfo(
	service path.ServiceType,
	tenant, user string,
	restoreCfg control.RestoreConfig,
	allInfo []ColInfo,
	backupVersion int,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte, error) {
	var (
		collections  = make([]data.RestoreCollection, 0, len(allInfo))
		expectedData = make(map[string]map[string][]byte, len(allInfo))
		totalItems   = 0
		kopiaEntries = 0
	)

	for _, info := range allInfo {
		pth, err := path.Build(
			tenant,
			user,
			service,
			info.Category,
			false,
			info.PathElements...)
		if err != nil {
			return totalItems, kopiaEntries, collections, expectedData, err
		}

		mc := exchMock.NewCollection(pth, pth, len(info.Items))

		baseDestPath, err := backupOutputPathFromRestore(restoreCfg, pth)
		if err != nil {
			return totalItems, kopiaEntries, collections, expectedData, err
		}

		baseExpected := expectedData[baseDestPath.String()]
		if baseExpected == nil {
			expectedData[baseDestPath.String()] = make(map[string][]byte, len(info.Items))
			baseExpected = expectedData[baseDestPath.String()]
		}

		for i := 0; i < len(info.Items); i++ {
			mc.Names[i] = info.Items[i].Name
			mc.Data[i] = info.Items[i].Data

			baseExpected[info.Items[i].LookupKey] = info.Items[i].Data

			// We do not count metadata files against item count
			if backupVersion > 0 &&
				(service == path.OneDriveService || service == path.SharePointService) &&
				metadata.HasMetaSuffix(info.Items[i].Name) {
				continue
			}

			totalItems++
		}

		c := mock.RestoreCollection{
			Collection: mc,
			AuxItems:   map[string]data.Stream{},
		}

		for _, aux := range info.AuxItems {
			c.AuxItems[aux.Name] = &exchMock.Data{
				ID:     aux.Name,
				Reader: io.NopCloser(bytes.NewReader(aux.Data)),
			}
		}

		collections = append(collections, c)
		kopiaEntries += len(info.Items)
	}

	return totalItems, kopiaEntries, collections, expectedData, nil
}

// backupOutputPathFromRestore returns a path.Path denoting the location in
// kopia the data will be placed at. The location is a data-type specific
// combination of the location the data was recently restored to and where the
// data was originally in the hierarchy.
func backupOutputPathFromRestore(
	restoreCfg control.RestoreConfig,
	inputPath path.Path,
) (path.Path, error) {
	base := []string{restoreCfg.Location}

	// OneDrive has leading information like the drive ID.
	if inputPath.Service() == path.OneDriveService || inputPath.Service() == path.SharePointService {
		folders := inputPath.Folders()
		base = append(append([]string{}, folders[:3]...), restoreCfg.Location)

		if len(folders) > 3 {
			base = append(base, folders[3:]...)
		}
	}

	if inputPath.Service() == path.ExchangeService && inputPath.Category() == path.EmailCategory {
		base = append(base, inputPath.Folders()...)
	}

	return path.Build(
		inputPath.Tenant(),
		inputPath.ProtectedResource(),
		inputPath.Service(),
		inputPath.Category(),
		false,
		base...)
}

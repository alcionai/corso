package connector

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/clues"
	exchMock "github.com/alcionai/corso/src/internal/connector/exchange/mock"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/stretchr/testify/require"
)

type ConfigInfo struct {
	Acct           account.Account
	Opts           control.Options
	Resource       Resource
	Service        path.ServiceType
	Tenant         string
	ResourceOwners []string
	Dest           control.RestoreDestination
}

func mustToDataLayerPath(
	t *testing.T,
	service path.ServiceType,
	tenant, resourceOwner string,
	category path.CategoryType,
	elements []string,
	isItem bool,
) path.Path {
	res, err := path.Build(tenant, resourceOwner, service, category, isItem, elements...)
	require.NoError(t, err, clues.ToCore(err))

	return res
}

// backupOutputPathFromRestore returns a path.Path denoting the location in
// kopia the data will be placed at. The location is a data-type specific
// combination of the location the data was recently restored to and where the
// data was originally in the hierarchy.
func backupOutputPathFromRestore(
	t *testing.T,
	restoreDest control.RestoreDestination,
	inputPath path.Path,
) path.Path {
	base := []string{restoreDest.ContainerName}

	// OneDrive has leading information like the drive ID.
	if inputPath.Service() == path.OneDriveService || inputPath.Service() == path.SharePointService {
		folders := inputPath.Folders()
		base = append(append([]string{}, folders[:3]...), restoreDest.ContainerName)

		if len(folders) > 3 {
			base = append(base, folders[3:]...)
		}
	}

	if inputPath.Service() == path.ExchangeService && inputPath.Category() == path.EmailCategory {
		base = append(base, inputPath.Folders()...)
	}

	return mustToDataLayerPath(
		t,
		inputPath.Service(),
		inputPath.Tenant(),
		inputPath.ResourceOwner(),
		inputPath.Category(),
		base,
		false,
	)
}

// TODO(ashmrtn): Make this an actual mock class that can be used in other
// packages.
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
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreDestination,
	allInfo []ColInfo,
	backupVersion int,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte) {
	var (
		collections  = make([]data.RestoreCollection, 0, len(allInfo))
		expectedData = make(map[string]map[string][]byte, len(allInfo))
		totalItems   = 0
		kopiaEntries = 0
	)

	for _, info := range allInfo {
		pth := mustToDataLayerPath(
			t,
			service,
			tenant,
			user,
			info.Category,
			info.PathElements,
			false)

		mc := exchMock.NewCollection(pth, pth, len(info.Items))
		baseDestPath := backupOutputPathFromRestore(t, dest, pth)

		baseExpected := expectedData[baseDestPath.String()]
		if baseExpected == nil {
			expectedData[baseDestPath.String()] = make(map[string][]byte, len(info.Items))
			baseExpected = expectedData[baseDestPath.String()]
		}

		for i := 0; i < len(info.Items); i++ {
			mc.Names[i] = info.Items[i].name
			mc.Data[i] = info.Items[i].data

			baseExpected[info.Items[i].lookupKey] = info.Items[i].data

			// We do not count metadata files against item count
			if backupVersion > 0 &&
				(service == path.OneDriveService || service == path.SharePointService) &&
				metadata.HasMetaSuffix(info.Items[i].name) {
				continue
			}

			totalItems++
		}

		c := mockRestoreCollection{Collection: mc, auxItems: map[string]data.Stream{}}

		for _, aux := range info.AuxItems {
			c.auxItems[aux.name] = &exchMock.Data{
				ID:     aux.name,
				Reader: io.NopCloser(bytes.NewReader(aux.data)),
			}
		}

		collections = append(collections, c)
		kopiaEntries += len(info.Items)
	}

	return totalItems, kopiaEntries, collections, expectedData
}

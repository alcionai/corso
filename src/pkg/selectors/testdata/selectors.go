package testdata

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func MakeSelector(
	t *testing.T,
	service path.ServiceType,
	resourceOwners []string,
	forRestore bool,
) selectors.Selector {
	switch service {
	case path.ExchangeService:
		if forRestore {
			return selectors.NewExchangeRestore(resourceOwners).Selector
		}

		return selectors.NewExchangeBackup(resourceOwners).Selector

	case path.OneDriveService:
		if forRestore {
			return selectors.NewOneDriveRestore(resourceOwners).Selector
		}

		return selectors.NewOneDriveBackup(resourceOwners).Selector

	case path.SharePointService:
		if forRestore {
			return selectors.NewSharePointRestore(resourceOwners).Selector
		}

		return selectors.NewSharePointBackup(resourceOwners).Selector

	default:
		require.FailNow(t, "unknown path service")
		return selectors.Selector{}
	}
}

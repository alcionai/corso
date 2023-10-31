package common

import (
	"context"
	"os"
	"strings"

	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type PermissionInfo struct {
	EntityID string
	Roles    []string
}

const (
	sanityBackupID             = "SANITY_BACKUP_ID"
	sanityTestSourceContainer  = "SANITY_TEST_SOURCE_CONTAINER"
	sanityTestRestoreContainer = "SANITY_TEST_RESTORE_CONTAINER"
	sanityTestUser             = "SANITY_TEST_USER"
)

type Envs struct {
	BackupID         string
	SourceContainer  string
	RestoreContainer string
	GroupID          string
	SiteID           string
	UserID           string
	TeamSiteID       string
}

func EnvVars(ctx context.Context) Envs {
	folder := strings.TrimSpace(os.Getenv(sanityTestRestoreContainer))

	e := Envs{
		BackupID:         os.Getenv(sanityBackupID),
		SourceContainer:  os.Getenv(sanityTestSourceContainer),
		RestoreContainer: folder,
		GroupID:          tconfig.GetM365TeamID(ctx),
		SiteID:           tconfig.GetM365SiteID(ctx),
		UserID:           tconfig.GetM365UserID(ctx),
		TeamSiteID:       tconfig.GetM365TeamSiteID(ctx),
	}

	if len(os.Getenv(sanityTestUser)) > 0 {
		e.UserID = os.Getenv(sanityTestUser)
	}

	Infof(ctx, "test env vars: %+v", e)

	return e
}

func GetAC() (api.Client, error) {
	creds := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: os.Getenv(account.AzureTenantID),
	}

	return api.NewClient(creds, control.DefaultOptions(), count.New())
}

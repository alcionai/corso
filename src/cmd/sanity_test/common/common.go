package common

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type PermissionInfo struct {
	EntityID string
	Roles    []string
}

const (
	sanityBaseBackup  = "SANITY_BASE_BACKUP"
	sanityTestData    = "SANITY_TEST_DATA"
	sanityTestFolder  = "SANITY_TEST_FOLDER"
	sanityTestService = "SANITY_TEST_SERVICE"
	sanityTestUser    = "SANITY_TEST_USER"
)

type Envs struct {
	BaseBackupFolder string
	DataFolder       string
	FolderName       string
	GroupID          string
	Service          string
	SiteID           string
	StartTime        time.Time
	UserID           string
}

func EnvVars(ctx context.Context) Envs {
	folder := strings.TrimSpace(os.Getenv(sanityTestFolder))
	startTime, _ := MustGetTimeFromName(ctx, folder)

	e := Envs{
		BaseBackupFolder: os.Getenv(sanityBaseBackup),
		DataFolder:       os.Getenv(sanityTestData),
		FolderName:       folder,
		GroupID:          tconfig.GetM365TeamID(ctx),
		SiteID:           tconfig.GetM365SiteID(ctx),
		Service:          os.Getenv(sanityTestService),
		StartTime:        startTime,
		UserID:           tconfig.GetM365UserID(ctx),
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

	return api.NewClient(creds, control.DefaultOptions())
}

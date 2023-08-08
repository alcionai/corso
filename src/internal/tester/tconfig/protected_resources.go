package tconfig

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/pkg/logger"
)

// M365TenantID returns a tenantID string representing the azureTenantID described
// by either the env var AZURE_TENANT_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365TenantID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving m365 tenant ID from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgAzureTenantID])
}

// M365TenantID returns a tenantID string representing the azureTenantID described
// by either the env var AZURE_TENANT_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func GetM365TenantID(ctx context.Context) string {
	cfg, err := ReadTestConfig()
	if err != nil {
		logger.Ctx(ctx).Error(err, "retrieving m365 tenant ID from test configuration")
	}

	return strings.ToLower(cfg[TestCfgAzureTenantID])
}

// M365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M365_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365UserID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving m365 user id from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgUserID])
}

// GetM365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M365_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func GetM365UserID(ctx context.Context) string {
	cfg, err := ReadTestConfig()
	if err != nil {
		logger.Ctx(ctx).Error(err, "retrieving m365 user id from test configuration")
	}

	return strings.ToLower(cfg[TestCfgUserID])
}

// SecondaryM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_SECONDARY_M365_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func SecondaryM365UserID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving secondary m365 user id from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgSecondaryUserID])
}

// TertiaryM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_TERTIARY_M365_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func TertiaryM365UserID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving tertiary m365 user id from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgTertiaryUserID])
}

// LoadTestM365SiteID returns a siteID string representing the m365SiteID
// described by either the env var CORSO_M365_LOAD_TEST_SITE_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365SiteID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving load test m365 site id from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgSiteID])
}

// LoadTestM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_M365_LOAD_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365UserID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving load test m365 user id from test configuration", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgLoadTestUserID])
}

// expects cfg value to be a string representing an array such as:
// ["site1\,uuid","site2\,uuid"]
// the delimiter must be a |.
func LoadTestM365OrgSites(t *testing.T) []string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving load test m365 org sites from test configuration %+v", clues.ToCore(err))

	// TODO: proper handling of site slice input.
	// sites := cfg[TestCfgLoadTestOrgSites]
	// sites = strings.TrimPrefix(sites, "[")
	// sites = strings.TrimSuffix(sites, "]")
	// sites = strings.ReplaceAll(sites, `"`, "")
	// sites = strings.ReplaceAll(sites, `'`, "")
	// sites = strings.ReplaceAll(sites, "|", ",")

	// return strings.Split(sites, ",")

	return []string{strings.ToLower(cfg[TestCfgSiteID])}
}

// expects cfg value to be a string representing an array such as:
// ["foo@example.com","bar@example.com"]
// the delimiter may be either a , or |.
func LoadTestM365OrgUsers(t *testing.T) []string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving load test m365 org users from test configuration %+v", clues.ToCore(err))

	users := cfg[TestCfgLoadTestOrgUsers]
	users = strings.TrimPrefix(users, "[")
	users = strings.TrimSuffix(users, "]")
	users = strings.ReplaceAll(users, `"`, "")
	users = strings.ReplaceAll(users, `'`, "")
	users = strings.ReplaceAll(users, "|", ",")

	// a hack to skip using certain users when those accounts are
	// temporarily being co-opted for non-testing purposes.
	sl := strings.Split(users, ",")
	remove := os.Getenv("IGNORE_LOAD_TEST_USER_ID")

	if len(remove) == 0 {
		return sl
	}

	idx := -1

	for i, s := range sl {
		if s == remove {
			idx = i
			break
		}
	}

	return append(sl[:idx], sl[idx+1:]...)
}

// M365SiteID returns a siteID string representing the m365SiteID described
// by either the env var CORSO_M365_TEST_SITE_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365SiteID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving m365 site id from test configuration: %+v", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgSiteID])
}

// M365SiteURL returns a site webURL string representing the m365SiteURL described
// by either the env var CORSO_M365_TEST_SITE_URL, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365SiteURL(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving m365 site url from test configuration: %+v", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgSiteURL])
}

// GetM365SiteID returns a siteID string representing the m365SitteID described
// by either the env var CORSO_M365_TEST_SITE_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func GetM365SiteID(ctx context.Context) string {
	cfg, err := ReadTestConfig()
	if err != nil {
		logger.Ctx(ctx).Error(err, "retrieving m365 user id from test configuration")
	}

	return strings.ToLower(cfg[TestCfgSiteID])
}

// UnlicensedM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_M365_TEST_UNLICENSED_USER, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func UnlicensedM365UserID(t *testing.T) string {
	cfg, err := ReadTestConfig()
	require.NoError(t, err, "retrieving unlicensed m365 user id from test configuration: %+v", clues.ToCore(err))

	return strings.ToLower(cfg[TestCfgSecondaryUserID])
}

package tester

import (
	"os"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
)

// M365TenantID returns a tenantID string representing the azureTenantID described
// by either the env var AZURE_TENANT_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365TenantID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 user id from test configuration", clues.ToCore(err))

	return cfg[TestCfgAzureTenantID]
}

// M365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M365_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 user id from test configuration", clues.ToCore(err))

	return cfg[TestCfgUserID]
}

// SecondaryM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_SECONDARY_M365_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func SecondaryM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving secondary m365 user id from test configuration", clues.ToCore(err))

	return cfg[TestCfgSecondaryUserID]
}

// LoadTestM365SiteID returns a siteID string representing the m365SiteID
// described by either the env var CORSO_M365_LOAD_TEST_SITE_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365SiteID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 site id from test configuration", clues.ToCore(err))

	// TODO: load test site id, not standard test site id
	return cfg[TestCfgSiteID]
}

// LoadTestM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_M365_LOAD_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 user id from test configuration", clues.ToCore(err))

	return cfg[TestCfgLoadTestUserID]
}

// expects cfg value to be a string representing an array such as:
// ["site1\,uuid","site2\,uuid"]
// the delimeter must be a |.
func LoadTestM365OrgSites(t *testing.T) []string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 org sites from test configuration", clues.ToCore(err))

	// TODO: proper handling of site slice input.
	// sites := cfg[TestCfgLoadTestOrgSites]
	// sites = strings.TrimPrefix(sites, "[")
	// sites = strings.TrimSuffix(sites, "]")
	// sites = strings.ReplaceAll(sites, `"`, "")
	// sites = strings.ReplaceAll(sites, `'`, "")
	// sites = strings.ReplaceAll(sites, "|", ",")

	// return strings.Split(sites, ",")

	return []string{cfg[TestCfgSiteID]}
}

// expects cfg value to be a string representing an array such as:
// ["foo@example.com","bar@example.com"]
// the delimeter may be either a , or |.
func LoadTestM365OrgUsers(t *testing.T) []string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 org users from test configuration", clues.ToCore(err))

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
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 site id from test configuration", clues.ToCore(err))

	return cfg[TestCfgSiteID]
}

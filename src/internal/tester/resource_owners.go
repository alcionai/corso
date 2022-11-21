package tester

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// M365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M365_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 user id from test configuration")

	return cfg[TestCfgUserID]
}

// SecondaryM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_SECONDARY_M365_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func SecondaryM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving secondary m365 user id from test configuration")

	return cfg[TestCfgSecondaryUserID]
}

// LoadTestM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_M365_LOAD_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 user id from test configuration")

	return cfg[TestCfgLoadTestUserID]
}

// expects cfg value to be a string representing an array like:
// "['foo@example.com','bar@example.com']"
func LoadTestM365OrgUsers(t *testing.T) []string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 org users from test configuration")

	users := cfg[TestCfgLoadTestOrgUsers]
	users = strings.TrimPrefix(users, "[")
	users = strings.TrimSuffix(users, "]")
	users = strings.ReplaceAll(users, `"`, "")
	users = strings.ReplaceAll(users, `'`, "")
	users = strings.ReplaceAll(users, "|", ",")

	return strings.Split(users, ",")
}

// M365SiteID returns a siteID string representing the m365SiteID described
// by either the env var CORSO_M365_TEST_SITE_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365SiteID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 site id from test configuration")

	return cfg[TestCfgSiteID]
}

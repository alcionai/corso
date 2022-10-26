package tester

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// M365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M356_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving m365 user id from test configuration")

	return cfg[TestCfgUserID]
}

// SecondaryM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_SECONDARY_M356_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func SecondaryM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving secondary m365 user id from test configuration")

	return cfg[TestCfgSecondaryUserID]
}

// LoadTestM365UserID returns an userID string representing the m365UserID
// described by either the env var CORSO_M356_LOAD_TEST_USER_ID, the
// corso_test.toml config file or the default value (in that order of priority).
// The default is a last-attempt fallback that will only work on alcion's
// testing org.
func LoadTestM365UserID(t *testing.T) string {
	cfg, err := readTestConfig()
	require.NoError(t, err, "retrieving load test m365 user id from test configuration")

	return cfg[TestCfgLoadTestUserID]
}

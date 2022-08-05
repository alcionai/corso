package tester

import (
	"github.com/pkg/errors"
)

// M365UserID returns an userID string representing the m365UserID described
// by either the env var CORSO_M356_TEST_USER_ID, the corso_test.toml config
// file or the default value (in that order of priority).  The default is a
// last-attempt fallback that will only work on alcion's testing org.
func M365UserID() (string, error) {
	cfg, err := readTestConfig()
	if err != nil {
		return "", errors.Wrap(err, "retrieving m365 user id from test configuration")
	}

	return cfg[TestCfgUserID], nil
}

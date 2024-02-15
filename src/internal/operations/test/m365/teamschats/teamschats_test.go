package teamschats_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	. "github.com/alcionai/corso/src/internal/operations/test/m365"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type BackupIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestBackupIntgSuite(t *testing.T) {
	suite.Run(t, &BackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *BackupIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *BackupIntgSuite) TestBackup_Run_basicBackup() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunBasicBackupTest(suite, sel.Selector)
}

// ---------------------------------------------------------------------------
// nightly tests
// ---------------------------------------------------------------------------

type BackupNightlyIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestsBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &BackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *BackupNightlyIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *BackupNightlyIntgSuite) TestBackup_Run_vVersion9MergeBase() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunMergeBaseGroupsUpdate(suite, sel.Selector, true)
}

func (suite *BackupNightlyIntgSuite) TestBackup_Run_version9AssistBases() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunDriveAssistBaseGroupsUpdate(suite, sel.Selector, true)
}

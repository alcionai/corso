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

type TeamsChatsBackupIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestTeamsChatsBackupIntgSuite(t *testing.T) {
	suite.Run(t, &TeamsChatsBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *TeamsChatsBackupIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *TeamsChatsBackupIntgSuite) TestBackup_Run_basicBackup() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunBasicBackupTest(suite, sel.Selector)
}

// ---------------------------------------------------------------------------
// nightly tests
// ---------------------------------------------------------------------------

type TeamsChatsBackupNightlyIntgSuite struct {
	tester.Suite
	its IntgTesterSetup
}

func TestTeamsChatsBackupNightlyIntgSuite(t *testing.T) {
	suite.Run(t, &TeamsChatsBackupNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *TeamsChatsBackupNightlyIntgSuite) SetupSuite() {
	suite.its = NewIntegrationTesterSetup(suite.T())
}

func (suite *TeamsChatsBackupNightlyIntgSuite) TestBackup_Run_teamschatsVersion9MergeBase() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunMergeBaseGroupsUpdate(suite, sel.Selector, false)
}

func (suite *TeamsChatsBackupNightlyIntgSuite) TestBackup_Run_teamschatsVersion9AssistBases() {
	sel := selectors.NewTeamsChatsBackup([]string{suite.its.User.ID})
	sel.Include(selTD.TeamsChatsBackupChatScope(sel))

	RunDriveAssistBaseGroupsUpdate(suite, sel.Selector, false)
}

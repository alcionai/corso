package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type GroupUnitSuite struct {
	tester.Suite
}

func TestGroupsUnitSuite(t *testing.T) {
	suite.Run(t, &GroupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GroupUnitSuite) TestValidateGroup() {
	group := models.NewGroup()
	group.SetDisplayName(ptr.To("testgroup"))
	group.SetId(ptr.To("testID"))

	tests := []struct {
		name           string
		args           models.Groupable
		expectErr      assert.ErrorAssertionFunc
		errIsSkippable bool
	}{
		{
			name: "Valid group ",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			expectErr: assert.NoError,
		},
		{
			name: "No name",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetId(ptr.To("id"))
				return s
			}(),
			expectErr: assert.Error,
		},
		{
			name: "No ID",
			args: func() *models.Group {
				s := models.NewGroup()
				s.SetDisplayName(ptr.To("testgroup"))
				return s
			}(),
			expectErr: assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := api.ValidateGroup(test.args)
			test.expectErr(t, err, clues.ToCore(err))

			if test.errIsSkippable {
				assert.ErrorIs(t, err, api.ErrKnownSkippableCase)
			}
		})
	}
}

type GroupsIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestGroupsIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *GroupsIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *GroupsIntgSuite) TestGetAll() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	groups, err := suite.its.ac.
		Groups().
		GetAll(ctx, fault.New(true))
	require.NoError(t, err)
	require.NotZero(t, len(groups), "must have at least one group")
}

func (suite *GroupsIntgSuite) TestGroups_GetByID() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		groupID   = suite.its.group.id
		groupsAPI = suite.its.ac.Groups()
	)

	grp, err := groupsAPI.GetByID(ctx, groupID, api.CallConfig{})
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		id        string
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name:      "valid id",
			id:        groupID,
			expectErr: assert.NoError,
		},
		{
			name:      "invalid id",
			id:        uuid.NewString(),
			expectErr: assert.Error,
		},
		{
			name:      "valid display name",
			id:        ptr.Val(grp.GetDisplayName()),
			expectErr: assert.NoError,
		},
		{
			name:      "invalid displayName",
			id:        "jabberwocky",
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := groupsAPI.GetByID(ctx, test.id, api.CallConfig{})
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

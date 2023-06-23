package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ postItemer[models.Contactable] = &mockContactRestorer{}

type mockContactRestorer struct {
	postItemErr error
}

func (m mockContactRestorer) PostItem(
	ctx context.Context,
	userID, containerID string,
	body models.Contactable,
) (models.Contactable, error) {
	return models.NewContact(), m.postItemErr
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type ContactsRestoreIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestContactsRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &ContactsRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *ContactsRestoreIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

// Testing to ensure that cache system works for in multiple different environments
func (suite *ContactsRestoreIntgSuite) TestCreateContainerDestination() {
	runCreateDestinationTest(
		suite.T(),
		newContactRestoreHandler(suite.its.ac),
		path.ContactsCategory,
		suite.its.creds.AzureTenantID,
		suite.its.userID,
		testdata.DefaultRestoreConfig("").Location,
		[]string{"Hufflepuff"},
		[]string{"Ravenclaw"})
}

func (suite *ContactsRestoreIntgSuite) TestRestoreContact() {
	body := mock.ContactBytes("middlename")

	stub, err := api.BytesToContactable(body)
	require.NoError(suite.T(), err, clues.ToCore(err))

	collisionKey := api.ContactCollisionKey(stub)

	table := []struct {
		name         string
		apiMock      postItemer[models.Contactable]
		collisionMap map[string]string
		onCollision  control.CollisionPolicy
		expectErr    func(*testing.T, error)
	}{
		{
			name:         "no collision: skip",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "no collision: copy",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "no collision: replace",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "collision: skip",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
		},
		{
			name:         "collision: copy",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "collision: replace",
			apiMock:      mockContactRestorer{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := restoreContact(
				ctx,
				test.apiMock,
				body,
				suite.its.userID,
				"destination",
				test.collisionMap,
				test.onCollision,
				fault.New(true))

			test.expectErr(t, err)
		})
	}
}

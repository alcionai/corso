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

var _ mailRestorer = &mockMailRestorer{}

type mockMailRestorer struct {
	postItemErr       error
	postAttachmentErr error
}

func (m mockMailRestorer) PostItem(
	ctx context.Context,
	userID, containerID string,
	body models.Messageable,
) (models.Messageable, error) {
	return models.NewMessage(), m.postItemErr
}

func (m mockMailRestorer) PostSmallAttachment(
	_ context.Context,
	_, _, _ string,
	_ models.Attachmentable,
) error {
	return m.postAttachmentErr
}

func (m mockMailRestorer) PostLargeAttachment(
	_ context.Context,
	_, _, _, _ string,
	_ int64,
	_ models.Attachmentable,
) (models.UploadSessionable, error) {
	return models.NewUploadSession(), m.postAttachmentErr
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type MailRestoreIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestMailRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &MailRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *MailRestoreIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *MailRestoreIntgSuite) TestCreateContainerDestination() {
	runCreateDestinationTest(
		suite.T(),
		newMailRestoreHandler(suite.its.ac),
		path.EmailCategory,
		suite.its.creds.AzureTenantID,
		suite.its.userID,
		testdata.DefaultRestoreConfig("").Location,
		[]string{"Griffindor", "Croix"},
		[]string{"Griffindor", "Felicius"})
}

func (suite *MailRestoreIntgSuite) TestRestoreMail() {
	body := mock.MessageBytes("subject")

	stub, err := api.BytesToMessageable(body)
	require.NoError(suite.T(), err, clues.ToCore(err))

	collisionKey := api.MailCollisionKey(stub)

	table := []struct {
		name         string
		apiMock      mailRestorer
		collisionMap map[string]string
		onCollision  control.CollisionPolicy
		expectErr    func(*testing.T, error)
	}{
		{
			name:         "no collision: skip",
			apiMock:      mockMailRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "no collision: copy",
			apiMock:      mockMailRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "no collision: replace",
			apiMock:      mockMailRestorer{},
			collisionMap: map[string]string{},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "collision: skip",
			apiMock:      mockMailRestorer{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
		},
		{
			name:         "collision: copy",
			apiMock:      mockMailRestorer{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name:         "collision: replace",
			apiMock:      mockMailRestorer{},
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

			_, err := restoreMail(
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

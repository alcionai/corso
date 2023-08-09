package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ mailRestorer = &mailRestoreMock{}

type mailRestoreMock struct {
	postItemErr       error
	calledPost        bool
	deleteItemErr     error
	calledDelete      bool
	postAttachmentErr error
}

func (m *mailRestoreMock) PostItem(
	_ context.Context,
	_, _ string,
	_ models.Messageable,
) (models.Messageable, error) {
	m.calledPost = true
	return models.NewMessage(), m.postItemErr
}

func (m *mailRestoreMock) DeleteItem(
	_ context.Context,
	_, _ string,
) error {
	m.calledDelete = true
	return m.deleteItemErr
}

func (m *mailRestoreMock) PostSmallAttachment(
	_ context.Context,
	_, _, _ string,
	_ models.Attachmentable,
) error {
	return m.postAttachmentErr
}

func (m *mailRestoreMock) PostLargeAttachment(
	_ context.Context,
	_, _, _, _ string,
	_ []byte,
) (string, error) {
	return uuid.NewString(), m.postAttachmentErr
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
			[][]string{tconfig.M365AcctCredEnvs}),
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

	type counts struct {
		skip    int64
		replace int64
		new     int64
	}

	table := []struct {
		name         string
		apiMock      *mailRestoreMock
		collisionMap map[string]string
		onCollision  control.CollisionPolicy
		expectErr    func(*testing.T, error)
		expectMock   func(*testing.T, *mailRestoreMock)
		expectCounts counts
	}{
		{
			name:         "no collision: skip",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "no collision: copy",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "no collision: replace",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "collision: skip",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.False(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{1, 0, 0},
		},
		{
			name:         "collision: copy",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "collision: replace",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.True(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 1, 0},
		},
		{
			name:         "collision: replace - err already deleted",
			apiMock:      &mailRestoreMock{deleteItemErr: graph.ErrDeletedInFlight},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.True(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 1, 0},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctr := count.New()

			_, err := restoreMail(
				ctx,
				test.apiMock,
				body,
				suite.its.userID,
				"destination",
				test.collisionMap,
				test.onCollision,
				fault.New(true),
				ctr)

			test.expectErr(t, err)
			test.expectMock(t, test.apiMock)
			assert.Equal(t, test.expectCounts.skip, ctr.Get(count.CollisionSkip), "skips")
			assert.Equal(t, test.expectCounts.replace, ctr.Get(count.CollisionReplace), "replaces")
			assert.Equal(t, test.expectCounts.new, ctr.Get(count.NewItemCreated), "new items")
		})
	}
}

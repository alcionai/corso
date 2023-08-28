package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type commonM365UnitSuite struct {
	tester.Suite
}

func TestM365UnitSuite(t *testing.T) {
	suite.Run(t, &commonM365UnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockDGDD struct {
	response models.Driveable
	err      error
}

func (m mockDGDD) GetDefaultDrive(context.Context, string) (models.Driveable, error) {
	return m.response, m.err
}

func (suite *commonM365UnitSuite) TestIsOneDriveServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getDefaultDriver
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getDefaultDriver {
				return mockDGDD{models.NewDrive(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mysite not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.MysiteNotFound)))
				odErr.SetErrorEscaped(merr)

				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mysite URL not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.MysiteURLNotFound)))
				odErr.SetErrorEscaped(merr)

				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.NoSPLicense)))
				odErr.SetErrorEscaped(merr)

				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To(string(graph.RequestResourceNotFound)))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			dgdd := test.mock(ctx)

			ok, err := IsOneDriveServiceEnabled(ctx, dgdd, "foo")
			test.expect(t, ok, "has drives flag")
			test.expectErr(t, err)
		})
	}
}

type mockGMI struct {
	response models.MailFolderable
	err      error
}

func (m mockGMI) GetMailInbox(context.Context, string) (models.MailFolderable, error) {
	return m.response, m.err
}

func (suite *commonM365UnitSuite) TestIsExchangeServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getMailInboxer
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getMailInboxer {
				return mockGMI{models.NewMailFolder(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user has no mailbox",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To(string(graph.ResourceNotFound)))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockGMI{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user not found",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To(string(graph.RequestResourceNotFound)))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockGMI{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "overlapping resourcenotfound",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To(string(graph.ResourceNotFound)))
				merr.SetMessage(ptr.To("User not found"))
				odErr.SetErrorEscaped(merr)

				return mockGMI{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockGMI{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gmi := test.mock(ctx)

			ok, err := IsExchangeServiceEnabled(ctx, gmi, "foo")
			test.expect(t, ok, "has mailbox flag")
			test.expectErr(t, err)
		})
	}
}

// Test IsSharePointServiceEnabled
type mockGSR struct {
	response models.Siteable
	err      error
}

func (m mockGSR) GetRoot(context.Context) (models.Siteable, error) {
	return m.response, m.err
}

func (suite *commonM365UnitSuite) TestIsSharePointServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getSiteRooter
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getSiteRooter {
				return mockGSR{models.NewSite(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.NoSPLicense)))
				odErr.SetErrorEscaped(merr)

				return mockGSR{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetErrorEscaped(merr)

				return mockGSR{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gsr := test.mock(ctx)

			ok, err := IsSharePointServiceEnabled(ctx, gsr, "foo")
			test.expect(t, ok, "has sites flag")
			test.expectErr(t, err)
		})
	}
}

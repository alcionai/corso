package graph

import (
	"context"
	"net/http"
	"strings"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/custom"
)

type GraphErrorsUnitSuite struct {
	tester.Suite
}

func TestGraphErrorsUnitSuite(t *testing.T) {
	suite.Run(t, &GraphErrorsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GraphErrorsUnitSuite) TestIsErrConnectionReset() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "matching",
			err:    syscall.ECONNRESET,
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrConnectionReset(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrApplicationThrottled() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "applicationThrottled oDataErr",
			err:    graphTD.ODataErr(string(ApplicationThrottled)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrApplicationThrottled(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrAuthenticationError() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "authenticationError oDataErr",
			err:    graphTD.ODataErr(string(AuthenticationError)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrAuthenticationError(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrInsufficientAuthorization() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "AuthorizationRequestDenied oDataErr",
			err:    graphTD.ODataErr(string(AuthorizationRequestDenied)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrInsufficientAuthorization(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrDeletedInFlight() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "as",
			err:    ErrDeletedInFlight,
			expect: assert.True,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "not-found oDataErr",
			err:    graphTD.ODataErr(string(ErrorItemNotFound)),
			expect: assert.True,
		},
		{
			name:   "sync-not-found oDataErr",
			err:    graphTD.ODataErr(string(syncFolderNotFound)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrDeletedInFlight(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) Test() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "invalid receipient oDataErr",
			err:    graphTD.ODataErr(string(ErrorInvalidRecipients)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrInvalidRecipients(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrInvalidDelta() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErrMsg",
			err:    graphTD.ODataErrWithMsg("fnords", "deltatoken not supported"),
			expect: assert.False,
		},
		{
			name:   "resync-required oDataErr",
			err:    graphTD.ODataErr(string(resyncRequired)),
			expect: assert.True,
		},
		{
			name:   "sync state invalid oDataErr",
			err:    graphTD.ODataErr(string(syncStateInvalid)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "resync-required oDataErr camelcase",
			err:    graphTD.ODataErr("resyncRequired"),
			expect: assert.True,
		},
		{
			name:   "resync-required oDataErr lowercase",
			err:    graphTD.ODataErr("resyncrequired"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrInvalidDelta(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrDeltaNotSupported() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErrMsg",
			err:    graphTD.ODataErrWithMsg("fnords", "deltatoken not supported"),
			expect: assert.False,
		},
		{
			name:   "deltatoken not supported oDataErrMsg",
			err:    graphTD.ODataErrWithMsg("fnords", string(ParameterDeltaTokenNotSupported)),
			expect: assert.True,
		},
		{
			name:   "deltatoken not supported oDataErrMsg with punctuation",
			err:    graphTD.ODataErrWithMsg("fnords", string(ParameterDeltaTokenNotSupported)+"."),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrDeltaNotSupported(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrQuotaExceeded() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "quota-exceeded oDataErr",
			err:    graphTD.ODataErr("ErrorQuotaExceeded"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrQuotaExceeded(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrUserNotFound() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name: "non-matching resource not found",
			err: func() error {
				res := graphTD.ODataErr(string(ResourceNotFound))
				res.GetErrorEscaped().SetMessage(ptr.To("Calendar not found"))

				return res
			}(),
			expect: assert.False,
		},
		{
			name:   "request resource not found oDataErr",
			err:    graphTD.ODataErr(string(RequestResourceNotFound)),
			expect: assert.True,
		},
		{
			name:   "invalid user oDataErr",
			err:    graphTD.ODataErr(string(invalidUser)),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr",
			err: func() error {
				res := graphTD.ODataErrWithMsg(string(ResourceNotFound), "User not found")
				return res
			}(),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr wrapped",
			err: func() error {
				res := graphTD.ODataErrWithMsg(string(ResourceNotFound), "User not found")
				return clues.Wrap(res, "getting mail folder")
			}(),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr stacked",
			err: func() error {
				res := graphTD.ODataErrWithMsg(string(ResourceNotFound), "User not found")
				return clues.Stack(res, assert.AnError)
			}(),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUserNotFound(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrTimeout() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "context deadline",
			err:    context.DeadlineExceeded,
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrTimeout(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrUnauthorizedOrBadToken() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("folder doesn't exist"),
			expect: assert.False,
		},
		{
			name: "graph 401",
			err: clues.Stack(assert.AnError).
				Label(LabelStatus(http.StatusUnauthorized)),
			expect: assert.True,
		},
		{
			name:   "err token expired",
			err:    clues.Stack(assert.AnError, ErrTokenExpired),
			expect: assert.True,
		},
		{
			name:   "oDataErr code invalid auth token ",
			err:    graphTD.ODataErr(string(invalidAuthenticationToken)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUnauthorizedOrBadToken(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrIsErrBadJWTToken() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("folder doesn't exist"),
			expect: assert.False,
		},
		{
			name: "graph 401",
			err: clues.Stack(assert.AnError).
				Label(LabelStatus(http.StatusUnauthorized)),
			expect: assert.False,
		},
		{
			name:   "err token expired",
			err:    clues.Stack(assert.AnError, ErrTokenExpired),
			expect: assert.False,
		},
		{
			name:   "oDataErr code invalid auth token ",
			err:    graphTD.ODataErr(string(invalidAuthenticationToken)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrBadJWTToken(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestMalwareInfo() {
	var (
		i         = models.NewDriveItem()
		createdBy = models.NewUser()
		cbID      = "created-by"
		lm        = models.NewUser()
		lmID      = "last-mod-by"
		ref       = models.NewItemReference()
		refCID    = "container-id"
		refCN     = "container-name"
		refCP     = "/drives/b!vF-sdsdsds-sdsdsa-sdsd/root:/Folder/container-name"
		refCPexp  = "/Folder/container-name"
		mal       = models.NewMalware()
		malDesc   = "malware-description"
	)

	createdBy.SetId(&cbID)
	i.SetCreatedByUser(createdBy)

	lm.SetId(&lmID)
	i.SetLastModifiedByUser(lm)

	ref.SetId(&refCID)
	ref.SetName(&refCN)
	ref.SetPath(&refCP)
	i.SetParentReference(ref)

	mal.SetDescription(&malDesc)
	i.SetMalware(mal)

	expect := map[string]any{
		fault.AddtlCreatedBy:     cbID,
		fault.AddtlLastModBy:     lmID,
		fault.AddtlContainerID:   refCID,
		fault.AddtlContainerName: refCN,
		fault.AddtlContainerPath: refCPexp,
		fault.AddtlMalwareDesc:   malDesc,
	}

	assert.Equal(suite.T(), expect, ItemInfo(custom.ToCustomDriveItem(i)))
}

func (suite *GraphErrorsUnitSuite) TestIsErrFolderExists() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("folder doesn't exist"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr msg",
			err:    graphTD.ODataErr(string(folderExists)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "oDataErr camelcase",
			err:    graphTD.ODataErr("ErrorFolderExists"),
			expect: assert.True,
		},
		{
			name:   "oDataErr lowercase",
			err:    graphTD.ODataErr("errorfolderexists"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrFolderExists(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrUsersCannotBeResolved() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", "cant resolve users"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr code",
			err:    graphTD.ODataErrWithMsg(string(noResolvedUsers), "usersCannotBeResolved"),
			expect: assert.True,
		},
		{
			name:   "matching oDataErr msg",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", string(usersCannotBeResolved)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "oDataErr uppercase",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", strings.ToUpper(string(usersCannotBeResolved))),
			expect: assert.True,
		},
		{
			name:   "oDataErr lowercase",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", strings.ToLower(string(usersCannotBeResolved))),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUsersCannotBeResolved(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrSiteCouldNotBeFound() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", "cant resolve sites"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr msg",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", string(requestedSiteCouldNotBeFound)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "oDataErr uppercase",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", strings.ToUpper(string(requestedSiteCouldNotBeFound))),
			expect: assert.True,
		},
		{
			name:   "oDataErr lowercase",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", strings.ToLower(string(requestedSiteCouldNotBeFound))),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrSiteNotFound(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrCannotOpenFileAttachment() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "quota-exceeded oDataErr",
			err:    graphTD.ODataErr(string(cannotOpenFileAttachment)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrCannotOpenFileAttachment(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestGraphStack_labels() {
	table := []struct {
		name   string
		err    error
		expect []string
	}{
		{
			name:   "nil",
			err:    nil,
			expect: []string{},
		},
		{
			name:   "not-odata",
			err:    assert.AnError,
			expect: []string{},
		},
		{
			name:   "oDataErr matches no labels",
			err:    graphTD.ODataErr("code"),
			expect: []string{},
		},
		{
			name:   "mysite not found",
			err:    graphTD.ODataErrWithMsg("code", string(MysiteNotFound)),
			expect: []string{LabelsMysiteNotFound},
		},
		{
			name:   "mysite url not found",
			err:    graphTD.ODataErrWithMsg("code", string(MysiteURLNotFound)),
			expect: []string{LabelsMysiteNotFound},
		},
		{
			name:   "no sp license",
			err:    graphTD.ODataErrWithMsg("code", string(NoSPLicense)),
			expect: []string{LabelsNoSharePointLicense},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result := Stack(ctx, test.err)

			for _, e := range test.expect {
				assert.True(t, clues.HasLabel(result, e), clues.ToCore(result))
			}

			labels := clues.Labels(result)
			assert.Equal(t,
				len(test.expect), len(labels),
				"result should have as many labels as expected")
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrItemNotFound() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "item not found oDataErr",
			err:    graphTD.ODataErr(string(ItemNotFound)),
			expect: assert.True,
		},
		{
			name:   "error item not found oDataErr",
			err:    graphTD.ODataErr(string(ErrorItemNotFound)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrItemNotFound(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrResourceLocked() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", "resource is locked"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr code",
			err:    graphTD.ODataErr(string(NotAllowed)),
			expect: assert.True,
		},
		{
			name:   "matching oDataErr inner code",
			err:    graphTD.ODataInner(string(ResourceLocked)),
			expect: assert.True,
		},
		{
			name: "matching oDataErr message",
			err: graphTD.ODataErrWithMsg(
				string(AuthenticationError),
				"AADSTS500014: The service principal for resource 'beefe6b7-f5df-413d-ac2d-abf1e3fd9c0b' "+
					"is disabled. This indicate that a subscription within the tenant has lapsed, or that the "+
					"administrator for this tenant has disabled the application, preventing tokens from being "+
					"issued for it. Trace ID: dead78e1-0830-4edf-bea7-f0a445620100 Correlation ID: "+
					"deadbeef-7f1e-4578-8215-36004a2c935c Timestamp: 2023-12-05 19:31:01Z"),
			expect: assert.True,
		},
		{
			name:   "matching err sentinel",
			err:    ErrResourceLocked,
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrResourceLocked(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrSharingDisabled() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    graphTD.ODataErrWithMsg("InvalidRequest", "resource is locked"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr inner code",
			err:    graphTD.ODataInner(string(sharingDisabled)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrSharingDisabled(test.err))
		})
	}
}

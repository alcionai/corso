package graph

import (
	"context"
	"encoding/json"
	"net/http"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
)

type GraphErrorsUnitSuite struct {
	tester.Suite
}

func TestGraphErrorsUnitSuite(t *testing.T) {
	suite.Run(t, &GraphErrorsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func odErr(code string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func odErrMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func parseableToMap(t *testing.T, thing serialization.Parsable) map[string]any {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "deserialize")

	var out map[string]any
	err = json.Unmarshal([]byte(content), &out)
	require.NoError(t, err, "unmarshall")

	return out
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
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "authenticationError oDataErr",
			err:    odErr(string(AuthenticationError)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrAuthenticationError(test.err))
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
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "not-found oDataErr",
			err:    odErr(string(errorItemNotFound)),
			expect: assert.True,
		},
		{
			name:   "sync-not-found oDataErr",
			err:    odErr(string(syncFolderNotFound)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrDeletedInFlight(test.err))
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
			name:   "as",
			err:    ErrInvalidDelta,
			expect: assert.True,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "resync-required oDataErr",
			err:    odErr(string(resyncRequired)),
			expect: assert.True,
		},
		{
			name:   "sync state invalid oDataErr",
			err:    odErr(string(syncStateInvalid)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "resync-required oDataErr camelcase",
			err:    odErr("resyncRequired"),
			expect: assert.True,
		},
		{
			name:   "resync-required oDataErr lowercase",
			err:    odErr("resyncrequired"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrInvalidDelta(test.err))
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
			name:   "as",
			err:    ErrInvalidDelta,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "quota-exceeded oDataErr",
			err:    odErr("ErrorQuotaExceeded"),
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
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name: "non-matching resource not found",
			err: func() error {
				res := odErr(string(ResourceNotFound))
				res.GetErrorEscaped().SetMessage(ptr.To("Calendar not found"))

				return res
			}(),
			expect: assert.False,
		},
		{
			name:   "request resource not found oDataErr",
			err:    odErr(string(RequestResourceNotFound)),
			expect: assert.True,
		},
		{
			name:   "invalid user oDataErr",
			err:    odErr(string(invalidUser)),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr",
			err: func() error {
				res := odErrMsg(string(ResourceNotFound), "User not found")
				return res
			}(),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr wrapped",
			err: func() error {
				res := odErrMsg(string(ResourceNotFound), "User not found")
				return clues.Wrap(res, "getting mail folder")
			}(),
			expect: assert.True,
		},
		{
			name: "resource not found oDataErr stacked",
			err: func() error {
				res := odErrMsg(string(ResourceNotFound), "User not found")
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

// TODO: Test for IsErrExchangeMailFolderNotFound
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
			name:   "as",
			err:    ErrTimeout,
			expect: assert.True,
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

func (suite *GraphErrorsUnitSuite) TestIsErrUnauthorized() {
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
			name: "as",
			err: clues.Stack(assert.AnError).
				Label(LabelStatus(http.StatusUnauthorized)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUnauthorized(test.err))
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

	assert.Equal(suite.T(), expect, ItemInfo(i))
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
			err:    odErr("folder doesn't exist"),
			expect: assert.False,
		},
		{
			name:   "matching oDataErr",
			err:    odErr(string(folderExists)),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "oDataErr camelcase",
			err:    odErr("ErrorFolderExists"),
			expect: assert.True,
		},
		{
			name:   "oDataErr lowercase",
			err:    odErr("errorfolderexists"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrFolderExists(test.err))
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
			name:   "as",
			err:    ErrInvalidDelta,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "quota-exceeded oDataErr",
			err:    odErr(string(cannotOpenFileAttachment)),
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
			err:    odErr("code"),
			expect: []string{},
		},
		{
			name:   "mysite not found",
			err:    odErrMsg("code", string(MysiteNotFound)),
			expect: []string{LabelsMysiteNotFound},
		},
		{
			name:   "mysite url not found",
			err:    odErrMsg("code", string(MysiteURLNotFound)),
			expect: []string{LabelsMysiteNotFound},
		},
		{
			name:   "no sp license",
			err:    odErrMsg("code", string(NoSPLicense)),
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
			name:   "as",
			err:    ErrInvalidDelta,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "item nott found oDataErr",
			err:    odErr(string(itemNotFound)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrItemNotFound(test.err))
		})
	}
}

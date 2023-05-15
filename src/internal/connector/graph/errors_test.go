package graph

import (
	"context"
	"net/http"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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
	odErr := &odataerrors.ODataError{}
	merr := odataerrors.MainError{}
	merr.SetCode(&code)
	odErr.SetError(&merr)

	return odErr
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
			err:    odErr(string(itemNotFound)),
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
			name:   "request resource not found oDataErr",
			err:    odErr(string(requestResourceNotFound)),
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
		i        = models.DriveItem{}
		cb       = models.User{}
		cbID     = "created-by"
		lm       = models.User{}
		lmID     = "last-mod-by"
		ref      = models.ItemReference{}
		refCID   = "container-id"
		refCN    = "container-name"
		refCP    = "/drives/b!vF-sdsdsds-sdsdsa-sdsd/root:/Folder/container-name"
		refCPexp = "/Folder/container-name"
		mal      = models.Malware{}
		malDesc  = "malware-description"
	)

	cb.SetId(&cbID)
	i.SetCreatedByUser(&cb)

	lm.SetId(&lmID)
	i.SetLastModifiedByUser(&lm)

	ref.SetId(&refCID)
	ref.SetName(&refCN)
	ref.SetPath(&refCP)
	i.SetParentReference(&ref)

	mal.SetDescription(&malDesc)
	i.SetMalware(&mal)

	expect := map[string]any{
		fault.AddtlCreatedBy:     cbID,
		fault.AddtlLastModBy:     lmID,
		fault.AddtlContainerID:   refCID,
		fault.AddtlContainerName: refCN,
		fault.AddtlContainerPath: refCPexp,
		fault.AddtlMalwareDesc:   malDesc,
	}

	assert.Equal(suite.T(), expect, ItemInfo(&i))
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

package exchange

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type MailBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestMailBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &MailBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MailBackupHandlerUnitSuite) TestHandler_CanSkipItemFailure() {
	resourceID := uuid.NewString()

	table := []struct {
		name        string
		err         error
		opts        control.Options
		expect      assert.BoolAssertionFunc
		expectCause fault.SkipCause
	}{
		{
			name:   "no config",
			err:    assert.AnError,
			opts:   control.Options{},
			expect: assert.False,
		},
		{
			name: "false when map is empty",
			err:  assert.AnError,
			opts: control.Options{
				SkipEventsOnInstance503ForResources: map[string]struct{}{},
			},
			expect: assert.False,
		},
		{
			name: "false on nil error",
			err:  nil,
			opts: control.Options{
				SkipEventsOnInstance503ForResources: map[string]struct{}{
					resourceID: {},
				},
			},
			expect: assert.False,
		},
		{
			name: "false even if resource matches",
			err:  assert.AnError,
			opts: control.Options{
				SkipEventsOnInstance503ForResources: map[string]struct{}{
					resourceID: {},
				},
			},
			expect: assert.False,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			h := newMailBackupHandler(api.Client{})
			cause, result := h.CanSkipItemFailure(
				test.err,
				resourceID,
				test.opts)

			test.expect(t, result)
			assert.Equal(t, test.expectCause, cause)
		})
	}
}

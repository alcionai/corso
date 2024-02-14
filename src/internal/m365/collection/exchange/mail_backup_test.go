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
	var (
		resourceID = uuid.NewString()
		itemID     = uuid.NewString()
	)

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
				SkipTheseEventsOnInstance503: map[string][]string{},
			},
			expect: assert.False,
		},
		{
			name: "false on nil error",
			err:  nil,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {"bar", itemID},
				},
			},
			expect: assert.False,
		},
		{
			name: "false even if item matches",
			err:  assert.AnError,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {itemID},
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
				itemID,
				test.opts)

			test.expect(t, result)
			assert.Equal(t, test.expectCause, cause)
		})
	}
}

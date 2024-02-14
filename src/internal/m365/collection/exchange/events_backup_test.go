package exchange

import (
	"net/http"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type EventsBackupHandlerUnitSuite struct {
	tester.Suite
}

func TestEventsBackupHandlerUnitSuite(t *testing.T) {
	suite.Run(t, &EventsBackupHandlerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *EventsBackupHandlerUnitSuite) TestHandler_CanSkipItemFailure() {
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
			err:    graph.ErrServiceUnavailableEmptyResp,
			opts:   control.Options{},
			expect: assert.False,
		},
		{
			name: "empty skip on 503",
			err:  graph.ErrServiceUnavailableEmptyResp,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{},
			},
			expect: assert.False,
		},
		{
			name: "nil error",
			err:  nil,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {"bar", itemID},
				},
			},
			expect: assert.False,
		},
		{
			name: "non-matching resource",
			err:  graph.ErrServiceUnavailableEmptyResp,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					"foo": {"bar", itemID},
				},
			},
			expect: assert.False,
		},
		{
			name: "non-matching item",
			err:  graph.ErrServiceUnavailableEmptyResp,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {"bar", "baz"},
				},
			},
			expect: assert.False,
			// the item won't match, but we still return this as the cause
			expectCause: fault.SkipKnownEventInstance503s,
		},
		{
			name: "match on instance 503 empty resp",
			err:  graph.ErrServiceUnavailableEmptyResp,
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {"bar", itemID},
				},
			},
			expect:      assert.True,
			expectCause: fault.SkipKnownEventInstance503s,
		},
		{
			name: "match on instance 503",
			err: clues.New("arbitrary error").
				Label(graph.LabelStatus(http.StatusServiceUnavailable)),
			opts: control.Options{
				SkipTheseEventsOnInstance503: map[string][]string{
					resourceID: {"bar", itemID},
				},
			},
			expect:      assert.True,
			expectCause: fault.SkipKnownEventInstance503s,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			h := newEventBackupHandler(api.Client{})
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

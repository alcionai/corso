package api

import (
	"fmt"
	"net/http"
	stdpath "path"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/count"
)

// ---------------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------------

type EventsPagerUnitSuite struct {
	tester.Suite
}

func TestEventsPagerUnitSuite(t *testing.T) {
	suite.Run(t, &EventsPagerUnitSuite{
		Suite: tester.NewUnitSuite(t),
	})
}

func (suite *EventsPagerUnitSuite) TestGetAddedAndRemovedItemIDs_SendsCorrectDeltaPageSize() {
	const (
		validEmptyResponse = `{
  "@odata.context": "https://graph.microsoft.com/beta/$metadata#Collection(event)",
  "value": [],
  "@odata.deltaLink": "link"
}`

		// deltaPath helps make gock matching a little easier since it splits out
		// the hostname from the remainder of the URL. Graph SDK uses the URL
		// directly though.
		deltaPath = "/prev-delta"
		prevDelta = graphAPIHostURL + deltaPath

		userID      = "user-id"
		containerID = "container-id"
	)

	deltaTests := []struct {
		name       string
		reqPath    string
		inputDelta string
	}{
		{
			name: "NoPrevDelta",
			reqPath: stdpath.Join(
				"/beta",
				"users",
				userID,
				"calendars",
				containerID,
				"events",
				"delta"),
		},
		{
			name:       "HasPrevDelta",
			reqPath:    deltaPath,
			inputDelta: prevDelta,
		},
	}

	for _, deltaTest := range deltaTests {
		suite.Run(deltaTest.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			t.Cleanup(flush)

			a := tconfig.NewFakeM365Account(t)
			creds, err := a.M365Config()
			require.NoError(t, err, clues.ToCore(err))

			client, err := gockClient(creds, count.New())
			require.NoError(t, err, clues.ToCore(err))

			t.Cleanup(gock.Off)

			// Number of retries and delay between retries is handled by a kiota
			// middleware. We can change the default config parameters when setting up
			// the mock in a later PR.
			gock.New(graphAPIHostURL).
				Get(deltaTest.reqPath).
				Times(4).
				Reply(http.StatusServiceUnavailable).
				BodyString("").
				Type("text/plain")

			gock.New(graphAPIHostURL).
				Get(deltaTest.reqPath).
				SetMatcher(gock.NewMatcher()).
				// Need a custom Matcher since the prefer header is also used for
				// immutable ID behavior.
				AddMatcher(func(got *http.Request, want *gock.Request) (bool, error) {
					var (
						found         bool
						preferHeaders = got.Header.Values("Prefer")
						expected      = fmt.Sprintf(
							"odata.maxpagesize=%d",
							minEventsDeltaPageSize)
					)

					for _, header := range preferHeaders {
						if strings.Contains(header, expected) {
							found = true
							break
						}
					}

					assert.Truef(
						t,
						found,
						"header %s not found in set %v",
						expected,
						preferHeaders)

					return true, nil
				}).
				Reply(http.StatusOK).
				JSON(validEmptyResponse)

			_, err = client.Events().GetAddedAndRemovedItemIDs(
				ctx,
				userID,
				containerID,
				deltaTest.inputDelta,
				CallConfig{
					CanMakeDeltaQueries: true,
				})
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

type EventsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestEventsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &EventsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *EventsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *EventsPagerIntgSuite) TestEvents_GetItemsInContainerByCollisionKey() {
	t := suite.T()
	ac := suite.its.ac.Events()

	ctx, flush := tester.NewContext(t)
	defer flush()

	container, err := ac.GetContainerByID(ctx, suite.its.user.id, "calendar")
	require.NoError(t, err, clues.ToCore(err))

	evts, err := ac.Stable.
		Client().
		Users().
		ByUserId(suite.its.user.id).
		Calendars().
		ByCalendarId(ptr.Val(container.GetId())).
		Events().
		Get(ctx, nil)
	require.NoError(t, err, clues.ToCore(err))

	es := evts.GetValue()
	expectM := map[string]struct{}{}

	for _, e := range es {
		expectM[EventCollisionKey(e)] = struct{}{}
	}

	expect := maps.Keys(expectM)

	results, err := suite.its.ac.
		Events().
		GetItemsInContainerByCollisionKey(ctx, suite.its.user.id, "calendar")
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(results), "requires at least one result")

	for k, v := range results {
		assert.NotEmpty(t, k, "all keys should be populated")
		assert.NotEmpty(t, v, "all values should be populated")
	}

	for _, k := range expect {
		t.Log("expects key", k)
	}

	for k := range results {
		t.Log("results key", k)
	}

	for _, e := range expect {
		_, ok := results[e]
		assert.Truef(t, ok, "expected results to contain collision key: %s", e)
	}
}

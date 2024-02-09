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

func (suite *EventsPagerUnitSuite) TestEventsList() {
	const (
		nextLinkPath = "/next-link"
		nextLinkURL  = graphAPIHostURL + nextLinkPath

		nextDeltaURL = graphAPIHostURL + "/next-delta"

		validEventsListSingleNextLinkResponse = `{
  "@odata.context": "https://graph.microsoft.com/beta/$metadata#Collection(event)",
  "value": [{"id":"foo"}],
  "@odata.nextLink": "` + nextLinkURL + `"
}`
		validEventsListEmptyResponse = `{
  "@odata.context": "https://graph.microsoft.com/beta/$metadata#Collection(event)",
  "value": [],
  "@odata.deltaLink": "` + nextDeltaURL + `"
}`

		// deltaPath helps make gock matching a little easier since it splits out
		// the hostname from the remainder of the URL. Graph SDK uses the URL
		// directly though.
		deltaPath = "/prev-delta"
		prevDelta = graphAPIHostURL + deltaPath

		userID      = "user-id"
		containerID = "container-id"
	)

	pageSizeMatcher := func(
		t *testing.T,
		expectedPageSize int32,
	) func(*http.Request, *gock.Request) (bool, error) {
		return func(got *http.Request, want *gock.Request) (bool, error) {
			var (
				found         bool
				preferHeaders = got.Header.Values("Prefer")
				expected      = fmt.Sprintf(
					"odata.maxpagesize=%d",
					expectedPageSize)
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
		}
	}

	// configureFailedRequests configures the HTTP mocker to return one successful
	// request with a single item and then enough failed requests to exhaust the
	// retries. This allows testing that results are not mixed between fallbacks.
	configureFailedRequests := func(
		t *testing.T,
		reqPath string,
		expectedPageSize int32,
	) {
		gock.New(graphAPIHostURL).
			Get(reqPath).
			AddMatcher(pageSizeMatcher(t, expectedPageSize)).
			Reply(http.StatusOK).
			JSON(validEventsListSingleNextLinkResponse)

		// Number of retries and delay between retries is handled by a kiota
		// middleware. We can change the default config parameters when setting
		// up the mock in a later PR.
		gock.New(graphAPIHostURL).
			Get(nextLinkPath).
			AddMatcher(pageSizeMatcher(t, expectedPageSize)).
			Times(4).
			Reply(http.StatusServiceUnavailable).
			BodyString("").
			Type("text/plain")
	}

	deltaTests := []struct {
		name               string
		inputDelta         string
		configureMocks     func(t *testing.T, userID string, containerID string)
		expectNextDeltaURL string
		expectDeltaReset   assert.BoolAssertionFunc
	}{
		{
			name: "NoPrevDelta DeltaFallback",
			configureMocks: func(t *testing.T, userID string, containerID string) {
				reqPath := stdpath.Join(
					"/beta",
					"users",
					userID,
					"calendars",
					containerID,
					"events",
					"delta")

				configureFailedRequests(t, reqPath, maxDeltaPageSize)

				gock.New(graphAPIHostURL).
					Get(reqPath).
					SetMatcher(gock.NewMatcher()).
					// Need a custom Matcher since the prefer header is also used for
					// immutable ID behavior.
					AddMatcher(pageSizeMatcher(t, minEventsDeltaPageSize)).
					Reply(http.StatusOK).
					JSON(validEventsListEmptyResponse)
			},
			expectNextDeltaURL: nextDeltaURL,
			// OK to be true for this since we didn't have a delta to start with.
			expectDeltaReset: assert.True,
		},
		{
			name:       "PrevDelta DeltaFallback",
			inputDelta: prevDelta,
			configureMocks: func(t *testing.T, userID string, containerID string) {
				// Number of retries and delay between retries is handled by a kiota
				// middleware. We can change the default config parameters when setting
				// up the mock in a later PR.
				configureFailedRequests(t, deltaPath, maxDeltaPageSize)

				gock.New(graphAPIHostURL).
					Get(deltaPath).
					SetMatcher(gock.NewMatcher()).
					// Need a custom Matcher since the prefer header is also used for
					// immutable ID behavior.
					AddMatcher(pageSizeMatcher(t, minEventsDeltaPageSize)).
					Reply(http.StatusOK).
					JSON(validEventsListEmptyResponse)
			},
			expectNextDeltaURL: nextDeltaURL,
			expectDeltaReset:   assert.False,
		},
		{
			name:       "PrevDelta SecondaryNonDeltaFallback",
			inputDelta: prevDelta,
			configureMocks: func(t *testing.T, userID string, containerID string) {
				// Number of retries and delay between retries is handled by a kiota
				// middleware. We can change the default config parameters when setting
				// up the mock in a later PR.
				configureFailedRequests(t, deltaPath, maxDeltaPageSize)

				// Smaller page size delta fallback.
				configureFailedRequests(t, deltaPath, minEventsDeltaPageSize)

				// Non delta endpoint fallback
				gock.New(graphAPIHostURL).
					Get(v1APIURLPath(
						"users",
						userID,
						"calendars",
						containerID,
						"events")).
					SetMatcher(gock.NewMatcher()).
					// Need a custom Matcher since the prefer header is also used for
					// immutable ID behavior.
					AddMatcher(pageSizeMatcher(t, maxNonDeltaPageSize)).
					Reply(http.StatusOK).
					JSON(validEventsListEmptyResponse)
			},
			expectDeltaReset: assert.True,
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

			deltaTest.configureMocks(t, userID, containerID)

			res, err := client.Events().GetAddedAndRemovedItemIDs(
				ctx,
				userID,
				containerID,
				deltaTest.inputDelta,
				CallConfig{
					CanMakeDeltaQueries: true,
				})

			require.NoError(t, err, clues.ToCore(err))
			assert.Empty(t, res.Added, "added items")
			assert.Empty(t, res.Removed, "removed items")
			assert.Equal(t, deltaTest.expectNextDeltaURL, res.DU.URL, "next delta URL")
			deltaTest.expectDeltaReset(t, res.DU.Reset, "delta reset")
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

package api_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/exchange/api/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type ContactsAPIUnitSuite struct {
	tester.Suite
}

func TestContactsAPIUnitSuite(t *testing.T) {
	suite.Run(t, &ContactsAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ContactsAPIUnitSuite) TestContactInfo() {
	initial := time.Now()

	tests := []struct {
		name         string
		contactAndRP func() (models.Contactable, *details.ExchangeInfo)
	}{
		{
			name: "Empty Contact",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				contact := models.NewContact()
				contact.SetCreatedDateTime(&initial)
				contact.SetLastModifiedDateTime(&initial)

				i := &details.ExchangeInfo{
					ItemType: details.ExchangeContact,
					Created:  initial,
					Modified: initial,
				}
				return contact, i
			},
		}, {
			name: "Only Name",
			contactAndRP: func() (models.Contactable, *details.ExchangeInfo) {
				aPerson := "Whole Person"
				contact := models.NewContact()
				contact.SetCreatedDateTime(&initial)
				contact.SetLastModifiedDateTime(&initial)
				contact.SetDisplayName(&aPerson)
				i := &details.ExchangeInfo{
					ItemType:    details.ExchangeContact,
					ContactName: aPerson,
					Created:     initial,
					Modified:    initial,
				}
				return contact, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			contact, expected := test.contactAndRP()
			assert.Equal(suite.T(), expected, api.ContactInfo(contact))
		})
	}
}

type ContactAPIE2ESuite struct {
	tester.Suite
	credentials account.M365Config
	ac          api.Client
	user        string
}

// We do end up mocking the actual request, but creating the rest
// similar to E2E suite
func TestContactAPIE2ESuite(t *testing.T) {
	suite.Run(t, &ContactAPIE2ESuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *ContactAPIE2ESuite) SetupSuite() {
	t := suite.T()

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365
	suite.ac, err = mock.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	suite.user = tester.M365UserID(suite.T())
}

func (suite *ContactAPIE2ESuite) TestPaginationErrorConditions() {
	did := "directory-id"

	type errResp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	tests := []struct {
		name      string
		prevDelta bool
		setupf    func()
	}{
		{
			name: "direct error on dleta",
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/" + suite.user + "/contactFolders/" + did + "/contacts/microsoft.graph.delta..$").
					Reply(404)
			},
		},
		{
			name:      "delta reset",
			prevDelta: true,
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/fakedelta").
					Reply(403).
					JSON(map[string]errResp{"error": {Code: "SyncStateNotFound", Message: "..."}})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/" + suite.user + "/contactFolders/" + did + "/contacts/microsoft.graph.delta..$").
					Reply(404)
			},
		},
		{
			name:      "box full",
			prevDelta: true,
			setupf: func() {
				gock.New("https://graph.microsoft.com").
					Get("/fakedelta").
					Reply(403).
					JSON(map[string]errResp{
						"error": {
							Code:    "ErrorQuotaExceeded",
							Message: "The process failed to get the correct properties.",
						},
					})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/" + suite.user + "/contactFolders/" + did + "/contacts$").
					Reply(404)
			},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			defer gock.Off()
			tt.setupf()

			delta := ""
			if tt.prevDelta {
				delta = "https://graph.microsoft.com/fakedelta"
			}

			pgr, err := api.NewContactPager(suite.ac.Stable, suite.user, did, delta, false, false)
			require.NoError(suite.T(), err, "create pager")

			_, _, _, err = api.GetAddedAndRemovedItemIDsFromPager(ctx, delta, &pgr)
			fmt.Println("shared_test.go:118 err:", err)

			// just a unique enough check
			assert.True(suite.T(), err.Error() == "The server returned an unexpected status code with no response body: 404", "get 404")
			assert.True(suite.T(), gock.IsDone(), "all mocks used")
		})
	}
}

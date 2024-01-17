package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	controlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type RestoreIntgSuite struct {
	tester.Suite
}

func TestRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &RestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

// TestRestoreCollections_HandlesEmptyRestoreLocation checks to make sure that
// even if the restore location is empty we fallback to using the collection
// path as the folder, resulting in an in-place restore. It doesn't attempt to
// retore any items because that would bloat the data set in the test user.
func (suite *RestoreIntgSuite) TestRestoreCollections_HandlesEmptyRestoreLocation() {
	t := suite.T()

	acct := tconfig.NewM365Account(t)

	table := []struct {
		service     path.ServiceType
		category    path.CategoryType
		selector    func(*testing.T) selectors.Selector
		defaultPath func(*testing.T) path.Path
		otherPath   func(t *testing.T, location string) path.Path
	}{
		{
			service:  path.ExchangeService,
			category: path.EmailCategory,
			selector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeRestore([]string{tconfig.M365UserID(t)})
				sel.Include(sel.Mails(selectors.Any(), selectors.Any()))

				return sel.Selector
			},
			defaultPath: func(t *testing.T) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.EmailCategory,
					false,
					api.MailInbox)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
			otherPath: func(t *testing.T, location string) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.EmailCategory,
					false,
					location)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
		},
		{
			service:  path.ExchangeService,
			category: path.EventsCategory,
			selector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeRestore([]string{tconfig.M365UserID(t)})
				sel.Include(sel.Events(selectors.Any(), selectors.Any()))

				return sel.Selector
			},
			defaultPath: func(t *testing.T) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.EventsCategory,
					false,
					api.DefaultCalendar)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
			otherPath: func(t *testing.T, location string) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.EventsCategory,
					false,
					location)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
		},
		{
			service:  path.ExchangeService,
			category: path.ContactsCategory,
			selector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeRestore([]string{tconfig.M365UserID(t)})
				sel.Include(sel.Contacts(selectors.Any(), selectors.Any()))

				return sel.Selector
			},
			defaultPath: func(t *testing.T) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.ContactsCategory,
					false,
					api.DefaultContacts)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
			otherPath: func(t *testing.T, location string) path.Path {
				res, err := path.Build(
					tconfig.M365TenantID(t),
					tconfig.M365UserID(t),
					path.ExchangeService,
					path.ContactsCategory,
					false,
					location)
				require.NoError(t, err, clues.ToCore(err))

				return res
			},
		},
	}

	for _, test := range table {
		suite.Run(test.service.HumanString()+test.category.HumanString(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			controller, err := NewController(
				ctx,
				acct,
				test.service,
				control.DefaultOptions(),
				count.New())
			require.NoError(t, err, clues.ToCore(err))

			handler, err := controller.NewServiceHandler(test.service)
			require.NoError(t, err, clues.ToCore(err))

			restoreConfig := controlTD.DefaultRestoreConfig("restore_in_place")
			restoreConfig.OnCollision = control.Copy

			// Create 2 empty collections so we don't bloat the data set.
			path1 := test.defaultPath(t)
			path2 := test.otherPath(t, restoreConfig.Location)
			cols := []data.RestoreCollection{
				data.NoFetchRestoreCollection{
					Collection: exchMock.NewCollection(
						path1,
						path1,
						0),
				},
				data.NoFetchRestoreCollection{
					Collection: exchMock.NewCollection(
						path2,
						path2,
						0),
				},
			}

			restoreConfig.Location = ""

			sel := test.selector(t)

			_, _, err = handler.ConsumeRestoreCollections(
				ctx,
				inject.RestoreConsumerConfig{
					BackupVersion:     version.Backup,
					Options:           control.DefaultOptions(),
					ProtectedResource: sel,
					RestoreConfig:     restoreConfig,
					Selector:          sel,
				},
				cols,
				fault.New(true),
				count.New())
			assert.NoError(t, err, clues.ToCore(err))
		})
	}
}

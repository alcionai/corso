package connector

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
	user      string
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(suite.T(), err)
	suite.connector = loadConnector(ctx, suite.T())
	suite.user = tester.M365UserID(suite.T())
	tester.LogTimeOfTest(suite.T())
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the users associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantUsers() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Users:       make(map[string]string, 0),
		credentials: suite.connector.credentials,
	}

	ctx, flush := tester.NewContext()
	defer flush()

	service, err := newConnector.createService(false)
	require.NoError(suite.T(), err)

	newConnector.graphService = *service

	suite.Equal(len(newConnector.Users), 0)
	err = newConnector.setTenantUsers(ctx)
	assert.NoError(suite.T(), err)
	suite.Greater(len(newConnector.Users), 0)
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the sites associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantSites() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Sites:       make(map[string]string, 0),
		credentials: suite.connector.credentials,
	}

	ctx, flush := tester.NewContext()
	defer flush()

	service, err := newConnector.createService(false)
	require.NoError(suite.T(), err)

	newConnector.graphService = *service

	suite.Equal(0, len(newConnector.Sites))
	err = newConnector.setTenantSites(ctx)
	assert.NoError(suite.T(), err)
	// TODO: should be non-zero once implemented.
	// suite.Greater(len(newConnector.Users), 0)
	suite.Equal(0, len(newConnector.Sites))
}

func (suite *GraphConnectorIntegrationSuite) TestEmptyCollections() {
	dest := tester.DefaultTestRestoreDestination()
	table := []struct {
		name string
		col  []data.Collection
		sel  selectors.Selector
	}{
		{
			name: "ExchangeNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "ExchangeEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "OneDriveNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
		{
			name: "OneDriveEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
		// TODO: SharePoint
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			deets, err := suite.connector.RestoreDataCollections(ctx, test.sel, dest, test.col)
			require.NoError(t, err)
			assert.NotNil(t, deets)

			stats := suite.connector.AwaitStatus()
			assert.Zero(t, stats.ObjectCount)
			assert.Zero(t, stats.FolderCount)
			assert.Zero(t, stats.Successful)
		})
	}
}

//-------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------------

func runRestoreBackupTest(
	t *testing.T,
	test restoreBackupInfo,
	tenant string,
	users []string,
) {
	var (
		collections  []data.Collection
		expectedData = map[string]map[string][]byte{}
		totalItems   = 0
		// Get a dest per test so they're independent.
		dest = tester.DefaultTestRestoreDestination()
	)

	ctx, flush := tester.NewContext()
	defer flush()

	for _, user := range users {
		numItems, userCollections, userExpectedData := collectionsForInfo(
			t,
			test.service,
			tenant,
			user,
			dest,
			test.collections,
		)

		collections = append(collections, userCollections...)
		totalItems += numItems

		for k, v := range userExpectedData {
			expectedData[k] = v
		}
	}

	t.Logf(
		"Restoring collections to %s for user(s) %v\n",
		dest.ContainerName,
		users,
	)

	start := time.Now()

	restoreGC := loadConnector(ctx, t)
	restoreSel := getSelectorWith(test.service)
	deets, err := restoreGC.RestoreDataCollections(ctx, restoreSel, dest, collections)
	require.NoError(t, err)
	assert.NotNil(t, deets)

	status := restoreGC.AwaitStatus()
	runTime := time.Now().Sub(start)

	assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
	assert.Equal(t, totalItems, status.Successful, "status.Successful")
	assert.Len(
		t,
		deets.Entries,
		totalItems,
		"details entries contains same item count as total successful items restored")

	t.Logf("Restore complete in %v\n", runTime)

	// Run a backup and compare its output with what we put in.

	backupGC := loadConnector(ctx, t)
	backupSel := backupSelectorForExpected(t, expectedData)
	t.Logf("Selective backup of %s\n", backupSel)

	start = time.Now()
	dcs, err := backupGC.DataCollections(ctx, backupSel)
	require.NoError(t, err)

	t.Logf("Backup enumeration complete in %v\n", time.Now().Sub(start))

	// Pull the data prior to waiting for the status as otherwise it will
	// deadlock.
	checkCollections(t, totalItems, expectedData, dcs)

	status = backupGC.AwaitStatus()
	assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
	assert.Equal(t, totalItems, status.Successful, "status.Successful")
}

func (suite *GraphConnectorIntegrationSuite) TestRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	table := []restoreBackupInfo{
		{
			name:    "EmailsWithAttachments",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithDirectAttachment(
								subjectText + "-1",
							),
							lookupKey: subjectText + "-1",
						},
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithTwoAttachments(
								subjectText + "-2",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:    "MultipleEmailsMultipleFolders",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
					},
				},
				{
					pathElements: []string{"Work"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
						{
							name: "someencodeditemID3",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-3",
								bodyText+" 3.",
								bodyText+" 3.",
							),
							lookupKey: subjectText + "-3",
						},
					},
				},
			},
		},
		{
			name:    "MultipleContactsSingleFolder",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Contacts"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
			},
		},
		{
			name:    "MultipleContactsMutlipleFolders",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID4",
							data:      mockconnector.GetMockContactBytes("Argon"),
							lookupKey: "Argon",
						},
						{
							name:      "someencodeditemID5",
							data:      mockconnector.GetMockContactBytes("Bernard"),
							lookupKey: "Bernard",
						},
					},
				},
			},
		},
		// {
		// 	name:    "MultipleEventsSingleCalendar",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
		// 					lookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name:    "MultipleEventsMultipleCalendars",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
		// 					lookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			pathElements: []string{"Personal"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID4",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Argon"),
		// 					lookupKey: "Argon",
		// 				},
		// 				{
		// 					name:      "someencodeditemID5",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Bernard"),
		// 					lookupKey: "Bernard",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(t, test, suite.connector.tenant, []string{suite.user})
		})
	}
}

func (suite *GraphConnectorIntegrationSuite) TestMultiFolderBackupDifferentNames() {
	table := []restoreBackupInfo{
		{
			name:    "Contacts",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
					},
				},
			},
		},
		// {
		// 	name:    "Events",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			pathElements: []string{"Personal"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			restoreSel := getSelectorWith(test.service)
			dests := make([]control.RestoreDestination, 0, len(test.collections))
			allItems := 0
			allExpectedData := map[string]map[string][]byte{}

			for i, collection := range test.collections {
				// Get a dest per collection so they're independent.
				dest := tester.DefaultTestRestoreDestination()
				dests = append(dests, dest)

				totalItems, collections, expectedData := collectionsForInfo(
					t,
					test.service,
					suite.connector.tenant,
					suite.user,
					dest,
					[]colInfo{collection},
				)
				allItems += totalItems

				for k, v := range expectedData {
					allExpectedData[k] = v
				}

				t.Logf(
					"Restoring %v/%v collections to %s\n",
					i+1,
					len(test.collections),
					dest.ContainerName,
				)

				restoreGC := loadConnector(ctx, t)
				deets, err := restoreGC.RestoreDataCollections(ctx, restoreSel, dest, collections)
				require.NoError(t, err)
				require.NotNil(t, deets)

				status := restoreGC.AwaitStatus()
				// Always just 1 because it's just 1 collection.
				assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
				assert.Equal(t, totalItems, status.Successful, "status.Successful")
				assert.Equal(
					t, totalItems, len(deets.Entries),
					"details entries contains same item count as total successful items restored")

				t.Log("Restore complete")
			}

			// Run a backup and compare its output with what we put in.

			backupGC := loadConnector(ctx, t)
			backupSel := backupSelectorForExpected(t, allExpectedData)
			t.Log("Selective backup of", backupSel)

			dcs, err := backupGC.DataCollections(ctx, backupSel)
			require.NoError(t, err)

			t.Log("Backup enumeration complete")

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			checkCollections(t, allItems, allExpectedData, dcs)

			status := backupGC.AwaitStatus()
			assert.Equal(t, allItems, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, allItems, status.Successful, "status.Successful")
		})
	}
}

func (suite *GraphConnectorIntegrationSuite) TestMultiuserRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	users := []string{
		suite.user,
		tester.SecondaryM365UserID(suite.T()),
	}
	table := []restoreBackupInfo{
		{
			name:    "Email",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
					},
				},
				{
					pathElements: []string{"Archive"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:    "Contacts",
			service: path.ExchangeService,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
					},
				},
			},
		},
		// {
		// 	name:    "Events",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			pathElements: []string{"Personal"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(t, test, suite.connector.tenant, users)
		})
	}
}

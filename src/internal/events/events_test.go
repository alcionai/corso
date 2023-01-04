package events_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

type EventsIntegrationSuite struct {
	suite.Suite
}

func TestMetricsIntegrationSuite(t *testing.T) {
	tester.RunOnAny(t, tester.CorsoCITests)
	suite.Run(t, new(EventsIntegrationSuite))
}

func (suite *EventsIntegrationSuite) TestNewBus() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()

	s, err := storage.NewStorage(
		storage.ProviderS3,
		storage.S3Config{
			Bucket: "bckt",
			Prefix: "prfx",
		},
	)
	require.NoError(t, err)

	a, err := account.NewAccount(
		account.ProviderM365,
		account.M365Config{
			M365: credentials.M365{
				AzureClientID:     "id",
				AzureClientSecret: "secret",
			},
			AzureTenantID: "tid",
		},
	)
	require.NoError(t, err)

	b, err := events.NewBus(ctx, s, a.ID(), control.Options{})
	require.NotEmpty(t, b)
	require.NoError(t, err)
	require.NoError(t, b.Close())

	b2, err := events.NewBus(ctx, s, a.ID(), control.Options{DisableMetrics: true})
	require.Empty(t, b2)
	require.NoError(t, err)
	require.NoError(t, b2.Close())
}

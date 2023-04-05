package utils

import (
	"context"
	"os"

	"github.com/pkg/errors"

	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
)

func GetGC(ctx context.Context, tenant string) (*connector.GraphConnector, account.M365Config, error) {
	// get account info
	m365Cfg := account.M365Config{
		M365:          credentials.GetM365(),
		AzureTenantID: common.First(tenant, os.Getenv(account.AzureTenantID)),
	}

	acct, err := account.NewAccount(account.ProviderM365, m365Cfg)
	if err != nil {
		return nil, m365Cfg, Only(ctx, errors.Wrap(err, "finding m365 account details"))
	}

	// TODO: log/print recoverable errors
	errs := fault.New(false)

	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Users, errs)
	if err != nil {
		return nil, m365Cfg, Only(ctx, errors.Wrap(err, "connecting to graph API"))
	}

	return gc, m365Cfg, nil
}

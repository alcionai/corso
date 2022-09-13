package m365

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/pkg/account"
)

// Users returns a list of users in the specified M365 tenant
// TODO: Implement paging support
func Users(m365Account account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(m365Account)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.GetUsers(), nil
}

package account

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common"
)

type accountProvider int

//go:generate stringer -type=accountProvider -linecomment
const (
	ProviderUnknown accountProvider = 0 // Unknown Provider
	ProviderM365    accountProvider = 1 // M365
)

// storage parsing errors
var (
	errMissingRequired = clues.New("missing required storage configuration")
)

// Account defines an account provider, along with any credentials
// and identifiers required to set up or communicate with that provider.
type Account struct {
	Provider accountProvider
	Config   map[string]string
}

type providerIDer interface {
	common.StringConfigurer

	providerID(accountProvider) string
}

// NewAccount aggregates all the supplied configurations into a single configuration
func NewAccount(p accountProvider, cfgs ...providerIDer) (Account, error) {
	var (
		pid string
		scs = make([]common.StringConfigurer, len(cfgs))
	)

	for i, c := range cfgs {
		scs[i] = c.(common.StringConfigurer)

		if len(c.providerID(p)) > 0 {
			pid = c.providerID(p)
		}
	}

	cs, err := common.UnionStringConfigs(scs...)

	a := Account{
		Provider: p,
		Config:   cs,
	}

	a = setProviderID(a, p, pid)

	return a, err
}

func setProviderID(a Account, p accountProvider, id string) Account {
	if len(a.Config) == 0 {
		a.Config = map[string]string{}
	}

	a.Config[p.String()+"-tenant-id"] = id

	return a
}

// ID returns the primary tenant ID held by its configuration.
// Ex: if the account uses an M365 provider, the M365 tenant ID
// is returned.  If the account contains no ID info, returns an
// empty string.
func (a Account) ID() string {
	if len(a.Config) == 0 {
		return ""
	}

	return a.Config[a.Provider.String()+"-tenant-id"]
}

package account

import (
	"errors"

	"github.com/alcionai/corso/src/internal/common"
)

type accountProvider int

//go:generate stringer -type=accountProvider -linecomment
const (
	ProviderUnknown accountProvider = iota // Unknown Provider
	ProviderM365                           // M365
)

// storage parsing errors
var (
	errMissingRequired = errors.New("missing required storage configuration")
)

// Account defines an account provider, along with any credentials
// and identifiers required to set up or communicate with that provider.
type Account struct {
	Provider accountProvider
	Config   map[string]string
}

// NewAccount aggregates all the supplied configurations into a single configuration
func NewAccount(p accountProvider, cfgs ...common.StringConfigurer) (Account, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Account{
		Provider: p,
		Config:   cs,
	}, err
}

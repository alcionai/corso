package account

import "errors"

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

type (
	config     map[string]string
	configurer interface {
		Config() (config, error)
	}
)

// Account defines an account provider, along with any credentials
// and identifiers requried to set up or communicate with that provider.
type Account struct {
	Provider accountProvider
	Config   config
}

// NewAccount aggregates all the supplied configurations into a single configuration
func NewAccount(p accountProvider, cfgs ...configurer) (Account, error) {
	cs, err := unionConfigs(cfgs...)
	return Account{
		Provider: p,
		Config:   cs,
	}, err
}

func unionConfigs(cfgs ...configurer) (config, error) {
	union := config{}
	for _, cfg := range cfgs {
		c, err := cfg.Config()
		if err != nil {
			return nil, err
		}
		for k, v := range c {
			union[k] = v
		}
	}
	return union, nil
}

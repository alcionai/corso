package storage

type storageProvider int

//go:generate stringer -type=storageProvider -linecomment
const (
	ProviderUnknown storageProvider = iota // Unknown Provider
	ProviderS3                             // S3
)

type (
	config     map[string]any
	configurer interface {
		Config() config
	}
)

// Storage defines a storage provider, along with any configuration
// requried to set up or communicate with that provider.
type Storage struct {
	Provider storageProvider
	Config   config
}

// NewStorage aggregates all the supplied configurations into a single configuration.
func NewStorage(p storageProvider, cfgs ...configurer) Storage {
	return Storage{
		Provider: p,
		Config:   unionConfigs(cfgs...),
	}
}

func unionConfigs(cfgs ...configurer) config {
	c := config{}
	for _, cfg := range cfgs {
		for k, v := range cfg.Config() {
			c[k] = v
		}
	}
	return c
}

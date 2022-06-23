package storage

import (
	"errors"
	"fmt"

	"github.com/alcionai/corso/internal/common"
)

type storageProvider int

//go:generate stringer -type=storageProvider -linecomment
const (
	ProviderUnknown storageProvider = iota // Unknown Provider
	ProviderS3                             // S3
)

// storage parsing errors
var (
	errMissingRequired = errors.New("missing required storage configuration")
)

// Storage defines a storage provider, along with any configuration
// requried to set up or communicate with that provider.
type Storage struct {
	Provider storageProvider
	Config   map[string]string
}

// NewStorage aggregates all the supplied configurations into a single configuration.
func NewStorage(p storageProvider, cfgs ...common.StringConfigurer) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)
	return Storage{
		Provider: p,
		Config:   cs,
	}, err
}

// Helper for parsing the values in a config object.
// If the value is nil or not a string, returns an empty string.
func orEmptyString(v any) string {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("panic recovery casting %v to string\n", v)
		}
	}()
	if v == nil {
		return ""
	}
	return v.(string)
}

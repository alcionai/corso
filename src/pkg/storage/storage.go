package storage

import (
	"fmt"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common"
)

type storageProvider int

//go:generate stringer -type=storageProvider -linecomment
const (
	ProviderUnknown storageProvider = 0 // Unknown Provider
	ProviderS3      storageProvider = 1 // S3
	ProviderFS      storageProvider = 2 // local or network attached file system
)

// storage parsing errors
var (
	errMissingRequired = clues.New("missing required storage configuration")
)

// Storage defines a storage provider, along with any configuration
// required to set up or communicate with that provider.
type Storage struct {
	Provider storageProvider
	Config   map[string]string
	// TODO: These are AWS S3 specific -> move these out
	SessionTags     map[string]string
	Role            string
	SessionName     string
	SessionDuration string
}

// NewStorage aggregates all the supplied configurations into a single configuration.
func NewStorage(p storageProvider, cfgs ...common.StringConfigurer) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Storage{
		Provider: p,
		Config:   cs,
	}, err
}

// NewStorageUsingRole supports specifying an AWS IAM role the storage provider
// should assume.
func NewStorageUsingRole(
	p storageProvider,
	roleARN string,
	sessionName string,
	sessionTags map[string]string,
	duration string,
	cfgs ...common.StringConfigurer,
) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Storage{
		Provider:        p,
		Config:          cs,
		Role:            roleARN,
		SessionTags:     sessionTags,
		SessionName:     sessionName,
		SessionDuration: duration,
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

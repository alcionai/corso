package storage

import (
	"github.com/alcionai/corso/src/internal/common"
)

type StorageProvider int

//go:generate stringer -type=StorageProvider -linecomment
const (
	ProviderUnknown StorageProvider = 0 // Unknown Provider
	ProviderS3      StorageProvider = 1 // S3
)

// Storage defines a storage provider, along with any configuration
// required to set up or communicate with that provider.
type Storage struct {
	Provider StorageProvider
	Config   map[string]string
	// TODO: These are AWS S3 specific -> move these out
	SessionTags     map[string]string
	Role            string
	SessionName     string
	SessionDuration string
}

// NewStorage aggregates all the supplied configurations into a single configuration.
func NewStorage(p StorageProvider, cfgs ...common.StringConfigurer) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Storage{
		Provider: p,
		Config:   cs,
	}, err
}

// NewStorageUsingRole supports specifying an AWS IAM role the storage provider
// should assume.
func NewStorageUsingRole(
	p StorageProvider,
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

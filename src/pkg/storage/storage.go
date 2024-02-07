package storage

import (
	"fmt"
	"slices"

	"github.com/alcionai/clues"
	"github.com/spf13/cast"

	"github.com/alcionai/canario/src/internal/common"
	"github.com/alcionai/canario/src/pkg/path"
)

var ErrVerifyingConfigStorage = clues.New("verifying configs in corso config file")

type ProviderType int

//go:generate stringer -type=ProviderType -linecomment
const (
	ProviderUnknown    ProviderType = 0 // Unknown Provider
	ProviderS3         ProviderType = 1 // S3
	ProviderFilesystem ProviderType = 2 // Filesystem
)

var StringToProviderType = map[string]ProviderType{
	ProviderUnknown.String():    ProviderUnknown,
	ProviderS3.String():         ProviderS3,
	ProviderFilesystem.String(): ProviderFilesystem,
}

const (
	StorageProviderTypeKey = "provider"
)

// storage parsing errors
var (
	errMissingRequired = clues.New("missing required storage configuration")
	errInvalidProvider = clues.New("unsupported storage provider")
)

// Storage defines a storage provider, along with any configuration
// required to set up or communicate with that provider.
type Storage struct {
	Provider ProviderType
	Config   map[string]string
	// TODO: These are AWS S3 specific -> move these out
	SessionTags     map[string]string
	Role            string
	SessionName     string
	SessionDuration string
	// QueryParams allows passing custom query parameters to S3 GET object
	// requests. Parameters prefixed by `x-` will be ignored by S3 (no functional
	// changes) but added to server access logs. These KV-pairs are session
	// specific and won't be persisted to or sourced from the corso config file.
	QueryParams map[string]string
}

// NewStorage aggregates all the supplied configurations into a single configuration.
func NewStorage(p ProviderType, cfgs ...common.StringConfigurer) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Storage{
		Provider: p,
		Config:   cs,
	}, err
}

// NewStorageUsingRole supports specifying an AWS IAM role the storage provider
// should assume.
func NewStorageUsingRole(
	p ProviderType,
	roleARN string,
	sessionName string,
	sessionTags map[string]string,
	queryParams map[string]string,
	duration string,
	cfgs ...common.StringConfigurer,
) (Storage, error) {
	cs, err := common.UnionStringConfigs(cfgs...)

	return Storage{
		Provider:        p,
		Config:          cs,
		Role:            roleARN,
		SessionTags:     sessionTags,
		QueryParams:     queryParams,
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

func (s Storage) StorageConfig() (Configurer, error) {
	switch s.Provider {
	case ProviderS3:
		return buildS3ConfigFromMap(s.Config)
	case ProviderFilesystem:
		return buildFilesystemConfigFromMap(s.Config)
	}

	return nil, errInvalidProvider.With("provider", s.Provider)
}

func (s Storage) GetStorageConfigHash() (string, error) {
	switch s.Provider {
	case ProviderS3:
		s3Cnf, err := s.ToS3Config()
		if err != nil {
			return "", err
		}

		return s3Cnf.configHash()

	case ProviderFilesystem:
		fsCnf, err := s.ToFilesystemConfig()
		if err != nil {
			return "", err
		}

		return fsCnf.configHash()
	}

	return "", errInvalidProvider.With("provider", s.Provider)
}

func NewStorageConfig(provider ProviderType) (Configurer, error) {
	switch provider {
	case ProviderS3:
		return &S3Config{}, nil
	case ProviderFilesystem:
		return &FilesystemConfig{}, nil
	}

	return nil, errInvalidProvider.With("provider", provider)
}

type Getter interface {
	Get(key string) any
}

type Setter interface {
	Set(key string, value any)
}

// WriteConfigToStorer writes config key value pairs to provided store.
type WriteConfigToStorer interface {
	WriteConfigToStore(
		s Setter,
	)
}

type Configurer interface {
	common.StringConfigurer

	// ApplyOverrides fetches config from file, processes overrides
	// from sources like environment variables and flags, and updates the
	// underlying configuration accordingly.
	ApplyConfigOverrides(
		g Getter,
		readConfigFromStore bool,
		matchFromConfig bool,
		overrides map[string]string,
	) error

	WriteConfigToStorer

	configHash() (string, error)
}

// mustMatchConfig compares the values of each key to their config file value in store.
// If any value differs from the store value, an error is returned.
// values in m that aren't stored in the config are ignored.
func mustMatchConfig(
	g Getter,
	tomlMap map[string]string,
	m map[string]string,
	pathKeys []string,
) error {
	for k, v := range m {
		if len(v) == 0 {
			continue // empty variables will get caught by configuration validators, if necessary
		}

		tomlK, ok := tomlMap[k]
		if !ok {
			continue // m may declare values which aren't stored in the config file
		}

		vv := cast.ToString(g.Get(tomlK))

		areEqual := false

		// some of the values maybe paths, hence they require more than just string equality check
		if len(pathKeys) > 0 && slices.Contains(pathKeys, k) {
			areEqual = path.ArePathsEquivalent(v, vv)
		} else {
			areEqual = v == vv
		}

		if !areEqual {
			err := clues.New("value of " + k + " (" + v + ") does not match corso configuration value (" + vv + ")")
			return clues.Stack(ErrVerifyingConfigStorage, err)
		}
	}

	return nil
}

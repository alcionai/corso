package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alcionai/clues"
	"github.com/spf13/cast"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/pkg/logger"
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

func (s Storage) StorageConfig() (Configurer, error) {
	switch s.Provider {
	case ProviderS3:
		return buildS3ConfigFromMap(s.Config)
	case ProviderFilesystem:
		return buildFilesystemConfigFromMap(s.Config)
	}

	return nil, clues.New("unsupported storage provider: [" + s.Provider.String() + "]")
}

func NewStorageConfig(provider ProviderType) (Configurer, error) {
	switch provider {
	case ProviderS3:
		return &S3Config{}, nil
	case ProviderFilesystem:
		return &FilesystemConfig{}, nil
	}

	return nil, clues.New("unsupported storage provider: [" + provider.String() + "]")
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
}

// mustMatchConfig compares the values of each key to their config file value in store.
// If any value differs from the store value, an error is returned.
// values in m that aren't stored in the config are ignored.
func mustMatchConfig(
	g Getter,
	tomlMap map[string]string,
	m map[string]string,
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

		if IsValidPath(v) && IsValidPath(vv) {
			areEqual, _ = ArePathsEquivalent(v, vv)
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

func ArePathsEquivalent(path1, path2 string) (bool, error) {
	normalizedPath1 := strings.TrimSpace(filepath.Clean(path1))
	normalizedPath2 := strings.TrimSpace(filepath.Clean(path2))

	normalizedPath1 = strings.TrimSuffix(normalizedPath1, string(filepath.Separator))
	normalizedPath2 = strings.TrimSuffix(normalizedPath2, string(filepath.Separator))

	return normalizedPath1 == normalizedPath2, nil
}

func IsValidPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	isDir := strings.TrimSpace(filepath.Ext(path)) == ""

	if isDir {
		err := os.Mkdir(path, 0644)
		if err == nil {
			os.Remove(path)
			return true
		}

		if os.IsPermission(err) {
			logger.Ctx(context.Background()).Info("directory could have been created but got permission error")
			return true
		}
	} else {
		dir := filepath.Dir(path)
		err := os.Mkdir(dir, 0644)
		if err != nil {
			return false
		}
		defer os.Remove(dir)

		_, err = os.Create(path)
		if err == nil {
			os.Remove(path)
			return true
		}

		if os.IsPermission(err) {
			logger.Ctx(context.Background()).Info("file could have been created but got permission error")
			return true
		}
	}
	return true
}

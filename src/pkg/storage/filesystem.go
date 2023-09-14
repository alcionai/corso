package storage

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cast"

	"github.com/alcionai/corso/src/internal/common/str"
)

const (
	FilesystemPath = "path"
)

var fsConstToTomlKeyMap = map[string]string{
	StorageProviderTypeKey: StorageProviderTypeKey,
	FilesystemPath:         FilesystemPath,
}

type FilesystemConfig struct {
	Path string
}

func buildFilesystemConfigFromMap(config map[string]string) (*FilesystemConfig, error) {
	c := &FilesystemConfig{}

	if len(config) > 0 {
		c.Path = orEmptyString(config[FilesystemPath])
	}

	return c, c.validate()
}

func (c *FilesystemConfig) validate() error {
	check := map[string]string{
		FilesystemPath: c.Path,
	}

	for k, v := range check {
		if len(v) == 0 {
			return clues.Stack(errMissingRequired, clues.New(k))
		}
	}

	return nil
}

func (c *FilesystemConfig) fsConfigsFromStore(kvg KVStoreGetter) {
	c.Path = cast.ToString(kvg.Get(FilesystemPath))
}

// TODO(pandeyabs): Remove this. It's not adding any value.
func fsOverrides(in map[string]string) map[string]string {
	return map[string]string{
		FilesystemPath: in[FilesystemPath],
	}
}

var _ Configurer = &FilesystemConfig{}

func (c *FilesystemConfig) ApplyConfigOverrides(
	kvg KVStoreGetter,
	readConfigFromStore bool,
	matchFromConfig bool,
	overrides map[string]string,
) error {
	if readConfigFromStore {
		c.fsConfigsFromStore(kvg)

		if matchFromConfig {
			providerType := cast.ToString(kvg.Get(StorageProviderTypeKey))
			if providerType != ProviderFilesystem.String() {
				return clues.New("unsupported storage provider in config file: " + providerType)
			}

			// This is matching override values from config file.
			if err := mustMatchConfig(kvg, fsConstToTomlKeyMap, fsOverrides(overrides)); err != nil {
				return clues.Wrap(err, "verifying storage configs in corso config file")
			}
		}
	}

	c.Path = str.First(overrides[FilesystemPath], c.Path)

	return c.validate()
}

// TODO(pandeyabs): We need to sanitize paths e.g. handle relative paths,
// make paths cross platform compatible, etc.
func (c FilesystemConfig) StringConfig() (map[string]string, error) {
	cfg := map[string]string{
		FilesystemPath: c.Path,
	}

	return cfg, c.validate()
}

var _ WriteConfigToStorer = FilesystemConfig{}

func (c FilesystemConfig) WriteConfigToStore(
	kvs KVStoreSetter,
) {
	kvs.Set(StorageProviderTypeKey, ProviderFilesystem.String())
	kvs.Set(FilesystemPath, c.Path)
}

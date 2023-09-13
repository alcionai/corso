package storage

import (
	"github.com/alcionai/clues"
	"github.com/spf13/cast"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/str"
)

const (
	// Make this better
	FilesystemPath = "path"
)

type FilesystemConfig struct {
	Path string
}

func (s Storage) FsConfig() (FilesystemConfig, error) {
	c := FilesystemConfig{}

	if len(s.Config) > 0 {
		c.Path = orEmptyString(s.Config[FilesystemPath])
	}

	return c, c.validate()
}

func (c FilesystemConfig) validate() error {
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

// S3Config retrieves the S3Config details from the Storage config.
// TODO(pandeyabs): Unexpose
func MakeFSConfigFromMap(config map[string]string) (FilesystemConfig, error) {
	c := FilesystemConfig{}

	if len(config) > 0 {
		c.Path = orEmptyString(config[FilesystemPath])
	}

	return c, c.validate()
}

// No need to return error here. Viper returns empty values.
func fsConfigsFromStore(kvg KVStoreGetter) FilesystemConfig {
	var fsConfig FilesystemConfig

	fsConfig.Path = cast.ToString(kvg.Get(FilesystemPath))

	return fsConfig
}

func fsOverrides(in map[string]string) map[string]string {
	return map[string]string{
		FilesystemPath: in[FilesystemPath],
	}
}

var _ Configurer = FilesystemConfig{}

func (c FilesystemConfig) FetchConfigFromStore(
	kvg KVStoreGetter,
	readConfigFromStore bool,
	matchFromConfig bool,
	overrides map[string]string,
) (Configurer, error) {
	var fsCfg FilesystemConfig

	if readConfigFromStore {
		fsCfg = fsConfigsFromStore(kvg)

		if b, ok := overrides[Bucket]; ok {
			overrides[Bucket] = common.NormalizeBucket(b)
		}

		if matchFromConfig {
			providerType := cast.ToString(kvg.Get(StorageProviderTypeKey))
			if providerType != ProviderFilesystem.String() {
				return FilesystemConfig{}, clues.New("unsupported storage provider: " + providerType)
			}

			// This is matching override values from config file.
			if err := mustMatchConfig(kvg, fsOverrides(overrides)); err != nil {
				return S3Config{}, clues.Wrap(err, "verifying s3 configs in corso config file")
			}
		}
	}

	fsCfg = FilesystemConfig{
		Path: str.First(overrides[FilesystemPath], fsCfg.Path),
	}

	return fsCfg, fsCfg.validate()
}

// TODO(pandeyabs): Do we need to sanitize path?
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

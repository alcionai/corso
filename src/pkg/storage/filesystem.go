package storage

import (
	"encoding/json"

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

func (s Storage) ToFilesystemConfig() (*FilesystemConfig, error) {
	return buildFilesystemConfigFromMap(s.Config)
}

func (s Storage) GenerateFilesystemHash() (string, error) {
	fsCfg, err := buildFilesystemConfigFromMap(s.Config)
	if err != nil {
		return "", err
	}

	fsCfgBytes, err := json.Marshal(fsCfg)
	if err != nil {
		return "", clues.New("failed to serialize filesystem config")
	}

	fsCfgHash := GenerateHash(fsCfgBytes, hashLength)
	if len(fsCfgHash) != hashLength {
		return "", clues.New("failed to generate hash of configured length")
	}

	return fsCfgHash, nil
}

func buildFilesystemConfigFromMap(config map[string]string) (*FilesystemConfig, error) {
	c := &FilesystemConfig{}

	if len(config) > 0 {
		c.Path = orEmptyString(config[FilesystemPath])
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

func (c *FilesystemConfig) fsConfigsFromStore(g Getter) {
	c.Path = cast.ToString(g.Get(FilesystemPath))
}

// TODO(pandeyabs): Remove this. It's not adding any value.
func fsOverrides(in map[string]string) map[string]string {
	return map[string]string{
		FilesystemPath: in[FilesystemPath],
	}
}

var _ Configurer = &FilesystemConfig{}

func (c *FilesystemConfig) ApplyConfigOverrides(
	g Getter,
	readConfigFromStore bool,
	matchFromConfig bool,
	overrides map[string]string,
) error {
	if readConfigFromStore {
		c.fsConfigsFromStore(g)

		if matchFromConfig {
			providerType := cast.ToString(g.Get(StorageProviderTypeKey))
			if providerType != ProviderFilesystem.String() {
				return clues.New("unsupported storage provider in config file: [" + providerType + "]")
			}

			// This is matching override values from config file.
			if err := mustMatchConfig(g, fsConstToTomlKeyMap, fsOverrides(overrides)); err != nil {
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
	s Setter,
) {
	s.Set(StorageProviderTypeKey, ProviderFilesystem.String())
	s.Set(FilesystemPath, c.Path)
}

package storage

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type testConfig struct {
	expect string
	err    error
}

func (c testConfig) StringConfig() (map[string]string, error) {
	return map[string]string{"expect": c.expect}, c.err
}

type StorageUnitSuite struct {
	tester.Suite
}

func TestStorageUnitSuite(t *testing.T) {
	suite.Run(t, &StorageUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *StorageUnitSuite) TestNewStorage() {
	table := []struct {
		name     string
		p        ProviderType
		c        testConfig
		errCheck assert.ErrorAssertionFunc
	}{
		{"unknown no error", ProviderUnknown, testConfig{"configVal", nil}, assert.NoError},
		{"s3 no error", ProviderS3, testConfig{"configVal", nil}, assert.NoError},
		{"unknown w/ error", ProviderUnknown, testConfig{"configVal", assert.AnError}, assert.Error},
		{"s3 w/ error", ProviderS3, testConfig{"configVal", assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			s, err := NewStorage(test.p, test.c)
			test.errCheck(t, err, clues.ToCore(err))

			// remaining tests are dependent upon error-free state
			if test.c.err != nil {
				return
			}

			assert.Equalf(t,
				test.p,
				s.Provider,
				"expected storage provider [%s], got [%s]", test.p, s.Provider)
			assert.Equalf(t,
				test.c.expect,
				s.Config["expect"],
				"expected storage config [%s], got [%s]", test.c.expect, s.Config["expect"])
		})
	}
}

func (suite *StorageUnitSuite) TestGetAccountConfigHash() {
	tests := []struct {
		name     string
		provider ProviderType
		config   any
	}{
		{
			name:     "s3 storage",
			provider: ProviderS3,
			config:   getTestS3Config("test-bucket", "https://aws.s3", "test-prefix"),
		},
		{
			name:     "filesystem storage",
			provider: ProviderFilesystem,
			config:   getTestFileSystemConfig("test/to/dir"),
		},
		{
			name:     "invalid account",
			provider: ProviderUnknown,
			config:   testConfig{"configVal", nil},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			if test.provider == ProviderUnknown {
				s, err := NewStorage(test.provider, test.config.(testConfig))
				require.NoError(t, err)

				_, err = s.GetStorageConfigHash()
				require.Error(t, err)
			}

			if test.provider == ProviderS3 {
				_, ok := test.config.(Configurer)
				require.True(t, ok)

				s3Cnf := test.config.(*S3Config)
				s, err := NewStorage(test.provider, s3Cnf)
				require.NoError(t, err)

				hash, err := s.GetStorageConfigHash()
				require.NoError(t, err)
				assert.True(t, len(hash) > 0)
			}

			if test.provider == ProviderFilesystem {
				_, ok := test.config.(Configurer)
				require.True(t, ok)

				fsCnf := test.config.(*FilesystemConfig)
				s, err := NewStorage(test.provider, fsCnf)
				require.NoError(t, err)

				hash, err := s.GetStorageConfigHash()
				require.NoError(t, err)
				assert.True(t, len(hash) > 0)
			}
		})
	}
}

func getTestS3Config(bucket, endpoint, prefix string) *S3Config {
	return &S3Config{
		Bucket:   bucket,
		Endpoint: endpoint,
		Prefix:   prefix,
	}
}

func getTestFileSystemConfig(path string) *FilesystemConfig {
	return &FilesystemConfig{
		Path: path,
	}
}

type testGetter struct {
	storeMap map[string]string
}

func (tg testGetter) Get(key string) any {
	val, ok := tg.storeMap[key]
	if ok {
		return val
	}

	return ""
}

func (suite *StorageUnitSuite) TestMustMatchConfig() {
	t := suite.T()

	table := []struct {
		name        string
		tomlMap     map[string]string
		overrideMap map[string]string
		getterMap   map[string]string
		pathKeys    []string
		errorCheck  assert.ErrorAssertionFunc
	}{
		{
			name:    "s3 config match",
			tomlMap: s3constToTomlKeyMap,
			overrideMap: map[string]string{
				Bucket:           "test-bucket",
				Endpoint:         "https://aws.s3",
				Prefix:           "test-prefix",
				"additional-key": "additional-value",
			},
			getterMap: map[string]string{
				"bucket":   "test-bucket",
				"endpoint": "https://aws.s3",
				"prefix":   "test-prefix",
			},
			errorCheck: assert.NoError,
		},
		{
			name:    "s3 config match - bucket mismatch",
			tomlMap: s3constToTomlKeyMap,
			overrideMap: map[string]string{
				Bucket:           "test-bucket",
				Endpoint:         "https://aws.s3",
				Prefix:           "test-prefix",
				"additional-key": "additional-value",
			},
			getterMap: map[string]string{
				"bucket":   "test-bucket-new",
				"endpoint": "https://aws.s3",
				"prefix":   "test-prefix",
			},
			errorCheck: assert.Error,
		},
		{
			name:    "s3 config match - endpoint mismatch",
			tomlMap: s3constToTomlKeyMap,
			overrideMap: map[string]string{
				Bucket:           "test-bucket",
				Endpoint:         "https://aws.s3",
				Prefix:           "test-prefix",
				"additional-key": "additional-value",
			},
			getterMap: map[string]string{
				"bucket":   "test-bucket",
				"endpoint": "https://aws.s3/new",
				"prefix":   "test-prefix",
			},
			errorCheck: assert.Error,
		},
		{
			name:    "s3 config match - prefix mismatch",
			tomlMap: s3constToTomlKeyMap,
			overrideMap: map[string]string{
				Bucket:           "test-bucket",
				Endpoint:         "https://aws.s3",
				Prefix:           "test-prefix",
				"additional-key": "additional-value",
			},
			getterMap: map[string]string{
				"bucket":   "test-bucket",
				"endpoint": "https://aws.s3",
				"prefix":   "test-prefix-new",
			},
			errorCheck: assert.Error,
		},
		{
			name:    "filesystem config match - success case",
			tomlMap: fsConstToTomlKeyMap,
			overrideMap: map[string]string{
				StorageProviderTypeKey: "filesystem",
				FilesystemPath:         "/path/to/dir",
				"additional-key":       "additional-value",
			},
			getterMap: map[string]string{
				"provider": "filesystem",
				"path":     "/path/to/dir",
			},
			pathKeys:   []string{FilesystemPath},
			errorCheck: assert.NoError,
		},
		{
			name:    "filesystem config match - path mismatch",
			tomlMap: fsConstToTomlKeyMap,
			overrideMap: map[string]string{
				StorageProviderTypeKey: "filesystem",
				FilesystemPath:         "/path/to/dir",
				"additional-key":       "additional-value",
			},
			getterMap: map[string]string{
				"provider": "filesystem",
				"path":     "/path/to/dir/new",
			},
			pathKeys:   []string{FilesystemPath},
			errorCheck: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			tg := testGetter{test.getterMap}
			err := mustMatchConfig(tg, test.tomlMap, test.overrideMap, test.pathKeys)
			test.errorCheck(t, err)
		})
	}
}

package storage_test

import (
	"encoding/json"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/storage"
)

type CommonCfgUnitSuite struct {
	tester.Suite
}

func TestCommonCfgUnitSuite(t *testing.T) {
	suite.Run(t, &CommonCfgUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var goodCommonConfig = storage.CommonConfig{
	Corso: credentials.Corso{
		CorsoPassphrase: "passph",
	},
}

func (suite *CommonCfgUnitSuite) TestCommonConfig_Config() {
	cfg := goodCommonConfig
	c, err := cfg.StringConfig()
	assert.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		key    string
		expect string
	}{
		{"common_corsoPassphrase", cfg.CorsoPassphrase},
	}
	for _, test := range table {
		suite.Run(test.key, func() {
			assert.Equal(suite.T(), test.expect, c[test.key])
		})
	}
}

func (suite *CommonCfgUnitSuite) TestStorage_CommonConfig() {
	t := suite.T()

	in := goodCommonConfig
	s, err := storage.NewStorage(storage.ProviderUnknown, in)
	assert.NoError(t, err, clues.ToCore(err))
	out, err := s.CommonConfig()
	assert.NoError(t, err, clues.ToCore(err))

	assert.Equal(t, in.CorsoPassphrase, out.CorsoPassphrase)
}

func (suite *CommonCfgUnitSuite) TestStorage_CommonConfig_InvalidCases() {
	// missing required properties
	table := []struct {
		name string
		cfg  storage.CommonConfig
	}{
		{"missing passphrase", storage.CommonConfig{}},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			_, err := storage.NewStorage(storage.ProviderUnknown, test.cfg)
			assert.Error(suite.T(), err)
		})
	}

	// required property not populated in storage
	table2 := []struct {
		name  string
		amend func(storage.Storage)
	}{
		{
			"missing passphrase",
			func(s storage.Storage) {
				s.Config["common_corsoPassphrase"] = ""
			},
		},
	}
	for _, test := range table2 {
		suite.Run(test.name, func() {
			t := suite.T()

			st, err := storage.NewStorage(storage.ProviderUnknown, goodCommonConfig)
			assert.NoError(t, err, clues.ToCore(err))
			test.amend(st)
			_, err = st.CommonConfig()
			assert.Error(t, err)
		})
	}
}

// Test GenerateHash
func (suite *CommonCfgUnitSuite) TestGenerateHash() {
	type testStruct struct {
		Text   string
		Number int
		Status bool
	}

	table := []struct {
		name       string
		input1     interface{}
		input2     interface{}
		sameCheck  bool
		hashLength int
	}{
		{
			name:       "check if same hash is generated for same string input",
			input1:     "test data",
			hashLength: 7,
			sameCheck:  true,
		},
		{
			name:       "check if same hash is generated for same struct input",
			input1:     testStruct{Text: "test text", Number: 1, Status: true},
			hashLength: 7,
			sameCheck:  true,
		},
		{
			name:       "check if different hash is generated for different string input",
			input1:     "test data",
			input2:     "test data 2",
			hashLength: 7,
			sameCheck:  false,
		},
		{
			name:       "check if different hash is generated for different struct input",
			input1:     testStruct{Text: "test text", Number: 1, Status: true},
			input2:     testStruct{Text: "test text 2", Number: 2, Status: false},
			hashLength: 7,
			sameCheck:  false,
		},
		{
			name:       "check if length of hash is 32 if hash length is not provided for string input",
			input1:     "test data",
			hashLength: 0,
			sameCheck:  false,
		},
		{
			name:       "check if length of hash is 32 if hash length is not provided for struct input",
			input1:     testStruct{Text: "test text", Number: 1, Status: true},
			hashLength: 0,
			sameCheck:  false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			var input1Bytes []byte
			var err error
			var hash1 string

			input1Bytes, err = json.Marshal(test.input1)
			require.NoError(t, err)

			hash1 = storage.GenerateHash(input1Bytes, test.hashLength)

			if test.hashLength == 0 {
				assert.Equal(t, 32, len(hash1))
			}

			if test.hashLength > 0 && test.sameCheck {
				hash2 := storage.GenerateHash(input1Bytes, test.hashLength)

				assert.Equal(t, test.hashLength, len(hash1))
				assert.Equal(t, test.hashLength, len(hash2))
				assert.Equal(t, hash1, hash2)
			}

			if test.hashLength > 0 && !test.sameCheck {
				input2Bytes, err := json.Marshal(test.input2)
				require.NoError(t, err)

				hash2 := storage.GenerateHash(input2Bytes, test.hashLength)

				assert.Equal(t, test.hashLength, len(hash1))
				assert.Equal(t, test.hashLength, len(hash2))
				assert.NotEqual(t, hash1, hash2)
			}

		})
	}
}

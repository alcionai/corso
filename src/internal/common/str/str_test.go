package str

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Warning: importing the corso tester.suite causes a circular import
// ---------------------------------------------------------------------------

func TestPreview(t *testing.T) {
	table := []struct {
		input  string
		size   int
		expect string
	}{
		{
			input:  "",
			size:   1,
			expect: "",
		},
		{
			input:  "yes",
			size:   1,
			expect: "yes",
		},
		{
			input:  "yes!",
			size:   5,
			expect: "yes!",
		},
		{
			input:  "however",
			size:   6,
			expect: "how...",
		},
		{
			input:  "negative",
			size:   -1,
			expect: "n...",
		},
	}
	for _, test := range table {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(
				t,
				test.expect,
				Preview(test.input, test.size))
		})
	}
}

// Test GenerateHash
func TestGenerateHash(t *testing.T) {
	type testStruct struct {
		Text   string
		Number int
		Status bool
	}

	table := []struct {
		name       string
		input1     any
		input2     any
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
		var input1Bytes []byte

		var err error

		var hash1 string

		input1Bytes, err = json.Marshal(test.input1)
		require.NoError(t, err)

		hash1 = GenerateHash(input1Bytes, test.hashLength)

		if test.hashLength == 0 {
			assert.Equal(t, 32, len(hash1))
		}

		if test.hashLength > 0 && test.sameCheck {
			hash2 := GenerateHash(input1Bytes, test.hashLength)

			assert.Equal(t, test.hashLength, len(hash1))
			assert.Equal(t, test.hashLength, len(hash2))
			assert.Equal(t, hash1, hash2)
		}

		if test.hashLength > 0 && !test.sameCheck {
			input2Bytes, err := json.Marshal(test.input2)
			require.NoError(t, err)

			hash2 := GenerateHash(input2Bytes, test.hashLength)

			assert.Equal(t, test.hashLength, len(hash1))
			assert.Equal(t, test.hashLength, len(hash2))
			assert.NotEqual(t, hash1, hash2)
		}
	}
}

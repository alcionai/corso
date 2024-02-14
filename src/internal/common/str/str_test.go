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
		name      string
		input1    any
		input2    any
		sameCheck bool
	}{
		{
			name:      "check if same hash is generated for same string input",
			input1:    "test data",
			sameCheck: true,
		},
		{
			name:      "check if same hash is generated for same struct input",
			input1:    testStruct{Text: "test text", Number: 1, Status: true},
			sameCheck: true,
		},
		{
			name:      "check if different hash is generated for different string input",
			input1:    "test data",
			input2:    "test data 2",
			sameCheck: false,
		},
		{
			name:      "check if different hash is generated for different struct input",
			input1:    testStruct{Text: "test text", Number: 1, Status: true},
			input2:    testStruct{Text: "test text 2", Number: 2, Status: false},
			sameCheck: false,
		},
	}

	for _, test := range table {
		var input1Bytes []byte

		var err error

		var hash1 string

		input1Bytes, err = json.Marshal(test.input1)
		require.NoError(t, err)

		hash1 = GenerateHash(input1Bytes)

		if test.sameCheck {
			hash2 := GenerateHash(input1Bytes)

			assert.Equal(t, hash1, hash2)
		} else {
			input2Bytes, err := json.Marshal(test.input2)
			require.NoError(t, err)

			hash2 := GenerateHash(input2Bytes)

			assert.NotEqual(t, hash1, hash2)
		}
	}
}

func TestFirstIn(t *testing.T) {
	table := []struct {
		name   string
		m      map[string]any
		keys   []string
		expect string
	}{
		{
			name:   "nil map",
			keys:   []string{"foo", "bar"},
			expect: "",
		},
		{
			name:   "empty map",
			m:      map[string]any{},
			keys:   []string{"foo", "bar"},
			expect: "",
		},
		{
			name: "no match",
			m: map[string]any{
				"baz": "baz",
			},
			keys:   []string{"foo", "bar"},
			expect: "",
		},
		{
			name: "no keys",
			m: map[string]any{
				"baz": "baz",
			},
			keys:   []string{},
			expect: "",
		},
		{
			name: "nil match",
			m: map[string]any{
				"foo": nil,
			},
			keys:   []string{"foo", "bar"},
			expect: "",
		},
		{
			name: "empty match",
			m: map[string]any{
				"foo": "",
			},
			keys:   []string{"foo", "bar"},
			expect: "",
		},
		{
			name: "matches first key",
			m: map[string]any{
				"foo": "fnords",
			},
			keys:   []string{"foo", "bar"},
			expect: "fnords",
		},
		{
			name: "matches second key",
			m: map[string]any{
				"bar": "smarf",
			},
			keys:   []string{"foo", "bar"},
			expect: "smarf",
		},
		{
			name: "matches second key with nil first match",
			m: map[string]any{
				"foo": nil,
				"bar": "smarf",
			},
			keys:   []string{"foo", "bar"},
			expect: "smarf",
		},
		{
			name: "matches second key with empty first match",
			m: map[string]any{
				"foo": "",
				"bar": "smarf",
			},
			keys:   []string{"foo", "bar"},
			expect: "smarf",
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			result := FirstIn(test.m, test.keys...)
			assert.Equal(t, test.expect, result)
		})
	}
}

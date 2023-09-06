package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

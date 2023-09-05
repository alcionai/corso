package str

import (
	"fmt"
	"strconv"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
)

// parseBool returns the bool value represented by the string
// or false on error
func ParseBool(v string) bool {
	s, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}

	return s
}

func AnyValueToString(k string, m map[string]any) (string, error) {
	if len(m) == 0 {
		return "", clues.New("missing entry").With("map_key", k)
	}

	return AnyToString(m[k])
}

func AnyToString(a any) (string, error) {
	if a == nil {
		return "", clues.New("missing value")
	}

	sp, ok := a.(*string)
	if ok {
		return ptr.Val(sp), nil
	}

	s, ok := a.(string)
	if ok {
		return s, nil
	}

	return "", clues.New(fmt.Sprintf("unexpected type: %T", a))
}

// First returns the first non-zero valued string
func First(vs ...string) string {
	for _, v := range vs {
		if len(v) > 0 {
			return v
		}
	}

	return ""
}

// Preview reduces the string to the specified size.
// If the string is longer than the size, the last three
// characters are replaced with an ellipsis.  Size < 4
// will default to 4.
// ex:
// Preview("123", 6) => "123"
// Preview("1234567", 6) "123..."
func Preview(s string, size int) string {
	if size < 4 {
		size = 4
	}

	if len(s) < size {
		return s
	}

	ss := s[:size]
	if len(s) > size {
		ss = s[:size-3] + "..."
	}

	return ss
}

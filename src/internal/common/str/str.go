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

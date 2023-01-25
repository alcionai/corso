package ediscovery
import (
    "errors"
)
// Provides operations to call the add method.
type ChildSelectability int

const (
    ONE_CHILDSELECTABILITY ChildSelectability = iota
    MANY_CHILDSELECTABILITY
)

func (i ChildSelectability) String() string {
    return []string{"One", "Many"}[i]
}
func ParseChildSelectability(v string) (interface{}, error) {
    result := ONE_CHILDSELECTABILITY
    switch v {
        case "One":
            result = ONE_CHILDSELECTABILITY
        case "Many":
            result = MANY_CHILDSELECTABILITY
        default:
            return 0, errors.New("Unknown ChildSelectability value: " + v)
    }
    return &result, nil
}
func SerializeChildSelectability(values []ChildSelectability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

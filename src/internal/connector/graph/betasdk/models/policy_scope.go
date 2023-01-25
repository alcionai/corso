package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type PolicyScope int

const (
    NONE_POLICYSCOPE PolicyScope = iota
    ALL_POLICYSCOPE
    SELECTED_POLICYSCOPE
    UNKNOWNFUTUREVALUE_POLICYSCOPE
)

func (i PolicyScope) String() string {
    return []string{"none", "all", "selected", "unknownFutureValue"}[i]
}
func ParsePolicyScope(v string) (interface{}, error) {
    result := NONE_POLICYSCOPE
    switch v {
        case "none":
            result = NONE_POLICYSCOPE
        case "all":
            result = ALL_POLICYSCOPE
        case "selected":
            result = SELECTED_POLICYSCOPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_POLICYSCOPE
        default:
            return 0, errors.New("Unknown PolicyScope value: " + v)
    }
    return &result, nil
}
func SerializePolicyScope(values []PolicyScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

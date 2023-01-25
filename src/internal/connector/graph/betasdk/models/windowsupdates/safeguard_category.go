package windowsupdates
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type SafeguardCategory int

const (
    LIKELYISSUES_SAFEGUARDCATEGORY SafeguardCategory = iota
    UNKNOWNFUTUREVALUE_SAFEGUARDCATEGORY
)

func (i SafeguardCategory) String() string {
    return []string{"likelyIssues", "unknownFutureValue"}[i]
}
func ParseSafeguardCategory(v string) (interface{}, error) {
    result := LIKELYISSUES_SAFEGUARDCATEGORY
    switch v {
        case "likelyIssues":
            result = LIKELYISSUES_SAFEGUARDCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SAFEGUARDCATEGORY
        default:
            return 0, errors.New("Unknown SafeguardCategory value: " + v)
    }
    return &result, nil
}
func SerializeSafeguardCategory(values []SafeguardCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package windowsupdates
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UpdateCategory int

const (
    FEATURE_UPDATECATEGORY UpdateCategory = iota
    QUALITY_UPDATECATEGORY
    UNKNOWNFUTUREVALUE_UPDATECATEGORY
)

func (i UpdateCategory) String() string {
    return []string{"feature", "quality", "unknownFutureValue"}[i]
}
func ParseUpdateCategory(v string) (interface{}, error) {
    result := FEATURE_UPDATECATEGORY
    switch v {
        case "feature":
            result = FEATURE_UPDATECATEGORY
        case "quality":
            result = QUALITY_UPDATECATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_UPDATECATEGORY
        default:
            return 0, errors.New("Unknown UpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeUpdateCategory(values []UpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

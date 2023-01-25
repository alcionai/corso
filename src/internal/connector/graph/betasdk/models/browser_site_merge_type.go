package models
import (
    "errors"
)
// Provides operations to call the add method.
type BrowserSiteMergeType int

const (
    // No merge type
    NOMERGE_BROWSERSITEMERGETYPE BrowserSiteMergeType = iota
    // Default merge type
    DEFAULT_ESCAPED_BROWSERSITEMERGETYPE
    // Placeholder for evolvable enum, but this enum is never returned to the caller, so it shouldn't be necessary.
    UNKNOWNFUTUREVALUE_BROWSERSITEMERGETYPE
)

func (i BrowserSiteMergeType) String() string {
    return []string{"noMerge", "default", "unknownFutureValue"}[i]
}
func ParseBrowserSiteMergeType(v string) (interface{}, error) {
    result := NOMERGE_BROWSERSITEMERGETYPE
    switch v {
        case "noMerge":
            result = NOMERGE_BROWSERSITEMERGETYPE
        case "default":
            result = DEFAULT_ESCAPED_BROWSERSITEMERGETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BROWSERSITEMERGETYPE
        default:
            return 0, errors.New("Unknown BrowserSiteMergeType value: " + v)
    }
    return &result, nil
}
func SerializeBrowserSiteMergeType(values []BrowserSiteMergeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

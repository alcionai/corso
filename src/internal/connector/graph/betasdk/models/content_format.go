package models
import (
    "errors"
)
// Provides operations to call the evaluateClassificationResults method.
type ContentFormat int

const (
    DEFAULT_ESCAPED_CONTENTFORMAT ContentFormat = iota
    EMAIL_CONTENTFORMAT
)

func (i ContentFormat) String() string {
    return []string{"default", "email"}[i]
}
func ParseContentFormat(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_CONTENTFORMAT
    switch v {
        case "default":
            result = DEFAULT_ESCAPED_CONTENTFORMAT
        case "email":
            result = EMAIL_CONTENTFORMAT
        default:
            return 0, errors.New("Unknown ContentFormat value: " + v)
    }
    return &result, nil
}
func SerializeContentFormat(values []ContentFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

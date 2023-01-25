package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EmailType int

const (
    UNKNOWN_EMAILTYPE EmailType = iota
    WORK_EMAILTYPE
    PERSONAL_EMAILTYPE
    MAIN_EMAILTYPE
    OTHER_EMAILTYPE
)

func (i EmailType) String() string {
    return []string{"unknown", "work", "personal", "main", "other"}[i]
}
func ParseEmailType(v string) (interface{}, error) {
    result := UNKNOWN_EMAILTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_EMAILTYPE
        case "work":
            result = WORK_EMAILTYPE
        case "personal":
            result = PERSONAL_EMAILTYPE
        case "main":
            result = MAIN_EMAILTYPE
        case "other":
            result = OTHER_EMAILTYPE
        default:
            return 0, errors.New("Unknown EmailType value: " + v)
    }
    return &result, nil
}
func SerializeEmailType(values []EmailType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

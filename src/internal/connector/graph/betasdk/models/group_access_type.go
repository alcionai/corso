package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type GroupAccessType int

const (
    NONE_GROUPACCESSTYPE GroupAccessType = iota
    PRIVATE_GROUPACCESSTYPE
    SECRET_GROUPACCESSTYPE
    PUBLIC_GROUPACCESSTYPE
)

func (i GroupAccessType) String() string {
    return []string{"none", "private", "secret", "public"}[i]
}
func ParseGroupAccessType(v string) (interface{}, error) {
    result := NONE_GROUPACCESSTYPE
    switch v {
        case "none":
            result = NONE_GROUPACCESSTYPE
        case "private":
            result = PRIVATE_GROUPACCESSTYPE
        case "secret":
            result = SECRET_GROUPACCESSTYPE
        case "public":
            result = PUBLIC_GROUPACCESSTYPE
        default:
            return 0, errors.New("Unknown GroupAccessType value: " + v)
    }
    return &result, nil
}
func SerializeGroupAccessType(values []GroupAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

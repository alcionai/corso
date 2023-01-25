package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type GroupPrivacy int

const (
    UNSPECIFIED_GROUPPRIVACY GroupPrivacy = iota
    PUBLIC_GROUPPRIVACY
    PRIVATE_GROUPPRIVACY
    UNKNOWNFUTUREVALUE_GROUPPRIVACY
)

func (i GroupPrivacy) String() string {
    return []string{"unspecified", "public", "private", "unknownFutureValue"}[i]
}
func ParseGroupPrivacy(v string) (interface{}, error) {
    result := UNSPECIFIED_GROUPPRIVACY
    switch v {
        case "unspecified":
            result = UNSPECIFIED_GROUPPRIVACY
        case "public":
            result = PUBLIC_GROUPPRIVACY
        case "private":
            result = PRIVATE_GROUPPRIVACY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GROUPPRIVACY
        default:
            return 0, errors.New("Unknown GroupPrivacy value: " + v)
    }
    return &result, nil
}
func SerializeGroupPrivacy(values []GroupPrivacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

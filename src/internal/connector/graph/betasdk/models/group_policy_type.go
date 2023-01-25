package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type GroupPolicyType int

const (
    // Group Policy administrative templates built-in to the Policy configuration service provider (CSP).
    ADMXBACKED_GROUPPOLICYTYPE GroupPolicyType = iota
    // Group Policy administrative templates installed using the Policy configuration service provider (CSP).
    ADMXINGESTED_GROUPPOLICYTYPE
)

func (i GroupPolicyType) String() string {
    return []string{"admxBacked", "admxIngested"}[i]
}
func ParseGroupPolicyType(v string) (interface{}, error) {
    result := ADMXBACKED_GROUPPOLICYTYPE
    switch v {
        case "admxBacked":
            result = ADMXBACKED_GROUPPOLICYTYPE
        case "admxIngested":
            result = ADMXINGESTED_GROUPPOLICYTYPE
        default:
            return 0, errors.New("Unknown GroupPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyType(values []GroupPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

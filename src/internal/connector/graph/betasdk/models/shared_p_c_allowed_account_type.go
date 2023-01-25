package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SharedPCAllowedAccountType int

const (
    // Not configured. Default value.
    NOTCONFIGURED_SHAREDPCALLOWEDACCOUNTTYPE SharedPCAllowedAccountType = iota
    // Only guest accounts.
    GUEST_SHAREDPCALLOWEDACCOUNTTYPE
    // Only domain-joined accounts.
    DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
)

func (i SharedPCAllowedAccountType) String() string {
    return []string{"notConfigured", "guest", "domain"}[i]
}
func ParseSharedPCAllowedAccountType(v string) (interface{}, error) {
    result := NOTCONFIGURED_SHAREDPCALLOWEDACCOUNTTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_SHAREDPCALLOWEDACCOUNTTYPE
        case "guest":
            result = GUEST_SHAREDPCALLOWEDACCOUNTTYPE
        case "domain":
            result = DOMAIN_SHAREDPCALLOWEDACCOUNTTYPE
        default:
            return 0, errors.New("Unknown SharedPCAllowedAccountType value: " + v)
    }
    return &result, nil
}
func SerializeSharedPCAllowedAccountType(values []SharedPCAllowedAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

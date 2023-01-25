package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AuthenticationAppAdminConfiguration int

const (
    NOTAPPLICABLE_AUTHENTICATIONAPPADMINCONFIGURATION AuthenticationAppAdminConfiguration = iota
    ENABLED_AUTHENTICATIONAPPADMINCONFIGURATION
    DISABLED_AUTHENTICATIONAPPADMINCONFIGURATION
    UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPADMINCONFIGURATION
)

func (i AuthenticationAppAdminConfiguration) String() string {
    return []string{"notApplicable", "enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseAuthenticationAppAdminConfiguration(v string) (interface{}, error) {
    result := NOTAPPLICABLE_AUTHENTICATIONAPPADMINCONFIGURATION
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_AUTHENTICATIONAPPADMINCONFIGURATION
        case "enabled":
            result = ENABLED_AUTHENTICATIONAPPADMINCONFIGURATION
        case "disabled":
            result = DISABLED_AUTHENTICATIONAPPADMINCONFIGURATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHENTICATIONAPPADMINCONFIGURATION
        default:
            return 0, errors.New("Unknown AuthenticationAppAdminConfiguration value: " + v)
    }
    return &result, nil
}
func SerializeAuthenticationAppAdminConfiguration(values []AuthenticationAppAdminConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

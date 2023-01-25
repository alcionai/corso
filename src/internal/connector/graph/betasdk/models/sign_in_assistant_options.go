package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SignInAssistantOptions int

const (
    // Not configured - wlidsvc Start will be set to SERVICE_DEMAND_START.
    NOTCONFIGURED_SIGNINASSISTANTOPTIONS SignInAssistantOptions = iota
    // Disabled - wlidsvc Start will be set to SERVICE_DISABLED.
    DISABLED_SIGNINASSISTANTOPTIONS
)

func (i SignInAssistantOptions) String() string {
    return []string{"notConfigured", "disabled"}[i]
}
func ParseSignInAssistantOptions(v string) (interface{}, error) {
    result := NOTCONFIGURED_SIGNINASSISTANTOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_SIGNINASSISTANTOPTIONS
        case "disabled":
            result = DISABLED_SIGNINASSISTANTOPTIONS
        default:
            return 0, errors.New("Unknown SignInAssistantOptions value: " + v)
    }
    return &result, nil
}
func SerializeSignInAssistantOptions(values []SignInAssistantOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

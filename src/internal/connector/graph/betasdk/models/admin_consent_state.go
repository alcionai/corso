package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AdminConsentState int

const (
    // Admin did not configure the item
    NOTCONFIGURED_ADMINCONSENTSTATE AdminConsentState = iota
    // Admin granted item
    GRANTED_ADMINCONSENTSTATE
    // Admin deos not grant item
    NOTGRANTED_ADMINCONSENTSTATE
)

func (i AdminConsentState) String() string {
    return []string{"notConfigured", "granted", "notGranted"}[i]
}
func ParseAdminConsentState(v string) (interface{}, error) {
    result := NOTCONFIGURED_ADMINCONSENTSTATE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ADMINCONSENTSTATE
        case "granted":
            result = GRANTED_ADMINCONSENTSTATE
        case "notGranted":
            result = NOTGRANTED_ADMINCONSENTSTATE
        default:
            return 0, errors.New("Unknown AdminConsentState value: " + v)
    }
    return &result, nil
}
func SerializeAdminConsentState(values []AdminConsentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

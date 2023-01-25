package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AppInstallControlType int

const (
    // Not configured
    NOTCONFIGURED_APPINSTALLCONTROLTYPE AppInstallControlType = iota
    // Turn off app recommendations
    ANYWHERE_APPINSTALLCONTROLTYPE
    // Allow apps from Store only
    STOREONLY_APPINSTALLCONTROLTYPE
    // Show me app recommendations
    RECOMMENDATIONS_APPINSTALLCONTROLTYPE
    // Warn me before installing apps from outside the Store
    PREFERSTORE_APPINSTALLCONTROLTYPE
)

func (i AppInstallControlType) String() string {
    return []string{"notConfigured", "anywhere", "storeOnly", "recommendations", "preferStore"}[i]
}
func ParseAppInstallControlType(v string) (interface{}, error) {
    result := NOTCONFIGURED_APPINSTALLCONTROLTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_APPINSTALLCONTROLTYPE
        case "anywhere":
            result = ANYWHERE_APPINSTALLCONTROLTYPE
        case "storeOnly":
            result = STOREONLY_APPINSTALLCONTROLTYPE
        case "recommendations":
            result = RECOMMENDATIONS_APPINSTALLCONTROLTYPE
        case "preferStore":
            result = PREFERSTORE_APPINSTALLCONTROLTYPE
        default:
            return 0, errors.New("Unknown AppInstallControlType value: " + v)
    }
    return &result, nil
}
func SerializeAppInstallControlType(values []AppInstallControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MobileAppActionType int

const (
    // Unknown result.
    UNKNOWN_MOBILEAPPACTIONTYPE MobileAppActionType = iota
    // Application install command was sent.
    INSTALLCOMMANDSENT_MOBILEAPPACTIONTYPE
    // Application installed.
    INSTALLED_MOBILEAPPACTIONTYPE
    // Application uninstalled.
    UNINSTALLED_MOBILEAPPACTIONTYPE
    // User requested installation
    USERREQUESTEDINSTALL_MOBILEAPPACTIONTYPE
)

func (i MobileAppActionType) String() string {
    return []string{"unknown", "installCommandSent", "installed", "uninstalled", "userRequestedInstall"}[i]
}
func ParseMobileAppActionType(v string) (interface{}, error) {
    result := UNKNOWN_MOBILEAPPACTIONTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_MOBILEAPPACTIONTYPE
        case "installCommandSent":
            result = INSTALLCOMMANDSENT_MOBILEAPPACTIONTYPE
        case "installed":
            result = INSTALLED_MOBILEAPPACTIONTYPE
        case "uninstalled":
            result = UNINSTALLED_MOBILEAPPACTIONTYPE
        case "userRequestedInstall":
            result = USERREQUESTEDINSTALL_MOBILEAPPACTIONTYPE
        default:
            return 0, errors.New("Unknown MobileAppActionType value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppActionType(values []MobileAppActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

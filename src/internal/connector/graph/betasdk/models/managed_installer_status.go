package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedInstallerStatus int

const (
    // Managed Installer is Disabled
    DISABLED_MANAGEDINSTALLERSTATUS ManagedInstallerStatus = iota
    // Managed Installer is Enabled
    ENABLED_MANAGEDINSTALLERSTATUS
)

func (i ManagedInstallerStatus) String() string {
    return []string{"disabled", "enabled"}[i]
}
func ParseManagedInstallerStatus(v string) (interface{}, error) {
    result := DISABLED_MANAGEDINSTALLERSTATUS
    switch v {
        case "disabled":
            result = DISABLED_MANAGEDINSTALLERSTATUS
        case "enabled":
            result = ENABLED_MANAGEDINSTALLERSTATUS
        default:
            return 0, errors.New("Unknown ManagedInstallerStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagedInstallerStatus(values []ManagedInstallerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

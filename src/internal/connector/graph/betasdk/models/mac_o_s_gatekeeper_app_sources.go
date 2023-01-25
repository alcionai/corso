package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSGatekeeperAppSources int

const (
    // Device default value, no intent.
    NOTCONFIGURED_MACOSGATEKEEPERAPPSOURCES MacOSGatekeeperAppSources = iota
    // Only apps from the Mac AppStore can be run.
    MACAPPSTORE_MACOSGATEKEEPERAPPSOURCES
    // Only apps from the Mac AppStore and identified developers can be run.
    MACAPPSTOREANDIDENTIFIEDDEVELOPERS_MACOSGATEKEEPERAPPSOURCES
    // Apps from anywhere can be run.
    ANYWHERE_MACOSGATEKEEPERAPPSOURCES
)

func (i MacOSGatekeeperAppSources) String() string {
    return []string{"notConfigured", "macAppStore", "macAppStoreAndIdentifiedDevelopers", "anywhere"}[i]
}
func ParseMacOSGatekeeperAppSources(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSGATEKEEPERAPPSOURCES
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSGATEKEEPERAPPSOURCES
        case "macAppStore":
            result = MACAPPSTORE_MACOSGATEKEEPERAPPSOURCES
        case "macAppStoreAndIdentifiedDevelopers":
            result = MACAPPSTOREANDIDENTIFIEDDEVELOPERS_MACOSGATEKEEPERAPPSOURCES
        case "anywhere":
            result = ANYWHERE_MACOSGATEKEEPERAPPSOURCES
        default:
            return 0, errors.New("Unknown MacOSGatekeeperAppSources value: " + v)
    }
    return &result, nil
}
func SerializeMacOSGatekeeperAppSources(values []MacOSGatekeeperAppSources) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type LocalSecurityOptionsSmartCardRemovalBehaviorType int

const (
    // No Action
    NOACTION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE LocalSecurityOptionsSmartCardRemovalBehaviorType = iota
    // Lock Workstation
    LOCKWORKSTATION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
    // Force Logoff
    FORCELOGOFF_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
    // Disconnect if a remote Remote Desktop Services session
    DISCONNECTREMOTEDESKTOPSESSION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
)

func (i LocalSecurityOptionsSmartCardRemovalBehaviorType) String() string {
    return []string{"noAction", "lockWorkstation", "forceLogoff", "disconnectRemoteDesktopSession"}[i]
}
func ParseLocalSecurityOptionsSmartCardRemovalBehaviorType(v string) (interface{}, error) {
    result := NOACTION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
    switch v {
        case "noAction":
            result = NOACTION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
        case "lockWorkstation":
            result = LOCKWORKSTATION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
        case "forceLogoff":
            result = FORCELOGOFF_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
        case "disconnectRemoteDesktopSession":
            result = DISCONNECTREMOTEDESKTOPSESSION_LOCALSECURITYOPTIONSSMARTCARDREMOVALBEHAVIORTYPE
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsSmartCardRemovalBehaviorType value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsSmartCardRemovalBehaviorType(values []LocalSecurityOptionsSmartCardRemovalBehaviorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

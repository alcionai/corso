package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedAppDeviceThreatLevel int

const (
    // Value not configured
    NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL ManagedAppDeviceThreatLevel = iota
    // Device needs to have no threat
    SECURED_MANAGEDAPPDEVICETHREATLEVEL
    // Device needs to have a low threat.
    LOW_MANAGEDAPPDEVICETHREATLEVEL
    // Device needs to have not more than medium threat.
    MEDIUM_MANAGEDAPPDEVICETHREATLEVEL
    // Device needs to have not more than high threat
    HIGH_MANAGEDAPPDEVICETHREATLEVEL
)

func (i ManagedAppDeviceThreatLevel) String() string {
    return []string{"notConfigured", "secured", "low", "medium", "high"}[i]
}
func ParseManagedAppDeviceThreatLevel(v string) (interface{}, error) {
    result := NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MANAGEDAPPDEVICETHREATLEVEL
        case "secured":
            result = SECURED_MANAGEDAPPDEVICETHREATLEVEL
        case "low":
            result = LOW_MANAGEDAPPDEVICETHREATLEVEL
        case "medium":
            result = MEDIUM_MANAGEDAPPDEVICETHREATLEVEL
        case "high":
            result = HIGH_MANAGEDAPPDEVICETHREATLEVEL
        default:
            return 0, errors.New("Unknown ManagedAppDeviceThreatLevel value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppDeviceThreatLevel(values []ManagedAppDeviceThreatLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

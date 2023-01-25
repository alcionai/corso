package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DiscoverySource int

const (
    // DiscoverySource is Unknown.
    UNKNOWN_DISCOVERYSOURCE DiscoverySource = iota
    // Device is imported by admin.
    ADMINIMPORT_DISCOVERYSOURCE
    // Device is added by Apple device enrollment program (Dep).
    DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE
)

func (i DiscoverySource) String() string {
    return []string{"unknown", "adminImport", "deviceEnrollmentProgram"}[i]
}
func ParseDiscoverySource(v string) (interface{}, error) {
    result := UNKNOWN_DISCOVERYSOURCE
    switch v {
        case "unknown":
            result = UNKNOWN_DISCOVERYSOURCE
        case "adminImport":
            result = ADMINIMPORT_DISCOVERYSOURCE
        case "deviceEnrollmentProgram":
            result = DEVICEENROLLMENTPROGRAM_DISCOVERYSOURCE
        default:
            return 0, errors.New("Unknown DiscoverySource value: " + v)
    }
    return &result, nil
}
func SerializeDiscoverySource(values []DiscoverySource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

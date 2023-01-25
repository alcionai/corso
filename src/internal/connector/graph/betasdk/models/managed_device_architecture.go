package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedDeviceArchitecture int

const (
    // Unknown architecture
    UNKNOWN_MANAGEDDEVICEARCHITECTURE ManagedDeviceArchitecture = iota
    // X86
    X86_MANAGEDDEVICEARCHITECTURE
    // X64
    X64_MANAGEDDEVICEARCHITECTURE
    // ARM
    ARM_MANAGEDDEVICEARCHITECTURE
    // ARM64
    ARM64_MANAGEDDEVICEARCHITECTURE
)

func (i ManagedDeviceArchitecture) String() string {
    return []string{"unknown", "x86", "x64", "arm", "arM64"}[i]
}
func ParseManagedDeviceArchitecture(v string) (interface{}, error) {
    result := UNKNOWN_MANAGEDDEVICEARCHITECTURE
    switch v {
        case "unknown":
            result = UNKNOWN_MANAGEDDEVICEARCHITECTURE
        case "x86":
            result = X86_MANAGEDDEVICEARCHITECTURE
        case "x64":
            result = X64_MANAGEDDEVICEARCHITECTURE
        case "arm":
            result = ARM_MANAGEDDEVICEARCHITECTURE
        case "arM64":
            result = ARM64_MANAGEDDEVICEARCHITECTURE
        default:
            return 0, errors.New("Unknown ManagedDeviceArchitecture value: " + v)
    }
    return &result, nil
}
func SerializeManagedDeviceArchitecture(values []ManagedDeviceArchitecture) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

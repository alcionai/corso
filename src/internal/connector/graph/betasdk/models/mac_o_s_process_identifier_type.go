package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSProcessIdentifierType int

const (
    // Indicates an app with a bundle ID.
    BUNDLEID_MACOSPROCESSIDENTIFIERTYPE MacOSProcessIdentifierType = iota
    // Indicates a file path for a process.
    PATH_MACOSPROCESSIDENTIFIERTYPE
)

func (i MacOSProcessIdentifierType) String() string {
    return []string{"bundleID", "path"}[i]
}
func ParseMacOSProcessIdentifierType(v string) (interface{}, error) {
    result := BUNDLEID_MACOSPROCESSIDENTIFIERTYPE
    switch v {
        case "bundleID":
            result = BUNDLEID_MACOSPROCESSIDENTIFIERTYPE
        case "path":
            result = PATH_MACOSPROCESSIDENTIFIERTYPE
        default:
            return 0, errors.New("Unknown MacOSProcessIdentifierType value: " + v)
    }
    return &result, nil
}
func SerializeMacOSProcessIdentifierType(values []MacOSProcessIdentifierType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

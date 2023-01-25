package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type WindowsQualityUpdateClassification int

const (
    // All update type
    ALL_WINDOWSQUALITYUPDATECLASSIFICATION WindowsQualityUpdateClassification = iota
    // Security only update type
    SECURITY_WINDOWSQUALITYUPDATECLASSIFICATION
    // Non security only update type
    NONSECURITY_WINDOWSQUALITYUPDATECLASSIFICATION
)

func (i WindowsQualityUpdateClassification) String() string {
    return []string{"all", "security", "nonSecurity"}[i]
}
func ParseWindowsQualityUpdateClassification(v string) (interface{}, error) {
    result := ALL_WINDOWSQUALITYUPDATECLASSIFICATION
    switch v {
        case "all":
            result = ALL_WINDOWSQUALITYUPDATECLASSIFICATION
        case "security":
            result = SECURITY_WINDOWSQUALITYUPDATECLASSIFICATION
        case "nonSecurity":
            result = NONSECURITY_WINDOWSQUALITYUPDATECLASSIFICATION
        default:
            return 0, errors.New("Unknown WindowsQualityUpdateClassification value: " + v)
    }
    return &result, nil
}
func SerializeWindowsQualityUpdateClassification(values []WindowsQualityUpdateClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

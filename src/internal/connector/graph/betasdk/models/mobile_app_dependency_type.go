package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MobileAppDependencyType int

const (
    // Indicates that the child app should be detected before installing the parent app.
    DETECT_MOBILEAPPDEPENDENCYTYPE MobileAppDependencyType = iota
    // Indicates that the child app should be installed before installing the parent app.
    AUTOINSTALL_MOBILEAPPDEPENDENCYTYPE
)

func (i MobileAppDependencyType) String() string {
    return []string{"detect", "autoInstall"}[i]
}
func ParseMobileAppDependencyType(v string) (interface{}, error) {
    result := DETECT_MOBILEAPPDEPENDENCYTYPE
    switch v {
        case "detect":
            result = DETECT_MOBILEAPPDEPENDENCYTYPE
        case "autoInstall":
            result = AUTOINSTALL_MOBILEAPPDEPENDENCYTYPE
        default:
            return 0, errors.New("Unknown MobileAppDependencyType value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppDependencyType(values []MobileAppDependencyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

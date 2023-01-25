package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSContentCachingType int

const (
    // Default. Both user iCloud data and non-iCloud data will be cached.
    NOTCONFIGURED_MACOSCONTENTCACHINGTYPE MacOSContentCachingType = iota
    // Allow Apple's content caching service to cache user iCloud data.
    USERCONTENTONLY_MACOSCONTENTCACHINGTYPE
    // Allow Apple's content caching service to cache non-iCloud data (e.g. app and software updates).
    SHAREDCONTENTONLY_MACOSCONTENTCACHINGTYPE
)

func (i MacOSContentCachingType) String() string {
    return []string{"notConfigured", "userContentOnly", "sharedContentOnly"}[i]
}
func ParseMacOSContentCachingType(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSCONTENTCACHINGTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSCONTENTCACHINGTYPE
        case "userContentOnly":
            result = USERCONTENTONLY_MACOSCONTENTCACHINGTYPE
        case "sharedContentOnly":
            result = SHAREDCONTENTONLY_MACOSCONTENTCACHINGTYPE
        default:
            return 0, errors.New("Unknown MacOSContentCachingType value: " + v)
    }
    return &result, nil
}
func SerializeMacOSContentCachingType(values []MacOSContentCachingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

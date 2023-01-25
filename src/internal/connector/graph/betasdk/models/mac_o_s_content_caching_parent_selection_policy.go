package models
import (
    "errors"
)
// Provides operations to call the add method.
type MacOSContentCachingParentSelectionPolicy int

const (
    // Defaults to round-robin strategy.
    NOTCONFIGURED_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY MacOSContentCachingParentSelectionPolicy = iota
    // Rotate through the parents in order. Use this policy for load balancing.
    ROUNDROBIN_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
    // Always use the first available parent in the Parents list. Use this policy to designate permanent primary, secondary, and subsequent parents.
    FIRSTAVAILABLE_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
    // Hash the path part of the requested URL so that the same parent is always used for the same URL. This is useful for maximizing the size of the combined caches of the parents.
    URLPATHHASH_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
    // Choose a parent at random. Use this policy for load balancing.
    RANDOM_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
    // Use the first available parent that is available in the Parents list until it becomes unavailable, then advance to the next one. Use this policy for designating floating primary, secondary, and subsequent parents.
    STICKYAVAILABLE_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
)

func (i MacOSContentCachingParentSelectionPolicy) String() string {
    return []string{"notConfigured", "roundRobin", "firstAvailable", "urlPathHash", "random", "stickyAvailable"}[i]
}
func ParseMacOSContentCachingParentSelectionPolicy(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        case "roundRobin":
            result = ROUNDROBIN_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        case "firstAvailable":
            result = FIRSTAVAILABLE_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        case "urlPathHash":
            result = URLPATHHASH_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        case "random":
            result = RANDOM_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        case "stickyAvailable":
            result = STICKYAVAILABLE_MACOSCONTENTCACHINGPARENTSELECTIONPOLICY
        default:
            return 0, errors.New("Unknown MacOSContentCachingParentSelectionPolicy value: " + v)
    }
    return &result, nil
}
func SerializeMacOSContentCachingParentSelectionPolicy(values []MacOSContentCachingParentSelectionPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

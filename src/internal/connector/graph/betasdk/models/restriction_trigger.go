package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RestrictionTrigger int

const (
    COPYPASTE_RESTRICTIONTRIGGER RestrictionTrigger = iota
    COPYTONETWORKSHARE_RESTRICTIONTRIGGER
    COPYTOREMOVABLEMEDIA_RESTRICTIONTRIGGER
    SCREENCAPTURE_RESTRICTIONTRIGGER
    PRINT_RESTRICTIONTRIGGER
    CLOUDEGRESS_RESTRICTIONTRIGGER
    UNALLOWEDAPPS_RESTRICTIONTRIGGER
)

func (i RestrictionTrigger) String() string {
    return []string{"copyPaste", "copyToNetworkShare", "copyToRemovableMedia", "screenCapture", "print", "cloudEgress", "unallowedApps"}[i]
}
func ParseRestrictionTrigger(v string) (interface{}, error) {
    result := COPYPASTE_RESTRICTIONTRIGGER
    switch v {
        case "copyPaste":
            result = COPYPASTE_RESTRICTIONTRIGGER
        case "copyToNetworkShare":
            result = COPYTONETWORKSHARE_RESTRICTIONTRIGGER
        case "copyToRemovableMedia":
            result = COPYTOREMOVABLEMEDIA_RESTRICTIONTRIGGER
        case "screenCapture":
            result = SCREENCAPTURE_RESTRICTIONTRIGGER
        case "print":
            result = PRINT_RESTRICTIONTRIGGER
        case "cloudEgress":
            result = CLOUDEGRESS_RESTRICTIONTRIGGER
        case "unallowedApps":
            result = UNALLOWEDAPPS_RESTRICTIONTRIGGER
        default:
            return 0, errors.New("Unknown RestrictionTrigger value: " + v)
    }
    return &result, nil
}
func SerializeRestrictionTrigger(values []RestrictionTrigger) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

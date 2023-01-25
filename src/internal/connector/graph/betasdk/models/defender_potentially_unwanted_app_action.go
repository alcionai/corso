package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DefenderPotentiallyUnwantedAppAction int

const (
    // PUA Protection is off. Defender will not protect against potentially unwanted applications.
    DEVICEDEFAULT_DEFENDERPOTENTIALLYUNWANTEDAPPACTION DefenderPotentiallyUnwantedAppAction = iota
    // PUA Protection is on. Detected items are blocked. They will show in history along with other threats.
    BLOCK_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
    // Audit mode. Defender will detect potentially unwanted applications, but take no actions. You can review information about applications Defender would have taken action against by searching for events created by Defender in the Event Viewer.
    AUDIT_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
)

func (i DefenderPotentiallyUnwantedAppAction) String() string {
    return []string{"deviceDefault", "block", "audit"}[i]
}
func ParseDefenderPotentiallyUnwantedAppAction(v string) (interface{}, error) {
    result := DEVICEDEFAULT_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
        case "block":
            result = BLOCK_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
        case "audit":
            result = AUDIT_DEFENDERPOTENTIALLYUNWANTEDAPPACTION
        default:
            return 0, errors.New("Unknown DefenderPotentiallyUnwantedAppAction value: " + v)
    }
    return &result, nil
}
func SerializeDefenderPotentiallyUnwantedAppAction(values []DefenderPotentiallyUnwantedAppAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

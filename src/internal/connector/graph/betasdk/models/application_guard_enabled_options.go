package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ApplicationGuardEnabledOptions int

const (
    // Not Configured
    NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS ApplicationGuardEnabledOptions = iota
    // Enabled For Edge
    ENABLEDFOREDGE_APPLICATIONGUARDENABLEDOPTIONS
    // Enabled For Office
    ENABLEDFOROFFICE_APPLICATIONGUARDENABLEDOPTIONS
    // Enabled For Edge And Office
    ENABLEDFOREDGEANDOFFICE_APPLICATIONGUARDENABLEDOPTIONS
)

func (i ApplicationGuardEnabledOptions) String() string {
    return []string{"notConfigured", "enabledForEdge", "enabledForOffice", "enabledForEdgeAndOffice"}[i]
}
func ParseApplicationGuardEnabledOptions(v string) (interface{}, error) {
    result := NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForEdge":
            result = ENABLEDFOREDGE_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForOffice":
            result = ENABLEDFOROFFICE_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForEdgeAndOffice":
            result = ENABLEDFOREDGEANDOFFICE_APPLICATIONGUARDENABLEDOPTIONS
        default:
            return 0, errors.New("Unknown ApplicationGuardEnabledOptions value: " + v)
    }
    return &result, nil
}
func SerializeApplicationGuardEnabledOptions(values []ApplicationGuardEnabledOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

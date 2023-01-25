package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedAppRemediationAction int

const (
    // app and the corresponding company data to be blocked
    BLOCK_MANAGEDAPPREMEDIATIONACTION ManagedAppRemediationAction = iota
    // app and the corresponding company data to be wiped
    WIPE_MANAGEDAPPREMEDIATIONACTION
    // app and the corresponding user to be warned
    WARN_MANAGEDAPPREMEDIATIONACTION
)

func (i ManagedAppRemediationAction) String() string {
    return []string{"block", "wipe", "warn"}[i]
}
func ParseManagedAppRemediationAction(v string) (interface{}, error) {
    result := BLOCK_MANAGEDAPPREMEDIATIONACTION
    switch v {
        case "block":
            result = BLOCK_MANAGEDAPPREMEDIATIONACTION
        case "wipe":
            result = WIPE_MANAGEDAPPREMEDIATIONACTION
        case "warn":
            result = WARN_MANAGEDAPPREMEDIATIONACTION
        default:
            return 0, errors.New("Unknown ManagedAppRemediationAction value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppRemediationAction(values []ManagedAppRemediationAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

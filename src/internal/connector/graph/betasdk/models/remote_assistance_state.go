package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RemoteAssistanceState int

const (
    // Remote assistance is disabled for the account. With this value, Quick Assist remote assistance sessions are not allowed for the account.
    DISABLED_REMOTEASSISTANCESTATE RemoteAssistanceState = iota
    // Remote assistance is enabled for the account. With this value, enterprise Quick Assist remote assistance features are provided.
    ENABLED_REMOTEASSISTANCESTATE
)

func (i RemoteAssistanceState) String() string {
    return []string{"disabled", "enabled"}[i]
}
func ParseRemoteAssistanceState(v string) (interface{}, error) {
    result := DISABLED_REMOTEASSISTANCESTATE
    switch v {
        case "disabled":
            result = DISABLED_REMOTEASSISTANCESTATE
        case "enabled":
            result = ENABLED_REMOTEASSISTANCESTATE
        default:
            return 0, errors.New("Unknown RemoteAssistanceState value: " + v)
    }
    return &result, nil
}
func SerializeRemoteAssistanceState(values []RemoteAssistanceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

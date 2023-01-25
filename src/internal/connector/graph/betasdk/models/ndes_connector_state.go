package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type NdesConnectorState int

const (
    // State not available yet for this connector.
    NONE_NDESCONNECTORSTATE NdesConnectorState = iota
    // Ndes connector has connected recently
    ACTIVE_NDESCONNECTORSTATE
    // No recent activity for the Ndes connector
    INACTIVE_NDESCONNECTORSTATE
)

func (i NdesConnectorState) String() string {
    return []string{"none", "active", "inactive"}[i]
}
func ParseNdesConnectorState(v string) (interface{}, error) {
    result := NONE_NDESCONNECTORSTATE
    switch v {
        case "none":
            result = NONE_NDESCONNECTORSTATE
        case "active":
            result = ACTIVE_NDESCONNECTORSTATE
        case "inactive":
            result = INACTIVE_NDESCONNECTORSTATE
        default:
            return 0, errors.New("Unknown NdesConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeNdesConnectorState(values []NdesConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

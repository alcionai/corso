package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type LostModeState int

const (
    // Lost mode is disabled.
    DISABLED_LOSTMODESTATE LostModeState = iota
    // Lost mode is enabled.
    ENABLED_LOSTMODESTATE
)

func (i LostModeState) String() string {
    return []string{"disabled", "enabled"}[i]
}
func ParseLostModeState(v string) (interface{}, error) {
    result := DISABLED_LOSTMODESTATE
    switch v {
        case "disabled":
            result = DISABLED_LOSTMODESTATE
        case "enabled":
            result = ENABLED_LOSTMODESTATE
        default:
            return 0, errors.New("Unknown LostModeState value: " + v)
    }
    return &result, nil
}
func SerializeLostModeState(values []LostModeState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

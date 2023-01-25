package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RestrictionAction int

const (
    WARN_RESTRICTIONACTION RestrictionAction = iota
    AUDIT_RESTRICTIONACTION
    BLOCK_RESTRICTIONACTION
)

func (i RestrictionAction) String() string {
    return []string{"warn", "audit", "block"}[i]
}
func ParseRestrictionAction(v string) (interface{}, error) {
    result := WARN_RESTRICTIONACTION
    switch v {
        case "warn":
            result = WARN_RESTRICTIONACTION
        case "audit":
            result = AUDIT_RESTRICTIONACTION
        case "block":
            result = BLOCK_RESTRICTIONACTION
        default:
            return 0, errors.New("Unknown RestrictionAction value: " + v)
    }
    return &result, nil
}
func SerializeRestrictionAction(values []RestrictionAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

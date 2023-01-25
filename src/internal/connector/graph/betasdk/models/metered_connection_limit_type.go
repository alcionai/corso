package models
import (
    "errors"
)
// Provides operations to call the add method.
type MeteredConnectionLimitType int

const (
    // Unrestricted
    UNRESTRICTED_METEREDCONNECTIONLIMITTYPE MeteredConnectionLimitType = iota
    // Fixed
    FIXED_METEREDCONNECTIONLIMITTYPE
    // Variable
    VARIABLE_METEREDCONNECTIONLIMITTYPE
)

func (i MeteredConnectionLimitType) String() string {
    return []string{"unrestricted", "fixed", "variable"}[i]
}
func ParseMeteredConnectionLimitType(v string) (interface{}, error) {
    result := UNRESTRICTED_METEREDCONNECTIONLIMITTYPE
    switch v {
        case "unrestricted":
            result = UNRESTRICTED_METEREDCONNECTIONLIMITTYPE
        case "fixed":
            result = FIXED_METEREDCONNECTIONLIMITTYPE
        case "variable":
            result = VARIABLE_METEREDCONNECTIONLIMITTYPE
        default:
            return 0, errors.New("Unknown MeteredConnectionLimitType value: " + v)
    }
    return &result, nil
}
func SerializeMeteredConnectionLimitType(values []MeteredConnectionLimitType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

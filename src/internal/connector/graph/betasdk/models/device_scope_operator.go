package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceScopeOperator int

const (
    // No operator set for the device scope configuration.
    NONE_DEVICESCOPEOPERATOR DeviceScopeOperator = iota
    // Operator for the device configuration query to be used (Equals).
    EQUALS_DEVICESCOPEOPERATOR
    // Placeholder value for future expansion enums such as notEquals, contains, notContains, greaterThan, lessThan.
    UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
)

func (i DeviceScopeOperator) String() string {
    return []string{"none", "equals", "unknownFutureValue"}[i]
}
func ParseDeviceScopeOperator(v string) (interface{}, error) {
    result := NONE_DEVICESCOPEOPERATOR
    switch v {
        case "none":
            result = NONE_DEVICESCOPEOPERATOR
        case "equals":
            result = EQUALS_DEVICESCOPEOPERATOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPEOPERATOR
        default:
            return 0, errors.New("Unknown DeviceScopeOperator value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeOperator(values []DeviceScopeOperator) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

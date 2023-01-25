package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MultiFactorAuthConfiguration int

const (
    NOTREQUIRED_MULTIFACTORAUTHCONFIGURATION MultiFactorAuthConfiguration = iota
    REQUIRED_MULTIFACTORAUTHCONFIGURATION
    UNKNOWNFUTUREVALUE_MULTIFACTORAUTHCONFIGURATION
)

func (i MultiFactorAuthConfiguration) String() string {
    return []string{"notRequired", "required", "unknownFutureValue"}[i]
}
func ParseMultiFactorAuthConfiguration(v string) (interface{}, error) {
    result := NOTREQUIRED_MULTIFACTORAUTHCONFIGURATION
    switch v {
        case "notRequired":
            result = NOTREQUIRED_MULTIFACTORAUTHCONFIGURATION
        case "required":
            result = REQUIRED_MULTIFACTORAUTHCONFIGURATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MULTIFACTORAUTHCONFIGURATION
        default:
            return 0, errors.New("Unknown MultiFactorAuthConfiguration value: " + v)
    }
    return &result, nil
}
func SerializeMultiFactorAuthConfiguration(values []MultiFactorAuthConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

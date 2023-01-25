package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RegistrationStatusType int

const (
    REGISTERED_REGISTRATIONSTATUSTYPE RegistrationStatusType = iota
    ENABLED_REGISTRATIONSTATUSTYPE
    CAPABLE_REGISTRATIONSTATUSTYPE
    MFAREGISTERED_REGISTRATIONSTATUSTYPE
    UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
)

func (i RegistrationStatusType) String() string {
    return []string{"registered", "enabled", "capable", "mfaRegistered", "unknownFutureValue"}[i]
}
func ParseRegistrationStatusType(v string) (interface{}, error) {
    result := REGISTERED_REGISTRATIONSTATUSTYPE
    switch v {
        case "registered":
            result = REGISTERED_REGISTRATIONSTATUSTYPE
        case "enabled":
            result = ENABLED_REGISTRATIONSTATUSTYPE
        case "capable":
            result = CAPABLE_REGISTRATIONSTATUSTYPE
        case "mfaRegistered":
            result = MFAREGISTERED_REGISTRATIONSTATUSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REGISTRATIONSTATUSTYPE
        default:
            return 0, errors.New("Unknown RegistrationStatusType value: " + v)
    }
    return &result, nil
}
func SerializeRegistrationStatusType(values []RegistrationStatusType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

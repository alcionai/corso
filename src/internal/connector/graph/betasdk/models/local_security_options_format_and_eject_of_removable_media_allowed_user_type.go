package models
import (
    "errors"
)
// Provides operations to call the add method.
type LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType = iota
    // Administrators
    ADMINISTRATORS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
    // Administrators and Power Users
    ADMINISTRATORSANDPOWERUSERS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
    // Administrators and Interactive Users 
    ADMINISTRATORSANDINTERACTIVEUSERS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
)

func (i LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) String() string {
    return []string{"notConfigured", "administrators", "administratorsAndPowerUsers", "administratorsAndInteractiveUsers"}[i]
}
func ParseLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType(v string) (interface{}, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
        case "administrators":
            result = ADMINISTRATORS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
        case "administratorsAndPowerUsers":
            result = ADMINISTRATORSANDPOWERUSERS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
        case "administratorsAndInteractiveUsers":
            result = ADMINISTRATORSANDINTERACTIVEUSERS_LOCALSECURITYOPTIONSFORMATANDEJECTOFREMOVABLEMEDIAALLOWEDUSERTYPE
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType(values []LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

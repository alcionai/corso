package models
import (
    "errors"
)
// Provides operations to call the add method.
type DefenderProtectionType int

const (
    // Device default value, no intent.
    USERDEFINED_DEFENDERPROTECTIONTYPE DefenderProtectionType = iota
    // Block functionality.
    ENABLE_DEFENDERPROTECTIONTYPE
    // Allow functionality but generate logs.
    AUDITMODE_DEFENDERPROTECTIONTYPE
    // Warning message to end user with ability to bypass block from attack surface reduction rule.
    WARN_DEFENDERPROTECTIONTYPE
    // Not configured.
    NOTCONFIGURED_DEFENDERPROTECTIONTYPE
)

func (i DefenderProtectionType) String() string {
    return []string{"userDefined", "enable", "auditMode", "warn", "notConfigured"}[i]
}
func ParseDefenderProtectionType(v string) (interface{}, error) {
    result := USERDEFINED_DEFENDERPROTECTIONTYPE
    switch v {
        case "userDefined":
            result = USERDEFINED_DEFENDERPROTECTIONTYPE
        case "enable":
            result = ENABLE_DEFENDERPROTECTIONTYPE
        case "auditMode":
            result = AUDITMODE_DEFENDERPROTECTIONTYPE
        case "warn":
            result = WARN_DEFENDERPROTECTIONTYPE
        case "notConfigured":
            result = NOTCONFIGURED_DEFENDERPROTECTIONTYPE
        default:
            return 0, errors.New("Unknown DefenderProtectionType value: " + v)
    }
    return &result, nil
}
func SerializeDefenderProtectionType(values []DefenderProtectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type SecurityBaselineComplianceState int

const (
    // Unknown state
    UNKNOWN_SECURITYBASELINECOMPLIANCESTATE SecurityBaselineComplianceState = iota
    // Secure state
    SECURE_SECURITYBASELINECOMPLIANCESTATE
    // Not applicable state
    NOTAPPLICABLE_SECURITYBASELINECOMPLIANCESTATE
    // Not secure state
    NOTSECURE_SECURITYBASELINECOMPLIANCESTATE
    // Error state
    ERROR_SECURITYBASELINECOMPLIANCESTATE
    // Conflict state
    CONFLICT_SECURITYBASELINECOMPLIANCESTATE
)

func (i SecurityBaselineComplianceState) String() string {
    return []string{"unknown", "secure", "notApplicable", "notSecure", "error", "conflict"}[i]
}
func ParseSecurityBaselineComplianceState(v string) (interface{}, error) {
    result := UNKNOWN_SECURITYBASELINECOMPLIANCESTATE
    switch v {
        case "unknown":
            result = UNKNOWN_SECURITYBASELINECOMPLIANCESTATE
        case "secure":
            result = SECURE_SECURITYBASELINECOMPLIANCESTATE
        case "notApplicable":
            result = NOTAPPLICABLE_SECURITYBASELINECOMPLIANCESTATE
        case "notSecure":
            result = NOTSECURE_SECURITYBASELINECOMPLIANCESTATE
        case "error":
            result = ERROR_SECURITYBASELINECOMPLIANCESTATE
        case "conflict":
            result = CONFLICT_SECURITYBASELINECOMPLIANCESTATE
        default:
            return 0, errors.New("Unknown SecurityBaselineComplianceState value: " + v)
    }
    return &result, nil
}
func SerializeSecurityBaselineComplianceState(values []SecurityBaselineComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

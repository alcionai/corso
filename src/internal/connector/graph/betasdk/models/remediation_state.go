package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RemediationState int

const (
    // Unknown result.
    UNKNOWN_REMEDIATIONSTATE RemediationState = iota
    // Remediation script execution was skipped
    SKIPPED_REMEDIATIONSTATE
    // Remediation script executed successfully and remediated the device state
    SUCCESS_REMEDIATIONSTATE
    // Remediation script executed successfully but failed to remediated the device state
    REMEDIATIONFAILED_REMEDIATIONSTATE
    // Remediation script execution encountered and error or timed out
    SCRIPTERROR_REMEDIATIONSTATE
)

func (i RemediationState) String() string {
    return []string{"unknown", "skipped", "success", "remediationFailed", "scriptError"}[i]
}
func ParseRemediationState(v string) (interface{}, error) {
    result := UNKNOWN_REMEDIATIONSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_REMEDIATIONSTATE
        case "skipped":
            result = SKIPPED_REMEDIATIONSTATE
        case "success":
            result = SUCCESS_REMEDIATIONSTATE
        case "remediationFailed":
            result = REMEDIATIONFAILED_REMEDIATIONSTATE
        case "scriptError":
            result = SCRIPTERROR_REMEDIATIONSTATE
        default:
            return 0, errors.New("Unknown RemediationState value: " + v)
    }
    return &result, nil
}
func SerializeRemediationState(values []RemediationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

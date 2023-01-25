package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsDefenderApplicationControlSupplementalPolicyStatuses int

const (
    // The status of the WindowsDefenderApplicationControl supplemental policy is not known.
    UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES WindowsDefenderApplicationControlSupplementalPolicyStatuses = iota
    // The WindowsDefenderApplicationControl supplemental policy is in effect.
    SUCCESS_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    // The WindowsDefenderApplicationControl supplemental policy is structurally okay but there is an error with authorizing the token.
    TOKENERROR_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    // The token does not authorize this WindowsDefenderApplicationControl supplemental policy.
    NOTAUTHORIZEDBYTOKEN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    // The WindowsDefenderApplicationControl supplemental policy is not found.
    POLICYNOTFOUND_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
)

func (i WindowsDefenderApplicationControlSupplementalPolicyStatuses) String() string {
    return []string{"unknown", "success", "tokenError", "notAuthorizedByToken", "policyNotFound"}[i]
}
func ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "success":
            result = SUCCESS_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "tokenError":
            result = TOKENERROR_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "notAuthorizedByToken":
            result = NOTAUTHORIZEDBYTOKEN_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        case "policyNotFound":
            result = POLICYNOTFOUND_WINDOWSDEFENDERAPPLICATIONCONTROLSUPPLEMENTALPOLICYSTATUSES
        default:
            return 0, errors.New("Unknown WindowsDefenderApplicationControlSupplementalPolicyStatuses value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDefenderApplicationControlSupplementalPolicyStatuses(values []WindowsDefenderApplicationControlSupplementalPolicyStatuses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

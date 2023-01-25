package models
import (
    "errors"
)
// Provides operations to call the add method.
type PolicySetStatus int

const (
    // Default Value.
    UNKNOWN_POLICYSETSTATUS PolicySetStatus = iota
    // All PolicySet items are now validating for corresponding settings of workloads.
    VALIDATING_POLICYSETSTATUS
    // Post process complete for all PolicySet items but there are failures.
    PARTIALSUCCESS_POLICYSETSTATUS
    // All PolicySet items are deployed. Doesn’t mean that all deployment succeeded. 
    SUCCESS_POLICYSETSTATUS
    // PolicySet processing completely failed.
    ERROR_POLICYSETSTATUS
    // PolicySet/PolicySetItem is not assigned to any group.
    NOTASSIGNED_POLICYSETSTATUS
)

func (i PolicySetStatus) String() string {
    return []string{"unknown", "validating", "partialSuccess", "success", "error", "notAssigned"}[i]
}
func ParsePolicySetStatus(v string) (interface{}, error) {
    result := UNKNOWN_POLICYSETSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_POLICYSETSTATUS
        case "validating":
            result = VALIDATING_POLICYSETSTATUS
        case "partialSuccess":
            result = PARTIALSUCCESS_POLICYSETSTATUS
        case "success":
            result = SUCCESS_POLICYSETSTATUS
        case "error":
            result = ERROR_POLICYSETSTATUS
        case "notAssigned":
            result = NOTASSIGNED_POLICYSETSTATUS
        default:
            return 0, errors.New("Unknown PolicySetStatus value: " + v)
    }
    return &result, nil
}
func SerializePolicySetStatus(values []PolicySetStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

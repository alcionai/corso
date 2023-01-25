package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type RunState int

const (
    // Unknown result.
    UNKNOWN_RUNSTATE RunState = iota
    // Script is run successfully.
    SUCCESS_RUNSTATE
    // Script failed to run.
    FAIL_RUNSTATE
    // Discovery script hits error.
    SCRIPTERROR_RUNSTATE
    // Script is pending to execute.
    PENDING_RUNSTATE
    // Script is not applicable for this device.
    NOTAPPLICABLE_RUNSTATE
)

func (i RunState) String() string {
    return []string{"unknown", "success", "fail", "scriptError", "pending", "notApplicable"}[i]
}
func ParseRunState(v string) (interface{}, error) {
    result := UNKNOWN_RUNSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_RUNSTATE
        case "success":
            result = SUCCESS_RUNSTATE
        case "fail":
            result = FAIL_RUNSTATE
        case "scriptError":
            result = SCRIPTERROR_RUNSTATE
        case "pending":
            result = PENDING_RUNSTATE
        case "notApplicable":
            result = NOTAPPLICABLE_RUNSTATE
        default:
            return 0, errors.New("Unknown RunState value: " + v)
    }
    return &result, nil
}
func SerializeRunState(values []RunState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

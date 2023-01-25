package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsAutopilotDeploymentState int

const (
    UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE WindowsAutopilotDeploymentState = iota
    SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
)

func (i WindowsAutopilotDeploymentState) String() string {
    return []string{"unknown", "success", "inProgress", "failure", "successWithTimeout", "notAttempted", "disabled"}[i]
}
func ParseWindowsAutopilotDeploymentState(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "success":
            result = SUCCESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "inProgress":
            result = INPROGRESS_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "failure":
            result = FAILURE_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "successWithTimeout":
            result = SUCCESSWITHTIMEOUT_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "notAttempted":
            result = NOTATTEMPTED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        case "disabled":
            result = DISABLED_WINDOWSAUTOPILOTDEPLOYMENTSTATE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeploymentState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeploymentState(values []WindowsAutopilotDeploymentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsAutopilotDeviceRemediationState int

const (
    // Unknown status.
    UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE WindowsAutopilotDeviceRemediationState = iota
    // No hardware change has been detected.
    NOREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    // Hardware change detected on client. Additional remediation is required.
    AUTOMATICREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    // Hardware change detected on client that could not resolved automatically. Additional remediation is required.
    MANUALREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    // Marks the end of known enum values, and allows for additional values in the future.
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
)

func (i WindowsAutopilotDeviceRemediationState) String() string {
    return []string{"unknown", "noRemediationRequired", "automaticRemediationRequired", "manualRemediationRequired", "unknownFutureValue"}[i]
}
func ParseWindowsAutopilotDeviceRemediationState(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "noRemediationRequired":
            result = NOREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "automaticRemediationRequired":
            result = AUTOMATICREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "manualRemediationRequired":
            result = MANUALREMEDIATIONREQUIRED_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICEREMEDIATIONSTATE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeviceRemediationState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeviceRemediationState(values []WindowsAutopilotDeviceRemediationState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

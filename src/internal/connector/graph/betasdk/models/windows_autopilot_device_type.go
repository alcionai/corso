package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsAutopilotDeviceType int

const (
    // Windows PC
    WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE WindowsAutopilotDeviceType = iota
    // Surface Hub 2
    SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
    // HoloLens
    HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
    // SurfaceHub2S
    SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
    // VirtualMachine
    VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
    // Placeholder for evolvable enum, but this enum is never returned to the caller, so it shouldn't be necessary.         
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
)

func (i WindowsAutopilotDeviceType) String() string {
    return []string{"windowsPc", "surfaceHub2", "holoLens", "surfaceHub2S", "virtualMachine", "unknownFutureValue"}[i]
}
func ParseWindowsAutopilotDeviceType(v string) (interface{}, error) {
    result := WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
    switch v {
        case "windowsPc":
            result = WINDOWSPC_WINDOWSAUTOPILOTDEVICETYPE
        case "surfaceHub2":
            result = SURFACEHUB2_WINDOWSAUTOPILOTDEVICETYPE
        case "holoLens":
            result = HOLOLENS_WINDOWSAUTOPILOTDEVICETYPE
        case "surfaceHub2S":
            result = SURFACEHUB2S_WINDOWSAUTOPILOTDEVICETYPE
        case "virtualMachine":
            result = VIRTUALMACHINE_WINDOWSAUTOPILOTDEVICETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTDEVICETYPE
        default:
            return 0, errors.New("Unknown WindowsAutopilotDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotDeviceType(values []WindowsAutopilotDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

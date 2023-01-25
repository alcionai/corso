package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsAutopilotProfileAssignmentDetailedStatus int

const (
    // No assignment detailed status
    NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS WindowsAutopilotProfileAssignmentDetailedStatus = iota
    // Hardware requirements are not met. This can happen if a self-deploying AutoPilot Profile is assigned to a device without TPM 2.0.
    HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    // Indicates that a Surface Hub AutoPilot Profile is assigned to a device that is not Surface Hub(Aruba).
    SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    // Indicates that a HoloLens AutoPilot Profile is assigned to a device that is not HoloLens.
    HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    // Indicates that a Windows PC AutoPilot Profile is assigned to a device that is not Windows PC.
    WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    // Indicates that a surface Hub 2S  AutoPilot Profile is assigned to a device that is not surface Hub 2S.
    SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    // Placeholder for evolvable enum, but this enum is never returned to the caller, so it shouldn't be necessary.
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
)

func (i WindowsAutopilotProfileAssignmentDetailedStatus) String() string {
    return []string{"none", "hardwareRequirementsNotMet", "surfaceHubProfileNotSupported", "holoLensProfileNotSupported", "windowsPcProfileNotSupported", "surfaceHub2SProfileNotSupported", "unknownFutureValue"}[i]
}
func ParseWindowsAutopilotProfileAssignmentDetailedStatus(v string) (interface{}, error) {
    result := NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    switch v {
        case "none":
            result = NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "hardwareRequirementsNotMet":
            result = HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "surfaceHubProfileNotSupported":
            result = SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "holoLensProfileNotSupported":
            result = HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "windowsPcProfileNotSupported":
            result = WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "surfaceHub2SProfileNotSupported":
            result = SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentDetailedStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotProfileAssignmentDetailedStatus(values []WindowsAutopilotProfileAssignmentDetailedStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

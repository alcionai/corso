package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsAutopilotProfileAssignmentStatus int

const (
    // Unknown assignment status
    UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS WindowsAutopilotProfileAssignmentStatus = iota
    // Assigned successfully in Intune and in sync with Windows auto pilot program
    ASSIGNEDINSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    // Assigned successfully in Intune and not in sync with Windows auto pilot program
    ASSIGNEDOUTOFSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    // Assigned successfully in Intune and either in-sync or out of sync with Windows auto pilot program
    ASSIGNEDUNKOWNSYNCSTATE_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    // Not assigned
    NOTASSIGNED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    // Pending assignment
    PENDING_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    //  Assignment failed
    FAILED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
)

func (i WindowsAutopilotProfileAssignmentStatus) String() string {
    return []string{"unknown", "assignedInSync", "assignedOutOfSync", "assignedUnkownSyncState", "notAssigned", "pending", "failed"}[i]
}
func ParseWindowsAutopilotProfileAssignmentStatus(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "assignedInSync":
            result = ASSIGNEDINSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "assignedOutOfSync":
            result = ASSIGNEDOUTOFSYNC_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "assignedUnkownSyncState":
            result = ASSIGNEDUNKOWNSYNCSTATE_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "notAssigned":
            result = NOTASSIGNED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "pending":
            result = PENDING_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        case "failed":
            result = FAILED_WINDOWSAUTOPILOTPROFILEASSIGNMENTSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotProfileAssignmentStatus(values []WindowsAutopilotProfileAssignmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type WindowsAutopilotSyncStatus int

const (
    // Unknown sync status
    UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS WindowsAutopilotSyncStatus = iota
    // Sync is in progress
    INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
    // Sync completed.
    COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
    // Sync failed.
    FAILED_WINDOWSAUTOPILOTSYNCSTATUS
)

func (i WindowsAutopilotSyncStatus) String() string {
    return []string{"unknown", "inProgress", "completed", "failed"}[i]
}
func ParseWindowsAutopilotSyncStatus(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTSYNCSTATUS
        case "inProgress":
            result = INPROGRESS_WINDOWSAUTOPILOTSYNCSTATUS
        case "completed":
            result = COMPLETED_WINDOWSAUTOPILOTSYNCSTATUS
        case "failed":
            result = FAILED_WINDOWSAUTOPILOTSYNCSTATUS
        default:
            return 0, errors.New("Unknown WindowsAutopilotSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotSyncStatus(values []WindowsAutopilotSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsUpdateStatus int

const (
    // There are no pending updates, no pending reboot updates and no failed updates.
    UPTODATE_WINDOWSUPDATESTATUS WindowsUpdateStatus = iota
    // There are updates that’s pending installation which includes updates that are not approved. There are no Pending reboot updates, no failed updates.
    PENDINGINSTALLATION_WINDOWSUPDATESTATUS
    // There are updates that requires reboot. There are not failed updates.
    PENDINGREBOOT_WINDOWSUPDATESTATUS
    // There are updates failed to install on the device.
    FAILED_WINDOWSUPDATESTATUS
)

func (i WindowsUpdateStatus) String() string {
    return []string{"upToDate", "pendingInstallation", "pendingReboot", "failed"}[i]
}
func ParseWindowsUpdateStatus(v string) (interface{}, error) {
    result := UPTODATE_WINDOWSUPDATESTATUS
    switch v {
        case "upToDate":
            result = UPTODATE_WINDOWSUPDATESTATUS
        case "pendingInstallation":
            result = PENDINGINSTALLATION_WINDOWSUPDATESTATUS
        case "pendingReboot":
            result = PENDINGREBOOT_WINDOWSUPDATESTATUS
        case "failed":
            result = FAILED_WINDOWSUPDATESTATUS
        default:
            return 0, errors.New("Unknown WindowsUpdateStatus value: " + v)
    }
    return &result, nil
}
func SerializeWindowsUpdateStatus(values []WindowsUpdateStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

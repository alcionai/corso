package models
import (
    "errors"
)
// Provides operations to call the add method.
type WinGetAppNotification int

const (
    // Show all notifications.
    SHOWALL_WINGETAPPNOTIFICATION WinGetAppNotification = iota
    // Only show restart notification and suppress other notifications.
    SHOWREBOOT_WINGETAPPNOTIFICATION
    // Hide all notifications.
    HIDEALL_WINGETAPPNOTIFICATION
    // Unknown future value, reserved for future usage as expandable enum.
    UNKNOWNFUTUREVALUE_WINGETAPPNOTIFICATION
)

func (i WinGetAppNotification) String() string {
    return []string{"showAll", "showReboot", "hideAll", "unknownFutureValue"}[i]
}
func ParseWinGetAppNotification(v string) (interface{}, error) {
    result := SHOWALL_WINGETAPPNOTIFICATION
    switch v {
        case "showAll":
            result = SHOWALL_WINGETAPPNOTIFICATION
        case "showReboot":
            result = SHOWREBOOT_WINGETAPPNOTIFICATION
        case "hideAll":
            result = HIDEALL_WINGETAPPNOTIFICATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINGETAPPNOTIFICATION
        default:
            return 0, errors.New("Unknown WinGetAppNotification value: " + v)
    }
    return &result, nil
}
func SerializeWinGetAppNotification(values []WinGetAppNotification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

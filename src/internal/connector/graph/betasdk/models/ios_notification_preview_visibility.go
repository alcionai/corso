package models
import (
    "errors"
)
// Provides operations to call the add method.
type IosNotificationPreviewVisibility int

const (
    // Notification preview settings will not be overwritten.
    NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY IosNotificationPreviewVisibility = iota
    // Always show notification previews.
    ALWAYSSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
    // Only show notification previews when the device is unlocked.
    HIDEWHENLOCKED_IOSNOTIFICATIONPREVIEWVISIBILITY
    // Never show notification previews.
    NEVERSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
)

func (i IosNotificationPreviewVisibility) String() string {
    return []string{"notConfigured", "alwaysShow", "hideWhenLocked", "neverShow"}[i]
}
func ParseIosNotificationPreviewVisibility(v string) (interface{}, error) {
    result := NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "alwaysShow":
            result = ALWAYSSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "hideWhenLocked":
            result = HIDEWHENLOCKED_IOSNOTIFICATIONPREVIEWVISIBILITY
        case "neverShow":
            result = NEVERSHOW_IOSNOTIFICATIONPREVIEWVISIBILITY
        default:
            return 0, errors.New("Unknown IosNotificationPreviewVisibility value: " + v)
    }
    return &result, nil
}
func SerializeIosNotificationPreviewVisibility(values []IosNotificationPreviewVisibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package devicemanagement
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type NotificationChannelType int

const (
    PORTAL_NOTIFICATIONCHANNELTYPE NotificationChannelType = iota
    EMAIL_NOTIFICATIONCHANNELTYPE
    PHONECALL_NOTIFICATIONCHANNELTYPE
    SMS_NOTIFICATIONCHANNELTYPE
    UNKNOWNFUTUREVALUE_NOTIFICATIONCHANNELTYPE
)

func (i NotificationChannelType) String() string {
    return []string{"portal", "email", "phoneCall", "sms", "unknownFutureValue"}[i]
}
func ParseNotificationChannelType(v string) (interface{}, error) {
    result := PORTAL_NOTIFICATIONCHANNELTYPE
    switch v {
        case "portal":
            result = PORTAL_NOTIFICATIONCHANNELTYPE
        case "email":
            result = EMAIL_NOTIFICATIONCHANNELTYPE
        case "phoneCall":
            result = PHONECALL_NOTIFICATIONCHANNELTYPE
        case "sms":
            result = SMS_NOTIFICATIONCHANNELTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NOTIFICATIONCHANNELTYPE
        default:
            return 0, errors.New("Unknown NotificationChannelType value: " + v)
    }
    return &result, nil
}
func SerializeNotificationChannelType(values []NotificationChannelType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

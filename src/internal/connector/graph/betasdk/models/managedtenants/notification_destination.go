package managedtenants
import (
    "errors"
)
// Provides operations to call the add method.
type NotificationDestination int

const (
    NONE_NOTIFICATIONDESTINATION NotificationDestination = iota
    API_NOTIFICATIONDESTINATION
    EMAIL_NOTIFICATIONDESTINATION
    SMS_NOTIFICATIONDESTINATION
    UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION
)

func (i NotificationDestination) String() string {
    return []string{"none", "api", "email", "sms", "unknownFutureValue"}[i]
}
func ParseNotificationDestination(v string) (interface{}, error) {
    result := NONE_NOTIFICATIONDESTINATION
    switch v {
        case "none":
            result = NONE_NOTIFICATIONDESTINATION
        case "api":
            result = API_NOTIFICATIONDESTINATION
        case "email":
            result = EMAIL_NOTIFICATIONDESTINATION
        case "sms":
            result = SMS_NOTIFICATIONDESTINATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NOTIFICATIONDESTINATION
        default:
            return 0, errors.New("Unknown NotificationDestination value: " + v)
    }
    return &result, nil
}
func SerializeNotificationDestination(values []NotificationDestination) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

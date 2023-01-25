package models
import (
    "errors"
)
// Provides operations to call the add method.
type VppTokenActionFailureReason int

const (
    // None.
    NONE_VPPTOKENACTIONFAILUREREASON VppTokenActionFailureReason = iota
    // There was an error on Apple's service.
    APPLEFAILURE_VPPTOKENACTIONFAILUREREASON
    // There was an internal error.
    INTERNALERROR_VPPTOKENACTIONFAILUREREASON
    // There was an error because the Apple Volume Purchase Program token was expired.
    EXPIREDVPPTOKEN_VPPTOKENACTIONFAILUREREASON
    // There was an error because the Apple Volume Purchase Program Push Notification certificate expired.
    EXPIREDAPPLEPUSHNOTIFICATIONCERTIFICATE_VPPTOKENACTIONFAILUREREASON
)

func (i VppTokenActionFailureReason) String() string {
    return []string{"none", "appleFailure", "internalError", "expiredVppToken", "expiredApplePushNotificationCertificate"}[i]
}
func ParseVppTokenActionFailureReason(v string) (interface{}, error) {
    result := NONE_VPPTOKENACTIONFAILUREREASON
    switch v {
        case "none":
            result = NONE_VPPTOKENACTIONFAILUREREASON
        case "appleFailure":
            result = APPLEFAILURE_VPPTOKENACTIONFAILUREREASON
        case "internalError":
            result = INTERNALERROR_VPPTOKENACTIONFAILUREREASON
        case "expiredVppToken":
            result = EXPIREDVPPTOKEN_VPPTOKENACTIONFAILUREREASON
        case "expiredApplePushNotificationCertificate":
            result = EXPIREDAPPLEPUSHNOTIFICATIONCERTIFICATE_VPPTOKENACTIONFAILUREREASON
        default:
            return 0, errors.New("Unknown VppTokenActionFailureReason value: " + v)
    }
    return &result, nil
}
func SerializeVppTokenActionFailureReason(values []VppTokenActionFailureReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

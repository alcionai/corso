package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type EnrollmentNotificationTemplateType int

const (
    // Email Notification
    EMAIL_ENROLLMENTNOTIFICATIONTEMPLATETYPE EnrollmentNotificationTemplateType = iota
    // Push Notification
    PUSH_ENROLLMENTNOTIFICATIONTEMPLATETYPE
    // Unknown Type
    UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONTEMPLATETYPE
)

func (i EnrollmentNotificationTemplateType) String() string {
    return []string{"email", "push", "unknownFutureValue"}[i]
}
func ParseEnrollmentNotificationTemplateType(v string) (interface{}, error) {
    result := EMAIL_ENROLLMENTNOTIFICATIONTEMPLATETYPE
    switch v {
        case "email":
            result = EMAIL_ENROLLMENTNOTIFICATIONTEMPLATETYPE
        case "push":
            result = PUSH_ENROLLMENTNOTIFICATIONTEMPLATETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONTEMPLATETYPE
        default:
            return 0, errors.New("Unknown EnrollmentNotificationTemplateType value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentNotificationTemplateType(values []EnrollmentNotificationTemplateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

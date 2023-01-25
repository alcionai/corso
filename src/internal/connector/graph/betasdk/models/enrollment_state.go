package models
import (
    "errors"
)
// Provides operations to call the add method.
type EnrollmentState int

const (
    // Device enrollment state is unknown
    UNKNOWN_ENROLLMENTSTATE EnrollmentState = iota
    // Device is Enrolled.
    ENROLLED_ENROLLMENTSTATE
    // Enrolled but it's enrolled via enrollment profile and the enrolled profile is different from the assigned profile.
    PENDINGRESET_ENROLLMENTSTATE
    // Not enrolled and there is enrollment failure record.
    FAILED_ENROLLMENTSTATE
    // Device is imported but not enrolled.
    NOTCONTACTED_ENROLLMENTSTATE
    // Device is enrolled as userless, but is blocked from moving to user enrollment because the app failed to install.
    BLOCKED_ENROLLMENTSTATE
)

func (i EnrollmentState) String() string {
    return []string{"unknown", "enrolled", "pendingReset", "failed", "notContacted", "blocked"}[i]
}
func ParseEnrollmentState(v string) (interface{}, error) {
    result := UNKNOWN_ENROLLMENTSTATE
    switch v {
        case "unknown":
            result = UNKNOWN_ENROLLMENTSTATE
        case "enrolled":
            result = ENROLLED_ENROLLMENTSTATE
        case "pendingReset":
            result = PENDINGRESET_ENROLLMENTSTATE
        case "failed":
            result = FAILED_ENROLLMENTSTATE
        case "notContacted":
            result = NOTCONTACTED_ENROLLMENTSTATE
        case "blocked":
            result = BLOCKED_ENROLLMENTSTATE
        default:
            return 0, errors.New("Unknown EnrollmentState value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentState(values []EnrollmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

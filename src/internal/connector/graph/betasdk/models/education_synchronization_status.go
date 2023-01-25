package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type EducationSynchronizationStatus int

const (
    PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS EducationSynchronizationStatus = iota
    INPROGRESS_EDUCATIONSYNCHRONIZATIONSTATUS
    SUCCESS_EDUCATIONSYNCHRONIZATIONSTATUS
    ERROR_EDUCATIONSYNCHRONIZATIONSTATUS
    VALIDATIONERROR_EDUCATIONSYNCHRONIZATIONSTATUS
    QUARANTINED_EDUCATIONSYNCHRONIZATIONSTATUS
    UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONSTATUS
    EXTRACTING_EDUCATIONSYNCHRONIZATIONSTATUS
    VALIDATING_EDUCATIONSYNCHRONIZATIONSTATUS
)

func (i EducationSynchronizationStatus) String() string {
    return []string{"paused", "inProgress", "success", "error", "validationError", "quarantined", "unknownFutureValue", "extracting", "validating"}[i]
}
func ParseEducationSynchronizationStatus(v string) (interface{}, error) {
    result := PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS
    switch v {
        case "paused":
            result = PAUSED_EDUCATIONSYNCHRONIZATIONSTATUS
        case "inProgress":
            result = INPROGRESS_EDUCATIONSYNCHRONIZATIONSTATUS
        case "success":
            result = SUCCESS_EDUCATIONSYNCHRONIZATIONSTATUS
        case "error":
            result = ERROR_EDUCATIONSYNCHRONIZATIONSTATUS
        case "validationError":
            result = VALIDATIONERROR_EDUCATIONSYNCHRONIZATIONSTATUS
        case "quarantined":
            result = QUARANTINED_EDUCATIONSYNCHRONIZATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EDUCATIONSYNCHRONIZATIONSTATUS
        case "extracting":
            result = EXTRACTING_EDUCATIONSYNCHRONIZATIONSTATUS
        case "validating":
            result = VALIDATING_EDUCATIONSYNCHRONIZATIONSTATUS
        default:
            return 0, errors.New("Unknown EducationSynchronizationStatus value: " + v)
    }
    return &result, nil
}
func SerializeEducationSynchronizationStatus(values []EducationSynchronizationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

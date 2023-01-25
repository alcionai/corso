package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type LongRunningOperationStatus int

const (
    NOTSTARTED_LONGRUNNINGOPERATIONSTATUS LongRunningOperationStatus = iota
    RUNNING_LONGRUNNINGOPERATIONSTATUS
    SUCCEEDED_LONGRUNNINGOPERATIONSTATUS
    FAILED_LONGRUNNINGOPERATIONSTATUS
    SKIPPED_LONGRUNNINGOPERATIONSTATUS
    UNKNOWNFUTUREVALUE_LONGRUNNINGOPERATIONSTATUS
)

func (i LongRunningOperationStatus) String() string {
    return []string{"notStarted", "running", "succeeded", "failed", "skipped", "unknownFutureValue"}[i]
}
func ParseLongRunningOperationStatus(v string) (interface{}, error) {
    result := NOTSTARTED_LONGRUNNINGOPERATIONSTATUS
    switch v {
        case "notStarted":
            result = NOTSTARTED_LONGRUNNINGOPERATIONSTATUS
        case "running":
            result = RUNNING_LONGRUNNINGOPERATIONSTATUS
        case "succeeded":
            result = SUCCEEDED_LONGRUNNINGOPERATIONSTATUS
        case "failed":
            result = FAILED_LONGRUNNINGOPERATIONSTATUS
        case "skipped":
            result = SKIPPED_LONGRUNNINGOPERATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LONGRUNNINGOPERATIONSTATUS
        default:
            return 0, errors.New("Unknown LongRunningOperationStatus value: " + v)
    }
    return &result, nil
}
func SerializeLongRunningOperationStatus(values []LongRunningOperationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

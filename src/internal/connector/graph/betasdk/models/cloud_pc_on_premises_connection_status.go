package models
import (
    "errors"
)
// Provides operations to call the add method.
type CloudPcOnPremisesConnectionStatus int

const (
    PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS CloudPcOnPremisesConnectionStatus = iota
    RUNNING_CLOUDPCONPREMISESCONNECTIONSTATUS
    PASSED_CLOUDPCONPREMISESCONNECTIONSTATUS
    FAILED_CLOUDPCONPREMISESCONNECTIONSTATUS
    WARNING_CLOUDPCONPREMISESCONNECTIONSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONSTATUS
)

func (i CloudPcOnPremisesConnectionStatus) String() string {
    return []string{"pending", "running", "passed", "failed", "warning", "unknownFutureValue"}[i]
}
func ParseCloudPcOnPremisesConnectionStatus(v string) (interface{}, error) {
    result := PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS
    switch v {
        case "pending":
            result = PENDING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "running":
            result = RUNNING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "passed":
            result = PASSED_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "failed":
            result = FAILED_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "warning":
            result = WARNING_CLOUDPCONPREMISESCONNECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCONPREMISESCONNECTIONSTATUS
        default:
            return 0, errors.New("Unknown CloudPcOnPremisesConnectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOnPremisesConnectionStatus(values []CloudPcOnPremisesConnectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

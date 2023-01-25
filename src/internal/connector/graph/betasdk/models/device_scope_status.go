package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceScopeStatus int

const (
    // Indicates the device scope is not enabled and there are no calculations in progress.
    NONE_DEVICESCOPESTATUS DeviceScopeStatus = iota
    // Indicates the device scope is enabled and report metrics data are being recalculated by the service.
    COMPUTING_DEVICESCOPESTATUS
    // Indicates the device scope is enabled but there is insufficient data to calculate results. The system requires information from at least 5 devices before calculations can occur.
    INSUFFICIENTDATA_DEVICESCOPESTATUS
    // Device scope is enabled and finished recalculating the report metric. Device scope is now ready to be used.
    COMPLETED_DEVICESCOPESTATUS
    // Placeholder value for future expansion.
    UNKNOWNFUTUREVALUE_DEVICESCOPESTATUS
)

func (i DeviceScopeStatus) String() string {
    return []string{"none", "computing", "insufficientData", "completed", "unknownFutureValue"}[i]
}
func ParseDeviceScopeStatus(v string) (interface{}, error) {
    result := NONE_DEVICESCOPESTATUS
    switch v {
        case "none":
            result = NONE_DEVICESCOPESTATUS
        case "computing":
            result = COMPUTING_DEVICESCOPESTATUS
        case "insufficientData":
            result = INSUFFICIENTDATA_DEVICESCOPESTATUS
        case "completed":
            result = COMPLETED_DEVICESCOPESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICESCOPESTATUS
        default:
            return 0, errors.New("Unknown DeviceScopeStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceScopeStatus(values []DeviceScopeStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

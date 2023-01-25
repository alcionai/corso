package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcDeviceImageOsStatus int

const (
    SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS CloudPcDeviceImageOsStatus = iota
    SUPPORTEDWITHWARNING_CLOUDPCDEVICEIMAGEOSSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGEOSSTATUS
)

func (i CloudPcDeviceImageOsStatus) String() string {
    return []string{"supported", "supportedWithWarning", "unknownFutureValue"}[i]
}
func ParseCloudPcDeviceImageOsStatus(v string) (interface{}, error) {
    result := SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS
    switch v {
        case "supported":
            result = SUPPORTED_CLOUDPCDEVICEIMAGEOSSTATUS
        case "supportedWithWarning":
            result = SUPPORTEDWITHWARNING_CLOUDPCDEVICEIMAGEOSSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCDEVICEIMAGEOSSTATUS
        default:
            return 0, errors.New("Unknown CloudPcDeviceImageOsStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcDeviceImageOsStatus(values []CloudPcDeviceImageOsStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

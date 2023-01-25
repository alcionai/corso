package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcOperatingSystem int

const (
    WINDOWS10_CLOUDPCOPERATINGSYSTEM CloudPcOperatingSystem = iota
    WINDOWS11_CLOUDPCOPERATINGSYSTEM
    UNKNOWNFUTUREVALUE_CLOUDPCOPERATINGSYSTEM
)

func (i CloudPcOperatingSystem) String() string {
    return []string{"windows10", "windows11", "unknownFutureValue"}[i]
}
func ParseCloudPcOperatingSystem(v string) (interface{}, error) {
    result := WINDOWS10_CLOUDPCOPERATINGSYSTEM
    switch v {
        case "windows10":
            result = WINDOWS10_CLOUDPCOPERATINGSYSTEM
        case "windows11":
            result = WINDOWS11_CLOUDPCOPERATINGSYSTEM
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCOPERATINGSYSTEM
        default:
            return 0, errors.New("Unknown CloudPcOperatingSystem value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcOperatingSystem(values []CloudPcOperatingSystem) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

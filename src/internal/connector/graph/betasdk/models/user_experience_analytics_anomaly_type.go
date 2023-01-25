package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UserExperienceAnalyticsAnomalyType int

const (
    // Indicates the detected anomaly is due to certain devices.
    DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE UserExperienceAnalyticsAnomalyType = iota
    // Indicates the detected anomaly is due to a specific application.
    APPLICATION_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the detected anomaly is due to a specific stop error.
    STOPERROR_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the detected anomaly is due to a specific driver.
    DRIVER_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Indicates the category of detected anomaly is undefined.
    OTHER_USEREXPERIENCEANALYTICSANOMALYTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYTYPE
)

func (i UserExperienceAnalyticsAnomalyType) String() string {
    return []string{"device", "application", "stopError", "driver", "other", "unknownFutureValue"}[i]
}
func ParseUserExperienceAnalyticsAnomalyType(v string) (interface{}, error) {
    result := DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE
    switch v {
        case "device":
            result = DEVICE_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "application":
            result = APPLICATION_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "stopError":
            result = STOPERROR_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "driver":
            result = DRIVER_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "other":
            result = OTHER_USEREXPERIENCEANALYTICSANOMALYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USEREXPERIENCEANALYTICSANOMALYTYPE
        default:
            return 0, errors.New("Unknown UserExperienceAnalyticsAnomalyType value: " + v)
    }
    return &result, nil
}
func SerializeUserExperienceAnalyticsAnomalyType(values []UserExperienceAnalyticsAnomalyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

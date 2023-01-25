package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcServicePlanType int

const (
    ENTERPRISE_CLOUDPCSERVICEPLANTYPE CloudPcServicePlanType = iota
    BUSINESS_CLOUDPCSERVICEPLANTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCSERVICEPLANTYPE
)

func (i CloudPcServicePlanType) String() string {
    return []string{"enterprise", "business", "unknownFutureValue"}[i]
}
func ParseCloudPcServicePlanType(v string) (interface{}, error) {
    result := ENTERPRISE_CLOUDPCSERVICEPLANTYPE
    switch v {
        case "enterprise":
            result = ENTERPRISE_CLOUDPCSERVICEPLANTYPE
        case "business":
            result = BUSINESS_CLOUDPCSERVICEPLANTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSERVICEPLANTYPE
        default:
            return 0, errors.New("Unknown CloudPcServicePlanType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcServicePlanType(values []CloudPcServicePlanType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

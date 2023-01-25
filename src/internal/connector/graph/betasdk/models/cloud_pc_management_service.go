package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcManagementService int

const (
    WINDOWS365_CLOUDPCMANAGEMENTSERVICE CloudPcManagementService = iota
    DEVBOX_CLOUDPCMANAGEMENTSERVICE
    UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
    RPABOX_CLOUDPCMANAGEMENTSERVICE
)

func (i CloudPcManagementService) String() string {
    return []string{"windows365", "devBox", "unknownFutureValue", "rpaBox"}[i]
}
func ParseCloudPcManagementService(v string) (interface{}, error) {
    result := WINDOWS365_CLOUDPCMANAGEMENTSERVICE
    switch v {
        case "windows365":
            result = WINDOWS365_CLOUDPCMANAGEMENTSERVICE
        case "devBox":
            result = DEVBOX_CLOUDPCMANAGEMENTSERVICE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCMANAGEMENTSERVICE
        case "rpaBox":
            result = RPABOX_CLOUDPCMANAGEMENTSERVICE
        default:
            return 0, errors.New("Unknown CloudPcManagementService value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcManagementService(values []CloudPcManagementService) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

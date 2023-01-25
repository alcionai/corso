package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcStatus int

const (
    NOTPROVISIONED_CLOUDPCSTATUS CloudPcStatus = iota
    PROVISIONING_CLOUDPCSTATUS
    PROVISIONED_CLOUDPCSTATUS
    INGRACEPERIOD_CLOUDPCSTATUS
    DEPROVISIONING_CLOUDPCSTATUS
    FAILED_CLOUDPCSTATUS
    PROVISIONEDWITHWARNINGS_CLOUDPCSTATUS
    RESIZING_CLOUDPCSTATUS
    RESTORING_CLOUDPCSTATUS
    PENDINGPROVISION_CLOUDPCSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCSTATUS
)

func (i CloudPcStatus) String() string {
    return []string{"notProvisioned", "provisioning", "provisioned", "inGracePeriod", "deprovisioning", "failed", "provisionedWithWarnings", "resizing", "restoring", "pendingProvision", "unknownFutureValue"}[i]
}
func ParseCloudPcStatus(v string) (interface{}, error) {
    result := NOTPROVISIONED_CLOUDPCSTATUS
    switch v {
        case "notProvisioned":
            result = NOTPROVISIONED_CLOUDPCSTATUS
        case "provisioning":
            result = PROVISIONING_CLOUDPCSTATUS
        case "provisioned":
            result = PROVISIONED_CLOUDPCSTATUS
        case "inGracePeriod":
            result = INGRACEPERIOD_CLOUDPCSTATUS
        case "deprovisioning":
            result = DEPROVISIONING_CLOUDPCSTATUS
        case "failed":
            result = FAILED_CLOUDPCSTATUS
        case "provisionedWithWarnings":
            result = PROVISIONEDWITHWARNINGS_CLOUDPCSTATUS
        case "resizing":
            result = RESIZING_CLOUDPCSTATUS
        case "restoring":
            result = RESTORING_CLOUDPCSTATUS
        case "pendingProvision":
            result = PENDINGPROVISION_CLOUDPCSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSTATUS
        default:
            return 0, errors.New("Unknown CloudPcStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcStatus(values []CloudPcStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

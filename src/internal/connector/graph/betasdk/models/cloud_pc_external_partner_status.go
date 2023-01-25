package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type CloudPcExternalPartnerStatus int

const (
    NOTAVAILABLE_CLOUDPCEXTERNALPARTNERSTATUS CloudPcExternalPartnerStatus = iota
    AVAILABLE_CLOUDPCEXTERNALPARTNERSTATUS
    HEALTHY_CLOUDPCEXTERNALPARTNERSTATUS
    UNHEALTHY_CLOUDPCEXTERNALPARTNERSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCEXTERNALPARTNERSTATUS
)

func (i CloudPcExternalPartnerStatus) String() string {
    return []string{"notAvailable", "available", "healthy", "unhealthy", "unknownFutureValue"}[i]
}
func ParseCloudPcExternalPartnerStatus(v string) (interface{}, error) {
    result := NOTAVAILABLE_CLOUDPCEXTERNALPARTNERSTATUS
    switch v {
        case "notAvailable":
            result = NOTAVAILABLE_CLOUDPCEXTERNALPARTNERSTATUS
        case "available":
            result = AVAILABLE_CLOUDPCEXTERNALPARTNERSTATUS
        case "healthy":
            result = HEALTHY_CLOUDPCEXTERNALPARTNERSTATUS
        case "unhealthy":
            result = UNHEALTHY_CLOUDPCEXTERNALPARTNERSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCEXTERNALPARTNERSTATUS
        default:
            return 0, errors.New("Unknown CloudPcExternalPartnerStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcExternalPartnerStatus(values []CloudPcExternalPartnerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

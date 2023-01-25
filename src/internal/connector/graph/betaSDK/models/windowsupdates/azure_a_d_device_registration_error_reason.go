package windowsupdates
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AzureADDeviceRegistrationErrorReason int

const (
    INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON AzureADDeviceRegistrationErrorReason = iota
    INVALIDAZUREADDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
    MISSINGTRUSTTYPE_AZUREADDEVICEREGISTRATIONERRORREASON
    INVALIDAZUREADJOIN_AZUREADDEVICEREGISTRATIONERRORREASON
    UNKNOWNFUTUREVALUE_AZUREADDEVICEREGISTRATIONERRORREASON
)

func (i AzureADDeviceRegistrationErrorReason) String() string {
    return []string{"invalidGlobalDeviceId", "invalidAzureADDeviceId", "missingTrustType", "invalidAzureADJoin", "unknownFutureValue"}[i]
}
func ParseAzureADDeviceRegistrationErrorReason(v string) (interface{}, error) {
    result := INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
    switch v {
        case "invalidGlobalDeviceId":
            result = INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
        case "invalidAzureADDeviceId":
            result = INVALIDAZUREADDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
        case "missingTrustType":
            result = MISSINGTRUSTTYPE_AZUREADDEVICEREGISTRATIONERRORREASON
        case "invalidAzureADJoin":
            result = INVALIDAZUREADJOIN_AZUREADDEVICEREGISTRATIONERRORREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREADDEVICEREGISTRATIONERRORREASON
        default:
            return 0, errors.New("Unknown AzureADDeviceRegistrationErrorReason value: " + v)
    }
    return &result, nil
}
func SerializeAzureADDeviceRegistrationErrorReason(values []AzureADDeviceRegistrationErrorReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

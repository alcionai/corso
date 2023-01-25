package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UsageAuthMethod int

const (
    EMAIL_USAGEAUTHMETHOD UsageAuthMethod = iota
    MOBILESMS_USAGEAUTHMETHOD
    MOBILECALL_USAGEAUTHMETHOD
    OFFICEPHONE_USAGEAUTHMETHOD
    SECURITYQUESTION_USAGEAUTHMETHOD
    APPNOTIFICATION_USAGEAUTHMETHOD
    APPCODE_USAGEAUTHMETHOD
    ALTERNATEMOBILECALL_USAGEAUTHMETHOD
    FIDO_USAGEAUTHMETHOD
    APPPASSWORD_USAGEAUTHMETHOD
    UNKNOWNFUTUREVALUE_USAGEAUTHMETHOD
)

func (i UsageAuthMethod) String() string {
    return []string{"email", "mobileSMS", "mobileCall", "officePhone", "securityQuestion", "appNotification", "appCode", "alternateMobileCall", "fido", "appPassword", "unknownFutureValue"}[i]
}
func ParseUsageAuthMethod(v string) (interface{}, error) {
    result := EMAIL_USAGEAUTHMETHOD
    switch v {
        case "email":
            result = EMAIL_USAGEAUTHMETHOD
        case "mobileSMS":
            result = MOBILESMS_USAGEAUTHMETHOD
        case "mobileCall":
            result = MOBILECALL_USAGEAUTHMETHOD
        case "officePhone":
            result = OFFICEPHONE_USAGEAUTHMETHOD
        case "securityQuestion":
            result = SECURITYQUESTION_USAGEAUTHMETHOD
        case "appNotification":
            result = APPNOTIFICATION_USAGEAUTHMETHOD
        case "appCode":
            result = APPCODE_USAGEAUTHMETHOD
        case "alternateMobileCall":
            result = ALTERNATEMOBILECALL_USAGEAUTHMETHOD
        case "fido":
            result = FIDO_USAGEAUTHMETHOD
        case "appPassword":
            result = APPPASSWORD_USAGEAUTHMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USAGEAUTHMETHOD
        default:
            return 0, errors.New("Unknown UsageAuthMethod value: " + v)
    }
    return &result, nil
}
func SerializeUsageAuthMethod(values []UsageAuthMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

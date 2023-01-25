package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EnrollmentNotificationBrandingOptions int

const (
    // Indicates that the template has no branding.
    NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS EnrollmentNotificationBrandingOptions = iota
    // Indicates that the Company Logo is included in the notification.
    INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Indicates that the Company Name is included in the notification.
    INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Indicates that the Contact Information is included in the notification.
    INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Indicates that the Company Portal Link is included in the notification.
    INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Indicates that the DeviceDetails is included in the notification.
    INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // unknownFutureValue for evolvable enums pattern.
    UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
)

func (i EnrollmentNotificationBrandingOptions) String() string {
    return []string{"none", "includeCompanyLogo", "includeCompanyName", "includeContactInformation", "includeCompanyPortalLink", "includeDeviceDetails", "unknownFutureValue"}[i]
}
func ParseEnrollmentNotificationBrandingOptions(v string) (interface{}, error) {
    result := NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    switch v {
        case "none":
            result = NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyLogo":
            result = INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyName":
            result = INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeContactInformation":
            result = INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyPortalLink":
            result = INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeDeviceDetails":
            result = INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        default:
            return 0, errors.New("Unknown EnrollmentNotificationBrandingOptions value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentNotificationBrandingOptions(values []EnrollmentNotificationBrandingOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

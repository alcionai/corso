package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type MobileAppIntent int

const (
    // Available
    AVAILABLE_MOBILEAPPINTENT MobileAppIntent = iota
    // Not Available
    NOTAVAILABLE_MOBILEAPPINTENT
    // Required Install
    REQUIREDINSTALL_MOBILEAPPINTENT
    // Required Uninstall
    REQUIREDUNINSTALL_MOBILEAPPINTENT
    // RequiredAndAvailableInstall
    REQUIREDANDAVAILABLEINSTALL_MOBILEAPPINTENT
    // AvailableInstallWithoutEnrollment
    AVAILABLEINSTALLWITHOUTENROLLMENT_MOBILEAPPINTENT
    // Exclude
    EXCLUDE_MOBILEAPPINTENT
)

func (i MobileAppIntent) String() string {
    return []string{"available", "notAvailable", "requiredInstall", "requiredUninstall", "requiredAndAvailableInstall", "availableInstallWithoutEnrollment", "exclude"}[i]
}
func ParseMobileAppIntent(v string) (interface{}, error) {
    result := AVAILABLE_MOBILEAPPINTENT
    switch v {
        case "available":
            result = AVAILABLE_MOBILEAPPINTENT
        case "notAvailable":
            result = NOTAVAILABLE_MOBILEAPPINTENT
        case "requiredInstall":
            result = REQUIREDINSTALL_MOBILEAPPINTENT
        case "requiredUninstall":
            result = REQUIREDUNINSTALL_MOBILEAPPINTENT
        case "requiredAndAvailableInstall":
            result = REQUIREDANDAVAILABLEINSTALL_MOBILEAPPINTENT
        case "availableInstallWithoutEnrollment":
            result = AVAILABLEINSTALLWITHOUTENROLLMENT_MOBILEAPPINTENT
        case "exclude":
            result = EXCLUDE_MOBILEAPPINTENT
        default:
            return 0, errors.New("Unknown MobileAppIntent value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppIntent(values []MobileAppIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

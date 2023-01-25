package models
import (
    "errors"
)
// Provides operations to call the add method.
type LocalSecurityOptionsInformationShownOnLockScreenType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE LocalSecurityOptionsInformationShownOnLockScreenType = iota
    // User display name, domain and user names
    USERDISPLAYNAMEDOMAINUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    // User display name only
    USERDISPLAYNAMEONLY_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    // Do not display user information
    DONOTDISPLAYUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
)

func (i LocalSecurityOptionsInformationShownOnLockScreenType) String() string {
    return []string{"notConfigured", "userDisplayNameDomainUser", "userDisplayNameOnly", "doNotDisplayUser"}[i]
}
func ParseLocalSecurityOptionsInformationShownOnLockScreenType(v string) (interface{}, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "userDisplayNameDomainUser":
            result = USERDISPLAYNAMEDOMAINUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "userDisplayNameOnly":
            result = USERDISPLAYNAMEONLY_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        case "doNotDisplayUser":
            result = DONOTDISPLAYUSER_LOCALSECURITYOPTIONSINFORMATIONSHOWNONLOCKSCREENTYPE
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsInformationShownOnLockScreenType value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsInformationShownOnLockScreenType(values []LocalSecurityOptionsInformationShownOnLockScreenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

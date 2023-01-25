package models
import (
    "errors"
)
// Provides operations to call the add method.
type LocalSecurityOptionsStandardUserElevationPromptBehaviorType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE LocalSecurityOptionsStandardUserElevationPromptBehaviorType = iota
    // Automatically deny elevation requests
    AUTOMATICALLYDENYELEVATIONREQUESTS_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for credentials on the secure desktop
    PROMPTFORCREDENTIALSONTHESECUREDESKTOP_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for credentials
    PROMPTFORCREDENTIALS_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
)

func (i LocalSecurityOptionsStandardUserElevationPromptBehaviorType) String() string {
    return []string{"notConfigured", "automaticallyDenyElevationRequests", "promptForCredentialsOnTheSecureDesktop", "promptForCredentials"}[i]
}
func ParseLocalSecurityOptionsStandardUserElevationPromptBehaviorType(v string) (interface{}, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
        case "automaticallyDenyElevationRequests":
            result = AUTOMATICALLYDENYELEVATIONREQUESTS_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
        case "promptForCredentialsOnTheSecureDesktop":
            result = PROMPTFORCREDENTIALSONTHESECUREDESKTOP_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
        case "promptForCredentials":
            result = PROMPTFORCREDENTIALS_LOCALSECURITYOPTIONSSTANDARDUSERELEVATIONPROMPTBEHAVIORTYPE
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsStandardUserElevationPromptBehaviorType value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsStandardUserElevationPromptBehaviorType(values []LocalSecurityOptionsStandardUserElevationPromptBehaviorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

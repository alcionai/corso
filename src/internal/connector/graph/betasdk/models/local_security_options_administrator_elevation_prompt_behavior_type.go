package models
import (
    "errors"
)
// Provides operations to call the add method.
type LocalSecurityOptionsAdministratorElevationPromptBehaviorType int

const (
    // Not Configured
    NOTCONFIGURED_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE LocalSecurityOptionsAdministratorElevationPromptBehaviorType = iota
    // Elevate without prompting.
    ELEVATEWITHOUTPROMPTING_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for credentials on the secure desktop
    PROMPTFORCREDENTIALSONTHESECUREDESKTOP_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for consent on the secure desktop
    PROMPTFORCONSENTONTHESECUREDESKTOP_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for credentials
    PROMPTFORCREDENTIALS_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for consent
    PROMPTFORCONSENT_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    // Prompt for consent for non-Windows binaries
    PROMPTFORCONSENTFORNONWINDOWSBINARIES_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
)

func (i LocalSecurityOptionsAdministratorElevationPromptBehaviorType) String() string {
    return []string{"notConfigured", "elevateWithoutPrompting", "promptForCredentialsOnTheSecureDesktop", "promptForConsentOnTheSecureDesktop", "promptForCredentials", "promptForConsent", "promptForConsentForNonWindowsBinaries"}[i]
}
func ParseLocalSecurityOptionsAdministratorElevationPromptBehaviorType(v string) (interface{}, error) {
    result := NOTCONFIGURED_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "elevateWithoutPrompting":
            result = ELEVATEWITHOUTPROMPTING_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "promptForCredentialsOnTheSecureDesktop":
            result = PROMPTFORCREDENTIALSONTHESECUREDESKTOP_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "promptForConsentOnTheSecureDesktop":
            result = PROMPTFORCONSENTONTHESECUREDESKTOP_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "promptForCredentials":
            result = PROMPTFORCREDENTIALS_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "promptForConsent":
            result = PROMPTFORCONSENT_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        case "promptForConsentForNonWindowsBinaries":
            result = PROMPTFORCONSENTFORNONWINDOWSBINARIES_LOCALSECURITYOPTIONSADMINISTRATORELEVATIONPROMPTBEHAVIORTYPE
        default:
            return 0, errors.New("Unknown LocalSecurityOptionsAdministratorElevationPromptBehaviorType value: " + v)
    }
    return &result, nil
}
func SerializeLocalSecurityOptionsAdministratorElevationPromptBehaviorType(values []LocalSecurityOptionsAdministratorElevationPromptBehaviorType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

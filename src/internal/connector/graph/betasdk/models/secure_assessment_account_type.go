package models
import (
    "errors"
)
// Provides operations to call the add method.
type SecureAssessmentAccountType int

const (
    // Indicates an Azure AD account in format of AzureAD\username@tenant.com.
    AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE SecureAssessmentAccountType = iota
    // Indicates a domain account in format of domain\user or user@domain.com.
    DOMAINACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    // Indicates a local account in format of username.
    LOCALACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    // Indicates a local guest account in format of test name.
    LOCALGUESTACCOUNT_SECUREASSESSMENTACCOUNTTYPE
)

func (i SecureAssessmentAccountType) String() string {
    return []string{"azureADAccount", "domainAccount", "localAccount", "localGuestAccount"}[i]
}
func ParseSecureAssessmentAccountType(v string) (interface{}, error) {
    result := AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE
    switch v {
        case "azureADAccount":
            result = AZUREADACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "domainAccount":
            result = DOMAINACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "localAccount":
            result = LOCALACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        case "localGuestAccount":
            result = LOCALGUESTACCOUNT_SECUREASSESSMENTACCOUNTTYPE
        default:
            return 0, errors.New("Unknown SecureAssessmentAccountType value: " + v)
    }
    return &result, nil
}
func SerializeSecureAssessmentAccountType(values []SecureAssessmentAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

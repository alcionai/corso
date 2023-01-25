package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EapFastConfiguration int

const (
    // Use EAP-FAST without Protected Access Credential (PAC).
    NOPROTECTEDACCESSCREDENTIAL_EAPFASTCONFIGURATION EapFastConfiguration = iota
    // Use Protected Access Credential (PAC).
    USEPROTECTEDACCESSCREDENTIAL_EAPFASTCONFIGURATION
    // Use Protected Access Credential (PAC) and Provision PAC.
    USEPROTECTEDACCESSCREDENTIALANDPROVISION_EAPFASTCONFIGURATION
    // Use Protected Access Credential (PAC), Provision PAC, and do so anonymously.
    USEPROTECTEDACCESSCREDENTIALANDPROVISIONANONYMOUSLY_EAPFASTCONFIGURATION
)

func (i EapFastConfiguration) String() string {
    return []string{"noProtectedAccessCredential", "useProtectedAccessCredential", "useProtectedAccessCredentialAndProvision", "useProtectedAccessCredentialAndProvisionAnonymously"}[i]
}
func ParseEapFastConfiguration(v string) (interface{}, error) {
    result := NOPROTECTEDACCESSCREDENTIAL_EAPFASTCONFIGURATION
    switch v {
        case "noProtectedAccessCredential":
            result = NOPROTECTEDACCESSCREDENTIAL_EAPFASTCONFIGURATION
        case "useProtectedAccessCredential":
            result = USEPROTECTEDACCESSCREDENTIAL_EAPFASTCONFIGURATION
        case "useProtectedAccessCredentialAndProvision":
            result = USEPROTECTEDACCESSCREDENTIALANDPROVISION_EAPFASTCONFIGURATION
        case "useProtectedAccessCredentialAndProvisionAnonymously":
            result = USEPROTECTEDACCESSCREDENTIALANDPROVISIONANONYMOUSLY_EAPFASTCONFIGURATION
        default:
            return 0, errors.New("Unknown EapFastConfiguration value: " + v)
    }
    return &result, nil
}
func SerializeEapFastConfiguration(values []EapFastConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

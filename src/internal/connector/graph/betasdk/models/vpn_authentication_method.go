package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnAuthenticationMethod int

const (
    // Authenticate with a certificate.
    CERTIFICATE_VPNAUTHENTICATIONMETHOD VpnAuthenticationMethod = iota
    // Use username and password for authentication.
    USERNAMEANDPASSWORD_VPNAUTHENTICATIONMETHOD
    // Use Shared Secret for Authentication.  Only valid for iOS IKEv2.
    SHAREDSECRET_VPNAUTHENTICATIONMETHOD
    // Use Derived Credential for Authentication.
    DERIVEDCREDENTIAL_VPNAUTHENTICATIONMETHOD
    // Use Azure AD for authentication.
    AZUREAD_VPNAUTHENTICATIONMETHOD
)

func (i VpnAuthenticationMethod) String() string {
    return []string{"certificate", "usernameAndPassword", "sharedSecret", "derivedCredential", "azureAD"}[i]
}
func ParseVpnAuthenticationMethod(v string) (interface{}, error) {
    result := CERTIFICATE_VPNAUTHENTICATIONMETHOD
    switch v {
        case "certificate":
            result = CERTIFICATE_VPNAUTHENTICATIONMETHOD
        case "usernameAndPassword":
            result = USERNAMEANDPASSWORD_VPNAUTHENTICATIONMETHOD
        case "sharedSecret":
            result = SHAREDSECRET_VPNAUTHENTICATIONMETHOD
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_VPNAUTHENTICATIONMETHOD
        case "azureAD":
            result = AZUREAD_VPNAUTHENTICATIONMETHOD
        default:
            return 0, errors.New("Unknown VpnAuthenticationMethod value: " + v)
    }
    return &result, nil
}
func SerializeVpnAuthenticationMethod(values []VpnAuthenticationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

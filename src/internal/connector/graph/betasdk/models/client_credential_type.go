package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ClientCredentialType int

const (
    NONE_CLIENTCREDENTIALTYPE ClientCredentialType = iota
    CLIENTSECRET_CLIENTCREDENTIALTYPE
    CLIENTASSERTION_CLIENTCREDENTIALTYPE
    FEDERATEDIDENTITYCREDENTIAL_CLIENTCREDENTIALTYPE
    MANAGEDIDENTITY_CLIENTCREDENTIALTYPE
    CERTIFICATE_CLIENTCREDENTIALTYPE
    UNKNOWNFUTUREVALUE_CLIENTCREDENTIALTYPE
)

func (i ClientCredentialType) String() string {
    return []string{"none", "clientSecret", "clientAssertion", "federatedIdentityCredential", "managedIdentity", "certificate", "unknownFutureValue"}[i]
}
func ParseClientCredentialType(v string) (interface{}, error) {
    result := NONE_CLIENTCREDENTIALTYPE
    switch v {
        case "none":
            result = NONE_CLIENTCREDENTIALTYPE
        case "clientSecret":
            result = CLIENTSECRET_CLIENTCREDENTIALTYPE
        case "clientAssertion":
            result = CLIENTASSERTION_CLIENTCREDENTIALTYPE
        case "federatedIdentityCredential":
            result = FEDERATEDIDENTITYCREDENTIAL_CLIENTCREDENTIALTYPE
        case "managedIdentity":
            result = MANAGEDIDENTITY_CLIENTCREDENTIALTYPE
        case "certificate":
            result = CERTIFICATE_CLIENTCREDENTIALTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLIENTCREDENTIALTYPE
        default:
            return 0, errors.New("Unknown ClientCredentialType value: " + v)
    }
    return &result, nil
}
func SerializeClientCredentialType(values []ClientCredentialType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

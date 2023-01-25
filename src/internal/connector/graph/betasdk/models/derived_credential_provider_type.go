package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DerivedCredentialProviderType int

const (
    // No Derived Credential Provider Configured.
    NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE DerivedCredentialProviderType = iota
    // Entrust.
    ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
    // Purebred - Defense Information Systems Agency.
    PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
    // Xtec - AuthentX.
    XTEC_DERIVEDCREDENTIALPROVIDERTYPE
    // Intercede.
    INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
)

func (i DerivedCredentialProviderType) String() string {
    return []string{"notConfigured", "entrustDataCard", "purebred", "xTec", "intercede"}[i]
}
func ParseDerivedCredentialProviderType(v string) (interface{}, error) {
    result := NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DERIVEDCREDENTIALPROVIDERTYPE
        case "entrustDataCard":
            result = ENTRUSTDATACARD_DERIVEDCREDENTIALPROVIDERTYPE
        case "purebred":
            result = PUREBRED_DERIVEDCREDENTIALPROVIDERTYPE
        case "xTec":
            result = XTEC_DERIVEDCREDENTIALPROVIDERTYPE
        case "intercede":
            result = INTERCEDE_DERIVEDCREDENTIALPROVIDERTYPE
        default:
            return 0, errors.New("Unknown DerivedCredentialProviderType value: " + v)
    }
    return &result, nil
}
func SerializeDerivedCredentialProviderType(values []DerivedCredentialProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

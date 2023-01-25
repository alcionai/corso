package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ExpirationRequirement int

const (
    REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES_EXPIRATIONREQUIREMENT ExpirationRequirement = iota
    TENANTTOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
    AUDIENCETOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
    SIGNINFREQUENCYPERIODICREAUTHENTICATION_EXPIRATIONREQUIREMENT
    NGCMFA_EXPIRATIONREQUIREMENT
    SIGNINFREQUENCYEVERYTIME_EXPIRATIONREQUIREMENT
    UNKNOWNFUTUREVALUE_EXPIRATIONREQUIREMENT
)

func (i ExpirationRequirement) String() string {
    return []string{"rememberMultifactorAuthenticationOnTrustedDevices", "tenantTokenLifetimePolicy", "audienceTokenLifetimePolicy", "signInFrequencyPeriodicReauthentication", "ngcMfa", "signInFrequencyEveryTime", "unknownFutureValue"}[i]
}
func ParseExpirationRequirement(v string) (interface{}, error) {
    result := REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES_EXPIRATIONREQUIREMENT
    switch v {
        case "rememberMultifactorAuthenticationOnTrustedDevices":
            result = REMEMBERMULTIFACTORAUTHENTICATIONONTRUSTEDDEVICES_EXPIRATIONREQUIREMENT
        case "tenantTokenLifetimePolicy":
            result = TENANTTOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
        case "audienceTokenLifetimePolicy":
            result = AUDIENCETOKENLIFETIMEPOLICY_EXPIRATIONREQUIREMENT
        case "signInFrequencyPeriodicReauthentication":
            result = SIGNINFREQUENCYPERIODICREAUTHENTICATION_EXPIRATIONREQUIREMENT
        case "ngcMfa":
            result = NGCMFA_EXPIRATIONREQUIREMENT
        case "signInFrequencyEveryTime":
            result = SIGNINFREQUENCYEVERYTIME_EXPIRATIONREQUIREMENT
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EXPIRATIONREQUIREMENT
        default:
            return 0, errors.New("Unknown ExpirationRequirement value: " + v)
    }
    return &result, nil
}
func SerializeExpirationRequirement(values []ExpirationRequirement) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

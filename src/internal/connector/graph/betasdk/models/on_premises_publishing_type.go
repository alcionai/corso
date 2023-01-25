package models
import (
    "errors"
)
// Provides operations to call the add method.
type OnPremisesPublishingType int

const (
    APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE OnPremisesPublishingType = iota
    EXCHANGEONLINE_ONPREMISESPUBLISHINGTYPE
    AUTHENTICATION_ONPREMISESPUBLISHINGTYPE
    PROVISIONING_ONPREMISESPUBLISHINGTYPE
    INTUNEPFX_ONPREMISESPUBLISHINGTYPE
    OFLINEDOMAINJOIN_ONPREMISESPUBLISHINGTYPE
    UNKNOWNFUTUREVALUE_ONPREMISESPUBLISHINGTYPE
)

func (i OnPremisesPublishingType) String() string {
    return []string{"applicationProxy", "exchangeOnline", "authentication", "provisioning", "intunePfx", "oflineDomainJoin", "unknownFutureValue"}[i]
}
func ParseOnPremisesPublishingType(v string) (interface{}, error) {
    result := APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE
    switch v {
        case "applicationProxy":
            result = APPLICATIONPROXY_ONPREMISESPUBLISHINGTYPE
        case "exchangeOnline":
            result = EXCHANGEONLINE_ONPREMISESPUBLISHINGTYPE
        case "authentication":
            result = AUTHENTICATION_ONPREMISESPUBLISHINGTYPE
        case "provisioning":
            result = PROVISIONING_ONPREMISESPUBLISHINGTYPE
        case "intunePfx":
            result = INTUNEPFX_ONPREMISESPUBLISHINGTYPE
        case "oflineDomainJoin":
            result = OFLINEDOMAINJOIN_ONPREMISESPUBLISHINGTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONPREMISESPUBLISHINGTYPE
        default:
            return 0, errors.New("Unknown OnPremisesPublishingType value: " + v)
    }
    return &result, nil
}
func SerializeOnPremisesPublishingType(values []OnPremisesPublishingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type NetworkType int

const (
    INTRANET_NETWORKTYPE NetworkType = iota
    EXTRANET_NETWORKTYPE
    NAMEDNETWORK_NETWORKTYPE
    TRUSTED_NETWORKTYPE
    TRUSTEDNAMEDLOCATION_NETWORKTYPE
    UNKNOWNFUTUREVALUE_NETWORKTYPE
)

func (i NetworkType) String() string {
    return []string{"intranet", "extranet", "namedNetwork", "trusted", "trustedNamedLocation", "unknownFutureValue"}[i]
}
func ParseNetworkType(v string) (interface{}, error) {
    result := INTRANET_NETWORKTYPE
    switch v {
        case "intranet":
            result = INTRANET_NETWORKTYPE
        case "extranet":
            result = EXTRANET_NETWORKTYPE
        case "namedNetwork":
            result = NAMEDNETWORK_NETWORKTYPE
        case "trusted":
            result = TRUSTED_NETWORKTYPE
        case "trustedNamedLocation":
            result = TRUSTEDNAMEDLOCATION_NETWORKTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NETWORKTYPE
        default:
            return 0, errors.New("Unknown NetworkType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkType(values []NetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ProtocolType int

const (
    NONE_PROTOCOLTYPE ProtocolType = iota
    OAUTH2_PROTOCOLTYPE
    ROPC_PROTOCOLTYPE
    WSFEDERATION_PROTOCOLTYPE
    SAML20_PROTOCOLTYPE
    DEVICECODE_PROTOCOLTYPE
    UNKNOWNFUTUREVALUE_PROTOCOLTYPE
)

func (i ProtocolType) String() string {
    return []string{"none", "oAuth2", "ropc", "wsFederation", "saml20", "deviceCode", "unknownFutureValue"}[i]
}
func ParseProtocolType(v string) (interface{}, error) {
    result := NONE_PROTOCOLTYPE
    switch v {
        case "none":
            result = NONE_PROTOCOLTYPE
        case "oAuth2":
            result = OAUTH2_PROTOCOLTYPE
        case "ropc":
            result = ROPC_PROTOCOLTYPE
        case "wsFederation":
            result = WSFEDERATION_PROTOCOLTYPE
        case "saml20":
            result = SAML20_PROTOCOLTYPE
        case "deviceCode":
            result = DEVICECODE_PROTOCOLTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PROTOCOLTYPE
        default:
            return 0, errors.New("Unknown ProtocolType value: " + v)
    }
    return &result, nil
}
func SerializeProtocolType(values []ProtocolType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

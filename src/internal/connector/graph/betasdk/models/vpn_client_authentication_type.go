package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnClientAuthenticationType int

const (
    // User Authentication
    USERAUTHENTICATION_VPNCLIENTAUTHENTICATIONTYPE VpnClientAuthenticationType = iota
    // Device Authentication
    DEVICEAUTHENTICATION_VPNCLIENTAUTHENTICATIONTYPE
)

func (i VpnClientAuthenticationType) String() string {
    return []string{"userAuthentication", "deviceAuthentication"}[i]
}
func ParseVpnClientAuthenticationType(v string) (interface{}, error) {
    result := USERAUTHENTICATION_VPNCLIENTAUTHENTICATIONTYPE
    switch v {
        case "userAuthentication":
            result = USERAUTHENTICATION_VPNCLIENTAUTHENTICATIONTYPE
        case "deviceAuthentication":
            result = DEVICEAUTHENTICATION_VPNCLIENTAUTHENTICATIONTYPE
        default:
            return 0, errors.New("Unknown VpnClientAuthenticationType value: " + v)
    }
    return &result, nil
}
func SerializeVpnClientAuthenticationType(values []VpnClientAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

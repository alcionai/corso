package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnServerCertificateType int

const (
    // RSA
    RSA_VPNSERVERCERTIFICATETYPE VpnServerCertificateType = iota
    // ECDSA256
    ECDSA256_VPNSERVERCERTIFICATETYPE
    // ECDSA384
    ECDSA384_VPNSERVERCERTIFICATETYPE
    // ECDSA521
    ECDSA521_VPNSERVERCERTIFICATETYPE
)

func (i VpnServerCertificateType) String() string {
    return []string{"rsa", "ecdsa256", "ecdsa384", "ecdsa521"}[i]
}
func ParseVpnServerCertificateType(v string) (interface{}, error) {
    result := RSA_VPNSERVERCERTIFICATETYPE
    switch v {
        case "rsa":
            result = RSA_VPNSERVERCERTIFICATETYPE
        case "ecdsa256":
            result = ECDSA256_VPNSERVERCERTIFICATETYPE
        case "ecdsa384":
            result = ECDSA384_VPNSERVERCERTIFICATETYPE
        case "ecdsa521":
            result = ECDSA521_VPNSERVERCERTIFICATETYPE
        default:
            return 0, errors.New("Unknown VpnServerCertificateType value: " + v)
    }
    return &result, nil
}
func SerializeVpnServerCertificateType(values []VpnServerCertificateType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnLocalIdentifier int

const (
    // Device Fully Qualified Domain Name
    DEVICEFQDN_VPNLOCALIDENTIFIER VpnLocalIdentifier = iota
    // Empty
    EMPTY_VPNLOCALIDENTIFIER
    // Client Certificate Subject Name
    CLIENTCERTIFICATESUBJECTNAME_VPNLOCALIDENTIFIER
)

func (i VpnLocalIdentifier) String() string {
    return []string{"deviceFQDN", "empty", "clientCertificateSubjectName"}[i]
}
func ParseVpnLocalIdentifier(v string) (interface{}, error) {
    result := DEVICEFQDN_VPNLOCALIDENTIFIER
    switch v {
        case "deviceFQDN":
            result = DEVICEFQDN_VPNLOCALIDENTIFIER
        case "empty":
            result = EMPTY_VPNLOCALIDENTIFIER
        case "clientCertificateSubjectName":
            result = CLIENTCERTIFICATESUBJECTNAME_VPNLOCALIDENTIFIER
        default:
            return 0, errors.New("Unknown VpnLocalIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeVpnLocalIdentifier(values []VpnLocalIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

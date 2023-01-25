package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type IntendedPurpose int

const (
    // Unassigned
    UNASSIGNED_INTENDEDPURPOSE IntendedPurpose = iota
    // SmimeEncryption
    SMIMEENCRYPTION_INTENDEDPURPOSE
    // SmimeSigning
    SMIMESIGNING_INTENDEDPURPOSE
    // VPN
    VPN_INTENDEDPURPOSE
    // Wifi
    WIFI_INTENDEDPURPOSE
)

func (i IntendedPurpose) String() string {
    return []string{"unassigned", "smimeEncryption", "smimeSigning", "vpn", "wifi"}[i]
}
func ParseIntendedPurpose(v string) (interface{}, error) {
    result := UNASSIGNED_INTENDEDPURPOSE
    switch v {
        case "unassigned":
            result = UNASSIGNED_INTENDEDPURPOSE
        case "smimeEncryption":
            result = SMIMEENCRYPTION_INTENDEDPURPOSE
        case "smimeSigning":
            result = SMIMESIGNING_INTENDEDPURPOSE
        case "vpn":
            result = VPN_INTENDEDPURPOSE
        case "wifi":
            result = WIFI_INTENDEDPURPOSE
        default:
            return 0, errors.New("Unknown IntendedPurpose value: " + v)
    }
    return &result, nil
}
func SerializeIntendedPurpose(values []IntendedPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

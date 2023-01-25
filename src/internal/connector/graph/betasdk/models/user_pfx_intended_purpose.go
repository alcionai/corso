package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type UserPfxIntendedPurpose int

const (
    // No roles/usages assigned.
    UNASSIGNED_USERPFXINTENDEDPURPOSE UserPfxIntendedPurpose = iota
    // Valid for S/MIME encryption.
    SMIMEENCRYPTION_USERPFXINTENDEDPURPOSE
    // Valid for S/MIME signing.
    SMIMESIGNING_USERPFXINTENDEDPURPOSE
    // Valid for use in VPN.
    VPN_USERPFXINTENDEDPURPOSE
    // Valid for use in WiFi.
    WIFI_USERPFXINTENDEDPURPOSE
)

func (i UserPfxIntendedPurpose) String() string {
    return []string{"unassigned", "smimeEncryption", "smimeSigning", "vpn", "wifi"}[i]
}
func ParseUserPfxIntendedPurpose(v string) (interface{}, error) {
    result := UNASSIGNED_USERPFXINTENDEDPURPOSE
    switch v {
        case "unassigned":
            result = UNASSIGNED_USERPFXINTENDEDPURPOSE
        case "smimeEncryption":
            result = SMIMEENCRYPTION_USERPFXINTENDEDPURPOSE
        case "smimeSigning":
            result = SMIMESIGNING_USERPFXINTENDEDPURPOSE
        case "vpn":
            result = VPN_USERPFXINTENDEDPURPOSE
        case "wifi":
            result = WIFI_USERPFXINTENDEDPURPOSE
        default:
            return 0, errors.New("Unknown UserPfxIntendedPurpose value: " + v)
    }
    return &result, nil
}
func SerializeUserPfxIntendedPurpose(values []UserPfxIntendedPurpose) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

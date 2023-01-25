package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsFirewallRuleNetworkProfileTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES WindowsFirewallRuleNetworkProfileTypes = iota
    // The profile for networks that are connected to domains.
    DOMAIN_WINDOWSFIREWALLRULENETWORKPROFILETYPES
    // The profile for private networks.
    PRIVATE_WINDOWSFIREWALLRULENETWORKPROFILETYPES
    // The profile for public networks.
    PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES
)

func (i WindowsFirewallRuleNetworkProfileTypes) String() string {
    return []string{"notConfigured", "domain", "private", "public"}[i]
}
func ParseWindowsFirewallRuleNetworkProfileTypes(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSFIREWALLRULENETWORKPROFILETYPES
        case "domain":
            result = DOMAIN_WINDOWSFIREWALLRULENETWORKPROFILETYPES
        case "private":
            result = PRIVATE_WINDOWSFIREWALLRULENETWORKPROFILETYPES
        case "public":
            result = PUBLIC_WINDOWSFIREWALLRULENETWORKPROFILETYPES
        default:
            return 0, errors.New("Unknown WindowsFirewallRuleNetworkProfileTypes value: " + v)
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleNetworkProfileTypes(values []WindowsFirewallRuleNetworkProfileTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsFirewallRuleInterfaceTypes int

const (
    // No flags set.
    NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES WindowsFirewallRuleInterfaceTypes = iota
    // The Remote Access interface type.
    REMOTEACCESS_WINDOWSFIREWALLRULEINTERFACETYPES
    // The Wireless interface type.
    WIRELESS_WINDOWSFIREWALLRULEINTERFACETYPES
    // The LAN interface type.
    LAN_WINDOWSFIREWALLRULEINTERFACETYPES
)

func (i WindowsFirewallRuleInterfaceTypes) String() string {
    return []string{"notConfigured", "remoteAccess", "wireless", "lan"}[i]
}
func ParseWindowsFirewallRuleInterfaceTypes(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSFIREWALLRULEINTERFACETYPES
        case "remoteAccess":
            result = REMOTEACCESS_WINDOWSFIREWALLRULEINTERFACETYPES
        case "wireless":
            result = WIRELESS_WINDOWSFIREWALLRULEINTERFACETYPES
        case "lan":
            result = LAN_WINDOWSFIREWALLRULEINTERFACETYPES
        default:
            return 0, errors.New("Unknown WindowsFirewallRuleInterfaceTypes value: " + v)
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleInterfaceTypes(values []WindowsFirewallRuleInterfaceTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

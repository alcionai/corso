package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsFirewallRuleTrafficDirectionType int

const (
    // Not configured.
    NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE WindowsFirewallRuleTrafficDirectionType = iota
    // The rule applies to outbound traffic.
    OUT_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
    // The rule applies to inbound traffic.
    IN_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
)

func (i WindowsFirewallRuleTrafficDirectionType) String() string {
    return []string{"notConfigured", "out", "in"}[i]
}
func ParseWindowsFirewallRuleTrafficDirectionType(v string) (interface{}, error) {
    result := NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        case "out":
            result = OUT_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        case "in":
            result = IN_WINDOWSFIREWALLRULETRAFFICDIRECTIONTYPE
        default:
            return 0, errors.New("Unknown WindowsFirewallRuleTrafficDirectionType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsFirewallRuleTrafficDirectionType(values []WindowsFirewallRuleTrafficDirectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type VpnTrafficRuleAppType int

const (
    // The traffic rule is not associated with an app.
    NONE_VPNTRAFFICRULEAPPTYPE VpnTrafficRuleAppType = iota
    // The traffic rule is associated with a desktop app.
    DESKTOP_VPNTRAFFICRULEAPPTYPE
    // The traffic rule is associated with a Universal app.
    UNIVERSAL_VPNTRAFFICRULEAPPTYPE
)

func (i VpnTrafficRuleAppType) String() string {
    return []string{"none", "desktop", "universal"}[i]
}
func ParseVpnTrafficRuleAppType(v string) (interface{}, error) {
    result := NONE_VPNTRAFFICRULEAPPTYPE
    switch v {
        case "none":
            result = NONE_VPNTRAFFICRULEAPPTYPE
        case "desktop":
            result = DESKTOP_VPNTRAFFICRULEAPPTYPE
        case "universal":
            result = UNIVERSAL_VPNTRAFFICRULEAPPTYPE
        default:
            return 0, errors.New("Unknown VpnTrafficRuleAppType value: " + v)
    }
    return &result, nil
}
func SerializeVpnTrafficRuleAppType(values []VpnTrafficRuleAppType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

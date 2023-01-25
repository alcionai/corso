package models
import (
    "errors"
)
// Provides operations to call the add method.
type VpnOnDemandRuleConnectionDomainAction int

const (
    // Connect if needed.
    CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION VpnOnDemandRuleConnectionDomainAction = iota
    // Never connect.
    NEVERCONNECT_VPNONDEMANDRULECONNECTIONDOMAINACTION
)

func (i VpnOnDemandRuleConnectionDomainAction) String() string {
    return []string{"connectIfNeeded", "neverConnect"}[i]
}
func ParseVpnOnDemandRuleConnectionDomainAction(v string) (interface{}, error) {
    result := CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION
    switch v {
        case "connectIfNeeded":
            result = CONNECTIFNEEDED_VPNONDEMANDRULECONNECTIONDOMAINACTION
        case "neverConnect":
            result = NEVERCONNECT_VPNONDEMANDRULECONNECTIONDOMAINACTION
        default:
            return 0, errors.New("Unknown VpnOnDemandRuleConnectionDomainAction value: " + v)
    }
    return &result, nil
}
func SerializeVpnOnDemandRuleConnectionDomainAction(values []VpnOnDemandRuleConnectionDomainAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

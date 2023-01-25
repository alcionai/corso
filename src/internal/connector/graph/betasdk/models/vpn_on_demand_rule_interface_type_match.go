package models
import (
    "errors"
)
// Provides operations to call the add method.
type VpnOnDemandRuleInterfaceTypeMatch int

const (
    // NotConfigured
    NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH VpnOnDemandRuleInterfaceTypeMatch = iota
    // Ethernet.
    ETHERNET_VPNONDEMANDRULEINTERFACETYPEMATCH
    // WiFi.
    WIFI_VPNONDEMANDRULEINTERFACETYPEMATCH
    // Cellular.
    CELLULAR_VPNONDEMANDRULEINTERFACETYPEMATCH
)

func (i VpnOnDemandRuleInterfaceTypeMatch) String() string {
    return []string{"notConfigured", "ethernet", "wiFi", "cellular"}[i]
}
func ParseVpnOnDemandRuleInterfaceTypeMatch(v string) (interface{}, error) {
    result := NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "ethernet":
            result = ETHERNET_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "wiFi":
            result = WIFI_VPNONDEMANDRULEINTERFACETYPEMATCH
        case "cellular":
            result = CELLULAR_VPNONDEMANDRULEINTERFACETYPEMATCH
        default:
            return 0, errors.New("Unknown VpnOnDemandRuleInterfaceTypeMatch value: " + v)
    }
    return &result, nil
}
func SerializeVpnOnDemandRuleInterfaceTypeMatch(values []VpnOnDemandRuleInterfaceTypeMatch) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

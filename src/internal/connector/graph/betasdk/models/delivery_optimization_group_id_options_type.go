package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DeliveryOptimizationGroupIdOptionsType int

const (
    // Not configured.
    NOTCONFIGURED_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE DeliveryOptimizationGroupIdOptionsType = iota
    // Active Directory site.
    ADSITE_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
    // Authenticated domain SID.
    AUTHENTICATEDDOMAINSID_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
    // DHCP user option.
    DHCPUSEROPTION_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
    // DNS suffix.
    DNSSUFFIX_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
)

func (i DeliveryOptimizationGroupIdOptionsType) String() string {
    return []string{"notConfigured", "adSite", "authenticatedDomainSid", "dhcpUserOption", "dnsSuffix"}[i]
}
func ParseDeliveryOptimizationGroupIdOptionsType(v string) (interface{}, error) {
    result := NOTCONFIGURED_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
        case "adSite":
            result = ADSITE_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
        case "authenticatedDomainSid":
            result = AUTHENTICATEDDOMAINSID_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
        case "dhcpUserOption":
            result = DHCPUSEROPTION_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
        case "dnsSuffix":
            result = DNSSUFFIX_DELIVERYOPTIMIZATIONGROUPIDOPTIONSTYPE
        default:
            return 0, errors.New("Unknown DeliveryOptimizationGroupIdOptionsType value: " + v)
    }
    return &result, nil
}
func SerializeDeliveryOptimizationGroupIdOptionsType(values []DeliveryOptimizationGroupIdOptionsType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

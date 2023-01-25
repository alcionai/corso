package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RoutingPolicy int

const (
    NONE_ROUTINGPOLICY RoutingPolicy = iota
    NOMISSEDCALL_ROUTINGPOLICY
    DISABLEFORWARDINGEXCEPTPHONE_ROUTINGPOLICY
    DISABLEFORWARDING_ROUTINGPOLICY
    PREFERSKYPEFORBUSINESS_ROUTINGPOLICY
    UNKNOWNFUTUREVALUE_ROUTINGPOLICY
)

func (i RoutingPolicy) String() string {
    return []string{"none", "noMissedCall", "disableForwardingExceptPhone", "disableForwarding", "preferSkypeForBusiness", "unknownFutureValue"}[i]
}
func ParseRoutingPolicy(v string) (interface{}, error) {
    result := NONE_ROUTINGPOLICY
    switch v {
        case "none":
            result = NONE_ROUTINGPOLICY
        case "noMissedCall":
            result = NOMISSEDCALL_ROUTINGPOLICY
        case "disableForwardingExceptPhone":
            result = DISABLEFORWARDINGEXCEPTPHONE_ROUTINGPOLICY
        case "disableForwarding":
            result = DISABLEFORWARDING_ROUTINGPOLICY
        case "preferSkypeForBusiness":
            result = PREFERSKYPEFORBUSINESS_ROUTINGPOLICY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ROUTINGPOLICY
        default:
            return 0, errors.New("Unknown RoutingPolicy value: " + v)
    }
    return &result, nil
}
func SerializeRoutingPolicy(values []RoutingPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

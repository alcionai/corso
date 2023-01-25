package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type RoutingMode int

const (
    ONETOONE_ROUTINGMODE RoutingMode = iota
    MULTICAST_ROUTINGMODE
)

func (i RoutingMode) String() string {
    return []string{"oneToOne", "multicast"}[i]
}
func ParseRoutingMode(v string) (interface{}, error) {
    result := ONETOONE_ROUTINGMODE
    switch v {
        case "oneToOne":
            result = ONETOONE_ROUTINGMODE
        case "multicast":
            result = MULTICAST_ROUTINGMODE
        default:
            return 0, errors.New("Unknown RoutingMode value: " + v)
    }
    return &result, nil
}
func SerializeRoutingMode(values []RoutingMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

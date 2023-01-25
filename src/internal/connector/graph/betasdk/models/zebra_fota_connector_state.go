package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ZebraFotaConnectorState int

const (
    // Default value when the connector has not been setup (the feature has not been used yet).
    NONE_ZEBRAFOTACONNECTORSTATE ZebraFotaConnectorState = iota
    // Connected state indicates that Intune is linked to Zebra Update Services for the current tenant.
    CONNECTED_ZEBRAFOTACONNECTORSTATE
    // Disconnected state indicates that the account was connected in the past and later disconnected.
    DISCONNECTED_ZEBRAFOTACONNECTORSTATE
    // Unknown future enum value.
    UNKNOWNFUTUREVALUE_ZEBRAFOTACONNECTORSTATE
)

func (i ZebraFotaConnectorState) String() string {
    return []string{"none", "connected", "disconnected", "unknownFutureValue"}[i]
}
func ParseZebraFotaConnectorState(v string) (interface{}, error) {
    result := NONE_ZEBRAFOTACONNECTORSTATE
    switch v {
        case "none":
            result = NONE_ZEBRAFOTACONNECTORSTATE
        case "connected":
            result = CONNECTED_ZEBRAFOTACONNECTORSTATE
        case "disconnected":
            result = DISCONNECTED_ZEBRAFOTACONNECTORSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTACONNECTORSTATE
        default:
            return 0, errors.New("Unknown ZebraFotaConnectorState value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaConnectorState(values []ZebraFotaConnectorState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

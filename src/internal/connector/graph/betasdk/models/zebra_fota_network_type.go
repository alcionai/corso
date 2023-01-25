package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type ZebraFotaNetworkType int

const (
    // The device will install the update regardless of current network type.
    ANY_ZEBRAFOTANETWORKTYPE ZebraFotaNetworkType = iota
    // The device will install the update only when connected to WiFi network.
    WIFI_ZEBRAFOTANETWORKTYPE
    // The device will install the update only when connected a Cellular network.
    CELLULAR_ZEBRAFOTANETWORKTYPE
    // The device will install the update when connected both WiFi and Cellular.
    WIFIANDCELLULAR_ZEBRAFOTANETWORKTYPE
    // Unknown future enum value.
    UNKNOWNFUTUREVALUE_ZEBRAFOTANETWORKTYPE
)

func (i ZebraFotaNetworkType) String() string {
    return []string{"any", "wifi", "cellular", "wifiAndCellular", "unknownFutureValue"}[i]
}
func ParseZebraFotaNetworkType(v string) (interface{}, error) {
    result := ANY_ZEBRAFOTANETWORKTYPE
    switch v {
        case "any":
            result = ANY_ZEBRAFOTANETWORKTYPE
        case "wifi":
            result = WIFI_ZEBRAFOTANETWORKTYPE
        case "cellular":
            result = CELLULAR_ZEBRAFOTANETWORKTYPE
        case "wifiAndCellular":
            result = WIFIANDCELLULAR_ZEBRAFOTANETWORKTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTANETWORKTYPE
        default:
            return 0, errors.New("Unknown ZebraFotaNetworkType value: " + v)
    }
    return &result, nil
}
func SerializeZebraFotaNetworkType(values []ZebraFotaNetworkType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

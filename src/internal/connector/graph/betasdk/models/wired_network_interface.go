package models
import (
    "errors"
)
// Provides operations to call the add method.
type WiredNetworkInterface int

const (
    // Any Ethernet.
    ANYETHERNET_WIREDNETWORKINTERFACE WiredNetworkInterface = iota
    // First active Ethernet.
    FIRSTACTIVEETHERNET_WIREDNETWORKINTERFACE
    // Second active Ethernet.
    SECONDACTIVEETHERNET_WIREDNETWORKINTERFACE
    // Third active Ethernet.
    THIRDACTIVEETHERNET_WIREDNETWORKINTERFACE
    // First Ethernet.
    FIRSTETHERNET_WIREDNETWORKINTERFACE
    // Second Ethernet.
    SECONDETHERNET_WIREDNETWORKINTERFACE
    // Third Ethernet.
    THIRDETHERNET_WIREDNETWORKINTERFACE
)

func (i WiredNetworkInterface) String() string {
    return []string{"anyEthernet", "firstActiveEthernet", "secondActiveEthernet", "thirdActiveEthernet", "firstEthernet", "secondEthernet", "thirdEthernet"}[i]
}
func ParseWiredNetworkInterface(v string) (interface{}, error) {
    result := ANYETHERNET_WIREDNETWORKINTERFACE
    switch v {
        case "anyEthernet":
            result = ANYETHERNET_WIREDNETWORKINTERFACE
        case "firstActiveEthernet":
            result = FIRSTACTIVEETHERNET_WIREDNETWORKINTERFACE
        case "secondActiveEthernet":
            result = SECONDACTIVEETHERNET_WIREDNETWORKINTERFACE
        case "thirdActiveEthernet":
            result = THIRDACTIVEETHERNET_WIREDNETWORKINTERFACE
        case "firstEthernet":
            result = FIRSTETHERNET_WIREDNETWORKINTERFACE
        case "secondEthernet":
            result = SECONDETHERNET_WIREDNETWORKINTERFACE
        case "thirdEthernet":
            result = THIRDETHERNET_WIREDNETWORKINTERFACE
        default:
            return 0, errors.New("Unknown WiredNetworkInterface value: " + v)
    }
    return &result, nil
}
func SerializeWiredNetworkInterface(values []WiredNetworkInterface) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

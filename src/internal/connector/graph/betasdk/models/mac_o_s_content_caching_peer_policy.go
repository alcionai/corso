package models
import (
    "errors"
)
// Provides operations to call the add method.
type MacOSContentCachingPeerPolicy int

const (
    // Defaults to peers in local network.
    NOTCONFIGURED_MACOSCONTENTCACHINGPEERPOLICY MacOSContentCachingPeerPolicy = iota
    // Content caches will only peer with caches in their immediate local network.
    PEERSINLOCALNETWORK_MACOSCONTENTCACHINGPEERPOLICY
    // Content caches will only peer with caches that share the same public IP address.
    PEERSWITHSAMEPUBLICIPADDRESS_MACOSCONTENTCACHINGPEERPOLICY
    // Content caches will use contentCachingPeerFilterRanges and contentCachingPeerListenRanges to determine which caches to peer with.
    PEERSINCUSTOMLOCALNETWORKS_MACOSCONTENTCACHINGPEERPOLICY
)

func (i MacOSContentCachingPeerPolicy) String() string {
    return []string{"notConfigured", "peersInLocalNetwork", "peersWithSamePublicIpAddress", "peersInCustomLocalNetworks"}[i]
}
func ParseMacOSContentCachingPeerPolicy(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSCONTENTCACHINGPEERPOLICY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSCONTENTCACHINGPEERPOLICY
        case "peersInLocalNetwork":
            result = PEERSINLOCALNETWORK_MACOSCONTENTCACHINGPEERPOLICY
        case "peersWithSamePublicIpAddress":
            result = PEERSWITHSAMEPUBLICIPADDRESS_MACOSCONTENTCACHINGPEERPOLICY
        case "peersInCustomLocalNetworks":
            result = PEERSINCUSTOMLOCALNETWORKS_MACOSCONTENTCACHINGPEERPOLICY
        default:
            return 0, errors.New("Unknown MacOSContentCachingPeerPolicy value: " + v)
    }
    return &result, nil
}
func SerializeMacOSContentCachingPeerPolicy(values []MacOSContentCachingPeerPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

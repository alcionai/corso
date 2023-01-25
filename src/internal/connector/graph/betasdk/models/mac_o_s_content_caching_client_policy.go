package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MacOSContentCachingClientPolicy int

const (
    // Defaults to clients in local network.
    NOTCONFIGURED_MACOSCONTENTCACHINGCLIENTPOLICY MacOSContentCachingClientPolicy = iota
    // Content caches will provide content to devices only in their immediate local network.
    CLIENTSINLOCALNETWORK_MACOSCONTENTCACHINGCLIENTPOLICY
    // Content caches will provide content to devices that share the same public IP address.
    CLIENTSWITHSAMEPUBLICIPADDRESS_MACOSCONTENTCACHINGCLIENTPOLICY
    // Content caches will provide content to devices in contentCachingClientListenRanges.
    CLIENTSINCUSTOMLOCALNETWORKS_MACOSCONTENTCACHINGCLIENTPOLICY
    // Content caches will provide content to devices in contentCachingClientListenRanges, contentCachingPeerListenRanges, and contentCachingParents.
    CLIENTSINCUSTOMLOCALNETWORKSWITHFALLBACK_MACOSCONTENTCACHINGCLIENTPOLICY
)

func (i MacOSContentCachingClientPolicy) String() string {
    return []string{"notConfigured", "clientsInLocalNetwork", "clientsWithSamePublicIpAddress", "clientsInCustomLocalNetworks", "clientsInCustomLocalNetworksWithFallback"}[i]
}
func ParseMacOSContentCachingClientPolicy(v string) (interface{}, error) {
    result := NOTCONFIGURED_MACOSCONTENTCACHINGCLIENTPOLICY
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_MACOSCONTENTCACHINGCLIENTPOLICY
        case "clientsInLocalNetwork":
            result = CLIENTSINLOCALNETWORK_MACOSCONTENTCACHINGCLIENTPOLICY
        case "clientsWithSamePublicIpAddress":
            result = CLIENTSWITHSAMEPUBLICIPADDRESS_MACOSCONTENTCACHINGCLIENTPOLICY
        case "clientsInCustomLocalNetworks":
            result = CLIENTSINCUSTOMLOCALNETWORKS_MACOSCONTENTCACHINGCLIENTPOLICY
        case "clientsInCustomLocalNetworksWithFallback":
            result = CLIENTSINCUSTOMLOCALNETWORKSWITHFALLBACK_MACOSCONTENTCACHINGCLIENTPOLICY
        default:
            return 0, errors.New("Unknown MacOSContentCachingClientPolicy value: " + v)
    }
    return &result, nil
}
func SerializeMacOSContentCachingClientPolicy(values []MacOSContentCachingClientPolicy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

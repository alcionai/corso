package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelConfigurationable 
type MicrosoftTunnelConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedSettings()([]KeyValuePairable)
    GetDefaultDomainSuffix()(*string)
    GetDescription()(*string)
    GetDisableUdpConnections()(*bool)
    GetDisplayName()(*string)
    GetDnsServers()([]string)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetListenPort()(*int32)
    GetNetwork()(*string)
    GetRoleScopeTagIds()([]string)
    GetRouteExcludes()([]string)
    GetRouteIncludes()([]string)
    GetRoutesExclude()([]string)
    GetRoutesInclude()([]string)
    GetSplitDNS()([]string)
    SetAdvancedSettings(value []KeyValuePairable)()
    SetDefaultDomainSuffix(value *string)()
    SetDescription(value *string)()
    SetDisableUdpConnections(value *bool)()
    SetDisplayName(value *string)()
    SetDnsServers(value []string)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetListenPort(value *int32)()
    SetNetwork(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetRouteExcludes(value []string)()
    SetRouteIncludes(value []string)()
    SetRoutesExclude(value []string)()
    SetRoutesInclude(value []string)()
    SetSplitDNS(value []string)()
}

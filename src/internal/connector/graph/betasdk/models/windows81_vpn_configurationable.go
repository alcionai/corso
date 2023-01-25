package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81VpnConfigurationable 
type Windows81VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsVpnConfigurationable
    GetApplyOnlyToWindows81()(*bool)
    GetConnectionType()(*WindowsVpnConnectionType)
    GetEnableSplitTunneling()(*bool)
    GetLoginGroupOrDomain()(*string)
    GetProxyServer()(Windows81VpnProxyServerable)
    SetApplyOnlyToWindows81(value *bool)()
    SetConnectionType(value *WindowsVpnConnectionType)()
    SetEnableSplitTunneling(value *bool)()
    SetLoginGroupOrDomain(value *string)()
    SetProxyServer(value Windows81VpnProxyServerable)()
}

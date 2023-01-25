package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleVpnAlwaysOnConfigurationable 
type AppleVpnAlwaysOnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAirPrintExceptionAction()(*VpnServiceExceptionAction)
    GetAllowAllCaptiveNetworkPlugins()(*bool)
    GetAllowCaptiveWebSheet()(*bool)
    GetAllowedCaptiveNetworkPlugins()(SpecifiedCaptiveNetworkPluginsable)
    GetCellularExceptionAction()(*VpnServiceExceptionAction)
    GetNatKeepAliveIntervalInSeconds()(*int32)
    GetNatKeepAliveOffloadEnable()(*bool)
    GetOdataType()(*string)
    GetTunnelConfiguration()(*VpnTunnelConfigurationType)
    GetUserToggleEnabled()(*bool)
    GetVoicemailExceptionAction()(*VpnServiceExceptionAction)
    SetAirPrintExceptionAction(value *VpnServiceExceptionAction)()
    SetAllowAllCaptiveNetworkPlugins(value *bool)()
    SetAllowCaptiveWebSheet(value *bool)()
    SetAllowedCaptiveNetworkPlugins(value SpecifiedCaptiveNetworkPluginsable)()
    SetCellularExceptionAction(value *VpnServiceExceptionAction)()
    SetNatKeepAliveIntervalInSeconds(value *int32)()
    SetNatKeepAliveOffloadEnable(value *bool)()
    SetOdataType(value *string)()
    SetTunnelConfiguration(value *VpnTunnelConfigurationType)()
    SetUserToggleEnabled(value *bool)()
    SetVoicemailExceptionAction(value *VpnServiceExceptionAction)()
}

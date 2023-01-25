package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWifiConfigurationable 
type WindowsWifiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectToPreferredNetwork()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetForceFIPSCompliance()(*bool)
    GetMeteredConnectionLimit()(*MeteredConnectionLimitType)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxyManualAddress()(*string)
    GetProxyManualPort()(*int32)
    GetProxySetting()(*WiFiProxySetting)
    GetSsid()(*string)
    GetWifiSecurityType()(*WiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectToPreferredNetwork(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetForceFIPSCompliance(value *bool)()
    SetMeteredConnectionLimit(value *MeteredConnectionLimitType)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxyManualAddress(value *string)()
    SetProxyManualPort(value *int32)()
    SetProxySetting(value *WiFiProxySetting)()
    SetSsid(value *string)()
    SetWifiSecurityType(value *WiFiSecurityType)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerWiFiConfigurationable 
type AospDeviceOwnerWiFiConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectAutomatically()(*bool)
    GetConnectWhenNetworkNameIsHidden()(*bool)
    GetNetworkName()(*string)
    GetPreSharedKey()(*string)
    GetPreSharedKeyIsSet()(*bool)
    GetSsid()(*string)
    GetWiFiSecurityType()(*AospDeviceOwnerWiFiSecurityType)
    SetConnectAutomatically(value *bool)()
    SetConnectWhenNetworkNameIsHidden(value *bool)()
    SetNetworkName(value *string)()
    SetPreSharedKey(value *string)()
    SetPreSharedKeyIsSet(value *bool)()
    SetSsid(value *string)()
    SetWiFiSecurityType(value *AospDeviceOwnerWiFiSecurityType)()
}

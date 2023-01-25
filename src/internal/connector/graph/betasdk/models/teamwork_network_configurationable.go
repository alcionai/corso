package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkNetworkConfigurationable 
type TeamworkNetworkConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultGateway()(*string)
    GetDomainName()(*string)
    GetHostName()(*string)
    GetIpAddress()(*string)
    GetIsDhcpEnabled()(*bool)
    GetIsPCPortEnabled()(*bool)
    GetOdataType()(*string)
    GetPrimaryDns()(*string)
    GetSecondaryDns()(*string)
    GetSubnetMask()(*string)
    SetDefaultGateway(value *string)()
    SetDomainName(value *string)()
    SetHostName(value *string)()
    SetIpAddress(value *string)()
    SetIsDhcpEnabled(value *bool)()
    SetIsPCPortEnabled(value *bool)()
    SetOdataType(value *string)()
    SetPrimaryDns(value *string)()
    SetSecondaryDns(value *string)()
    SetSubnetMask(value *string)()
}

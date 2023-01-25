package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnTrafficRuleable 
type VpnTrafficRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*string)
    GetAppType()(*VpnTrafficRuleAppType)
    GetClaims()(*string)
    GetLocalAddressRanges()([]IPv4Rangeable)
    GetLocalPortRanges()([]NumberRangeable)
    GetName()(*string)
    GetOdataType()(*string)
    GetProtocols()(*int32)
    GetRemoteAddressRanges()([]IPv4Rangeable)
    GetRemotePortRanges()([]NumberRangeable)
    GetRoutingPolicyType()(*VpnTrafficRuleRoutingPolicyType)
    SetAppId(value *string)()
    SetAppType(value *VpnTrafficRuleAppType)()
    SetClaims(value *string)()
    SetLocalAddressRanges(value []IPv4Rangeable)()
    SetLocalPortRanges(value []NumberRangeable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetProtocols(value *int32)()
    SetRemoteAddressRanges(value []IPv4Rangeable)()
    SetRemotePortRanges(value []NumberRangeable)()
    SetRoutingPolicyType(value *VpnTrafficRuleRoutingPolicyType)()
}

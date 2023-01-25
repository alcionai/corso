package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFirewallRuleable 
type WindowsFirewallRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*StateManagementSetting)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEdgeTraversal()(*StateManagementSetting)
    GetFilePath()(*string)
    GetInterfaceTypes()(*WindowsFirewallRuleInterfaceTypes)
    GetLocalAddressRanges()([]string)
    GetLocalPortRanges()([]string)
    GetLocalUserAuthorizations()(*string)
    GetOdataType()(*string)
    GetPackageFamilyName()(*string)
    GetProfileTypes()(*WindowsFirewallRuleNetworkProfileTypes)
    GetProtocol()(*int32)
    GetRemoteAddressRanges()([]string)
    GetRemotePortRanges()([]string)
    GetServiceName()(*string)
    GetTrafficDirection()(*WindowsFirewallRuleTrafficDirectionType)
    SetAction(value *StateManagementSetting)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEdgeTraversal(value *StateManagementSetting)()
    SetFilePath(value *string)()
    SetInterfaceTypes(value *WindowsFirewallRuleInterfaceTypes)()
    SetLocalAddressRanges(value []string)()
    SetLocalPortRanges(value []string)()
    SetLocalUserAuthorizations(value *string)()
    SetOdataType(value *string)()
    SetPackageFamilyName(value *string)()
    SetProfileTypes(value *WindowsFirewallRuleNetworkProfileTypes)()
    SetProtocol(value *int32)()
    SetRemoteAddressRanges(value []string)()
    SetRemotePortRanges(value []string)()
    SetServiceName(value *string)()
    SetTrafficDirection(value *WindowsFirewallRuleTrafficDirectionType)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityProfileable 
type HostSecurityProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureSubscriptionId()(*string)
    GetAzureTenantId()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFqdn()(*string)
    GetIsAzureAdJoined()(*bool)
    GetIsAzureAdRegistered()(*bool)
    GetIsHybridAzureDomainJoined()(*bool)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLogonUsers()([]LogonUserable)
    GetNetBiosName()(*string)
    GetNetworkInterfaces()([]NetworkInterfaceable)
    GetOs()(*string)
    GetOsVersion()(*string)
    GetParentHost()(*string)
    GetRelatedHostIds()([]string)
    GetRiskScore()(*string)
    GetTags()([]string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetAzureSubscriptionId(value *string)()
    SetAzureTenantId(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFqdn(value *string)()
    SetIsAzureAdJoined(value *bool)()
    SetIsAzureAdRegistered(value *bool)()
    SetIsHybridAzureDomainJoined(value *bool)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLogonUsers(value []LogonUserable)()
    SetNetBiosName(value *string)()
    SetNetworkInterfaces(value []NetworkInterfaceable)()
    SetOs(value *string)()
    SetOsVersion(value *string)()
    SetParentHost(value *string)()
    SetRelatedHostIds(value []string)()
    SetRiskScore(value *string)()
    SetTags(value []string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}

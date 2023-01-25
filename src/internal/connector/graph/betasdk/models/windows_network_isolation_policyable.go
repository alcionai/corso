package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsNetworkIsolationPolicyable 
type WindowsNetworkIsolationPolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnterpriseCloudResources()([]ProxiedDomainable)
    GetEnterpriseInternalProxyServers()([]string)
    GetEnterpriseIPRanges()([]IpRangeable)
    GetEnterpriseIPRangesAreAuthoritative()(*bool)
    GetEnterpriseNetworkDomainNames()([]string)
    GetEnterpriseProxyServers()([]string)
    GetEnterpriseProxyServersAreAuthoritative()(*bool)
    GetNeutralDomainResources()([]string)
    GetOdataType()(*string)
    SetEnterpriseCloudResources(value []ProxiedDomainable)()
    SetEnterpriseInternalProxyServers(value []string)()
    SetEnterpriseIPRanges(value []IpRangeable)()
    SetEnterpriseIPRangesAreAuthoritative(value *bool)()
    SetEnterpriseNetworkDomainNames(value []string)()
    SetEnterpriseProxyServers(value []string)()
    SetEnterpriseProxyServersAreAuthoritative(value *bool)()
    SetNeutralDomainResources(value []string)()
    SetOdataType(value *string)()
}

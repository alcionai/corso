package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOnPremisesConnectionable 
type CloudPcOnPremisesConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdDomainName()(*string)
    GetAdDomainPassword()(*string)
    GetAdDomainUsername()(*string)
    GetAlternateResourceUrl()(*string)
    GetDisplayName()(*string)
    GetHealthCheckStatus()(*CloudPcOnPremisesConnectionStatus)
    GetHealthCheckStatusDetails()(CloudPcOnPremisesConnectionStatusDetailsable)
    GetInUse()(*bool)
    GetManagedBy()(*CloudPcManagementService)
    GetOrganizationalUnit()(*string)
    GetResourceGroupId()(*string)
    GetSubnetId()(*string)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetType()(*CloudPcOnPremisesConnectionType)
    GetVirtualNetworkId()(*string)
    GetVirtualNetworkLocation()(*string)
    SetAdDomainName(value *string)()
    SetAdDomainPassword(value *string)()
    SetAdDomainUsername(value *string)()
    SetAlternateResourceUrl(value *string)()
    SetDisplayName(value *string)()
    SetHealthCheckStatus(value *CloudPcOnPremisesConnectionStatus)()
    SetHealthCheckStatusDetails(value CloudPcOnPremisesConnectionStatusDetailsable)()
    SetInUse(value *bool)()
    SetManagedBy(value *CloudPcManagementService)()
    SetOrganizationalUnit(value *string)()
    SetResourceGroupId(value *string)()
    SetSubnetId(value *string)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetType(value *CloudPcOnPremisesConnectionType)()
    SetVirtualNetworkId(value *string)()
    SetVirtualNetworkLocation(value *string)()
}

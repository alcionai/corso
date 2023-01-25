package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Domainable 
type Domainable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationType()(*string)
    GetAvailabilityStatus()(*string)
    GetDomainNameReferences()([]DirectoryObjectable)
    GetFederationConfiguration()([]InternalDomainFederationable)
    GetIsAdminManaged()(*bool)
    GetIsDefault()(*bool)
    GetIsInitial()(*bool)
    GetIsRoot()(*bool)
    GetIsVerified()(*bool)
    GetPasswordNotificationWindowInDays()(*int32)
    GetPasswordValidityPeriodInDays()(*int32)
    GetServiceConfigurationRecords()([]DomainDnsRecordable)
    GetSharedEmailDomainInvitations()([]SharedEmailDomainInvitationable)
    GetState()(DomainStateable)
    GetSupportedServices()([]string)
    GetVerificationDnsRecords()([]DomainDnsRecordable)
    SetAuthenticationType(value *string)()
    SetAvailabilityStatus(value *string)()
    SetDomainNameReferences(value []DirectoryObjectable)()
    SetFederationConfiguration(value []InternalDomainFederationable)()
    SetIsAdminManaged(value *bool)()
    SetIsDefault(value *bool)()
    SetIsInitial(value *bool)()
    SetIsRoot(value *bool)()
    SetIsVerified(value *bool)()
    SetPasswordNotificationWindowInDays(value *int32)()
    SetPasswordValidityPeriodInDays(value *int32)()
    SetServiceConfigurationRecords(value []DomainDnsRecordable)()
    SetSharedEmailDomainInvitations(value []SharedEmailDomainInvitationable)()
    SetState(value DomainStateable)()
    SetSupportedServices(value []string)()
    SetVerificationDnsRecords(value []DomainDnsRecordable)()
}

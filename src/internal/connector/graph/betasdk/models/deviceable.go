package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deviceable 
type Deviceable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountEnabled()(*bool)
    GetAlternativeSecurityIds()([]AlternativeSecurityIdable)
    GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCommands()([]Commandable)
    GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceCategory()(*string)
    GetDeviceId()(*string)
    GetDeviceMetadata()(*string)
    GetDeviceOwnership()(*string)
    GetDeviceVersion()(*int32)
    GetDisplayName()(*string)
    GetDomainName()(*string)
    GetEnrollmentProfileName()(*string)
    GetEnrollmentType()(*string)
    GetExtensionAttributes()(OnPremisesExtensionAttributesable)
    GetExtensions()([]Extensionable)
    GetHostnames()([]string)
    GetIsCompliant()(*bool)
    GetIsManaged()(*bool)
    GetIsManagementRestricted()(*bool)
    GetIsRooted()(*bool)
    GetKind()(*string)
    GetManagementType()(*string)
    GetManufacturer()(*string)
    GetMdmAppId()(*string)
    GetMemberOf()([]DirectoryObjectable)
    GetModel()(*string)
    GetName()(*string)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesSyncEnabled()(*bool)
    GetOperatingSystem()(*string)
    GetOperatingSystemVersion()(*string)
    GetPhysicalIds()([]string)
    GetPlatform()(*string)
    GetProfileType()(*string)
    GetRegisteredOwners()([]DirectoryObjectable)
    GetRegisteredUsers()([]DirectoryObjectable)
    GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*string)
    GetSystemLabels()([]string)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetTrustType()(*string)
    GetUsageRights()([]UsageRightable)
    SetAccountEnabled(value *bool)()
    SetAlternativeSecurityIds(value []AlternativeSecurityIdable)()
    SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCommands(value []Commandable)()
    SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceCategory(value *string)()
    SetDeviceId(value *string)()
    SetDeviceMetadata(value *string)()
    SetDeviceOwnership(value *string)()
    SetDeviceVersion(value *int32)()
    SetDisplayName(value *string)()
    SetDomainName(value *string)()
    SetEnrollmentProfileName(value *string)()
    SetEnrollmentType(value *string)()
    SetExtensionAttributes(value OnPremisesExtensionAttributesable)()
    SetExtensions(value []Extensionable)()
    SetHostnames(value []string)()
    SetIsCompliant(value *bool)()
    SetIsManaged(value *bool)()
    SetIsManagementRestricted(value *bool)()
    SetIsRooted(value *bool)()
    SetKind(value *string)()
    SetManagementType(value *string)()
    SetManufacturer(value *string)()
    SetMdmAppId(value *string)()
    SetMemberOf(value []DirectoryObjectable)()
    SetModel(value *string)()
    SetName(value *string)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOperatingSystem(value *string)()
    SetOperatingSystemVersion(value *string)()
    SetPhysicalIds(value []string)()
    SetPlatform(value *string)()
    SetProfileType(value *string)()
    SetRegisteredOwners(value []DirectoryObjectable)()
    SetRegisteredUsers(value []DirectoryObjectable)()
    SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *string)()
    SetSystemLabels(value []string)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetTrustType(value *string)()
    SetUsageRights(value []UsageRightable)()
}

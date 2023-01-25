package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppRegistrationable 
type ManagedAppRegistrationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppIdentifier()(MobileAppIdentifierable)
    GetApplicationVersion()(*string)
    GetAppliedPolicies()([]ManagedAppPolicyable)
    GetAzureADDeviceId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetDeviceTag()(*string)
    GetDeviceType()(*string)
    GetFlaggedReasons()([]ManagedAppFlaggedReason)
    GetIntendedPolicies()([]ManagedAppPolicyable)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetManagementSdkVersion()(*string)
    GetOperations()([]ManagedAppOperationable)
    GetPlatformVersion()(*string)
    GetUserId()(*string)
    GetVersion()(*string)
    SetAppIdentifier(value MobileAppIdentifierable)()
    SetApplicationVersion(value *string)()
    SetAppliedPolicies(value []ManagedAppPolicyable)()
    SetAzureADDeviceId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetDeviceTag(value *string)()
    SetDeviceType(value *string)()
    SetFlaggedReasons(value []ManagedAppFlaggedReason)()
    SetIntendedPolicies(value []ManagedAppPolicyable)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetManagementSdkVersion(value *string)()
    SetOperations(value []ManagedAppOperationable)()
    SetPlatformVersion(value *string)()
    SetUserId(value *string)()
    SetVersion(value *string)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAccountEnterpriseSettingsable 
type AndroidManagedStoreAccountEnterpriseSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidDeviceOwnerFullyManagedEnrollmentEnabled()(*bool)
    GetBindStatus()(*AndroidManagedStoreAccountBindStatus)
    GetCompanyCodes()([]AndroidEnrollmentCompanyCodeable)
    GetDeviceOwnerManagementEnabled()(*bool)
    GetEnrollmentTarget()(*AndroidManagedStoreAccountEnrollmentTarget)
    GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastAppSyncStatus()(*AndroidManagedStoreAccountAppSyncStatus)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedGooglePlayInitialScopeTagIds()([]string)
    GetOwnerOrganizationName()(*string)
    GetOwnerUserPrincipalName()(*string)
    GetTargetGroupIds()([]string)
    SetAndroidDeviceOwnerFullyManagedEnrollmentEnabled(value *bool)()
    SetBindStatus(value *AndroidManagedStoreAccountBindStatus)()
    SetCompanyCodes(value []AndroidEnrollmentCompanyCodeable)()
    SetDeviceOwnerManagementEnabled(value *bool)()
    SetEnrollmentTarget(value *AndroidManagedStoreAccountEnrollmentTarget)()
    SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastAppSyncStatus(value *AndroidManagedStoreAccountAppSyncStatus)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedGooglePlayInitialScopeTagIds(value []string)()
    SetOwnerOrganizationName(value *string)()
    SetOwnerUserPrincipalName(value *string)()
    SetTargetGroupIds(value []string)()
}

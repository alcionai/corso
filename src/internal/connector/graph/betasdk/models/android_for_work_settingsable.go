package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkSettingsable 
type AndroidForWorkSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBindStatus()(*AndroidForWorkBindStatus)
    GetDeviceOwnerManagementEnabled()(*bool)
    GetEnrollmentTarget()(*AndroidForWorkEnrollmentTarget)
    GetLastAppSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastAppSyncStatus()(*AndroidForWorkSyncStatus)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOwnerOrganizationName()(*string)
    GetOwnerUserPrincipalName()(*string)
    GetTargetGroupIds()([]string)
    SetBindStatus(value *AndroidForWorkBindStatus)()
    SetDeviceOwnerManagementEnabled(value *bool)()
    SetEnrollmentTarget(value *AndroidForWorkEnrollmentTarget)()
    SetLastAppSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastAppSyncStatus(value *AndroidForWorkSyncStatus)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOwnerOrganizationName(value *string)()
    SetOwnerUserPrincipalName(value *string)()
    SetTargetGroupIds(value []string)()
}

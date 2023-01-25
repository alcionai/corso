package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateProfileable 
type WindowsDriverUpdateProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalType()(*DriverUpdateProfileApprovalType)
    GetAssignments()([]WindowsDriverUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploymentDeferralInDays()(*int32)
    GetDescription()(*string)
    GetDeviceReporting()(*int32)
    GetDisplayName()(*string)
    GetDriverInventories()([]WindowsDriverUpdateInventoryable)
    GetInventorySyncStatus()(WindowsDriverUpdateProfileInventorySyncStatusable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNewUpdates()(*int32)
    GetRoleScopeTagIds()([]string)
    SetApprovalType(value *DriverUpdateProfileApprovalType)()
    SetAssignments(value []WindowsDriverUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploymentDeferralInDays(value *int32)()
    SetDescription(value *string)()
    SetDeviceReporting(value *int32)()
    SetDisplayName(value *string)()
    SetDriverInventories(value []WindowsDriverUpdateInventoryable)()
    SetInventorySyncStatus(value WindowsDriverUpdateProfileInventorySyncStatusable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNewUpdates(value *int32)()
    SetRoleScopeTagIds(value []string)()
}

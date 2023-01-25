package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCustomAttributeShellScriptable 
type DeviceCustomAttributeShellScriptable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceManagementScriptAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomAttributeName()(*string)
    GetCustomAttributeType()(*DeviceCustomAttributeValueType)
    GetDescription()(*string)
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetDisplayName()(*string)
    GetFileName()(*string)
    GetGroupAssignments()([]DeviceManagementScriptGroupAssignmentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceManagementScriptRunSummaryable)
    GetScriptContent()([]byte)
    GetUserRunStates()([]DeviceManagementScriptUserStateable)
    SetAssignments(value []DeviceManagementScriptAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomAttributeName(value *string)()
    SetCustomAttributeType(value *DeviceCustomAttributeValueType)()
    SetDescription(value *string)()
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetFileName(value *string)()
    SetGroupAssignments(value []DeviceManagementScriptGroupAssignmentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceManagementScriptRunSummaryable)()
    SetScriptContent(value []byte)()
    SetUserRunStates(value []DeviceManagementScriptUserStateable)()
}

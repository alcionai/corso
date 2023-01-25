package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptable 
type DeviceComplianceScriptable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]DeviceHealthScriptAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetectionScriptContent()([]byte)
    GetDeviceRunStates()([]DeviceComplianceScriptDeviceStateable)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublisher()(*string)
    GetRoleScopeTagIds()([]string)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetRunSummary()(DeviceComplianceScriptRunSummaryable)
    GetVersion()(*string)
    SetAssignments(value []DeviceHealthScriptAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetectionScriptContent(value []byte)()
    SetDeviceRunStates(value []DeviceComplianceScriptDeviceStateable)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublisher(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetRunSummary(value DeviceComplianceScriptRunSummaryable)()
    SetVersion(value *string)()
}

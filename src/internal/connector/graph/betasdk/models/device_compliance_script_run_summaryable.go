package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptRunSummaryable 
type DeviceComplianceScriptRunSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetectionScriptErrorDeviceCount()(*int32)
    GetDetectionScriptPendingDeviceCount()(*int32)
    GetIssueDetectedDeviceCount()(*int32)
    GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNoIssueDetectedDeviceCount()(*int32)
    SetDetectionScriptErrorDeviceCount(value *int32)()
    SetDetectionScriptPendingDeviceCount(value *int32)()
    SetIssueDetectedDeviceCount(value *int32)()
    SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNoIssueDetectedDeviceCount(value *int32)()
}

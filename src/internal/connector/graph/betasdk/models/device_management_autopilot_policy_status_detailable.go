package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotPolicyStatusDetailable 
type DeviceManagementAutopilotPolicyStatusDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComplianceStatus()(*DeviceManagementAutopilotPolicyComplianceStatus)
    GetDisplayName()(*string)
    GetErrorCode()(*int32)
    GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyType()(*DeviceManagementAutopilotPolicyType)
    GetTrackedOnEnrollmentStatus()(*bool)
    SetComplianceStatus(value *DeviceManagementAutopilotPolicyComplianceStatus)()
    SetDisplayName(value *string)()
    SetErrorCode(value *int32)()
    SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyType(value *DeviceManagementAutopilotPolicyType)()
    SetTrackedOnEnrollmentStatus(value *bool)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionWipeActionable 
type WindowsInformationProtectionWipeActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastCheckInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*ActionState)
    GetTargetedDeviceMacAddress()(*string)
    GetTargetedDeviceName()(*string)
    GetTargetedDeviceRegistrationId()(*string)
    GetTargetedUserId()(*string)
    SetLastCheckInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *ActionState)()
    SetTargetedDeviceMacAddress(value *string)()
    SetTargetedDeviceName(value *string)()
    SetTargetedDeviceRegistrationId(value *string)()
    SetTargetedUserId(value *string)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagementTaskable 
type DeviceAppManagementTaskable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTo()(*string)
    GetCategory()(*DeviceAppManagementTaskCategory)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreator()(*string)
    GetCreatorNotes()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*DeviceAppManagementTaskPriority)
    GetStatus()(*DeviceAppManagementTaskStatus)
    SetAssignedTo(value *string)()
    SetCategory(value *DeviceAppManagementTaskCategory)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreator(value *string)()
    SetCreatorNotes(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *DeviceAppManagementTaskPriority)()
    SetStatus(value *DeviceAppManagementTaskStatus)()
}

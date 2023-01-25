package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateInventoryable 
type WindowsDriverUpdateInventoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableDeviceCount()(*int32)
    GetApprovalStatus()(*DriverApprovalStatus)
    GetCategory()(*DriverCategory)
    GetDeployDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDriverClass()(*string)
    GetManufacturer()(*string)
    GetName()(*string)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetApplicableDeviceCount(value *int32)()
    SetApprovalStatus(value *DriverApprovalStatus)()
    SetCategory(value *DriverCategory)()
    SetDeployDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDriverClass(value *string)()
    SetManufacturer(value *string)()
    SetName(value *string)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcDeviceImageable 
type CloudPcDeviceImageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperatingSystem()(*string)
    GetOsBuildNumber()(*string)
    GetOsStatus()(*CloudPcDeviceImageOsStatus)
    GetSourceImageResourceId()(*string)
    GetStatus()(*CloudPcDeviceImageStatus)
    GetStatusDetails()(*CloudPcDeviceImageStatusDetails)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperatingSystem(value *string)()
    SetOsBuildNumber(value *string)()
    SetOsStatus(value *CloudPcDeviceImageOsStatus)()
    SetSourceImageResourceId(value *string)()
    SetStatus(value *CloudPcDeviceImageStatus)()
    SetStatusDetails(value *CloudPcDeviceImageStatusDetails)()
    SetVersion(value *string)()
}

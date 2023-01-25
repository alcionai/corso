package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateStateable 
type WindowsUpdateStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceDisplayName()(*string)
    GetDeviceId()(*string)
    GetFeatureUpdateVersion()(*string)
    GetLastScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQualityUpdateVersion()(*string)
    GetStatus()(*WindowsUpdateStatus)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    SetDeviceDisplayName(value *string)()
    SetDeviceId(value *string)()
    SetFeatureUpdateVersion(value *string)()
    SetLastScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQualityUpdateVersion(value *string)()
    SetStatus(value *WindowsUpdateStatus)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
}

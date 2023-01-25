package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallStatusable 
type MobileAppInstallStatusable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(MobileAppable)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDisplayVersion()(*string)
    GetErrorCode()(*int32)
    GetInstallState()(*ResultantAppState)
    GetInstallStateDetail()(*ResultantAppStateDetail)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMobileAppInstallStatusValue()(*ResultantAppState)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetApp(value MobileAppable)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDisplayVersion(value *string)()
    SetErrorCode(value *int32)()
    SetInstallState(value *ResultantAppState)()
    SetInstallStateDetail(value *ResultantAppStateDetail)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMobileAppInstallStatusValue(value *ResultantAppState)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceLogCollectionResponseable 
type DeviceLogCollectionResponseable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnrolledByUser()(*string)
    GetErrorCode()(*int64)
    GetExpirationDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInitiatedByUserPrincipalName()(*string)
    GetManagedDeviceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetReceivedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedDateTimeUTC()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSize()(*float64)
    GetStatus()(*string)
    SetEnrolledByUser(value *string)()
    SetErrorCode(value *int64)()
    SetExpirationDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInitiatedByUserPrincipalName(value *string)()
    SetManagedDeviceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetReceivedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedDateTimeUTC(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSize(value *float64)()
    SetStatus(value *string)()
}

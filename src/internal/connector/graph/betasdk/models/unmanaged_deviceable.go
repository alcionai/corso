package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmanagedDeviceable 
type UnmanagedDeviceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceName()(*string)
    GetDomain()(*string)
    GetIpAddress()(*string)
    GetLastLoggedOnUser()(*string)
    GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocation()(*string)
    GetMacAddress()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetOdataType()(*string)
    GetOs()(*string)
    GetOsVersion()(*string)
    SetDeviceName(value *string)()
    SetDomain(value *string)()
    SetIpAddress(value *string)()
    SetLastLoggedOnUser(value *string)()
    SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocation(value *string)()
    SetMacAddress(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetOdataType(value *string)()
    SetOs(value *string)()
    SetOsVersion(value *string)()
}

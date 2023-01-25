package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Employeeable 
type Employeeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetEmploymentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetGivenName()(*string)
    GetJobTitle()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMiddleName()(*string)
    GetMobilePhone()(*string)
    GetNumber()(*string)
    GetPersonalEmail()(*string)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetStatisticsGroupCode()(*string)
    GetStatus()(*string)
    GetSurname()(*string)
    GetTerminationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAddress(value PostalAddressTypeable)()
    SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetEmploymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetGivenName(value *string)()
    SetJobTitle(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMiddleName(value *string)()
    SetMobilePhone(value *string)()
    SetNumber(value *string)()
    SetPersonalEmail(value *string)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetStatisticsGroupCode(value *string)()
    SetStatus(value *string)()
    SetSurname(value *string)()
    SetTerminationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}

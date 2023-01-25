package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyInformationable 
type CompanyInformationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetCurrencyCode()(*string)
    GetCurrentFiscalYearStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetFaxNumber()(*string)
    GetIndustry()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPhoneNumber()(*string)
    GetPicture()([]byte)
    GetTaxRegistrationNumber()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetCurrencyCode(value *string)()
    SetCurrentFiscalYearStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetFaxNumber(value *string)()
    SetIndustry(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPhoneNumber(value *string)()
    SetPicture(value []byte)()
    SetTaxRegistrationNumber(value *string)()
    SetWebsite(value *string)()
}

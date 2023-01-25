package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Vendor_escapedable 
type Vendor_escapedable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBalance()(*float64)
    GetBlocked()(*string)
    GetCurrency()(Currencyable)
    GetCurrencyCode()(*string)
    GetCurrencyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPaymentMethod()(PaymentMethodable)
    GetPaymentMethodId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetPaymentTerm()(PaymentTermable)
    GetPaymentTermsId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetTaxLiable()(*bool)
    GetTaxRegistrationNumber()(*string)
    GetWebsite()(*string)
    SetAddress(value PostalAddressTypeable)()
    SetBalance(value *float64)()
    SetBlocked(value *string)()
    SetCurrency(value Currencyable)()
    SetCurrencyCode(value *string)()
    SetCurrencyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPaymentMethod(value PaymentMethodable)()
    SetPaymentMethodId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetPaymentTerm(value PaymentTermable)()
    SetPaymentTermsId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetTaxLiable(value *bool)()
    SetTaxRegistrationNumber(value *string)()
    SetWebsite(value *string)()
}

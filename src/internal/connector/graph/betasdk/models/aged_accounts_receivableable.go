package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgedAccountsReceivableable 
type AgedAccountsReceivableable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetBalanceDue()(*float64)
    GetCurrencyCode()(*string)
    GetCurrentAmount()(*float64)
    GetCustomerNumber()(*string)
    GetName()(*string)
    GetPeriod1Amount()(*float64)
    GetPeriod2Amount()(*float64)
    GetPeriod3Amount()(*float64)
    GetPeriodLengthFilter()(*string)
    SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetBalanceDue(value *float64)()
    SetCurrencyCode(value *string)()
    SetCurrentAmount(value *float64)()
    SetCustomerNumber(value *string)()
    SetName(value *string)()
    SetPeriod1Amount(value *float64)()
    SetPeriod2Amount(value *float64)()
    SetPeriod3Amount(value *float64)()
    SetPeriodLengthFilter(value *string)()
}

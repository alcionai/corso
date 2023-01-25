package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesQuoteLineable 
type SalesQuoteLineable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(Accountable)
    GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAmountExcludingTax()(*float64)
    GetAmountIncludingTax()(*float64)
    GetDescription()(*string)
    GetDiscountAmount()(*float64)
    GetDiscountAppliedBeforeTax()(*bool)
    GetDiscountPercent()(*float64)
    GetDocumentId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetItem()(Itemable)
    GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLineType()(*string)
    GetNetAmount()(*float64)
    GetNetAmountIncludingTax()(*float64)
    GetNetTaxAmount()(*float64)
    GetQuantity()(*float64)
    GetSequence()(*int32)
    GetTaxCode()(*string)
    GetTaxPercent()(*float64)
    GetTotalTaxAmount()(*float64)
    GetUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetUnitPrice()(*float64)
    SetAccount(value Accountable)()
    SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAmountExcludingTax(value *float64)()
    SetAmountIncludingTax(value *float64)()
    SetDescription(value *string)()
    SetDiscountAmount(value *float64)()
    SetDiscountAppliedBeforeTax(value *bool)()
    SetDiscountPercent(value *float64)()
    SetDocumentId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetItem(value Itemable)()
    SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLineType(value *string)()
    SetNetAmount(value *float64)()
    SetNetAmountIncludingTax(value *float64)()
    SetNetTaxAmount(value *float64)()
    SetQuantity(value *float64)()
    SetSequence(value *int32)()
    SetTaxCode(value *string)()
    SetTaxPercent(value *float64)()
    SetTotalTaxAmount(value *float64)()
    SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetUnitPrice(value *float64)()
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomerPaymentable 
type CustomerPaymentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmount()(*float64)
    GetAppliesToInvoiceId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetAppliesToInvoiceNumber()(*string)
    GetComment()(*string)
    GetContactId()(*string)
    GetCustomer()(Customerable)
    GetCustomerId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCustomerNumber()(*string)
    GetDescription()(*string)
    GetDocumentNumber()(*string)
    GetExternalDocumentNumber()(*string)
    GetJournalDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLineNumber()(*int32)
    GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAmount(value *float64)()
    SetAppliesToInvoiceId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetAppliesToInvoiceNumber(value *string)()
    SetComment(value *string)()
    SetContactId(value *string)()
    SetCustomer(value Customerable)()
    SetCustomerId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCustomerNumber(value *string)()
    SetDescription(value *string)()
    SetDocumentNumber(value *string)()
    SetExternalDocumentNumber(value *string)()
    SetJournalDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLineNumber(value *int32)()
    SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessScheduleRequestable 
type PrivilegedAccessScheduleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Requestable
    GetAction()(*ScheduleRequestActions)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetScheduleInfo()(RequestScheduleable)
    GetTicketInfo()(TicketInfoable)
    SetAction(value *ScheduleRequestActions)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetScheduleInfo(value RequestScheduleable)()
    SetTicketInfo(value TicketInfoable)()
}

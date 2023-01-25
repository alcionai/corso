package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessageRequestable 
type EventMessageRequestable interface {
    EventMessageable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowNewTimeProposals()(*bool)
    GetPreviousEndDateTime()(DateTimeTimeZoneable)
    GetPreviousLocation()(Locationable)
    GetPreviousStartDateTime()(DateTimeTimeZoneable)
    GetResponseRequested()(*bool)
    SetAllowNewTimeProposals(value *bool)()
    SetPreviousEndDateTime(value DateTimeTimeZoneable)()
    SetPreviousLocation(value Locationable)()
    SetPreviousStartDateTime(value DateTimeTimeZoneable)()
    SetResponseRequested(value *bool)()
}

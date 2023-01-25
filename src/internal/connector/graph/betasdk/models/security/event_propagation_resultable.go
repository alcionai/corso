package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventPropagationResultable 
type EventPropagationResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLocation()(*string)
    GetOdataType()(*string)
    GetServiceName()(*string)
    GetStatus()(*EventPropagationStatus)
    GetStatusInformation()(*string)
    SetLocation(value *string)()
    SetOdataType(value *string)()
    SetServiceName(value *string)()
    SetStatus(value *EventPropagationStatus)()
    SetStatusInformation(value *string)()
}

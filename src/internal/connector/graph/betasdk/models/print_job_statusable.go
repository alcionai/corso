package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintJobStatusable 
type PrintJobStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcquiredByPrinter()(*bool)
    GetDescription()(*string)
    GetDetails()([]PrintJobStateDetail)
    GetIsAcquiredByPrinter()(*bool)
    GetOdataType()(*string)
    GetProcessingState()(*PrintJobProcessingState)
    GetProcessingStateDescription()(*string)
    GetState()(*PrintJobProcessingState)
    SetAcquiredByPrinter(value *bool)()
    SetDescription(value *string)()
    SetDetails(value []PrintJobStateDetail)()
    SetIsAcquiredByPrinter(value *bool)()
    SetOdataType(value *string)()
    SetProcessingState(value *PrintJobProcessingState)()
    SetProcessingStateDescription(value *string)()
    SetState(value *PrintJobProcessingState)()
}

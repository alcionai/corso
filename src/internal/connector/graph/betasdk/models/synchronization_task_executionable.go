package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTaskExecutionable 
type SynchronizationTaskExecutionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityIdentifier()(*string)
    GetCountEntitled()(*int64)
    GetCountEntitledForProvisioning()(*int64)
    GetCountEscrowed()(*int64)
    GetCountEscrowedRaw()(*int64)
    GetCountExported()(*int64)
    GetCountExports()(*int64)
    GetCountImported()(*int64)
    GetCountImportedDeltas()(*int64)
    GetCountImportedReferenceDeltas()(*int64)
    GetError()(SynchronizationErrorable)
    GetOdataType()(*string)
    GetState()(*SynchronizationTaskExecutionResult)
    GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActivityIdentifier(value *string)()
    SetCountEntitled(value *int64)()
    SetCountEntitledForProvisioning(value *int64)()
    SetCountEscrowed(value *int64)()
    SetCountEscrowedRaw(value *int64)()
    SetCountExported(value *int64)()
    SetCountExports(value *int64)()
    SetCountImported(value *int64)()
    SetCountImportedDeltas(value *int64)()
    SetCountImportedReferenceDeltas(value *int64)()
    SetError(value SynchronizationErrorable)()
    SetOdataType(value *string)()
    SetState(value *SynchronizationTaskExecutionResult)()
    SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}

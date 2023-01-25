package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationOperationable 
type AttackSimulationOperationable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPercentageCompleted()(*int32)
    GetTenantId()(*string)
    GetType()(*AttackSimulationOperationType)
    SetPercentageCompleted(value *int32)()
    SetTenantId(value *string)()
    SetType(value *AttackSimulationOperationType)()
}

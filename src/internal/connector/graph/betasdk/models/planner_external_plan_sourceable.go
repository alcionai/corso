package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalPlanSourceable 
type PlannerExternalPlanSourceable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerPlanCreationable
    GetContextScenarioId()(*string)
    GetExternalContextId()(*string)
    GetExternalObjectId()(*string)
    SetContextScenarioId(value *string)()
    SetExternalContextId(value *string)()
    SetExternalObjectId(value *string)()
}

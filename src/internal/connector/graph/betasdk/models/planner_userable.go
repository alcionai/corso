package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerUserable 
type PlannerUserable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetAll()([]PlannerDeltaable)
    GetFavoritePlanReferences()(PlannerFavoritePlanReferenceCollectionable)
    GetFavoritePlans()([]PlannerPlanable)
    GetPlans()([]PlannerPlanable)
    GetRecentPlanReferences()(PlannerRecentPlanReferenceCollectionable)
    GetRecentPlans()([]PlannerPlanable)
    GetRosterPlans()([]PlannerPlanable)
    GetTasks()([]PlannerTaskable)
    SetAll(value []PlannerDeltaable)()
    SetFavoritePlanReferences(value PlannerFavoritePlanReferenceCollectionable)()
    SetFavoritePlans(value []PlannerPlanable)()
    SetPlans(value []PlannerPlanable)()
    SetRecentPlanReferences(value PlannerRecentPlanReferenceCollectionable)()
    SetRecentPlans(value []PlannerPlanable)()
    SetRosterPlans(value []PlannerPlanable)()
    SetTasks(value []PlannerTaskable)()
}

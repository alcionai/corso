package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanable 
type PlannerPlanable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetBuckets()([]PlannerBucketable)
    GetContainer()(PlannerPlanContainerable)
    GetContexts()(PlannerPlanContextCollectionable)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreationSource()(PlannerPlanCreationable)
    GetDetails()(PlannerPlanDetailsable)
    GetOwner()(*string)
    GetTasks()([]PlannerTaskable)
    GetTitle()(*string)
    SetBuckets(value []PlannerBucketable)()
    SetContainer(value PlannerPlanContainerable)()
    SetContexts(value PlannerPlanContextCollectionable)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreationSource(value PlannerPlanCreationable)()
    SetDetails(value PlannerPlanDetailsable)()
    SetOwner(value *string)()
    SetTasks(value []PlannerTaskable)()
    SetTitle(value *string)()
}

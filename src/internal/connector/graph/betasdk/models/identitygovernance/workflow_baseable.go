package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WorkflowBaseable 
type WorkflowBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*LifecycleWorkflowCategory)
    GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExecutionConditions()(WorkflowExecutionConditionsable)
    GetIsEnabled()(*bool)
    GetIsSchedulingEnabled()(*bool)
    GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetTasks()([]Taskable)
    SetCategory(value *LifecycleWorkflowCategory)()
    SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExecutionConditions(value WorkflowExecutionConditionsable)()
    SetIsEnabled(value *bool)()
    SetIsSchedulingEnabled(value *bool)()
    SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetTasks(value []Taskable)()
}

package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Workflowable 
type Workflowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WorkflowBaseable
    GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExecutionScope()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)
    GetId()(*string)
    GetNextScheduleRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRuns()([]Runable)
    GetTaskReports()([]TaskReportable)
    GetUserProcessingResults()([]UserProcessingResultable)
    GetVersion()(*int32)
    GetVersions()([]WorkflowVersionable)
    SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExecutionScope(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)()
    SetId(value *string)()
    SetNextScheduleRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRuns(value []Runable)()
    SetTaskReports(value []TaskReportable)()
    SetUserProcessingResults(value []UserProcessingResultable)()
    SetVersion(value *int32)()
    SetVersions(value []WorkflowVersionable)()
}

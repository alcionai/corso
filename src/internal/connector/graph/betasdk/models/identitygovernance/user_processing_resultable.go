package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// UserProcessingResultable 
type UserProcessingResultable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailedTasksCount()(*int32)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    GetTotalTasksCount()(*int32)
    GetTotalUnprocessedTasksCount()(*int32)
    GetWorkflowExecutionType()(*WorkflowExecutionType)
    GetWorkflowVersion()(*int32)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailedTasksCount(value *int32)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
    SetTotalTasksCount(value *int32)()
    SetTotalUnprocessedTasksCount(value *int32)()
    SetWorkflowExecutionType(value *WorkflowExecutionType)()
    SetWorkflowVersion(value *int32)()
}

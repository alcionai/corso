package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Runable 
type Runable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailedTasksCount()(*int32)
    GetFailedUsersCount()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulUsersCount()(*int32)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    GetTotalTasksCount()(*int32)
    GetTotalUnprocessedTasksCount()(*int32)
    GetTotalUsersCount()(*int32)
    GetUserProcessingResults()([]UserProcessingResultable)
    GetWorkflowExecutionType()(*WorkflowExecutionType)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailedTasksCount(value *int32)()
    SetFailedUsersCount(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulUsersCount(value *int32)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
    SetTotalTasksCount(value *int32)()
    SetTotalUnprocessedTasksCount(value *int32)()
    SetTotalUsersCount(value *int32)()
    SetUserProcessingResults(value []UserProcessingResultable)()
    SetWorkflowExecutionType(value *WorkflowExecutionType)()
}

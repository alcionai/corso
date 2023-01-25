package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TaskReport provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TaskReport struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The date time that the associated run completed. Value is null if the run has not completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of users in the run execution for which the associated task failed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    failedUsersCount *int32
    // The date and time that the task report was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The unique identifier of the associated run.
    runId *string
    // The date time that the associated run started. Value is null if the run has not started.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of users in the run execution for which the associated task succeeded.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    successfulUsersCount *int32
    // The task property
    task Taskable
    // The taskDefinition property
    taskDefinition TaskDefinitionable
    // The related lifecycle workflow taskProcessingResults.
    taskProcessingResults []TaskProcessingResultable
    // The total number of users in the run execution for which the associated task was scheduled to execute.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    totalUsersCount *int32
    // The number of users in the run execution for which the associated task is queued, in progress, or canceled.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    unprocessedUsersCount *int32
}
// NewTaskReport instantiates a new taskReport and sets the default values.
func NewTaskReport()(*TaskReport) {
    m := &TaskReport{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTaskReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskReport(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the associated run completed. Value is null if the run has not completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedUsersCount gets the failedUsersCount property value. The number of users in the run execution for which the associated task failed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) GetFailedUsersCount()(*int32) {
    return m.failedUsersCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["failedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUsersCount(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["processingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleWorkflowProcessingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStatus(val.(*LifecycleWorkflowProcessingStatus))
        }
        return nil
    }
    res["runId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunId(val)
        }
        return nil
    }
    res["startedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedDateTime(val)
        }
        return nil
    }
    res["successfulUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUsersCount(val)
        }
        return nil
    }
    res["task"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTask(val.(Taskable))
        }
        return nil
    }
    res["taskDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTaskDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskDefinition(val.(TaskDefinitionable))
        }
        return nil
    }
    res["taskProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskProcessingResultable, len(val))
            for i, v := range val {
                res[i] = v.(TaskProcessingResultable)
            }
            m.SetTaskProcessingResults(res)
        }
        return nil
    }
    res["totalUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsersCount(val)
        }
        return nil
    }
    res["unprocessedUsersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnprocessedUsersCount(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date and time that the task report was last updated.
func (m *TaskReport) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *TaskReport) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetRunId gets the runId property value. The unique identifier of the associated run.
func (m *TaskReport) GetRunId()(*string) {
    return m.runId
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the associated run started. Value is null if the run has not started.
func (m *TaskReport) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSuccessfulUsersCount gets the successfulUsersCount property value. The number of users in the run execution for which the associated task succeeded.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) GetSuccessfulUsersCount()(*int32) {
    return m.successfulUsersCount
}
// GetTask gets the task property value. The task property
func (m *TaskReport) GetTask()(Taskable) {
    return m.task
}
// GetTaskDefinition gets the taskDefinition property value. The taskDefinition property
func (m *TaskReport) GetTaskDefinition()(TaskDefinitionable) {
    return m.taskDefinition
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The related lifecycle workflow taskProcessingResults.
func (m *TaskReport) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the run execution for which the associated task was scheduled to execute.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) GetTotalUsersCount()(*int32) {
    return m.totalUsersCount
}
// GetUnprocessedUsersCount gets the unprocessedUsersCount property value. The number of users in the run execution for which the associated task is queued, in progress, or canceled.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) GetUnprocessedUsersCount()(*int32) {
    return m.unprocessedUsersCount
}
// Serialize serializes information the current object
func (m *TaskReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUsersCount", m.GetFailedUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err = writer.WriteStringValue("processingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("runId", m.GetRunId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successfulUsersCount", m.GetSuccessfulUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("task", m.GetTask())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("taskDefinition", m.GetTaskDefinition())
        if err != nil {
            return err
        }
    }
    if m.GetTaskProcessingResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskProcessingResults()))
        for i, v := range m.GetTaskProcessingResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUsersCount", m.GetTotalUsersCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unprocessedUsersCount", m.GetUnprocessedUsersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the associated run completed. Value is null if the run has not completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedUsersCount sets the failedUsersCount property value. The number of users in the run execution for which the associated task failed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) SetFailedUsersCount(value *int32)() {
    m.failedUsersCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date and time that the task report was last updated.
func (m *TaskReport) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *TaskReport) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetRunId sets the runId property value. The unique identifier of the associated run.
func (m *TaskReport) SetRunId(value *string)() {
    m.runId = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the associated run started. Value is null if the run has not started.
func (m *TaskReport) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSuccessfulUsersCount sets the successfulUsersCount property value. The number of users in the run execution for which the associated task succeeded.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) SetSuccessfulUsersCount(value *int32)() {
    m.successfulUsersCount = value
}
// SetTask sets the task property value. The task property
func (m *TaskReport) SetTask(value Taskable)() {
    m.task = value
}
// SetTaskDefinition sets the taskDefinition property value. The taskDefinition property
func (m *TaskReport) SetTaskDefinition(value TaskDefinitionable)() {
    m.taskDefinition = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The related lifecycle workflow taskProcessingResults.
func (m *TaskReport) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the run execution for which the associated task was scheduled to execute.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) SetTotalUsersCount(value *int32)() {
    m.totalUsersCount = value
}
// SetUnprocessedUsersCount sets the unprocessedUsersCount property value. The number of users in the run execution for which the associated task is queued, in progress, or canceled.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskReport) SetUnprocessedUsersCount(value *int32)() {
    m.unprocessedUsersCount = value
}

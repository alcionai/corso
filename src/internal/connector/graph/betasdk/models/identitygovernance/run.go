package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Run provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Run struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The date time that the run completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of tasks that failed in the run execution.
    failedTasksCount *int32
    // The number of users that failed in the run execution.
    failedUsersCount *int32
    // The datetime that the run was last updated.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time that the run is scheduled to be executed for a workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    scheduledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time that the run execution started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of successfully completed users in the run.
    successfulUsersCount *int32
    // The related taskProcessingResults.
    taskProcessingResults []TaskProcessingResultable
    // The totalTasksCount property
    totalTasksCount *int32
    // The total number of unprocessed tasks in the run execution.
    totalUnprocessedTasksCount *int32
    // The total number of users in the workflow execution.
    totalUsersCount *int32
    // The associated individual user execution.
    userProcessingResults []UserProcessingResultable
    // The workflowExecutionType property
    workflowExecutionType *WorkflowExecutionType
}
// NewRun instantiates a new run and sets the default values.
func NewRun()(*Run) {
    m := &Run{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRun(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time that the run completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetFailedTasksCount gets the failedTasksCount property value. The number of tasks that failed in the run execution.
func (m *Run) GetFailedTasksCount()(*int32) {
    return m.failedTasksCount
}
// GetFailedUsersCount gets the failedUsersCount property value. The number of users that failed in the run execution.
func (m *Run) GetFailedUsersCount()(*int32) {
    return m.failedUsersCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Run) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["failedTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedTasksCount(val)
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
    res["scheduledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledDateTime(val)
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
    res["totalTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTasksCount(val)
        }
        return nil
    }
    res["totalUnprocessedTasksCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUnprocessedTasksCount(val)
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
    res["userProcessingResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserProcessingResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserProcessingResultable, len(val))
            for i, v := range val {
                res[i] = v.(UserProcessingResultable)
            }
            m.SetUserProcessingResults(res)
        }
        return nil
    }
    res["workflowExecutionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkflowExecutionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowExecutionType(val.(*WorkflowExecutionType))
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The datetime that the run was last updated.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *Run) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetScheduledDateTime gets the scheduledDateTime property value. The date time that the run is scheduled to be executed for a workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.scheduledDateTime
}
// GetStartedDateTime gets the startedDateTime property value. The date time that the run execution started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSuccessfulUsersCount gets the successfulUsersCount property value. The number of successfully completed users in the run.
func (m *Run) GetSuccessfulUsersCount()(*int32) {
    return m.successfulUsersCount
}
// GetTaskProcessingResults gets the taskProcessingResults property value. The related taskProcessingResults.
func (m *Run) GetTaskProcessingResults()([]TaskProcessingResultable) {
    return m.taskProcessingResults
}
// GetTotalTasksCount gets the totalTasksCount property value. The totalTasksCount property
func (m *Run) GetTotalTasksCount()(*int32) {
    return m.totalTasksCount
}
// GetTotalUnprocessedTasksCount gets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks in the run execution.
func (m *Run) GetTotalUnprocessedTasksCount()(*int32) {
    return m.totalUnprocessedTasksCount
}
// GetTotalUsersCount gets the totalUsersCount property value. The total number of users in the workflow execution.
func (m *Run) GetTotalUsersCount()(*int32) {
    return m.totalUsersCount
}
// GetUserProcessingResults gets the userProcessingResults property value. The associated individual user execution.
func (m *Run) GetUserProcessingResults()([]UserProcessingResultable) {
    return m.userProcessingResults
}
// GetWorkflowExecutionType gets the workflowExecutionType property value. The workflowExecutionType property
func (m *Run) GetWorkflowExecutionType()(*WorkflowExecutionType) {
    return m.workflowExecutionType
}
// Serialize serializes information the current object
func (m *Run) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("failedTasksCount", m.GetFailedTasksCount())
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
        err = writer.WriteTimeValue("scheduledDateTime", m.GetScheduledDateTime())
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
        err = writer.WriteInt32Value("totalTasksCount", m.GetTotalTasksCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUnprocessedTasksCount", m.GetTotalUnprocessedTasksCount())
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
    if m.GetUserProcessingResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserProcessingResults()))
        for i, v := range m.GetUserProcessingResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userProcessingResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowExecutionType() != nil {
        cast := (*m.GetWorkflowExecutionType()).String()
        err = writer.WriteStringValue("workflowExecutionType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time that the run completed. Value is null if the workflow hasn't completed.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetFailedTasksCount sets the failedTasksCount property value. The number of tasks that failed in the run execution.
func (m *Run) SetFailedTasksCount(value *int32)() {
    m.failedTasksCount = value
}
// SetFailedUsersCount sets the failedUsersCount property value. The number of users that failed in the run execution.
func (m *Run) SetFailedUsersCount(value *int32)() {
    m.failedUsersCount = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The datetime that the run was last updated.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *Run) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetScheduledDateTime sets the scheduledDateTime property value. The date time that the run is scheduled to be executed for a workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.scheduledDateTime = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time that the run execution started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Run) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSuccessfulUsersCount sets the successfulUsersCount property value. The number of successfully completed users in the run.
func (m *Run) SetSuccessfulUsersCount(value *int32)() {
    m.successfulUsersCount = value
}
// SetTaskProcessingResults sets the taskProcessingResults property value. The related taskProcessingResults.
func (m *Run) SetTaskProcessingResults(value []TaskProcessingResultable)() {
    m.taskProcessingResults = value
}
// SetTotalTasksCount sets the totalTasksCount property value. The totalTasksCount property
func (m *Run) SetTotalTasksCount(value *int32)() {
    m.totalTasksCount = value
}
// SetTotalUnprocessedTasksCount sets the totalUnprocessedTasksCount property value. The total number of unprocessed tasks in the run execution.
func (m *Run) SetTotalUnprocessedTasksCount(value *int32)() {
    m.totalUnprocessedTasksCount = value
}
// SetTotalUsersCount sets the totalUsersCount property value. The total number of users in the workflow execution.
func (m *Run) SetTotalUsersCount(value *int32)() {
    m.totalUsersCount = value
}
// SetUserProcessingResults sets the userProcessingResults property value. The associated individual user execution.
func (m *Run) SetUserProcessingResults(value []UserProcessingResultable)() {
    m.userProcessingResults = value
}
// SetWorkflowExecutionType sets the workflowExecutionType property value. The workflowExecutionType property
func (m *Run) SetWorkflowExecutionType(value *WorkflowExecutionType)() {
    m.workflowExecutionType = value
}

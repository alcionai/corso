package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Workflow 
type Workflow struct {
    WorkflowBase
    // When the workflow was deleted.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Azure AD identity that last modified the workflow object.
    executionScope []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // Identifier used for individually addressing a specific workflow.Supports $filter(eq, ne) and $orderby.
    id *string
    // The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderBy.
    nextScheduleRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The runs property
    runs []Runable
    // Represents the aggregation of task execution data for tasks within a workflow object.
    taskReports []TaskReportable
    // The userProcessingResults property
    userProcessingResults []UserProcessingResultable
    // The current version number of the workflow. Value is 1 when the workflow is first created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    version *int32
    // The workflow versions that are available.
    versions []WorkflowVersionable
}
// NewWorkflow instantiates a new Workflow and sets the default values.
func NewWorkflow()(*Workflow) {
    m := &Workflow{
        WorkflowBase: *NewWorkflowBase(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.workflow";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkflowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflow(), nil
}
// GetDeletedDateTime gets the deletedDateTime property value. When the workflow was deleted.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Workflow) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetExecutionScope gets the executionScope property value. The unique identifier of the Azure AD identity that last modified the workflow object.
func (m *Workflow) GetExecutionScope()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.executionScope
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workflow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowBase.GetFieldDeserializers()
    res["deletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    res["executionScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable, len(val))
            for i, v := range val {
                res[i] = v.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)
            }
            m.SetExecutionScope(res)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["nextScheduleRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextScheduleRunDateTime(val)
        }
        return nil
    }
    res["runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Runable, len(val))
            for i, v := range val {
                res[i] = v.(Runable)
            }
            m.SetRuns(res)
        }
        return nil
    }
    res["taskReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TaskReportable, len(val))
            for i, v := range val {
                res[i] = v.(TaskReportable)
            }
            m.SetTaskReports(res)
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["versions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkflowVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkflowVersionable, len(val))
            for i, v := range val {
                res[i] = v.(WorkflowVersionable)
            }
            m.SetVersions(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Identifier used for individually addressing a specific workflow.Supports $filter(eq, ne) and $orderby.
func (m *Workflow) GetId()(*string) {
    return m.id
}
// GetNextScheduleRunDateTime gets the nextScheduleRunDateTime property value. The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderBy.
func (m *Workflow) GetNextScheduleRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextScheduleRunDateTime
}
// GetRuns gets the runs property value. The runs property
func (m *Workflow) GetRuns()([]Runable) {
    return m.runs
}
// GetTaskReports gets the taskReports property value. Represents the aggregation of task execution data for tasks within a workflow object.
func (m *Workflow) GetTaskReports()([]TaskReportable) {
    return m.taskReports
}
// GetUserProcessingResults gets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) GetUserProcessingResults()([]UserProcessingResultable) {
    return m.userProcessingResults
}
// GetVersion gets the version property value. The current version number of the workflow. Value is 1 when the workflow is first created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Workflow) GetVersion()(*int32) {
    return m.version
}
// GetVersions gets the versions property value. The workflow versions that are available.
func (m *Workflow) GetVersions()([]WorkflowVersionable) {
    return m.versions
}
// Serialize serializes information the current object
func (m *Workflow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetExecutionScope() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExecutionScope()))
        for i, v := range m.GetExecutionScope() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("executionScope", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextScheduleRunDateTime", m.GetNextScheduleRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuns()))
        for i, v := range m.GetRuns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("runs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTaskReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaskReports()))
        for i, v := range m.GetTaskReports() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taskReports", cast)
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
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetVersions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVersions()))
        for i, v := range m.GetVersions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("versions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeletedDateTime sets the deletedDateTime property value. When the workflow was deleted.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Workflow) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetExecutionScope sets the executionScope property value. The unique identifier of the Azure AD identity that last modified the workflow object.
func (m *Workflow) SetExecutionScope(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.executionScope = value
}
// SetId sets the id property value. Identifier used for individually addressing a specific workflow.Supports $filter(eq, ne) and $orderby.
func (m *Workflow) SetId(value *string)() {
    m.id = value
}
// SetNextScheduleRunDateTime sets the nextScheduleRunDateTime property value. The date time when the workflow is expected to run next based on the schedule interval, if there are any users matching the execution conditions. Supports $filter(lt,gt) and $orderBy.
func (m *Workflow) SetNextScheduleRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextScheduleRunDateTime = value
}
// SetRuns sets the runs property value. The runs property
func (m *Workflow) SetRuns(value []Runable)() {
    m.runs = value
}
// SetTaskReports sets the taskReports property value. Represents the aggregation of task execution data for tasks within a workflow object.
func (m *Workflow) SetTaskReports(value []TaskReportable)() {
    m.taskReports = value
}
// SetUserProcessingResults sets the userProcessingResults property value. The userProcessingResults property
func (m *Workflow) SetUserProcessingResults(value []UserProcessingResultable)() {
    m.userProcessingResults = value
}
// SetVersion sets the version property value. The current version number of the workflow. Value is 1 when the workflow is first created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *Workflow) SetVersion(value *int32)() {
    m.version = value
}
// SetVersions sets the versions property value. The workflow versions that are available.
func (m *Workflow) SetVersions(value []WorkflowVersionable)() {
    m.versions = value
}

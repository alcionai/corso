package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TaskProcessingResult provides operations to call the add method.
type TaskProcessingResult struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Describes why the taskProcessingResult has failed.
    failureReason *string
    // The processingStatus property
    processingStatus *LifecycleWorkflowProcessingStatus
    // The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // The task property
    task Taskable
}
// NewTaskProcessingResult instantiates a new taskProcessingResult and sets the default values.
func NewTaskProcessingResult()(*TaskProcessingResult) {
    m := &TaskProcessingResult{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTaskProcessingResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskProcessingResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskProcessingResult(), nil
}
// GetCompletedDateTime gets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFailureReason gets the failureReason property value. Describes why the taskProcessingResult has failed.
func (m *TaskProcessingResult) GetFailureReason()(*string) {
    return m.failureReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskProcessingResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["failureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
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
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable))
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
    return res
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *TaskProcessingResult) GetProcessingStatus()(*LifecycleWorkflowProcessingStatus) {
    return m.processingStatus
}
// GetStartedDateTime gets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *TaskProcessingResult) GetSubject()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.subject
}
// GetTask gets the task property value. The task property
func (m *TaskProcessingResult) GetTask()(Taskable) {
    return m.task
}
// Serialize serializes information the current object
func (m *TaskProcessingResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
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
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subject", m.GetSubject())
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
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. The date time when taskProcessingResult execution ended. Value is null if task execution is still in progress.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date time when the taskProcessingResult was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetFailureReason sets the failureReason property value. Describes why the taskProcessingResult has failed.
func (m *TaskProcessingResult) SetFailureReason(value *string)() {
    m.failureReason = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *TaskProcessingResult) SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)() {
    m.processingStatus = value
}
// SetStartedDateTime sets the startedDateTime property value. The date time when taskProcessingResult execution started. Value is null if task execution has not yet started.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *TaskProcessingResult) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *TaskProcessingResult) SetSubject(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.subject = value
}
// SetTask sets the task property value. The task property
func (m *TaskProcessingResult) SetTask(value Taskable)() {
    m.task = value
}

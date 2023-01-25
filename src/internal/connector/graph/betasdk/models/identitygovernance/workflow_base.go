package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WorkflowBase 
type WorkflowBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The category property
    category *LifecycleWorkflowCategory
    // The user who created the workflow.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // When a workflow was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A string that describes the purpose of the workflow.
    description *string
    // A string to identify the workflow.
    displayName *string
    // Defines when and for who the workflow will run.
    executionConditions WorkflowExecutionConditionsable
    // Whether the workflow is enabled or disabled. If this setting is true, the workflow can be run on demand or on schedule when isSchedulingEnabled is true.
    isEnabled *bool
    // If true, the Lifecycle Workflow engine executes the workflow based on the schedule defined by tenant settings. Cannot be true for a disabled workflow (where isEnabled is false).
    isSchedulingEnabled *bool
    // The user who last modified the workflow.
    lastModifiedBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // When the workflow was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The tasks in the workflow.
    tasks []Taskable
}
// NewWorkflowBase instantiates a new workflowBase and sets the default values.
func NewWorkflowBase()(*WorkflowBase) {
    m := &WorkflowBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkflowBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.identityGovernance.workflow":
                        return NewWorkflow(), nil
                    case "#microsoft.graph.identityGovernance.workflowVersion":
                        return NewWorkflowVersion(), nil
                }
            }
        }
    }
    return NewWorkflowBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkflowBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCategory gets the category property value. The category property
func (m *WorkflowBase) GetCategory()(*LifecycleWorkflowCategory) {
    return m.category
}
// GetCreatedBy gets the createdBy property value. The user who created the workflow.
func (m *WorkflowBase) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When a workflow was created.
func (m *WorkflowBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. A string that describes the purpose of the workflow.
func (m *WorkflowBase) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A string to identify the workflow.
func (m *WorkflowBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionConditions gets the executionConditions property value. Defines when and for who the workflow will run.
func (m *WorkflowBase) GetExecutionConditions()(WorkflowExecutionConditionsable) {
    return m.executionConditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkflowBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLifecycleWorkflowCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*LifecycleWorkflowCategory))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["executionConditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowExecutionConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExecutionConditions(val.(WorkflowExecutionConditionsable))
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isSchedulingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSchedulingEnabled(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Taskable, len(val))
            for i, v := range val {
                res[i] = v.(Taskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Whether the workflow is enabled or disabled. If this setting is true, the workflow can be run on demand or on schedule when isSchedulingEnabled is true.
func (m *WorkflowBase) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsSchedulingEnabled gets the isSchedulingEnabled property value. If true, the Lifecycle Workflow engine executes the workflow based on the schedule defined by tenant settings. Cannot be true for a disabled workflow (where isEnabled is false).
func (m *WorkflowBase) GetIsSchedulingEnabled()(*bool) {
    return m.isSchedulingEnabled
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the workflow.
func (m *WorkflowBase) GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the workflow was last modified.
func (m *WorkflowBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkflowBase) GetOdataType()(*string) {
    return m.odataType
}
// GetTasks gets the tasks property value. The tasks in the workflow.
func (m *WorkflowBase) GetTasks()([]Taskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *WorkflowBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("executionConditions", m.GetExecutionConditions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSchedulingEnabled", m.GetIsSchedulingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkflowBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCategory sets the category property value. The category property
func (m *WorkflowBase) SetCategory(value *LifecycleWorkflowCategory)() {
    m.category = value
}
// SetCreatedBy sets the createdBy property value. The user who created the workflow.
func (m *WorkflowBase) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When a workflow was created.
func (m *WorkflowBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. A string that describes the purpose of the workflow.
func (m *WorkflowBase) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A string to identify the workflow.
func (m *WorkflowBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionConditions sets the executionConditions property value. Defines when and for who the workflow will run.
func (m *WorkflowBase) SetExecutionConditions(value WorkflowExecutionConditionsable)() {
    m.executionConditions = value
}
// SetIsEnabled sets the isEnabled property value. Whether the workflow is enabled or disabled. If this setting is true, the workflow can be run on demand or on schedule when isSchedulingEnabled is true.
func (m *WorkflowBase) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsSchedulingEnabled sets the isSchedulingEnabled property value. If true, the Lifecycle Workflow engine executes the workflow based on the schedule defined by tenant settings. Cannot be true for a disabled workflow (where isEnabled is false).
func (m *WorkflowBase) SetIsSchedulingEnabled(value *bool)() {
    m.isSchedulingEnabled = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the workflow.
func (m *WorkflowBase) SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the workflow was last modified.
func (m *WorkflowBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkflowBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTasks sets the tasks property value. The tasks in the workflow.
func (m *WorkflowBase) SetTasks(value []Taskable)() {
    m.tasks = value
}

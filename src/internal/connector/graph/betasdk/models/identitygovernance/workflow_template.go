package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WorkflowTemplate provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WorkflowTemplate struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The category property
    category *LifecycleWorkflowCategory
    // The description of the workflowTemplate.
    description *string
    // The display name of the workflowTemplate.Supports $filter(eq, ne) and $orderby.
    displayName *string
    // Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
    executionConditions WorkflowExecutionConditionsable
    // Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
    tasks []Taskable
}
// NewWorkflowTemplate instantiates a new workflowTemplate and sets the default values.
func NewWorkflowTemplate()(*WorkflowTemplate) {
    m := &WorkflowTemplate{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateWorkflowTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowTemplate(), nil
}
// GetCategory gets the category property value. The category property
func (m *WorkflowTemplate) GetCategory()(*LifecycleWorkflowCategory) {
    return m.category
}
// GetDescription gets the description property value. The description of the workflowTemplate.
func (m *WorkflowTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the workflowTemplate.Supports $filter(eq, ne) and $orderby.
func (m *WorkflowTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetExecutionConditions gets the executionConditions property value. Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
func (m *WorkflowTemplate) GetExecutionConditions()(WorkflowExecutionConditionsable) {
    return m.executionConditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkflowTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
// GetTasks gets the tasks property value. Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
func (m *WorkflowTemplate) GetTasks()([]Taskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *WorkflowTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("executionConditions", m.GetExecutionConditions())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *WorkflowTemplate) SetCategory(value *LifecycleWorkflowCategory)() {
    m.category = value
}
// SetDescription sets the description property value. The description of the workflowTemplate.
func (m *WorkflowTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the workflowTemplate.Supports $filter(eq, ne) and $orderby.
func (m *WorkflowTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExecutionConditions sets the executionConditions property value. Conditions describing when to execute the workflow and the criteria to identify in-scope subject set.
func (m *WorkflowTemplate) SetExecutionConditions(value WorkflowExecutionConditionsable)() {
    m.executionConditions = value
}
// SetTasks sets the tasks property value. Represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
func (m *WorkflowTemplate) SetTasks(value []Taskable)() {
    m.tasks = value
}

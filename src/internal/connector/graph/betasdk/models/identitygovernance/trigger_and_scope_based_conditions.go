package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TriggerAndScopeBasedConditions 
type TriggerAndScopeBasedConditions struct {
    WorkflowExecutionConditions
    // Defines who the workflow runs for.
    scope ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SubjectSetable
    // What triggers a workflow to run.
    trigger WorkflowExecutionTriggerable
}
// NewTriggerAndScopeBasedConditions instantiates a new TriggerAndScopeBasedConditions and sets the default values.
func NewTriggerAndScopeBasedConditions()(*TriggerAndScopeBasedConditions) {
    m := &TriggerAndScopeBasedConditions{
        WorkflowExecutionConditions: *NewWorkflowExecutionConditions(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.triggerAndScopeBasedConditions";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTriggerAndScopeBasedConditionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerAndScopeBasedConditions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggerAndScopeBasedConditions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionConditions.GetFieldDeserializers()
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SubjectSetable))
        }
        return nil
    }
    res["trigger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkflowExecutionTriggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrigger(val.(WorkflowExecutionTriggerable))
        }
        return nil
    }
    return res
}
// GetScope gets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) GetScope()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SubjectSetable) {
    return m.scope
}
// GetTrigger gets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) GetTrigger()(WorkflowExecutionTriggerable) {
    return m.trigger
}
// Serialize serializes information the current object
func (m *TriggerAndScopeBasedConditions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionConditions.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trigger", m.GetTrigger())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetScope sets the scope property value. Defines who the workflow runs for.
func (m *TriggerAndScopeBasedConditions) SetScope(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SubjectSetable)() {
    m.scope = value
}
// SetTrigger sets the trigger property value. What triggers a workflow to run.
func (m *TriggerAndScopeBasedConditions) SetTrigger(value WorkflowExecutionTriggerable)() {
    m.trigger = value
}

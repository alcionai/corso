package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioTask 
type BusinessScenarioTask struct {
    PlannerTask
    // Scenario-specific properties of the task. externalObjectId and externalBucketId properties must be specified when creating a task.
    businessScenarioProperties BusinessScenarioPropertiesable
    // Target of the task that specifies where the task should be placed. Must be specified when creating a task.
    target BusinessScenarioTaskTargetBaseable
}
// NewBusinessScenarioTask instantiates a new BusinessScenarioTask and sets the default values.
func NewBusinessScenarioTask()(*BusinessScenarioTask) {
    m := &BusinessScenarioTask{
        PlannerTask: *NewPlannerTask(),
    }
    return m
}
// CreateBusinessScenarioTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioTask(), nil
}
// GetBusinessScenarioProperties gets the businessScenarioProperties property value. Scenario-specific properties of the task. externalObjectId and externalBucketId properties must be specified when creating a task.
func (m *BusinessScenarioTask) GetBusinessScenarioProperties()(BusinessScenarioPropertiesable) {
    return m.businessScenarioProperties
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerTask.GetFieldDeserializers()
    res["businessScenarioProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBusinessScenarioPropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBusinessScenarioProperties(val.(BusinessScenarioPropertiesable))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(BusinessScenarioTaskTargetBaseable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. Target of the task that specifies where the task should be placed. Must be specified when creating a task.
func (m *BusinessScenarioTask) GetTarget()(BusinessScenarioTaskTargetBaseable) {
    return m.target
}
// Serialize serializes information the current object
func (m *BusinessScenarioTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerTask.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("businessScenarioProperties", m.GetBusinessScenarioProperties())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessScenarioProperties sets the businessScenarioProperties property value. Scenario-specific properties of the task. externalObjectId and externalBucketId properties must be specified when creating a task.
func (m *BusinessScenarioTask) SetBusinessScenarioProperties(value BusinessScenarioPropertiesable)() {
    m.businessScenarioProperties = value
}
// SetTarget sets the target property value. Target of the task that specifies where the task should be placed. Must be specified when creating a task.
func (m *BusinessScenarioTask) SetTarget(value BusinessScenarioTaskTargetBaseable)() {
    m.target = value
}

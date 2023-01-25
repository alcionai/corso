package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskConfiguration 
type PlannerTaskConfiguration struct {
    Entity
    // Policy configuration for tasks created for the businessScenario when they are being changed outside of the scenario.
    editPolicy PlannerTaskPolicyable
}
// NewPlannerTaskConfiguration instantiates a new plannerTaskConfiguration and sets the default values.
func NewPlannerTaskConfiguration()(*PlannerTaskConfiguration) {
    m := &PlannerTaskConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerTaskConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskConfiguration(), nil
}
// GetEditPolicy gets the editPolicy property value. Policy configuration for tasks created for the businessScenario when they are being changed outside of the scenario.
func (m *PlannerTaskConfiguration) GetEditPolicy()(PlannerTaskPolicyable) {
    return m.editPolicy
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["editPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditPolicy(val.(PlannerTaskPolicyable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerTaskConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("editPolicy", m.GetEditPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEditPolicy sets the editPolicy property value. Policy configuration for tasks created for the businessScenario when they are being changed outside of the scenario.
func (m *PlannerTaskConfiguration) SetEditPolicy(value PlannerTaskPolicyable)() {
    m.editPolicy = value
}

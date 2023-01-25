package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerExternalPlanSource 
type PlannerExternalPlanSource struct {
    PlannerPlanCreation
    // Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
    contextScenarioId *string
    // Nullable. The id of the external entity's containing entity or context.
    externalContextId *string
    // Nullable. The id of the entity that an external service associates with a plan.
    externalObjectId *string
}
// NewPlannerExternalPlanSource instantiates a new PlannerExternalPlanSource and sets the default values.
func NewPlannerExternalPlanSource()(*PlannerExternalPlanSource) {
    m := &PlannerExternalPlanSource{
        PlannerPlanCreation: *NewPlannerPlanCreation(),
    }
    odataTypeValue := "#microsoft.graph.plannerExternalPlanSource";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePlannerExternalPlanSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerExternalPlanSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerExternalPlanSource(), nil
}
// GetContextScenarioId gets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
func (m *PlannerExternalPlanSource) GetContextScenarioId()(*string) {
    return m.contextScenarioId
}
// GetExternalContextId gets the externalContextId property value. Nullable. The id of the external entity's containing entity or context.
func (m *PlannerExternalPlanSource) GetExternalContextId()(*string) {
    return m.externalContextId
}
// GetExternalObjectId gets the externalObjectId property value. Nullable. The id of the entity that an external service associates with a plan.
func (m *PlannerExternalPlanSource) GetExternalObjectId()(*string) {
    return m.externalObjectId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerExternalPlanSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerPlanCreation.GetFieldDeserializers()
    res["contextScenarioId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContextScenarioId(val)
        }
        return nil
    }
    res["externalContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalContextId(val)
        }
        return nil
    }
    res["externalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PlannerExternalPlanSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerPlanCreation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contextScenarioId", m.GetContextScenarioId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalContextId", m.GetExternalContextId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalObjectId", m.GetExternalObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContextScenarioId sets the contextScenarioId property value. Nullable. An identifier for the scenario associated with this external source. This should be in reverse DNS format. For example, Contoso company owned application for customer support would have a value like 'com.constoso.customerSupport'.
func (m *PlannerExternalPlanSource) SetContextScenarioId(value *string)() {
    m.contextScenarioId = value
}
// SetExternalContextId sets the externalContextId property value. Nullable. The id of the external entity's containing entity or context.
func (m *PlannerExternalPlanSource) SetExternalContextId(value *string)() {
    m.externalContextId = value
}
// SetExternalObjectId sets the externalObjectId property value. Nullable. The id of the entity that an external service associates with a plan.
func (m *PlannerExternalPlanSource) SetExternalObjectId(value *string)() {
    m.externalObjectId = value
}

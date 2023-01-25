package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskPolicy 
type PlannerTaskPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The rules that should be enforced on the tasks when they are being changed outside of the scenario, based on the role of the caller.
    rules []PlannerTaskRoleBasedRuleable
}
// NewPlannerTaskPolicy instantiates a new plannerTaskPolicy and sets the default values.
func NewPlannerTaskPolicy()(*PlannerTaskPolicy) {
    m := &PlannerTaskPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerTaskPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskPolicy(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskPolicy) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTaskRoleBasedRuleable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerTaskRoleBasedRuleable)
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskPolicy) GetOdataType()(*string) {
    return m.odataType
}
// GetRules gets the rules property value. The rules that should be enforced on the tasks when they are being changed outside of the scenario, based on the role of the caller.
func (m *PlannerTaskPolicy) GetRules()([]PlannerTaskRoleBasedRuleable) {
    return m.rules
}
// Serialize serializes information the current object
func (m *PlannerTaskPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRules()))
        for i, v := range m.GetRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("rules", cast)
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
func (m *PlannerTaskPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskPolicy) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRules sets the rules property value. The rules that should be enforced on the tasks when they are being changed outside of the scenario, based on the role of the caller.
func (m *PlannerTaskPolicy) SetRules(value []PlannerTaskRoleBasedRuleable)() {
    m.rules = value
}

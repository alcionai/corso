package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskRoleBasedRule 
type PlannerTaskRoleBasedRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
    defaultRule *string
    // The OdataType property
    odataType *string
    // Rules for specific properties and actions.
    propertyRule PlannerTaskPropertyRuleable
    // The role these rules apply to.
    role PlannerTaskConfigurationRoleBaseable
}
// NewPlannerTaskRoleBasedRule instantiates a new plannerTaskRoleBasedRule and sets the default values.
func NewPlannerTaskRoleBasedRule()(*PlannerTaskRoleBasedRule) {
    m := &PlannerTaskRoleBasedRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskRoleBasedRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskRoleBasedRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskRoleBasedRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultRule gets the defaultRule property value. Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
func (m *PlannerTaskRoleBasedRule) GetDefaultRule()(*string) {
    return m.defaultRule
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskRoleBasedRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRule(val)
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
    res["propertyRule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskPropertyRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyRule(val.(PlannerTaskPropertyRuleable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTaskConfigurationRoleBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(PlannerTaskConfigurationRoleBaseable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskRoleBasedRule) GetOdataType()(*string) {
    return m.odataType
}
// GetPropertyRule gets the propertyRule property value. Rules for specific properties and actions.
func (m *PlannerTaskRoleBasedRule) GetPropertyRule()(PlannerTaskPropertyRuleable) {
    return m.propertyRule
}
// GetRole gets the role property value. The role these rules apply to.
func (m *PlannerTaskRoleBasedRule) GetRole()(PlannerTaskConfigurationRoleBaseable) {
    return m.role
}
// Serialize serializes information the current object
func (m *PlannerTaskRoleBasedRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultRule", m.GetDefaultRule())
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
    {
        err := writer.WriteObjectValue("propertyRule", m.GetPropertyRule())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("role", m.GetRole())
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
func (m *PlannerTaskRoleBasedRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultRule sets the defaultRule property value. Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
func (m *PlannerTaskRoleBasedRule) SetDefaultRule(value *string)() {
    m.defaultRule = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskRoleBasedRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPropertyRule sets the propertyRule property value. Rules for specific properties and actions.
func (m *PlannerTaskRoleBasedRule) SetPropertyRule(value PlannerTaskPropertyRuleable)() {
    m.propertyRule = value
}
// SetRole sets the role property value. The role these rules apply to.
func (m *PlannerTaskRoleBasedRule) SetRole(value PlannerTaskConfigurationRoleBaseable)() {
    m.role = value
}

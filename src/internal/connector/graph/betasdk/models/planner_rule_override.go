package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerRuleOverride 
type PlannerRuleOverride struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the override. Allowed override values will be dependent on the property affected by the rule.
    name *string
    // The OdataType property
    odataType *string
    // Overridden rules. These are used as rules for the override instead of the default rules.
    rules []string
}
// NewPlannerRuleOverride instantiates a new plannerRuleOverride and sets the default values.
func NewPlannerRuleOverride()(*PlannerRuleOverride) {
    m := &PlannerRuleOverride{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerRuleOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerRuleOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerRuleOverride(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerRuleOverride) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerRuleOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["rules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRules(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the override. Allowed override values will be dependent on the property affected by the rule.
func (m *PlannerRuleOverride) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerRuleOverride) GetOdataType()(*string) {
    return m.odataType
}
// GetRules gets the rules property value. Overridden rules. These are used as rules for the override instead of the default rules.
func (m *PlannerRuleOverride) GetRules()([]string) {
    return m.rules
}
// Serialize serializes information the current object
func (m *PlannerRuleOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetRules() != nil {
        err := writer.WriteCollectionOfStringValues("rules", m.GetRules())
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
func (m *PlannerRuleOverride) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the override. Allowed override values will be dependent on the property affected by the rule.
func (m *PlannerRuleOverride) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerRuleOverride) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRules sets the rules property value. Overridden rules. These are used as rules for the override instead of the default rules.
func (m *PlannerRuleOverride) SetRules(value []string)() {
    m.rules = value
}

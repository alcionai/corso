package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRuleSetting 
type GovernanceRuleSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The id of the rule. For example, ExpirationRule and MfaRule.
    ruleIdentifier *string
    // The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
    setting *string
}
// NewGovernanceRuleSetting instantiates a new governanceRuleSetting and sets the default values.
func NewGovernanceRuleSetting()(*GovernanceRuleSetting) {
    m := &GovernanceRuleSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceRuleSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRuleSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRuleSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRuleSetting) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRuleSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["ruleIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleIdentifier(val)
        }
        return nil
    }
    res["setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernanceRuleSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleIdentifier gets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
func (m *GovernanceRuleSetting) GetRuleIdentifier()(*string) {
    return m.ruleIdentifier
}
// GetSetting gets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
func (m *GovernanceRuleSetting) GetSetting()(*string) {
    return m.setting
}
// Serialize serializes information the current object
func (m *GovernanceRuleSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleIdentifier", m.GetRuleIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("setting", m.GetSetting())
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
func (m *GovernanceRuleSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernanceRuleSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleIdentifier sets the ruleIdentifier property value. The id of the rule. For example, ExpirationRule and MfaRule.
func (m *GovernanceRuleSetting) SetRuleIdentifier(value *string)() {
    m.ruleIdentifier = value
}
// SetSetting sets the setting property value. The settings of the rule. The value is a JSON string with a list of pairs in the format of Parameter_Name:Parameter_Value. For example, {'permanentAssignment':false,'maximumGrantPeriodInMinutes':129600}
func (m *GovernanceRuleSetting) SetSetting(value *string)() {
    m.setting = value
}

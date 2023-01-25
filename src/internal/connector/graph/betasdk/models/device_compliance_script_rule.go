package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptRule 
type DeviceComplianceScriptRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Data types for rules.
    dataType *DataType
    // Data types for rules.
    deviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDataType
    // Operator for rules.
    deviceComplianceScriptRulOperator *DeviceComplianceScriptRulOperator
    // The OdataType property
    odataType *string
    // Operand specified in the rule.
    operand *string
    // Operator for rules.
    operator *Operator
    // Setting name specified in the rule.
    settingName *string
}
// NewDeviceComplianceScriptRule instantiates a new deviceComplianceScriptRule and sets the default values.
func NewDeviceComplianceScriptRule()(*DeviceComplianceScriptRule) {
    m := &DeviceComplianceScriptRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceComplianceScriptRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptRule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDataType gets the dataType property value. Data types for rules.
func (m *DeviceComplianceScriptRule) GetDataType()(*DataType) {
    return m.dataType
}
// GetDeviceComplianceScriptRuleDataType gets the deviceComplianceScriptRuleDataType property value. Data types for rules.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType) {
    return m.deviceComplianceScriptRuleDataType
}
// GetDeviceComplianceScriptRulOperator gets the deviceComplianceScriptRulOperator property value. Operator for rules.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator) {
    return m.deviceComplianceScriptRulOperator
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*DataType))
        }
        return nil
    }
    res["deviceComplianceScriptRuleDataType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRuleDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptRuleDataType(val.(*DeviceComplianceScriptRuleDataType))
        }
        return nil
    }
    res["deviceComplianceScriptRulOperator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptRulOperator(val.(*DeviceComplianceScriptRulOperator))
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
    res["operand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperand(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Operator))
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptRule) GetOdataType()(*string) {
    return m.odataType
}
// GetOperand gets the operand property value. Operand specified in the rule.
func (m *DeviceComplianceScriptRule) GetOperand()(*string) {
    return m.operand
}
// GetOperator gets the operator property value. Operator for rules.
func (m *DeviceComplianceScriptRule) GetOperator()(*Operator) {
    return m.operator
}
// GetSettingName gets the settingName property value. Setting name specified in the rule.
func (m *DeviceComplianceScriptRule) GetSettingName()(*string) {
    return m.settingName
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := (*m.GetDataType()).String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRuleDataType() != nil {
        cast := (*m.GetDeviceComplianceScriptRuleDataType()).String()
        err := writer.WriteStringValue("deviceComplianceScriptRuleDataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulOperator() != nil {
        cast := (*m.GetDeviceComplianceScriptRulOperator()).String()
        err := writer.WriteStringValue("deviceComplianceScriptRulOperator", &cast)
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
        err := writer.WriteStringValue("operand", m.GetOperand())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
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
func (m *DeviceComplianceScriptRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDataType sets the dataType property value. Data types for rules.
func (m *DeviceComplianceScriptRule) SetDataType(value *DataType)() {
    m.dataType = value
}
// SetDeviceComplianceScriptRuleDataType sets the deviceComplianceScriptRuleDataType property value. Data types for rules.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)() {
    m.deviceComplianceScriptRuleDataType = value
}
// SetDeviceComplianceScriptRulOperator sets the deviceComplianceScriptRulOperator property value. Operator for rules.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)() {
    m.deviceComplianceScriptRulOperator = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptRule) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperand sets the operand property value. Operand specified in the rule.
func (m *DeviceComplianceScriptRule) SetOperand(value *string)() {
    m.operand = value
}
// SetOperator sets the operator property value. Operator for rules.
func (m *DeviceComplianceScriptRule) SetOperator(value *Operator)() {
    m.operator = value
}
// SetSettingName sets the settingName property value. Setting name specified in the rule.
func (m *DeviceComplianceScriptRule) SetSettingName(value *string)() {
    m.settingName = value
}

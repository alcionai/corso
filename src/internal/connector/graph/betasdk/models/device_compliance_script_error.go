package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptError 
type DeviceComplianceScriptError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Error code for rule validation.
    code *Code
    // Error code for rule validation.
    deviceComplianceScriptRulesValidationError *DeviceComplianceScriptRulesValidationError
    // Error message.
    message *string
    // The OdataType property
    odataType *string
}
// NewDeviceComplianceScriptError instantiates a new deviceComplianceScriptError and sets the default values.
func NewDeviceComplianceScriptError()(*DeviceComplianceScriptError) {
    m := &DeviceComplianceScriptError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceComplianceScriptErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceComplianceScriptRuleError":
                        return NewDeviceComplianceScriptRuleError(), nil
                }
            }
        }
    }
    return NewDeviceComplianceScriptError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptError) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCode gets the code property value. Error code for rule validation.
func (m *DeviceComplianceScriptError) GetCode()(*Code) {
    return m.code
}
// GetDeviceComplianceScriptRulesValidationError gets the deviceComplianceScriptRulesValidationError property value. Error code for rule validation.
func (m *DeviceComplianceScriptError) GetDeviceComplianceScriptRulesValidationError()(*DeviceComplianceScriptRulesValidationError) {
    return m.deviceComplianceScriptRulesValidationError
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val.(*Code))
        }
        return nil
    }
    res["deviceComplianceScriptRulesValidationError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulesValidationError)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptRulesValidationError(val.(*DeviceComplianceScriptRulesValidationError))
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    return res
}
// GetMessage gets the message property value. Error message.
func (m *DeviceComplianceScriptError) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptError) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := (*m.GetCode()).String()
        err := writer.WriteStringValue("code", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulesValidationError() != nil {
        cast := (*m.GetDeviceComplianceScriptRulesValidationError()).String()
        err := writer.WriteStringValue("deviceComplianceScriptRulesValidationError", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCode sets the code property value. Error code for rule validation.
func (m *DeviceComplianceScriptError) SetCode(value *Code)() {
    m.code = value
}
// SetDeviceComplianceScriptRulesValidationError sets the deviceComplianceScriptRulesValidationError property value. Error code for rule validation.
func (m *DeviceComplianceScriptError) SetDeviceComplianceScriptRulesValidationError(value *DeviceComplianceScriptRulesValidationError)() {
    m.deviceComplianceScriptRulesValidationError = value
}
// SetMessage sets the message property value. Error message.
func (m *DeviceComplianceScriptError) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceComplianceScriptError) SetOdataType(value *string)() {
    m.odataType = value
}

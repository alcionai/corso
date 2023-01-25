package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementApplicabilityRuleDeviceMode 
type DeviceManagementApplicabilityRuleDeviceMode struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Windows 10 Device Mode type.
    deviceMode *Windows10DeviceModeType
    // Name for object.
    name *string
    // The OdataType property
    odataType *string
    // Supported Applicability rule types for Device Configuration
    ruleType *DeviceManagementApplicabilityRuleType
}
// NewDeviceManagementApplicabilityRuleDeviceMode instantiates a new deviceManagementApplicabilityRuleDeviceMode and sets the default values.
func NewDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    m := &DeviceManagementApplicabilityRuleDeviceMode{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementApplicabilityRuleDeviceMode(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceMode gets the deviceMode property value. Windows 10 Device Mode type.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetDeviceMode()(*Windows10DeviceModeType) {
    return m.deviceMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10DeviceModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMode(val.(*Windows10DeviceModeType))
        }
        return nil
    }
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
    res["ruleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementApplicabilityRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val.(*DeviceManagementApplicabilityRuleType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleType gets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    return m.ruleType
}
// Serialize serializes information the current object
func (m *DeviceManagementApplicabilityRuleDeviceMode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceMode() != nil {
        cast := (*m.GetDeviceMode()).String()
        err := writer.WriteStringValue("deviceMode", &cast)
        if err != nil {
            return err
        }
    }
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
    if m.GetRuleType() != nil {
        cast := (*m.GetRuleType()).String()
        err := writer.WriteStringValue("ruleType", &cast)
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
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceMode sets the deviceMode property value. Windows 10 Device Mode type.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetDeviceMode(value *Windows10DeviceModeType)() {
    m.deviceMode = value
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleType sets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}

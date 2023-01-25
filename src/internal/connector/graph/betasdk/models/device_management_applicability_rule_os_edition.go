package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementApplicabilityRuleOsEdition 
type DeviceManagementApplicabilityRuleOsEdition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name for object.
    name *string
    // The OdataType property
    odataType *string
    // Applicability rule OS edition type.
    osEditionTypes []Windows10EditionType
    // Supported Applicability rule types for Device Configuration
    ruleType *DeviceManagementApplicabilityRuleType
}
// NewDeviceManagementApplicabilityRuleOsEdition instantiates a new deviceManagementApplicabilityRuleOsEdition and sets the default values.
func NewDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    m := &DeviceManagementApplicabilityRuleOsEdition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementApplicabilityRuleOsEditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementApplicabilityRuleOsEditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementApplicabilityRuleOsEdition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleOsEdition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["osEditionTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseWindows10EditionType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10EditionType, len(val))
            for i, v := range val {
                res[i] = *(v.(*Windows10EditionType))
            }
            m.SetOsEditionTypes(res)
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
func (m *DeviceManagementApplicabilityRuleOsEdition) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleOsEdition) GetOdataType()(*string) {
    return m.odataType
}
// GetOsEditionTypes gets the osEditionTypes property value. Applicability rule OS edition type.
func (m *DeviceManagementApplicabilityRuleOsEdition) GetOsEditionTypes()([]Windows10EditionType) {
    return m.osEditionTypes
}
// GetRuleType gets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleOsEdition) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    return m.ruleType
}
// Serialize serializes information the current object
func (m *DeviceManagementApplicabilityRuleOsEdition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetOsEditionTypes() != nil {
        err := writer.WriteCollectionOfStringValues("osEditionTypes", SerializeWindows10EditionType(m.GetOsEditionTypes()))
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
func (m *DeviceManagementApplicabilityRuleOsEdition) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleOsEdition) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOsEditionTypes sets the osEditionTypes property value. Applicability rule OS edition type.
func (m *DeviceManagementApplicabilityRuleOsEdition) SetOsEditionTypes(value []Windows10EditionType)() {
    m.osEditionTypes = value
}
// SetRuleType sets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleOsEdition) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}

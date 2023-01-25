package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeDeviceClass device Class in Exchange.
type DeviceManagementExchangeDeviceClass struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the device class which will be impacted by this rule.
    name *string
    // The OdataType property
    odataType *string
    // Criteria which defines the type of device this access rule will apply to
    type_escaped *DeviceManagementExchangeAccessRuleType
}
// NewDeviceManagementExchangeDeviceClass instantiates a new deviceManagementExchangeDeviceClass and sets the default values.
func NewDeviceManagementExchangeDeviceClass()(*DeviceManagementExchangeDeviceClass) {
    m := &DeviceManagementExchangeDeviceClass{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementExchangeDeviceClassFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeDeviceClassFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExchangeDeviceClass(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeDeviceClass) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeDeviceClass) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*DeviceManagementExchangeAccessRuleType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the device class which will be impacted by this rule.
func (m *DeviceManagementExchangeDeviceClass) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementExchangeDeviceClass) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Criteria which defines the type of device this access rule will apply to
func (m *DeviceManagementExchangeDeviceClass) GetType()(*DeviceManagementExchangeAccessRuleType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeDeviceClass) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *DeviceManagementExchangeDeviceClass) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. Name of the device class which will be impacted by this rule.
func (m *DeviceManagementExchangeDeviceClass) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementExchangeDeviceClass) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Criteria which defines the type of device this access rule will apply to
func (m *DeviceManagementExchangeDeviceClass) SetType(value *DeviceManagementExchangeAccessRuleType)() {
    m.type_escaped = value
}

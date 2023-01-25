package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementExchangeAccessRule device Access Rules in Exchange.
type DeviceManagementExchangeAccessRule struct {
    // Access Level in Exchange.
    accessLevel *DeviceManagementExchangeAccessLevel
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Device Class which will be impacted by this rule.
    deviceClass DeviceManagementExchangeDeviceClassable
    // The OdataType property
    odataType *string
}
// NewDeviceManagementExchangeAccessRule instantiates a new deviceManagementExchangeAccessRule and sets the default values.
func NewDeviceManagementExchangeAccessRule()(*DeviceManagementExchangeAccessRule) {
    m := &DeviceManagementExchangeAccessRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementExchangeAccessRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementExchangeAccessRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementExchangeAccessRule(), nil
}
// GetAccessLevel gets the accessLevel property value. Access Level in Exchange.
func (m *DeviceManagementExchangeAccessRule) GetAccessLevel()(*DeviceManagementExchangeAccessLevel) {
    return m.accessLevel
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeAccessRule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceClass gets the deviceClass property value. Device Class which will be impacted by this rule.
func (m *DeviceManagementExchangeAccessRule) GetDeviceClass()(DeviceManagementExchangeDeviceClassable) {
    return m.deviceClass
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementExchangeAccessRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessLevel(val.(*DeviceManagementExchangeAccessLevel))
        }
        return nil
    }
    res["deviceClass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementExchangeDeviceClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceClass(val.(DeviceManagementExchangeDeviceClassable))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementExchangeAccessRule) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementExchangeAccessRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAccessLevel() != nil {
        cast := (*m.GetAccessLevel()).String()
        err := writer.WriteStringValue("accessLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceClass", m.GetDeviceClass())
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
// SetAccessLevel sets the accessLevel property value. Access Level in Exchange.
func (m *DeviceManagementExchangeAccessRule) SetAccessLevel(value *DeviceManagementExchangeAccessLevel)() {
    m.accessLevel = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementExchangeAccessRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceClass sets the deviceClass property value. Device Class which will be impacted by this rule.
func (m *DeviceManagementExchangeAccessRule) SetDeviceClass(value DeviceManagementExchangeDeviceClassable)() {
    m.deviceClass = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementExchangeAccessRule) SetOdataType(value *string)() {
    m.odataType = value
}

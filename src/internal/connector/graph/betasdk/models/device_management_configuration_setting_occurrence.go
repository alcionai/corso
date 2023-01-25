package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingOccurrence 
type DeviceManagementConfigurationSettingOccurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Maximum times setting can be set on device.
    maxDeviceOccurrence *int32
    // Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
    minDeviceOccurrence *int32
    // The OdataType property
    odataType *string
}
// NewDeviceManagementConfigurationSettingOccurrence instantiates a new deviceManagementConfigurationSettingOccurrence and sets the default values.
func NewDeviceManagementConfigurationSettingOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    m := &DeviceManagementConfigurationSettingOccurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationSettingOccurrenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingOccurrenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSettingOccurrence(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingOccurrence) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingOccurrence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxDeviceOccurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxDeviceOccurrence(val)
        }
        return nil
    }
    res["minDeviceOccurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinDeviceOccurrence(val)
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
// GetMaxDeviceOccurrence gets the maxDeviceOccurrence property value. Maximum times setting can be set on device.
func (m *DeviceManagementConfigurationSettingOccurrence) GetMaxDeviceOccurrence()(*int32) {
    return m.maxDeviceOccurrence
}
// GetMinDeviceOccurrence gets the minDeviceOccurrence property value. Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
func (m *DeviceManagementConfigurationSettingOccurrence) GetMinDeviceOccurrence()(*int32) {
    return m.minDeviceOccurrence
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingOccurrence) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingOccurrence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maxDeviceOccurrence", m.GetMaxDeviceOccurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minDeviceOccurrence", m.GetMinDeviceOccurrence())
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
func (m *DeviceManagementConfigurationSettingOccurrence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMaxDeviceOccurrence sets the maxDeviceOccurrence property value. Maximum times setting can be set on device.
func (m *DeviceManagementConfigurationSettingOccurrence) SetMaxDeviceOccurrence(value *int32)() {
    m.maxDeviceOccurrence = value
}
// SetMinDeviceOccurrence sets the minDeviceOccurrence property value. Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
func (m *DeviceManagementConfigurationSettingOccurrence) SetMinDeviceOccurrence(value *int32)() {
    m.minDeviceOccurrence = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationSettingOccurrence) SetOdataType(value *string)() {
    m.odataType = value
}

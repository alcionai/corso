package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceCleanupSettings define the rule when the admin wants the devices to be cleaned up.
type ManagedDeviceCleanupSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Number of days when the device has not contacted Intune.
    deviceInactivityBeforeRetirementInDays *string
    // The OdataType property
    odataType *string
}
// NewManagedDeviceCleanupSettings instantiates a new managedDeviceCleanupSettings and sets the default values.
func NewManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    m := &ManagedDeviceCleanupSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagedDeviceCleanupSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceCleanupSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceCleanupSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceCleanupSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceInactivityBeforeRetirementInDays gets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
func (m *ManagedDeviceCleanupSettings) GetDeviceInactivityBeforeRetirementInDays()(*string) {
    return m.deviceInactivityBeforeRetirementInDays
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceCleanupSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceInactivityBeforeRetirementInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDays(val)
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
func (m *ManagedDeviceCleanupSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ManagedDeviceCleanupSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceInactivityBeforeRetirementInDays", m.GetDeviceInactivityBeforeRetirementInDays())
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
func (m *ManagedDeviceCleanupSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceInactivityBeforeRetirementInDays sets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
func (m *ManagedDeviceCleanupSettings) SetDeviceInactivityBeforeRetirementInDays(value *string)() {
    m.deviceInactivityBeforeRetirementInDays = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDeviceCleanupSettings) SetOdataType(value *string)() {
    m.odataType = value
}

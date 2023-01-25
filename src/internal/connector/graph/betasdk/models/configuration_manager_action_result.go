package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerActionResult 
type ConfigurationManagerActionResult struct {
    DeviceActionResult
    // Delivery state of Configuration Manager device action
    actionDeliveryStatus *ConfigurationManagerActionDeliveryStatus
    // Error code of Configuration Manager action from client
    errorCode *int32
}
// NewConfigurationManagerActionResult instantiates a new ConfigurationManagerActionResult and sets the default values.
func NewConfigurationManagerActionResult()(*ConfigurationManagerActionResult) {
    m := &ConfigurationManagerActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateConfigurationManagerActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerActionResult(), nil
}
// GetActionDeliveryStatus gets the actionDeliveryStatus property value. Delivery state of Configuration Manager device action
func (m *ConfigurationManagerActionResult) GetActionDeliveryStatus()(*ConfigurationManagerActionDeliveryStatus) {
    return m.actionDeliveryStatus
}
// GetErrorCode gets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["actionDeliveryStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConfigurationManagerActionDeliveryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionDeliveryStatus(val.(*ConfigurationManagerActionDeliveryStatus))
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConfigurationManagerActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionDeliveryStatus() != nil {
        cast := (*m.GetActionDeliveryStatus()).String()
        err = writer.WriteStringValue("actionDeliveryStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionDeliveryStatus sets the actionDeliveryStatus property value. Delivery state of Configuration Manager device action
func (m *ConfigurationManagerActionResult) SetActionDeliveryStatus(value *ConfigurationManagerActionDeliveryStatus)() {
    m.actionDeliveryStatus = value
}
// SetErrorCode sets the errorCode property value. Error code of Configuration Manager action from client
func (m *ConfigurationManagerActionResult) SetErrorCode(value *int32)() {
    m.errorCode = value
}

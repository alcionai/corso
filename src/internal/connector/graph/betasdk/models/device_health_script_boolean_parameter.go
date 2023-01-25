package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptBooleanParameter 
type DeviceHealthScriptBooleanParameter struct {
    DeviceHealthScriptParameter
    // The default value of boolean param
    defaultValue *bool
}
// NewDeviceHealthScriptBooleanParameter instantiates a new DeviceHealthScriptBooleanParameter and sets the default values.
func NewDeviceHealthScriptBooleanParameter()(*DeviceHealthScriptBooleanParameter) {
    m := &DeviceHealthScriptBooleanParameter{
        DeviceHealthScriptParameter: *NewDeviceHealthScriptParameter(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptBooleanParameter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptBooleanParameterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptBooleanParameterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptBooleanParameter(), nil
}
// GetDefaultValue gets the defaultValue property value. The default value of boolean param
func (m *DeviceHealthScriptBooleanParameter) GetDefaultValue()(*bool) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptBooleanParameter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptParameter.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptBooleanParameter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptParameter.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. The default value of boolean param
func (m *DeviceHealthScriptBooleanParameter) SetDefaultValue(value *bool)() {
    m.defaultValue = value
}

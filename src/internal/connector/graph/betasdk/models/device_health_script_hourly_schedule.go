package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptHourlySchedule 
type DeviceHealthScriptHourlySchedule struct {
    DeviceHealthScriptRunSchedule
}
// NewDeviceHealthScriptHourlySchedule instantiates a new DeviceHealthScriptHourlySchedule and sets the default values.
func NewDeviceHealthScriptHourlySchedule()(*DeviceHealthScriptHourlySchedule) {
    m := &DeviceHealthScriptHourlySchedule{
        DeviceHealthScriptRunSchedule: *NewDeviceHealthScriptRunSchedule(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptHourlySchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptHourlyScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptHourlyScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptHourlySchedule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptHourlySchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptRunSchedule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptHourlySchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptRunSchedule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

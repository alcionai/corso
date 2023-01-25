package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptDailySchedule 
type DeviceHealthScriptDailySchedule struct {
    DeviceHealthScriptTimeSchedule
}
// NewDeviceHealthScriptDailySchedule instantiates a new DeviceHealthScriptDailySchedule and sets the default values.
func NewDeviceHealthScriptDailySchedule()(*DeviceHealthScriptDailySchedule) {
    m := &DeviceHealthScriptDailySchedule{
        DeviceHealthScriptTimeSchedule: *NewDeviceHealthScriptTimeSchedule(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptDailySchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptDailyScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptDailyScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptDailySchedule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptDailySchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptTimeSchedule.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptDailySchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptTimeSchedule.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

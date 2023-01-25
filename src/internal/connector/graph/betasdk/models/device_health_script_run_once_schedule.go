package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRunOnceSchedule 
type DeviceHealthScriptRunOnceSchedule struct {
    DeviceHealthScriptTimeSchedule
    // The date the script is scheduled to run. This collection can contain a maximum of 20 elements.
    date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewDeviceHealthScriptRunOnceSchedule instantiates a new DeviceHealthScriptRunOnceSchedule and sets the default values.
func NewDeviceHealthScriptRunOnceSchedule()(*DeviceHealthScriptRunOnceSchedule) {
    m := &DeviceHealthScriptRunOnceSchedule{
        DeviceHealthScriptTimeSchedule: *NewDeviceHealthScriptTimeSchedule(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptRunOnceSchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptRunOnceScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRunOnceScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRunOnceSchedule(), nil
}
// GetDate gets the date property value. The date the script is scheduled to run. This collection can contain a maximum of 20 elements.
func (m *DeviceHealthScriptRunOnceSchedule) GetDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.date
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRunOnceSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptTimeSchedule.GetFieldDeserializers()
    res["date"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDate(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRunOnceSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptTimeSchedule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("date", m.GetDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDate sets the date property value. The date the script is scheduled to run. This collection can contain a maximum of 20 elements.
func (m *DeviceHealthScriptRunOnceSchedule) SetDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.date = value
}

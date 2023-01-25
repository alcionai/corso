package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptTimeSchedule 
type DeviceHealthScriptTimeSchedule struct {
    DeviceHealthScriptRunSchedule
    // At what time the script is scheduled to run. This collection can contain a maximum of 20 elements.
    time *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Indicate if the time is Utc or client local time.
    useUtc *bool
}
// NewDeviceHealthScriptTimeSchedule instantiates a new DeviceHealthScriptTimeSchedule and sets the default values.
func NewDeviceHealthScriptTimeSchedule()(*DeviceHealthScriptTimeSchedule) {
    m := &DeviceHealthScriptTimeSchedule{
        DeviceHealthScriptRunSchedule: *NewDeviceHealthScriptRunSchedule(),
    }
    odataTypeValue := "#microsoft.graph.deviceHealthScriptTimeSchedule";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceHealthScriptTimeScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptTimeScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceHealthScriptDailySchedule":
                        return NewDeviceHealthScriptDailySchedule(), nil
                    case "#microsoft.graph.deviceHealthScriptRunOnceSchedule":
                        return NewDeviceHealthScriptRunOnceSchedule(), nil
                }
            }
        }
    }
    return NewDeviceHealthScriptTimeSchedule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptTimeSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceHealthScriptRunSchedule.GetFieldDeserializers()
    res["time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTime(val)
        }
        return nil
    }
    res["useUtc"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseUtc(val)
        }
        return nil
    }
    return res
}
// GetTime gets the time property value. At what time the script is scheduled to run. This collection can contain a maximum of 20 elements.
func (m *DeviceHealthScriptTimeSchedule) GetTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.time
}
// GetUseUtc gets the useUtc property value. Indicate if the time is Utc or client local time.
func (m *DeviceHealthScriptTimeSchedule) GetUseUtc()(*bool) {
    return m.useUtc
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptTimeSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceHealthScriptRunSchedule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeOnlyValue("time", m.GetTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useUtc", m.GetUseUtc())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTime sets the time property value. At what time the script is scheduled to run. This collection can contain a maximum of 20 elements.
func (m *DeviceHealthScriptTimeSchedule) SetTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.time = value
}
// SetUseUtc sets the useUtc property value. Indicate if the time is Utc or client local time.
func (m *DeviceHealthScriptTimeSchedule) SetUseUtc(value *bool)() {
    m.useUtc = value
}

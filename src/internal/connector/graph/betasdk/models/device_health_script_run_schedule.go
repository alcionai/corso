package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRunSchedule base type of Device health script run schedule.
type DeviceHealthScriptRunSchedule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The x value of every x hours for hourly schedule, every x days for Daily Schedule, every x weeks for weekly schedule, every x months for Monthly Schedule. Valid values 1 to 23
    interval *int32
    // The OdataType property
    odataType *string
}
// NewDeviceHealthScriptRunSchedule instantiates a new deviceHealthScriptRunSchedule and sets the default values.
func NewDeviceHealthScriptRunSchedule()(*DeviceHealthScriptRunSchedule) {
    m := &DeviceHealthScriptRunSchedule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceHealthScriptRunScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRunScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceHealthScriptHourlySchedule":
                        return NewDeviceHealthScriptHourlySchedule(), nil
                    case "#microsoft.graph.deviceHealthScriptRunOnceSchedule":
                        return NewDeviceHealthScriptRunOnceSchedule(), nil
                    case "#microsoft.graph.deviceHealthScriptTimeSchedule":
                        return NewDeviceHealthScriptTimeSchedule(), nil
                }
            }
        }
    }
    return NewDeviceHealthScriptRunSchedule(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptRunSchedule) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRunSchedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterval(val)
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
// GetInterval gets the interval property value. The x value of every x hours for hourly schedule, every x days for Daily Schedule, every x weeks for weekly schedule, every x months for Monthly Schedule. Valid values 1 to 23
func (m *DeviceHealthScriptRunSchedule) GetInterval()(*int32) {
    return m.interval
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRunSchedule) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRunSchedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("interval", m.GetInterval())
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
func (m *DeviceHealthScriptRunSchedule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetInterval sets the interval property value. The x value of every x hours for hourly schedule, every x days for Daily Schedule, every x weeks for weekly schedule, every x months for Monthly Schedule. Valid values 1 to 23
func (m *DeviceHealthScriptRunSchedule) SetInterval(value *int32)() {
    m.interval = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthScriptRunSchedule) SetOdataType(value *string)() {
    m.odataType = value
}

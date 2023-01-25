package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomUpdateTimeWindow custom update time window
type CustomUpdateTimeWindow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The endDay property
    endDay *DayOfWeek
    // End time of the time window
    endTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The OdataType property
    odataType *string
    // The startDay property
    startDay *DayOfWeek
    // Start time of the time window
    startTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewCustomUpdateTimeWindow instantiates a new customUpdateTimeWindow and sets the default values.
func NewCustomUpdateTimeWindow()(*CustomUpdateTimeWindow) {
    m := &CustomUpdateTimeWindow{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomUpdateTimeWindowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomUpdateTimeWindowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomUpdateTimeWindow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomUpdateTimeWindow) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEndDay gets the endDay property value. The endDay property
func (m *CustomUpdateTimeWindow) GetEndDay()(*DayOfWeek) {
    return m.endDay
}
// GetEndTime gets the endTime property value. End time of the time window
func (m *CustomUpdateTimeWindow) GetEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.endTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomUpdateTimeWindow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["endDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDay(val.(*DayOfWeek))
        }
        return nil
    }
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
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
    res["startDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDay(val.(*DayOfWeek))
        }
        return nil
    }
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomUpdateTimeWindow) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDay gets the startDay property value. The startDay property
func (m *CustomUpdateTimeWindow) GetStartDay()(*DayOfWeek) {
    return m.startDay
}
// GetStartTime gets the startTime property value. Start time of the time window
func (m *CustomUpdateTimeWindow) GetStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.startTime
}
// Serialize serializes information the current object
func (m *CustomUpdateTimeWindow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEndDay() != nil {
        cast := (*m.GetEndDay()).String()
        err := writer.WriteStringValue("endDay", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("endTime", m.GetEndTime())
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
    if m.GetStartDay() != nil {
        cast := (*m.GetStartDay()).String()
        err := writer.WriteStringValue("startDay", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("startTime", m.GetStartTime())
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
func (m *CustomUpdateTimeWindow) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEndDay sets the endDay property value. The endDay property
func (m *CustomUpdateTimeWindow) SetEndDay(value *DayOfWeek)() {
    m.endDay = value
}
// SetEndTime sets the endTime property value. End time of the time window
func (m *CustomUpdateTimeWindow) SetEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.endTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomUpdateTimeWindow) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDay sets the startDay property value. The startDay property
func (m *CustomUpdateTimeWindow) SetStartDay(value *DayOfWeek)() {
    m.startDay = value
}
// SetStartTime sets the startTime property value. Start time of the time window
func (m *CustomUpdateTimeWindow) SetStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.startTime = value
}

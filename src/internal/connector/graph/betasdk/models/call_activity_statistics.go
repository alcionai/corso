package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallActivityStatistics 
type CallActivityStatistics struct {
    ActivityStatistics
    // Time spent on calls outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
    afterHours *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewCallActivityStatistics instantiates a new CallActivityStatistics and sets the default values.
func NewCallActivityStatistics()(*CallActivityStatistics) {
    m := &CallActivityStatistics{
        ActivityStatistics: *NewActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.callActivityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCallActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallActivityStatistics(), nil
}
// GetAfterHours gets the afterHours property value. Time spent on calls outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *CallActivityStatistics) GetAfterHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.afterHours
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ActivityStatistics.GetFieldDeserializers()
    res["afterHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAfterHours(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CallActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ActivityStatistics.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("afterHours", m.GetAfterHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAfterHours sets the afterHours property value. Time spent on calls outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *CallActivityStatistics) SetAfterHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.afterHours = value
}

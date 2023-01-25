package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatActivityStatistics 
type ChatActivityStatistics struct {
    ActivityStatistics
    // Time spent on chats outside of working hours, which is based on the user's Microsoft Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
    afterHours *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewChatActivityStatistics instantiates a new ChatActivityStatistics and sets the default values.
func NewChatActivityStatistics()(*ChatActivityStatistics) {
    m := &ChatActivityStatistics{
        ActivityStatistics: *NewActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.chatActivityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateChatActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatActivityStatistics(), nil
}
// GetAfterHours gets the afterHours property value. Time spent on chats outside of working hours, which is based on the user's Microsoft Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *ChatActivityStatistics) GetAfterHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.afterHours
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ChatActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAfterHours sets the afterHours property value. Time spent on chats outside of working hours, which is based on the user's Microsoft Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *ChatActivityStatistics) SetAfterHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.afterHours = value
}

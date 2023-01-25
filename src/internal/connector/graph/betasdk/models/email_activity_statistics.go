package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailActivityStatistics 
type EmailActivityStatistics struct {
    ActivityStatistics
    // Total hours spent on email outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
    afterHours *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Total hours spent reading email. The value is represented in ISO 8601 format for durations.
    readEmail *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Total hours spent writing and sending email. The value is represented in ISO 8601 format for durations.
    sentEmail *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewEmailActivityStatistics instantiates a new EmailActivityStatistics and sets the default values.
func NewEmailActivityStatistics()(*EmailActivityStatistics) {
    m := &EmailActivityStatistics{
        ActivityStatistics: *NewActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.emailActivityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmailActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailActivityStatistics(), nil
}
// GetAfterHours gets the afterHours property value. Total hours spent on email outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) GetAfterHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.afterHours
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["readEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadEmail(val)
        }
        return nil
    }
    res["sentEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentEmail(val)
        }
        return nil
    }
    return res
}
// GetReadEmail gets the readEmail property value. Total hours spent reading email. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) GetReadEmail()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.readEmail
}
// GetSentEmail gets the sentEmail property value. Total hours spent writing and sending email. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) GetSentEmail()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.sentEmail
}
// Serialize serializes information the current object
func (m *EmailActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteISODurationValue("readEmail", m.GetReadEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("sentEmail", m.GetSentEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAfterHours sets the afterHours property value. Total hours spent on email outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) SetAfterHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.afterHours = value
}
// SetReadEmail sets the readEmail property value. Total hours spent reading email. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) SetReadEmail(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.readEmail = value
}
// SetSentEmail sets the sentEmail property value. Total hours spent writing and sending email. The value is represented in ISO 8601 format for durations.
func (m *EmailActivityStatistics) SetSentEmail(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.sentEmail = value
}

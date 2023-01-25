package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingActivityStatistics 
type MeetingActivityStatistics struct {
    ActivityStatistics
    // Time spent on meetings outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
    afterHours *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Time spent in conflicting meetings (meetings that overlap with other meetings that the person accepted and where the person’s status is set to Busy). The value is represented in ISO 8601 format for durations.
    conflicting *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Time spent in long meetings (more than an hour in duration). The value is represented in ISO 8601 format for durations.
    long *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Time spent in meetings where the person was multitasking (read/sent more than a minimum number of emails and/or sent more than a minimum number of messages in Teams or in Skype for Business). The value is represented in ISO 8601 format for durations.
    multitasking *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Time spent in meetings organized by the user. The value is represented in ISO 8601 format for durations.
    organized *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Time spent on recurring meetings. The value is represented in ISO 8601 format for durations.
    recurring *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewMeetingActivityStatistics instantiates a new MeetingActivityStatistics and sets the default values.
func NewMeetingActivityStatistics()(*MeetingActivityStatistics) {
    m := &MeetingActivityStatistics{
        ActivityStatistics: *NewActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.meetingActivityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMeetingActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingActivityStatistics(), nil
}
// GetAfterHours gets the afterHours property value. Time spent on meetings outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetAfterHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.afterHours
}
// GetConflicting gets the conflicting property value. Time spent in conflicting meetings (meetings that overlap with other meetings that the person accepted and where the person’s status is set to Busy). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetConflicting()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.conflicting
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["conflicting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflicting(val)
        }
        return nil
    }
    res["long"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLong(val)
        }
        return nil
    }
    res["multitasking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMultitasking(val)
        }
        return nil
    }
    res["organized"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganized(val)
        }
        return nil
    }
    res["recurring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurring(val)
        }
        return nil
    }
    return res
}
// GetLong gets the long property value. Time spent in long meetings (more than an hour in duration). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetLong()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.long
}
// GetMultitasking gets the multitasking property value. Time spent in meetings where the person was multitasking (read/sent more than a minimum number of emails and/or sent more than a minimum number of messages in Teams or in Skype for Business). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetMultitasking()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.multitasking
}
// GetOrganized gets the organized property value. Time spent in meetings organized by the user. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetOrganized()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.organized
}
// GetRecurring gets the recurring property value. Time spent on recurring meetings. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) GetRecurring()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.recurring
}
// Serialize serializes information the current object
func (m *MeetingActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteISODurationValue("conflicting", m.GetConflicting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("long", m.GetLong())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("multitasking", m.GetMultitasking())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("organized", m.GetOrganized())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("recurring", m.GetRecurring())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAfterHours sets the afterHours property value. Time spent on meetings outside of working hours, which is based on the user's Outlook calendar setting for work hours. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetAfterHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.afterHours = value
}
// SetConflicting sets the conflicting property value. Time spent in conflicting meetings (meetings that overlap with other meetings that the person accepted and where the person’s status is set to Busy). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetConflicting(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.conflicting = value
}
// SetLong sets the long property value. Time spent in long meetings (more than an hour in duration). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetLong(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.long = value
}
// SetMultitasking sets the multitasking property value. Time spent in meetings where the person was multitasking (read/sent more than a minimum number of emails and/or sent more than a minimum number of messages in Teams or in Skype for Business). The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetMultitasking(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.multitasking = value
}
// SetOrganized sets the organized property value. Time spent in meetings organized by the user. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetOrganized(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.organized = value
}
// SetRecurring sets the recurring property value. Time spent on recurring meetings. The value is represented in ISO 8601 format for durations.
func (m *MeetingActivityStatistics) SetRecurring(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.recurring = value
}

package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceTimelineEvents the user experience analytics device events entity contains NRT device events details.
type UserExperienceAnalyticsDeviceTimelineEvents struct {
    Entity
    // The id of the device where the event occurred.
    deviceId *string
    // Placeholder value for future expansion.
    eventAdditionalInformation *string
    // The time the event occured.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The details provided by the event, format depends on event type.
    eventDetails *string
    // Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
    eventLevel *DeviceEventLevel
    // The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
    eventName *string
    // The source of the event. Examples include: Intune, Sccm.
    eventSource *string
}
// NewUserExperienceAnalyticsDeviceTimelineEvents instantiates a new userExperienceAnalyticsDeviceTimelineEvents and sets the default values.
func NewUserExperienceAnalyticsDeviceTimelineEvents()(*UserExperienceAnalyticsDeviceTimelineEvents) {
    m := &UserExperienceAnalyticsDeviceTimelineEvents{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceTimelineEventsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceTimelineEventsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceTimelineEvents(), nil
}
// GetDeviceId gets the deviceId property value. The id of the device where the event occurred.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetDeviceId()(*string) {
    return m.deviceId
}
// GetEventAdditionalInformation gets the eventAdditionalInformation property value. Placeholder value for future expansion.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventAdditionalInformation()(*string) {
    return m.eventAdditionalInformation
}
// GetEventDateTime gets the eventDateTime property value. The time the event occured.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetEventDetails gets the eventDetails property value. The details provided by the event, format depends on event type.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventDetails()(*string) {
    return m.eventDetails
}
// GetEventLevel gets the eventLevel property value. Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventLevel()(*DeviceEventLevel) {
    return m.eventLevel
}
// GetEventName gets the eventName property value. The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventName()(*string) {
    return m.eventName
}
// GetEventSource gets the eventSource property value. The source of the event. Examples include: Intune, Sccm.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetEventSource()(*string) {
    return m.eventSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceTimelineEvents) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["eventAdditionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventAdditionalInformation(val)
        }
        return nil
    }
    res["eventDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["eventDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetails(val)
        }
        return nil
    }
    res["eventLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEventLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventLevel(val.(*DeviceEventLevel))
        }
        return nil
    }
    res["eventName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
    res["eventSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventSource(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceTimelineEvents) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventAdditionalInformation", m.GetEventAdditionalInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventDetails", m.GetEventDetails())
        if err != nil {
            return err
        }
    }
    if m.GetEventLevel() != nil {
        cast := (*m.GetEventLevel()).String()
        err = writer.WriteStringValue("eventLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("eventSource", m.GetEventSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The id of the device where the event occurred.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetEventAdditionalInformation sets the eventAdditionalInformation property value. Placeholder value for future expansion.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventAdditionalInformation(value *string)() {
    m.eventAdditionalInformation = value
}
// SetEventDateTime sets the eventDateTime property value. The time the event occured.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetEventDetails sets the eventDetails property value. The details provided by the event, format depends on event type.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventDetails(value *string)() {
    m.eventDetails = value
}
// SetEventLevel sets the eventLevel property value. Indicates device event level. Possible values are: None, Verbose, Information, Warning, Error, Critical
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventLevel(value *DeviceEventLevel)() {
    m.eventLevel = value
}
// SetEventName sets the eventName property value. The name of the event. Examples include: BootEvent, LogonEvent, AppCrashEvent, AppHangEvent.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventName(value *string)() {
    m.eventName = value
}
// SetEventSource sets the eventSource property value. The source of the event. Examples include: Intune, Sccm.
func (m *UserExperienceAnalyticsDeviceTimelineEvents) SetEventSource(value *string)() {
    m.eventSource = value
}

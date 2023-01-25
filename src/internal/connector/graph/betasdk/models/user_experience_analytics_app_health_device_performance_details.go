package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthDevicePerformanceDetails the user experience analytics device performance entity contains device performance details.
type UserExperienceAnalyticsAppHealthDevicePerformanceDetails struct {
    Entity
    // The friendly name of the application for which the event occurred.
    appDisplayName *string
    // The publisher of the application.
    appPublisher *string
    // The version of the application.
    appVersion *string
    // The name of the device.
    deviceDisplayName *string
    // The id of the device.
    deviceId *string
    // The time the event occurred.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of the event.
    eventType *string
}
// NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails instantiates a new userExperienceAnalyticsAppHealthDevicePerformanceDetails and sets the default values.
func NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserExperienceAnalyticsAppHealthDevicePerformanceDetails) {
    m := &UserExperienceAnalyticsAppHealthDevicePerformanceDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthDevicePerformanceDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformanceDetails(), nil
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application for which the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetAppVersion()(*string) {
    return m.appVersion
}
// GetDeviceDisplayName gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetDeviceId()(*string) {
    return m.deviceId
}
// GetEventDateTime gets the eventDateTime property value. The time the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.eventDateTime
}
// GetEventType gets the eventType property value. The type of the event.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetEventType()(*string) {
    return m.eventType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["appVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppVersion(val)
        }
        return nil
    }
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
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
    res["eventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
        err = writer.WriteStringValue("eventType", m.GetEventType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application for which the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetDeviceDisplayName sets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetEventDateTime sets the eventDateTime property value. The time the event occurred.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// SetEventType sets the eventType property value. The type of the event.
func (m *UserExperienceAnalyticsAppHealthDevicePerformanceDetails) SetEventType(value *string)() {
    m.eventType = value
}

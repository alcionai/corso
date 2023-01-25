package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsMetricHistory the user experience analytics metric history.
type UserExperienceAnalyticsMetricHistory struct {
    Entity
    // The user experience analytics device id.
    deviceId *string
    // The user experience analytics metric date time.
    metricDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user experience analytics metric type.
    metricType *string
}
// NewUserExperienceAnalyticsMetricHistory instantiates a new userExperienceAnalyticsMetricHistory and sets the default values.
func NewUserExperienceAnalyticsMetricHistory()(*UserExperienceAnalyticsMetricHistory) {
    m := &UserExperienceAnalyticsMetricHistory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsMetricHistoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsMetricHistory(), nil
}
// GetDeviceId gets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsMetricHistory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["metricDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricDateTime(val)
        }
        return nil
    }
    res["metricType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetricType(val)
        }
        return nil
    }
    return res
}
// GetMetricDateTime gets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.metricDateTime
}
// GetMetricType gets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) GetMetricType()(*string) {
    return m.metricType
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsMetricHistory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("metricDateTime", m.GetMetricDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metricType", m.GetMetricType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsMetricHistory) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetMetricDateTime sets the metricDateTime property value. The user experience analytics metric date time.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.metricDateTime = value
}
// SetMetricType sets the metricType property value. The user experience analytics metric type.
func (m *UserExperienceAnalyticsMetricHistory) SetMetricType(value *string)() {
    m.metricType = value
}

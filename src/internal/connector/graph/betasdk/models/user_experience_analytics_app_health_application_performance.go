package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthApplicationPerformance the user experience analytics application performance entity contains app performance details.
type UserExperienceAnalyticsAppHealthApplicationPerformance struct {
    Entity
    // The number of devices where the app has been active. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32
    // The friendly name of the application.
    appDisplayName *string
    // The number of hangs for the app. Valid values -2147483648 to 2147483647
    appHangCount *int32
    // The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appHealthScore *float64
    // The overall health status of the app.
    appHealthStatus *string
    // The name of the application.
    appName *string
    // The publisher of the application.
    appPublisher *string
    // The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
    appUsageDuration *int32
    // The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
}
// NewUserExperienceAnalyticsAppHealthApplicationPerformance instantiates a new userExperienceAnalyticsAppHealthApplicationPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthApplicationPerformance()(*UserExperienceAnalyticsAppHealthApplicationPerformance) {
    m := &UserExperienceAnalyticsAppHealthApplicationPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthApplicationPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthApplicationPerformance(), nil
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetActiveDeviceCount()(*int32) {
    return m.activeDeviceCount
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppHangCount gets the appHangCount property value. The number of hangs for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHangCount()(*int32) {
    return m.appHangCount
}
// GetAppHealthScore gets the appHealthScore property value. The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthScore()(*float64) {
    return m.appHealthScore
}
// GetAppHealthStatus gets the appHealthStatus property value. The overall health status of the app.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppHealthStatus()(*string) {
    return m.appHealthStatus
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppUsageDuration gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetAppUsageDuration()(*int32) {
    return m.appUsageDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDeviceCount(val)
        }
        return nil
    }
    res["appCrashCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppCrashCount(val)
        }
        return nil
    }
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
    res["appHangCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHangCount(val)
        }
        return nil
    }
    res["appHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHealthScore(val)
        }
        return nil
    }
    res["appHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppHealthStatus(val)
        }
        return nil
    }
    res["appName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
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
    res["appUsageDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUsageDuration(val)
        }
        return nil
    }
    res["meanTimeToFailureInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeanTimeToFailureInMinutes(val)
        }
        return nil
    }
    return res
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDeviceCount", m.GetActiveDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("appCrashCount", m.GetAppCrashCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("appHangCount", m.GetAppHangCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("appHealthScore", m.GetAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appHealthStatus", m.GetAppHealthStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
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
        err = writer.WriteInt32Value("appUsageDuration", m.GetAppUsageDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("meanTimeToFailureInMinutes", m.GetMeanTimeToFailureInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDeviceCount sets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppHangCount sets the appHangCount property value. The number of hangs for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHangCount(value *int32)() {
    m.appHangCount = value
}
// SetAppHealthScore sets the appHealthScore property value. The health score of the app. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthScore(value *float64)() {
    m.appHealthScore = value
}
// SetAppHealthStatus sets the appHealthStatus property value. The overall health status of the app.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppHealthStatus(value *string)() {
    m.appHealthStatus = value
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppUsageDuration sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthApplicationPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}

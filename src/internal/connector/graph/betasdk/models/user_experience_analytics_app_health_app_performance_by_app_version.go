package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion the user experience analytics application performance entity contains app performance details by app version.
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion struct {
    Entity
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32
    // The friendly name of the application.
    appDisplayName *string
    // The name of the application.
    appName *string
    // The publisher of the application.
    appPublisher *string
    // The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
    appUsageDuration *int32
    // The version of the application.
    appVersion *string
    // The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
}
// NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersion and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersion(), nil
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppUsageDuration gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppUsageDuration()(*int32) {
    return m.appUsageDuration
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetAppVersion()(*string) {
    return m.appVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
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
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppUsageDuration sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}

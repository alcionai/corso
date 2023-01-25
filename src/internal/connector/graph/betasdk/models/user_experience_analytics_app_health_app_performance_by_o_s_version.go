package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion the user experience analytics application performance entity contains app performance details by OS version.
type UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion struct {
    Entity
    // The number of devices where the app has been active. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32
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
    // The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
    // The os build number of the application.
    osBuildNumber *string
    // The os version of the application.
    osVersion *string
}
// NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByOSVersion and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion()(*UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthAppPerformanceByOSVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByOSVersion(), nil
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetActiveDeviceCount()(*int32) {
    return m.activeDeviceCount
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppUsageDuration gets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetAppUsageDuration()(*int32) {
    return m.appUsageDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    return res
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// GetOsBuildNumber gets the osBuildNumber property value. The os build number of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetOsBuildNumber()(*string) {
    return m.osBuildNumber
}
// GetOsVersion gets the osVersion property value. The os version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) GetOsVersion()(*string) {
    return m.osVersion
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDeviceCount sets the activeDeviceCount property value. The number of devices where the app has been active. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppUsageDuration sets the appUsageDuration property value. The total usage time of the application in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetAppUsageDuration(value *int32)() {
    m.appUsageDuration = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the app in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// SetOsBuildNumber sets the osBuildNumber property value. The os build number of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// SetOsVersion sets the osVersion property value. The os version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion) SetOsVersion(value *string)() {
    m.osVersion = value
}

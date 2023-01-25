package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails the user experience analytics application performance entity contains app performance by app version details.
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails struct {
    Entity
    // The number of crashes for the app. Valid values -2147483648 to 2147483647
    appCrashCount *int32
    // The friendly name of the application.
    appDisplayName *string
    // The name of the application.
    appName *string
    // The publisher of the application.
    appPublisher *string
    // The version of the application.
    appVersion *string
    // The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
    deviceCountWithCrashes *int32
    // Is the version of application the latest version for that app that is in use.
    isLatestUsedVersion *bool
    // Is the version of application the most used version for that app.
    isMostUsedVersion *bool
}
// NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails instantiates a new userExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails and sets the default values.
func NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails()(*UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) {
    m := &UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails(), nil
}
// GetAppCrashCount gets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppDisplayName gets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetAppVersion gets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetAppVersion()(*string) {
    return m.appVersion
}
// GetDeviceCountWithCrashes gets the deviceCountWithCrashes property value. The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetDeviceCountWithCrashes()(*int32) {
    return m.deviceCountWithCrashes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["deviceCountWithCrashes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCountWithCrashes(val)
        }
        return nil
    }
    res["isLatestUsedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLatestUsedVersion(val)
        }
        return nil
    }
    res["isMostUsedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMostUsedVersion(val)
        }
        return nil
    }
    return res
}
// GetIsLatestUsedVersion gets the isLatestUsedVersion property value. Is the version of application the latest version for that app that is in use.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetIsLatestUsedVersion()(*bool) {
    return m.isLatestUsedVersion
}
// GetIsMostUsedVersion gets the isMostUsedVersion property value. Is the version of application the most used version for that app.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) GetIsMostUsedVersion()(*bool) {
    return m.isMostUsedVersion
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("appVersion", m.GetAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCountWithCrashes", m.GetDeviceCountWithCrashes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLatestUsedVersion", m.GetIsLatestUsedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMostUsedVersion", m.GetIsMostUsedVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppCrashCount sets the appCrashCount property value. The number of crashes for the app. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppDisplayName sets the appDisplayName property value. The friendly name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. The name of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. The publisher of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetAppVersion sets the appVersion property value. The version of the application.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetAppVersion(value *string)() {
    m.appVersion = value
}
// SetDeviceCountWithCrashes sets the deviceCountWithCrashes property value. The total number of devices that have reported one or more application crashes for this application and version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetDeviceCountWithCrashes(value *int32)() {
    m.deviceCountWithCrashes = value
}
// SetIsLatestUsedVersion sets the isLatestUsedVersion property value. Is the version of application the latest version for that app that is in use.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetIsLatestUsedVersion(value *bool)() {
    m.isLatestUsedVersion = value
}
// SetIsMostUsedVersion sets the isMostUsedVersion property value. Is the version of application the most used version for that app.
func (m *UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails) SetIsMostUsedVersion(value *bool)() {
    m.isMostUsedVersion = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthAppImpact the user experience analytics battery health app impact entity contains battery usage related information at an app level for the tenant.
type UserExperienceAnalyticsBatteryHealthAppImpact struct {
    Entity
    // Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
    activeDevices *int32
    // User friendly display name for the app. Eg: Outlook
    appDisplayName *string
    // App name. Eg: oltk.exe
    appName *string
    // App publisher. Eg: Microsoft Corporation
    appPublisher *string
    // The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryUsagePercentage *float64
    // true if the user had active interaction with the app.
    isForegroundApp *bool
}
// NewUserExperienceAnalyticsBatteryHealthAppImpact instantiates a new userExperienceAnalyticsBatteryHealthAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthAppImpact()(*UserExperienceAnalyticsBatteryHealthAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthAppImpactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthAppImpactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthAppImpact(), nil
}
// GetActiveDevices gets the activeDevices property value. Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetActiveDevices()(*int32) {
    return m.activeDevices
}
// GetAppDisplayName gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetBatteryUsagePercentage gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetBatteryUsagePercentage()(*float64) {
    return m.batteryUsagePercentage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
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
    res["batteryUsagePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryUsagePercentage(val)
        }
        return nil
    }
    res["isForegroundApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsForegroundApp(val)
        }
        return nil
    }
    return res
}
// GetIsForegroundApp gets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) GetIsForegroundApp()(*bool) {
    return m.isForegroundApp
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDevices", m.GetActiveDevices())
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
        err = writer.WriteFloat64Value("batteryUsagePercentage", m.GetBatteryUsagePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isForegroundApp", m.GetIsForegroundApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices for using that app over a 14-day period. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// SetAppDisplayName sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetBatteryUsagePercentage sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days computed across all devices in the tenant. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetBatteryUsagePercentage(value *float64)() {
    m.batteryUsagePercentage = value
}
// SetIsForegroundApp sets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthAppImpact) SetIsForegroundApp(value *bool)() {
    m.isForegroundApp = value
}

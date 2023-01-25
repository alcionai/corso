package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpact the user experience analytics battery health device app impact entity contains battery usage related information at an app level for a given device.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpact struct {
    Entity
    // User friendly display name for the app. Eg: Outlook
    appDisplayName *string
    // App name. Eg: oltk.exe
    appName *string
    // App publisher. Eg: Microsoft Corporation
    appPublisher *string
    // The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryUsagePercentage *float64
    // The unique identifier of the device, Intune DeviceID or SCCM device id.
    deviceId *string
    // true if the user had active interaction with the app.
    isForegroundApp *bool
}
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact instantiates a new userExperienceAnalyticsBatteryHealthDeviceAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact(), nil
}
// GetAppDisplayName gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppName gets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppName()(*string) {
    return m.appName
}
// GetAppPublisher gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppPublisher()(*string) {
    return m.appPublisher
}
// GetBatteryUsagePercentage gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetBatteryUsagePercentage()(*float64) {
    return m.batteryUsagePercentage
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetIsForegroundApp()(*bool) {
    return m.isForegroundApp
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
// SetAppDisplayName sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppName sets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppName(value *string)() {
    m.appName = value
}
// SetAppPublisher sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppPublisher(value *string)() {
    m.appPublisher = value
}
// SetBatteryUsagePercentage sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetBatteryUsagePercentage(value *float64)() {
    m.batteryUsagePercentage = value
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetIsForegroundApp sets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetIsForegroundApp(value *bool)() {
    m.isForegroundApp = value
}

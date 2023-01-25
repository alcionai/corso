package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthDevicePerformance the user experience analytics device performance entity contains device performance details.
type UserExperienceAnalyticsAppHealthDevicePerformance struct {
    Entity
    // The number of app crashes for the device. Valid values -2147483648 to 2147483647
    appCrashCount *int32
    // The number of app hangs for the device. Valid values -2147483648 to 2147483647
    appHangCount *int32
    // The number of distinct app crashes for the device. Valid values -2147483648 to 2147483647
    crashedAppCount *int32
    // The app health score of the device. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    deviceAppHealthScore *float64
    // The overall app health status of the device.
    deviceAppHealthStatus *string
    // The name of the device.
    deviceDisplayName *string
    // The id of the device.
    deviceId *string
    // The manufacturer name of the device.
    deviceManufacturer *string
    // The model name of the device.
    deviceModel *string
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The mean time to failure for the device in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
    // The date and time when the statistics were last computed.
    processedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewUserExperienceAnalyticsAppHealthDevicePerformance instantiates a new userExperienceAnalyticsAppHealthDevicePerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthDevicePerformance()(*UserExperienceAnalyticsAppHealthDevicePerformance) {
    m := &UserExperienceAnalyticsAppHealthDevicePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthDevicePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthDevicePerformance(), nil
}
// GetAppCrashCount gets the appCrashCount property value. The number of app crashes for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetAppCrashCount()(*int32) {
    return m.appCrashCount
}
// GetAppHangCount gets the appHangCount property value. The number of app hangs for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetAppHangCount()(*int32) {
    return m.appHangCount
}
// GetCrashedAppCount gets the crashedAppCount property value. The number of distinct app crashes for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetCrashedAppCount()(*int32) {
    return m.crashedAppCount
}
// GetDeviceAppHealthScore gets the deviceAppHealthScore property value. The app health score of the device. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceAppHealthScore()(*float64) {
    return m.deviceAppHealthScore
}
// GetDeviceAppHealthStatus gets the deviceAppHealthStatus property value. The overall app health status of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceAppHealthStatus()(*string) {
    return m.deviceAppHealthStatus
}
// GetDeviceDisplayName gets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceManufacturer()(*string) {
    return m.deviceManufacturer
}
// GetDeviceModel gets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["crashedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrashedAppCount(val)
        }
        return nil
    }
    res["deviceAppHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAppHealthScore(val)
        }
        return nil
    }
    res["deviceAppHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAppHealthStatus(val)
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
    res["deviceManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturer(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*UserExperienceAnalyticsHealthState))
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
    res["processedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessedDateTime(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// GetProcessedDateTime gets the processedDateTime property value. The date and time when the statistics were last computed.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) GetProcessedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.processedDateTime
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("appHangCount", m.GetAppHangCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("crashedAppCount", m.GetCrashedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("deviceAppHealthScore", m.GetDeviceAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceAppHealthStatus", m.GetDeviceAppHealthStatus())
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
        err = writer.WriteStringValue("deviceManufacturer", m.GetDeviceManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
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
        err = writer.WriteTimeValue("processedDateTime", m.GetProcessedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppCrashCount sets the appCrashCount property value. The number of app crashes for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetAppCrashCount(value *int32)() {
    m.appCrashCount = value
}
// SetAppHangCount sets the appHangCount property value. The number of app hangs for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetAppHangCount(value *int32)() {
    m.appHangCount = value
}
// SetCrashedAppCount sets the crashedAppCount property value. The number of distinct app crashes for the device. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetCrashedAppCount(value *int32)() {
    m.crashedAppCount = value
}
// SetDeviceAppHealthScore sets the deviceAppHealthScore property value. The app health score of the device. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceAppHealthScore(value *float64)() {
    m.deviceAppHealthScore = value
}
// SetDeviceAppHealthStatus sets the deviceAppHealthStatus property value. The overall app health status of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceAppHealthStatus(value *string)() {
    m.deviceAppHealthStatus = value
}
// SetDeviceDisplayName sets the deviceDisplayName property value. The name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceManufacturer(value *string)() {
    m.deviceManufacturer = value
}
// SetDeviceModel sets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the device in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// SetProcessedDateTime sets the processedDateTime property value. The date and time when the statistics were last computed.
func (m *UserExperienceAnalyticsAppHealthDevicePerformance) SetProcessedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processedDateTime = value
}

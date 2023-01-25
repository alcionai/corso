package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScores the user experience analytics device scores entity consolidates the various endpoint analytics scores.
type UserExperienceAnalyticsDeviceScores struct {
    Entity
    // The user experience analytics device app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appReliabilityScore *float64
    // The user experience analytics device battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryHealthScore *float64
    // The user experience analytics device name.
    deviceName *string
    // The user experience analytics device score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    endpointAnalyticsScore *float64
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The user experience analytics device manufacturer.
    manufacturer *string
    // The user experience analytics device model.
    model *string
    // The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64
    // The user experience analytics device work From anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsDeviceScores instantiates a new userExperienceAnalyticsDeviceScores and sets the default values.
func NewUserExperienceAnalyticsDeviceScores()(*UserExperienceAnalyticsDeviceScores) {
    m := &UserExperienceAnalyticsDeviceScores{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScores(), nil
}
// GetAppReliabilityScore gets the appReliabilityScore property value. The user experience analytics device app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) GetAppReliabilityScore()(*float64) {
    return m.appReliabilityScore
}
// GetBatteryHealthScore gets the batteryHealthScore property value. The user experience analytics device battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) GetBatteryHealthScore()(*float64) {
    return m.batteryHealthScore
}
// GetDeviceName gets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDeviceScores) GetDeviceName()(*string) {
    return m.deviceName
}
// GetEndpointAnalyticsScore gets the endpointAnalyticsScore property value. The user experience analytics device score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) GetEndpointAnalyticsScore()(*float64) {
    return m.endpointAnalyticsScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appReliabilityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppReliabilityScore(val)
        }
        return nil
    }
    res["batteryHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthScore(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["endpointAnalyticsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointAnalyticsScore(val)
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
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["startupPerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupPerformanceScore(val)
        }
        return nil
    }
    res["workFromAnywhereScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFromAnywhereScore(val)
        }
        return nil
    }
    return res
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDeviceScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDeviceScores) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDeviceScores) GetModel()(*string) {
    return m.model
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) GetStartupPerformanceScore()(*float64) {
    return m.startupPerformanceScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The user experience analytics device work From anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("appReliabilityScore", m.GetAppReliabilityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("batteryHealthScore", m.GetBatteryHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("endpointAnalyticsScore", m.GetEndpointAnalyticsScore())
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
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("startupPerformanceScore", m.GetStartupPerformanceScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("workFromAnywhereScore", m.GetWorkFromAnywhereScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppReliabilityScore sets the appReliabilityScore property value. The user experience analytics device app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetAppReliabilityScore(value *float64)() {
    m.appReliabilityScore = value
}
// SetBatteryHealthScore sets the batteryHealthScore property value. The user experience analytics device battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetBatteryHealthScore(value *float64)() {
    m.batteryHealthScore = value
}
// SetDeviceName sets the deviceName property value. The user experience analytics device name.
func (m *UserExperienceAnalyticsDeviceScores) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetEndpointAnalyticsScore sets the endpointAnalyticsScore property value. The user experience analytics device score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetEndpointAnalyticsScore(value *float64)() {
    m.endpointAnalyticsScore = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsDeviceScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsDeviceScores) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsDeviceScores) SetModel(value *string)() {
    m.model = value
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The user experience analytics device startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetStartupPerformanceScore(value *float64)() {
    m.startupPerformanceScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The user experience analytics device work From anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsDeviceScores) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}

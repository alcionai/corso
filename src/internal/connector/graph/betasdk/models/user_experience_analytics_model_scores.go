package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsModelScores the user experience analytics model scores entity consolidates the various endpoint analytics scores.
type UserExperienceAnalyticsModelScores struct {
    Entity
    // The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    appReliabilityScore *float64
    // The user experience analytics model battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    batteryHealthScore *float64
    // The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    endpointAnalyticsScore *float64
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // A unique identifier of the user experience analytics model scores: device manufacturer.
    manufacturer *string
    // A unique identifier of the user experience analytics model scores: device model.
    model *string
    // The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
    modelDeviceCount *int64
    // The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    startupPerformanceScore *float64
    // The user experience analytics model work from anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsModelScores instantiates a new userExperienceAnalyticsModelScores and sets the default values.
func NewUserExperienceAnalyticsModelScores()(*UserExperienceAnalyticsModelScores) {
    m := &UserExperienceAnalyticsModelScores{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsModelScoresFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsModelScores(), nil
}
// GetAppReliabilityScore gets the appReliabilityScore property value. The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetAppReliabilityScore()(*float64) {
    return m.appReliabilityScore
}
// GetBatteryHealthScore gets the batteryHealthScore property value. The user experience analytics model battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetBatteryHealthScore()(*float64) {
    return m.batteryHealthScore
}
// GetEndpointAnalyticsScore gets the endpointAnalyticsScore property value. The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetEndpointAnalyticsScore()(*float64) {
    return m.endpointAnalyticsScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsModelScores) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["modelDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelDeviceCount(val)
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
func (m *UserExperienceAnalyticsModelScores) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetManufacturer gets the manufacturer property value. A unique identifier of the user experience analytics model scores: device manufacturer.
func (m *UserExperienceAnalyticsModelScores) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. A unique identifier of the user experience analytics model scores: device model.
func (m *UserExperienceAnalyticsModelScores) GetModel()(*string) {
    return m.model
}
// GetModelDeviceCount gets the modelDeviceCount property value. The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *UserExperienceAnalyticsModelScores) GetModelDeviceCount()(*int64) {
    return m.modelDeviceCount
}
// GetStartupPerformanceScore gets the startupPerformanceScore property value. The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetStartupPerformanceScore()(*float64) {
    return m.startupPerformanceScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The user experience analytics model work from anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsModelScores) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("modelDeviceCount", m.GetModelDeviceCount())
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
// SetAppReliabilityScore sets the appReliabilityScore property value. The user experience analytics model app reliability score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetAppReliabilityScore(value *float64)() {
    m.appReliabilityScore = value
}
// SetBatteryHealthScore sets the batteryHealthScore property value. The user experience analytics model battery health score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetBatteryHealthScore(value *float64)() {
    m.batteryHealthScore = value
}
// SetEndpointAnalyticsScore sets the endpointAnalyticsScore property value. The user experience analytics model score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetEndpointAnalyticsScore(value *float64)() {
    m.endpointAnalyticsScore = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsModelScores) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetManufacturer sets the manufacturer property value. A unique identifier of the user experience analytics model scores: device manufacturer.
func (m *UserExperienceAnalyticsModelScores) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. A unique identifier of the user experience analytics model scores: device model.
func (m *UserExperienceAnalyticsModelScores) SetModel(value *string)() {
    m.model = value
}
// SetModelDeviceCount sets the modelDeviceCount property value. The user experience analytics model device count. Valid values -9.22337203685478E+18 to 9.22337203685478E+18
func (m *UserExperienceAnalyticsModelScores) SetModelDeviceCount(value *int64)() {
    m.modelDeviceCount = value
}
// SetStartupPerformanceScore sets the startupPerformanceScore property value. The user experience analytics model startup performance score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetStartupPerformanceScore(value *float64)() {
    m.startupPerformanceScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The user experience analytics model work from anywhere score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsModelScores) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}

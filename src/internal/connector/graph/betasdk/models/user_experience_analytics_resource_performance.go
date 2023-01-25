package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsResourcePerformance the user experience analytics resource performance entity.
type UserExperienceAnalyticsResourcePerformance struct {
    Entity
    // AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
    averageSpikeTimeScore *int32
    // CPU spike time in percentage. Valid values 0 to 100
    cpuSpikeTimePercentage *float64
    // Threshold of cpuSpikeTimeScore. Valid values 0 to 100
    cpuSpikeTimePercentageThreshold *float64
    // The user experience analytics device CPU spike time score. Valid values 0 to 100
    cpuSpikeTimeScore *int32
    // User experience analytics summarized device count.
    deviceCount *int64
    // The id of the device.
    deviceId *string
    // The name of the device.
    deviceName *string
    // Resource performance score of a specific device. Valid values 0 to 100
    deviceResourcePerformanceScore *int32
    // The user experience analytics device manufacturer.
    manufacturer *string
    // The user experience analytics device model.
    model *string
    // RAM spike time in percentage. Valid values 0 to 100
    ramSpikeTimePercentage *float64
    // Threshold of ramSpikeTimeScore. Valid values 0 to 100
    ramSpikeTimePercentageThreshold *float64
    // The user experience analytics device RAM spike time score. Valid values 0 to 100
    ramSpikeTimeScore *int32
}
// NewUserExperienceAnalyticsResourcePerformance instantiates a new userExperienceAnalyticsResourcePerformance and sets the default values.
func NewUserExperienceAnalyticsResourcePerformance()(*UserExperienceAnalyticsResourcePerformance) {
    m := &UserExperienceAnalyticsResourcePerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsResourcePerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsResourcePerformance(), nil
}
// GetAverageSpikeTimeScore gets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetAverageSpikeTimeScore()(*int32) {
    return m.averageSpikeTimeScore
}
// GetCpuSpikeTimePercentage gets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentage()(*float64) {
    return m.cpuSpikeTimePercentage
}
// GetCpuSpikeTimePercentageThreshold gets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimePercentageThreshold()(*float64) {
    return m.cpuSpikeTimePercentageThreshold
}
// GetCpuSpikeTimeScore gets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetCpuSpikeTimeScore()(*int32) {
    return m.cpuSpikeTimeScore
}
// GetDeviceCount gets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceCount()(*int64) {
    return m.deviceCount
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceResourcePerformanceScore gets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetDeviceResourcePerformanceScore()(*int32) {
    return m.deviceResourcePerformanceScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsResourcePerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["averageSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageSpikeTimeScore(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentage(val)
        }
        return nil
    }
    res["cpuSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["cpuSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuSpikeTimeScore(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
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
    res["deviceResourcePerformanceScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceResourcePerformanceScore(val)
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
    res["ramSpikeTimePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentage(val)
        }
        return nil
    }
    res["ramSpikeTimePercentageThreshold"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimePercentageThreshold(val)
        }
        return nil
    }
    res["ramSpikeTimeScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamSpikeTimeScore(val)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) GetModel()(*string) {
    return m.model
}
// GetRamSpikeTimePercentage gets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentage()(*float64) {
    return m.ramSpikeTimePercentage
}
// GetRamSpikeTimePercentageThreshold gets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimePercentageThreshold()(*float64) {
    return m.ramSpikeTimePercentageThreshold
}
// GetRamSpikeTimeScore gets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) GetRamSpikeTimeScore()(*int32) {
    return m.ramSpikeTimeScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsResourcePerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("averageSpikeTimeScore", m.GetAverageSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cpuSpikeTimePercentage", m.GetCpuSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cpuSpikeTimePercentageThreshold", m.GetCpuSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cpuSpikeTimeScore", m.GetCpuSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
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
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceResourcePerformanceScore", m.GetDeviceResourcePerformanceScore())
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
        err = writer.WriteFloat64Value("ramSpikeTimePercentage", m.GetRamSpikeTimePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("ramSpikeTimePercentageThreshold", m.GetRamSpikeTimePercentageThreshold())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ramSpikeTimeScore", m.GetRamSpikeTimeScore())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageSpikeTimeScore sets the averageSpikeTimeScore property value. AverageSpikeTimeScore of a device or a model type. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetAverageSpikeTimeScore(value *int32)() {
    m.averageSpikeTimeScore = value
}
// SetCpuSpikeTimePercentage sets the cpuSpikeTimePercentage property value. CPU spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentage(value *float64)() {
    m.cpuSpikeTimePercentage = value
}
// SetCpuSpikeTimePercentageThreshold sets the cpuSpikeTimePercentageThreshold property value. Threshold of cpuSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimePercentageThreshold(value *float64)() {
    m.cpuSpikeTimePercentageThreshold = value
}
// SetCpuSpikeTimeScore sets the cpuSpikeTimeScore property value. The user experience analytics device CPU spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetCpuSpikeTimeScore(value *int32)() {
    m.cpuSpikeTimeScore = value
}
// SetDeviceCount sets the deviceCount property value. User experience analytics summarized device count.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceResourcePerformanceScore sets the deviceResourcePerformanceScore property value. Resource performance score of a specific device. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetDeviceResourcePerformanceScore(value *int32)() {
    m.deviceResourcePerformanceScore = value
}
// SetManufacturer sets the manufacturer property value. The user experience analytics device manufacturer.
func (m *UserExperienceAnalyticsResourcePerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsResourcePerformance) SetModel(value *string)() {
    m.model = value
}
// SetRamSpikeTimePercentage sets the ramSpikeTimePercentage property value. RAM spike time in percentage. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentage(value *float64)() {
    m.ramSpikeTimePercentage = value
}
// SetRamSpikeTimePercentageThreshold sets the ramSpikeTimePercentageThreshold property value. Threshold of ramSpikeTimeScore. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimePercentageThreshold(value *float64)() {
    m.ramSpikeTimePercentageThreshold = value
}
// SetRamSpikeTimeScore sets the ramSpikeTimeScore property value. The user experience analytics device RAM spike time score. Valid values 0 to 100
func (m *UserExperienceAnalyticsResourcePerformance) SetRamSpikeTimeScore(value *int32)() {
    m.ramSpikeTimeScore = value
}

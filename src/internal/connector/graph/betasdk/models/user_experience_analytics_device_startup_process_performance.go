package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessPerformance the user experience analytics device startup process performance.
type UserExperienceAnalyticsDeviceStartupProcessPerformance struct {
    Entity
    // User experience analytics device startup process summarized count.
    deviceCount *int64
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs *int32
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs2 *int64
    // User experience analytics device startup process name.
    processName *string
    // The user experience analytics device startup process product name.
    productName *string
    // The User experience analytics device startup process publisher.
    publisher *string
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs *int32
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs2 *int64
}
// NewUserExperienceAnalyticsDeviceStartupProcessPerformance instantiates a new userExperienceAnalyticsDeviceStartupProcessPerformance and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformance) {
    m := &UserExperienceAnalyticsDeviceStartupProcessPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformance(), nil
}
// GetDeviceCount gets the deviceCount property value. User experience analytics device startup process summarized count.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetDeviceCount()(*int64) {
    return m.deviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["medianImpactInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianImpactInMs(val)
        }
        return nil
    }
    res["medianImpactInMs2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianImpactInMs2(val)
        }
        return nil
    }
    res["processName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["totalImpactInMs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalImpactInMs(val)
        }
        return nil
    }
    res["totalImpactInMs2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalImpactInMs2(val)
        }
        return nil
    }
    return res
}
// GetMedianImpactInMs gets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs()(*int32) {
    return m.medianImpactInMs
}
// GetMedianImpactInMs2 gets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs2()(*int64) {
    return m.medianImpactInMs2
}
// GetProcessName gets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProcessName()(*string) {
    return m.processName
}
// GetProductName gets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProductName()(*string) {
    return m.productName
}
// GetPublisher gets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetPublisher()(*string) {
    return m.publisher
}
// GetTotalImpactInMs gets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs()(*int32) {
    return m.totalImpactInMs
}
// GetTotalImpactInMs2 gets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs2()(*int64) {
    return m.totalImpactInMs2
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianImpactInMs", m.GetMedianImpactInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("medianImpactInMs2", m.GetMedianImpactInMs2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalImpactInMs", m.GetTotalImpactInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalImpactInMs2", m.GetTotalImpactInMs2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCount sets the deviceCount property value. User experience analytics device startup process summarized count.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
// SetMedianImpactInMs sets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs(value *int32)() {
    m.medianImpactInMs = value
}
// SetMedianImpactInMs2 sets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs2(value *int64)() {
    m.medianImpactInMs2 = value
}
// SetProcessName sets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProcessName(value *string)() {
    m.processName = value
}
// SetProductName sets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProductName(value *string)() {
    m.productName = value
}
// SetPublisher sets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetPublisher(value *string)() {
    m.publisher = value
}
// SetTotalImpactInMs sets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs(value *int32)() {
    m.totalImpactInMs = value
}
// SetTotalImpactInMs2 sets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs2(value *int64)() {
    m.totalImpactInMs2 = value
}

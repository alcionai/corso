package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthOSVersionPerformance the user experience analytics device OS version performance entity contains OS version performance details.
type UserExperienceAnalyticsAppHealthOSVersionPerformance struct {
    Entity
    // The number of active devices for the OS version. Valid values -2147483648 to 2147483647
    activeDeviceCount *int32
    // The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
    meanTimeToFailureInMinutes *int32
    // The OS build number installed on the device.
    osBuildNumber *string
    // The OS version installed on the device.
    osVersion *string
    // The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    osVersionAppHealthScore *float64
    // The overall app health status of the OS version.
    osVersionAppHealthStatus *string
}
// NewUserExperienceAnalyticsAppHealthOSVersionPerformance instantiates a new userExperienceAnalyticsAppHealthOSVersionPerformance and sets the default values.
func NewUserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserExperienceAnalyticsAppHealthOSVersionPerformance) {
    m := &UserExperienceAnalyticsAppHealthOSVersionPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsAppHealthOSVersionPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAppHealthOSVersionPerformance(), nil
}
// GetActiveDeviceCount gets the activeDeviceCount property value. The number of active devices for the OS version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetActiveDeviceCount()(*int32) {
    return m.activeDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["osVersionAppHealthScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersionAppHealthScore(val)
        }
        return nil
    }
    res["osVersionAppHealthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersionAppHealthStatus(val)
        }
        return nil
    }
    return res
}
// GetMeanTimeToFailureInMinutes gets the meanTimeToFailureInMinutes property value. The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetMeanTimeToFailureInMinutes()(*int32) {
    return m.meanTimeToFailureInMinutes
}
// GetOsBuildNumber gets the osBuildNumber property value. The OS build number installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsBuildNumber()(*string) {
    return m.osBuildNumber
}
// GetOsVersion gets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOsVersionAppHealthScore gets the osVersionAppHealthScore property value. The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersionAppHealthScore()(*float64) {
    return m.osVersionAppHealthScore
}
// GetOsVersionAppHealthStatus gets the osVersionAppHealthStatus property value. The overall app health status of the OS version.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) GetOsVersionAppHealthStatus()(*string) {
    return m.osVersionAppHealthStatus
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteFloat64Value("osVersionAppHealthScore", m.GetOsVersionAppHealthScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersionAppHealthStatus", m.GetOsVersionAppHealthStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDeviceCount sets the activeDeviceCount property value. The number of active devices for the OS version. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetActiveDeviceCount(value *int32)() {
    m.activeDeviceCount = value
}
// SetMeanTimeToFailureInMinutes sets the meanTimeToFailureInMinutes property value. The mean time to failure for the OS version in minutes. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetMeanTimeToFailureInMinutes(value *int32)() {
    m.meanTimeToFailureInMinutes = value
}
// SetOsBuildNumber sets the osBuildNumber property value. The OS build number installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// SetOsVersion sets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOsVersionAppHealthScore sets the osVersionAppHealthScore property value. The app health score of the OS version. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersionAppHealthScore(value *float64)() {
    m.osVersionAppHealthScore = value
}
// SetOsVersionAppHealthStatus sets the osVersionAppHealthStatus property value. The overall app health status of the OS version.
func (m *UserExperienceAnalyticsAppHealthOSVersionPerformance) SetOsVersionAppHealthStatus(value *string)() {
    m.osVersionAppHealthStatus = value
}

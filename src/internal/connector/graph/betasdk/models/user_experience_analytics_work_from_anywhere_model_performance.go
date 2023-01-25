package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereModelPerformance the user experience analytics work from anywhere model performance.
type UserExperienceAnalyticsWorkFromAnywhereModelPerformance struct {
    Entity
    // The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudIdentityScore *float64
    // The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudManagementScore *float64
    // The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudProvisioningScore *float64
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The user experience work from anywhere's manufacturer name of the devices.
    manufacturer *string
    // The user experience work from anywhere's model name of the devices.
    model *string
    // The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
    modelDeviceCount *int32
    // The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    windowsScore *float64
    // The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance instantiates a new userExperienceAnalyticsWorkFromAnywhereModelPerformance and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserExperienceAnalyticsWorkFromAnywhereModelPerformance) {
    m := &UserExperienceAnalyticsWorkFromAnywhereModelPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereModelPerformanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereModelPerformance(), nil
}
// GetCloudIdentityScore gets the cloudIdentityScore property value. The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudIdentityScore()(*float64) {
    return m.cloudIdentityScore
}
// GetCloudManagementScore gets the cloudManagementScore property value. The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudManagementScore()(*float64) {
    return m.cloudManagementScore
}
// GetCloudProvisioningScore gets the cloudProvisioningScore property value. The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetCloudProvisioningScore()(*float64) {
    return m.cloudProvisioningScore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudIdentityScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudIdentityScore(val)
        }
        return nil
    }
    res["cloudManagementScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudManagementScore(val)
        }
        return nil
    }
    res["cloudProvisioningScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudProvisioningScore(val)
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
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModelDeviceCount(val)
        }
        return nil
    }
    res["windowsScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsScore(val)
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
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetManufacturer gets the manufacturer property value. The user experience work from anywhere's manufacturer name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience work from anywhere's model name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModel()(*string) {
    return m.model
}
// GetModelDeviceCount gets the modelDeviceCount property value. The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetModelDeviceCount()(*int32) {
    return m.modelDeviceCount
}
// GetWindowsScore gets the windowsScore property value. The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWindowsScore()(*float64) {
    return m.windowsScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("cloudIdentityScore", m.GetCloudIdentityScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudManagementScore", m.GetCloudManagementScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudProvisioningScore", m.GetCloudProvisioningScore())
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
        err = writer.WriteInt32Value("modelDeviceCount", m.GetModelDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("windowsScore", m.GetWindowsScore())
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
// SetCloudIdentityScore sets the cloudIdentityScore property value. The user experience work from anywhere's cloud identity score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudIdentityScore(value *float64)() {
    m.cloudIdentityScore = value
}
// SetCloudManagementScore sets the cloudManagementScore property value. The user experience work from anywhere's cloud management score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudManagementScore(value *float64)() {
    m.cloudManagementScore = value
}
// SetCloudProvisioningScore sets the cloudProvisioningScore property value. The user experience work from anywhere's cloud provisioning score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetCloudProvisioningScore(value *float64)() {
    m.cloudProvisioningScore = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetManufacturer sets the manufacturer property value. The user experience work from anywhere's manufacturer name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience work from anywhere's model name of the devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModel(value *string)() {
    m.model = value
}
// SetModelDeviceCount sets the modelDeviceCount property value. The user experience work from anywhere's devices count for the model. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetModelDeviceCount(value *int32)() {
    m.modelDeviceCount = value
}
// SetWindowsScore sets the windowsScore property value. The user experience work from anywhere windows score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWindowsScore(value *float64)() {
    m.windowsScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The user experience work from anywhere overall score for the model. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereModelPerformance) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}

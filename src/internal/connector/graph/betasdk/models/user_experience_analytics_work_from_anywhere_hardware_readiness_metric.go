package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric 
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric struct {
    Entity
    // The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    osCheckFailedPercentage *float64
    // The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processor64BitCheckFailedPercentage *float64
    // The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorCoreCountCheckFailedPercentage *float64
    // The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorFamilyCheckFailedPercentage *float64
    // The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    processorSpeedCheckFailedPercentage *float64
    // The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    ramCheckFailedPercentage *float64
    // The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    secureBootCheckFailedPercentage *float64
    // The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    storageCheckFailedPercentage *float64
    // The count of total devices in an organization. Valid values -2147483648 to 2147483647
    totalDeviceCount *int32
    // The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    tpmCheckFailedPercentage *float64
    // The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
    upgradeEligibleDeviceCount *int32
}
// NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric instantiates a new userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) {
    m := &UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["osCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsCheckFailedPercentage(val)
        }
        return nil
    }
    res["processor64BitCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessor64BitCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorCoreCountCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCoreCountCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorFamilyCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorFamilyCheckFailedPercentage(val)
        }
        return nil
    }
    res["processorSpeedCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedCheckFailedPercentage(val)
        }
        return nil
    }
    res["ramCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamCheckFailedPercentage(val)
        }
        return nil
    }
    res["secureBootCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootCheckFailedPercentage(val)
        }
        return nil
    }
    res["storageCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageCheckFailedPercentage(val)
        }
        return nil
    }
    res["totalDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalDeviceCount(val)
        }
        return nil
    }
    res["tpmCheckFailedPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmCheckFailedPercentage(val)
        }
        return nil
    }
    res["upgradeEligibleDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeEligibleDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetOsCheckFailedPercentage gets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetOsCheckFailedPercentage()(*float64) {
    return m.osCheckFailedPercentage
}
// GetProcessor64BitCheckFailedPercentage gets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessor64BitCheckFailedPercentage()(*float64) {
    return m.processor64BitCheckFailedPercentage
}
// GetProcessorCoreCountCheckFailedPercentage gets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorCoreCountCheckFailedPercentage()(*float64) {
    return m.processorCoreCountCheckFailedPercentage
}
// GetProcessorFamilyCheckFailedPercentage gets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorFamilyCheckFailedPercentage()(*float64) {
    return m.processorFamilyCheckFailedPercentage
}
// GetProcessorSpeedCheckFailedPercentage gets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetProcessorSpeedCheckFailedPercentage()(*float64) {
    return m.processorSpeedCheckFailedPercentage
}
// GetRamCheckFailedPercentage gets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetRamCheckFailedPercentage()(*float64) {
    return m.ramCheckFailedPercentage
}
// GetSecureBootCheckFailedPercentage gets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetSecureBootCheckFailedPercentage()(*float64) {
    return m.secureBootCheckFailedPercentage
}
// GetStorageCheckFailedPercentage gets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetStorageCheckFailedPercentage()(*float64) {
    return m.storageCheckFailedPercentage
}
// GetTotalDeviceCount gets the totalDeviceCount property value. The count of total devices in an organization. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTotalDeviceCount()(*int32) {
    return m.totalDeviceCount
}
// GetTpmCheckFailedPercentage gets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetTpmCheckFailedPercentage()(*float64) {
    return m.tpmCheckFailedPercentage
}
// GetUpgradeEligibleDeviceCount gets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) GetUpgradeEligibleDeviceCount()(*int32) {
    return m.upgradeEligibleDeviceCount
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("osCheckFailedPercentage", m.GetOsCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processor64BitCheckFailedPercentage", m.GetProcessor64BitCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorCoreCountCheckFailedPercentage", m.GetProcessorCoreCountCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorFamilyCheckFailedPercentage", m.GetProcessorFamilyCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("processorSpeedCheckFailedPercentage", m.GetProcessorSpeedCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("ramCheckFailedPercentage", m.GetRamCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("secureBootCheckFailedPercentage", m.GetSecureBootCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("storageCheckFailedPercentage", m.GetStorageCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalDeviceCount", m.GetTotalDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("tpmCheckFailedPercentage", m.GetTpmCheckFailedPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upgradeEligibleDeviceCount", m.GetUpgradeEligibleDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOsCheckFailedPercentage sets the osCheckFailedPercentage property value. The percentage of devices for which OS check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetOsCheckFailedPercentage(value *float64)() {
    m.osCheckFailedPercentage = value
}
// SetProcessor64BitCheckFailedPercentage sets the processor64BitCheckFailedPercentage property value. The percentage of devices for which processor hardware 64-bit architecture check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessor64BitCheckFailedPercentage(value *float64)() {
    m.processor64BitCheckFailedPercentage = value
}
// SetProcessorCoreCountCheckFailedPercentage sets the processorCoreCountCheckFailedPercentage property value. The percentage of devices for which processor hardware core count check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorCoreCountCheckFailedPercentage(value *float64)() {
    m.processorCoreCountCheckFailedPercentage = value
}
// SetProcessorFamilyCheckFailedPercentage sets the processorFamilyCheckFailedPercentage property value. The percentage of devices for which processor hardware family check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorFamilyCheckFailedPercentage(value *float64)() {
    m.processorFamilyCheckFailedPercentage = value
}
// SetProcessorSpeedCheckFailedPercentage sets the processorSpeedCheckFailedPercentage property value. The percentage of devices for which processor hardware speed check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetProcessorSpeedCheckFailedPercentage(value *float64)() {
    m.processorSpeedCheckFailedPercentage = value
}
// SetRamCheckFailedPercentage sets the ramCheckFailedPercentage property value. The percentage of devices for which RAM hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetRamCheckFailedPercentage(value *float64)() {
    m.ramCheckFailedPercentage = value
}
// SetSecureBootCheckFailedPercentage sets the secureBootCheckFailedPercentage property value. The percentage of devices for which secure boot hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetSecureBootCheckFailedPercentage(value *float64)() {
    m.secureBootCheckFailedPercentage = value
}
// SetStorageCheckFailedPercentage sets the storageCheckFailedPercentage property value. The percentage of devices for which storage hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetStorageCheckFailedPercentage(value *float64)() {
    m.storageCheckFailedPercentage = value
}
// SetTotalDeviceCount sets the totalDeviceCount property value. The count of total devices in an organization. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTotalDeviceCount(value *int32)() {
    m.totalDeviceCount = value
}
// SetTpmCheckFailedPercentage sets the tpmCheckFailedPercentage property value. The percentage of devices for which Trusted Platform Module (TPM) hardware check has failed. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetTpmCheckFailedPercentage(value *float64)() {
    m.tpmCheckFailedPercentage = value
}
// SetUpgradeEligibleDeviceCount sets the upgradeEligibleDeviceCount property value. The count of devices in an organization eligible for windows upgrade. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric) SetUpgradeEligibleDeviceCount(value *int32)() {
    m.upgradeEligibleDeviceCount = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereDevice the user experience analytics Device for work from anywhere report
type UserExperienceAnalyticsWorkFromAnywhereDevice struct {
    Entity
    // The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
    autoPilotProfileAssigned *bool
    // The user experience work from anywhere intune device's autopilotRegistered.
    autoPilotRegistered *bool
    // The user experience work from anywhere azure Ad device Id.
    azureAdDeviceId *string
    // The user experience work from anywhere device's azure Ad joinType.
    azureAdJoinType *string
    // The user experience work from anywhere device's azureAdRegistered.
    azureAdRegistered *bool
    // The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudIdentityScore *float64
    // The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudManagementScore *float64
    // The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    cloudProvisioningScore *float64
    // The user experience work from anywhere device's compliancePolicySetToIntune.
    compliancePolicySetToIntune *bool
    // The user experience work from anywhere device Id.
    deviceId *string
    // The work from anywhere device's name.
    deviceName *string
    // The healthStatus property
    healthStatus *UserExperienceAnalyticsHealthState
    // The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
    isCloudManagedGatewayEnabled *bool
    // The user experience work from anywhere management agent of the device.
    managedBy *string
    // The user experience work from anywhere device's manufacturer.
    manufacturer *string
    // The user experience work from anywhere device's model.
    model *string
    // The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
    osCheckFailed *bool
    // The user experience work from anywhere device's OS Description.
    osDescription *string
    // The user experience work from anywhere device's OS Version.
    osVersion *string
    // The user experience work from anywhere device's otherWorkloadsSetToIntune.
    otherWorkloadsSetToIntune *bool
    // The user experience work from anywhere device's ownership.
    ownership *string
    // The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
    processor64BitCheckFailed *bool
    // The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
    processorCoreCountCheckFailed *bool
    // The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
    processorFamilyCheckFailed *bool
    // The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
    processorSpeedCheckFailed *bool
    // Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
    ramCheckFailed *bool
    // The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
    secureBootCheckFailed *bool
    // The user experience work from anywhere device's serial number.
    serialNumber *string
    // The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
    storageCheckFailed *bool
    // The user experience work from anywhere device's tenantAttached.
    tenantAttached *bool
    // The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
    tpmCheckFailed *bool
    // Work From Anywhere windows device upgrade eligibility status
    upgradeEligibility *OperatingSystemUpgradeEligibility
    // The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    windowsScore *float64
    // The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
    workFromAnywhereScore *float64
}
// NewUserExperienceAnalyticsWorkFromAnywhereDevice instantiates a new userExperienceAnalyticsWorkFromAnywhereDevice and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDevice()(*UserExperienceAnalyticsWorkFromAnywhereDevice) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsWorkFromAnywhereDevice(), nil
}
// GetAutoPilotProfileAssigned gets the autoPilotProfileAssigned property value. The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotProfileAssigned()(*bool) {
    return m.autoPilotProfileAssigned
}
// GetAutoPilotRegistered gets the autoPilotRegistered property value. The user experience work from anywhere intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAutoPilotRegistered()(*bool) {
    return m.autoPilotRegistered
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. The user experience work from anywhere azure Ad device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetAzureAdJoinType gets the azureAdJoinType property value. The user experience work from anywhere device's azure Ad joinType.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdJoinType()(*string) {
    return m.azureAdJoinType
}
// GetAzureAdRegistered gets the azureAdRegistered property value. The user experience work from anywhere device's azureAdRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetAzureAdRegistered()(*bool) {
    return m.azureAdRegistered
}
// GetCloudIdentityScore gets the cloudIdentityScore property value. The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudIdentityScore()(*float64) {
    return m.cloudIdentityScore
}
// GetCloudManagementScore gets the cloudManagementScore property value. The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudManagementScore()(*float64) {
    return m.cloudManagementScore
}
// GetCloudProvisioningScore gets the cloudProvisioningScore property value. The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCloudProvisioningScore()(*float64) {
    return m.cloudProvisioningScore
}
// GetCompliancePolicySetToIntune gets the compliancePolicySetToIntune property value. The user experience work from anywhere device's compliancePolicySetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetCompliancePolicySetToIntune()(*bool) {
    return m.compliancePolicySetToIntune
}
// GetDeviceId gets the deviceId property value. The user experience work from anywhere device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The work from anywhere device's name.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotProfileAssigned(val)
        }
        return nil
    }
    res["autoPilotRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotRegistered(val)
        }
        return nil
    }
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["azureAdJoinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdJoinType(val)
        }
        return nil
    }
    res["azureAdRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdRegistered(val)
        }
        return nil
    }
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
    res["compliancePolicySetToIntune"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliancePolicySetToIntune(val)
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
    res["isCloudManagedGatewayEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCloudManagedGatewayEnabled(val)
        }
        return nil
    }
    res["managedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val)
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
    res["osCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsCheckFailed(val)
        }
        return nil
    }
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
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
    res["otherWorkloadsSetToIntune"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherWorkloadsSetToIntune(val)
        }
        return nil
    }
    res["ownership"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnership(val)
        }
        return nil
    }
    res["processor64BitCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessor64BitCheckFailed(val)
        }
        return nil
    }
    res["processorCoreCountCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorCoreCountCheckFailed(val)
        }
        return nil
    }
    res["processorFamilyCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorFamilyCheckFailed(val)
        }
        return nil
    }
    res["processorSpeedCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorSpeedCheckFailed(val)
        }
        return nil
    }
    res["ramCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRamCheckFailed(val)
        }
        return nil
    }
    res["secureBootCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootCheckFailed(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["storageCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageCheckFailed(val)
        }
        return nil
    }
    res["tenantAttached"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttached(val)
        }
        return nil
    }
    res["tpmCheckFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmCheckFailed(val)
        }
        return nil
    }
    res["upgradeEligibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperatingSystemUpgradeEligibility)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeEligibility(val.(*OperatingSystemUpgradeEligibility))
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
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetHealthStatus()(*UserExperienceAnalyticsHealthState) {
    return m.healthStatus
}
// GetIsCloudManagedGatewayEnabled gets the isCloudManagedGatewayEnabled property value. The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetIsCloudManagedGatewayEnabled()(*bool) {
    return m.isCloudManagedGatewayEnabled
}
// GetManagedBy gets the managedBy property value. The user experience work from anywhere management agent of the device.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManagedBy()(*string) {
    return m.managedBy
}
// GetManufacturer gets the manufacturer property value. The user experience work from anywhere device's manufacturer.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience work from anywhere device's model.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetModel()(*string) {
    return m.model
}
// GetOsCheckFailed gets the osCheckFailed property value. The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsCheckFailed()(*bool) {
    return m.osCheckFailed
}
// GetOsDescription gets the osDescription property value. The user experience work from anywhere device's OS Description.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. The user experience work from anywhere device's OS Version.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOtherWorkloadsSetToIntune gets the otherWorkloadsSetToIntune property value. The user experience work from anywhere device's otherWorkloadsSetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOtherWorkloadsSetToIntune()(*bool) {
    return m.otherWorkloadsSetToIntune
}
// GetOwnership gets the ownership property value. The user experience work from anywhere device's ownership.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetOwnership()(*string) {
    return m.ownership
}
// GetProcessor64BitCheckFailed gets the processor64BitCheckFailed property value. The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessor64BitCheckFailed()(*bool) {
    return m.processor64BitCheckFailed
}
// GetProcessorCoreCountCheckFailed gets the processorCoreCountCheckFailed property value. The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorCoreCountCheckFailed()(*bool) {
    return m.processorCoreCountCheckFailed
}
// GetProcessorFamilyCheckFailed gets the processorFamilyCheckFailed property value. The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorFamilyCheckFailed()(*bool) {
    return m.processorFamilyCheckFailed
}
// GetProcessorSpeedCheckFailed gets the processorSpeedCheckFailed property value. The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetProcessorSpeedCheckFailed()(*bool) {
    return m.processorSpeedCheckFailed
}
// GetRamCheckFailed gets the ramCheckFailed property value. Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetRamCheckFailed()(*bool) {
    return m.ramCheckFailed
}
// GetSecureBootCheckFailed gets the secureBootCheckFailed property value. The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSecureBootCheckFailed()(*bool) {
    return m.secureBootCheckFailed
}
// GetSerialNumber gets the serialNumber property value. The user experience work from anywhere device's serial number.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetStorageCheckFailed gets the storageCheckFailed property value. The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetStorageCheckFailed()(*bool) {
    return m.storageCheckFailed
}
// GetTenantAttached gets the tenantAttached property value. The user experience work from anywhere device's tenantAttached.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTenantAttached()(*bool) {
    return m.tenantAttached
}
// GetTpmCheckFailed gets the tpmCheckFailed property value. The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetTpmCheckFailed()(*bool) {
    return m.tpmCheckFailed
}
// GetUpgradeEligibility gets the upgradeEligibility property value. Work From Anywhere windows device upgrade eligibility status
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetUpgradeEligibility()(*OperatingSystemUpgradeEligibility) {
    return m.upgradeEligibility
}
// GetWindowsScore gets the windowsScore property value. The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWindowsScore()(*float64) {
    return m.windowsScore
}
// GetWorkFromAnywhereScore gets the workFromAnywhereScore property value. The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) GetWorkFromAnywhereScore()(*float64) {
    return m.workFromAnywhereScore
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoPilotProfileAssigned", m.GetAutoPilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoPilotRegistered", m.GetAutoPilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdJoinType", m.GetAzureAdJoinType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureAdRegistered", m.GetAzureAdRegistered())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteBoolValue("compliancePolicySetToIntune", m.GetCompliancePolicySetToIntune())
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
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCloudManagedGatewayEnabled", m.GetIsCloudManagedGatewayEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedBy", m.GetManagedBy())
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
        err = writer.WriteBoolValue("osCheckFailed", m.GetOsCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
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
        err = writer.WriteBoolValue("otherWorkloadsSetToIntune", m.GetOtherWorkloadsSetToIntune())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownership", m.GetOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processor64BitCheckFailed", m.GetProcessor64BitCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorCoreCountCheckFailed", m.GetProcessorCoreCountCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorFamilyCheckFailed", m.GetProcessorFamilyCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("processorSpeedCheckFailed", m.GetProcessorSpeedCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ramCheckFailed", m.GetRamCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("secureBootCheckFailed", m.GetSecureBootCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageCheckFailed", m.GetStorageCheckFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tenantAttached", m.GetTenantAttached())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tpmCheckFailed", m.GetTpmCheckFailed())
        if err != nil {
            return err
        }
    }
    if m.GetUpgradeEligibility() != nil {
        cast := (*m.GetUpgradeEligibility()).String()
        err = writer.WriteStringValue("upgradeEligibility", &cast)
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
// SetAutoPilotProfileAssigned sets the autoPilotProfileAssigned property value. The user experience analytics work from anywhere intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotProfileAssigned(value *bool)() {
    m.autoPilotProfileAssigned = value
}
// SetAutoPilotRegistered sets the autoPilotRegistered property value. The user experience work from anywhere intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAutoPilotRegistered(value *bool)() {
    m.autoPilotRegistered = value
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. The user experience work from anywhere azure Ad device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// SetAzureAdJoinType sets the azureAdJoinType property value. The user experience work from anywhere device's azure Ad joinType.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdJoinType(value *string)() {
    m.azureAdJoinType = value
}
// SetAzureAdRegistered sets the azureAdRegistered property value. The user experience work from anywhere device's azureAdRegistered.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetAzureAdRegistered(value *bool)() {
    m.azureAdRegistered = value
}
// SetCloudIdentityScore sets the cloudIdentityScore property value. The user experience work from anywhere per device cloud identity score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudIdentityScore(value *float64)() {
    m.cloudIdentityScore = value
}
// SetCloudManagementScore sets the cloudManagementScore property value. The user experience work from anywhere per device cloud management score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudManagementScore(value *float64)() {
    m.cloudManagementScore = value
}
// SetCloudProvisioningScore sets the cloudProvisioningScore property value. The user experience work from anywhere per device cloud provisioning score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCloudProvisioningScore(value *float64)() {
    m.cloudProvisioningScore = value
}
// SetCompliancePolicySetToIntune sets the compliancePolicySetToIntune property value. The user experience work from anywhere device's compliancePolicySetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetCompliancePolicySetToIntune(value *bool)() {
    m.compliancePolicySetToIntune = value
}
// SetDeviceId sets the deviceId property value. The user experience work from anywhere device Id.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The work from anywhere device's name.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetHealthStatus(value *UserExperienceAnalyticsHealthState)() {
    m.healthStatus = value
}
// SetIsCloudManagedGatewayEnabled sets the isCloudManagedGatewayEnabled property value. The user experience work from anywhere device's Cloud Management Gateway for Configuration Manager is enabled.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetIsCloudManagedGatewayEnabled(value *bool)() {
    m.isCloudManagedGatewayEnabled = value
}
// SetManagedBy sets the managedBy property value. The user experience work from anywhere management agent of the device.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManagedBy(value *string)() {
    m.managedBy = value
}
// SetManufacturer sets the manufacturer property value. The user experience work from anywhere device's manufacturer.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience work from anywhere device's model.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetModel(value *string)() {
    m.model = value
}
// SetOsCheckFailed sets the osCheckFailed property value. The user experience work from anywhere device, Is OS check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsCheckFailed(value *bool)() {
    m.osCheckFailed = value
}
// SetOsDescription sets the osDescription property value. The user experience work from anywhere device's OS Description.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. The user experience work from anywhere device's OS Version.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOtherWorkloadsSetToIntune sets the otherWorkloadsSetToIntune property value. The user experience work from anywhere device's otherWorkloadsSetToIntune.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOtherWorkloadsSetToIntune(value *bool)() {
    m.otherWorkloadsSetToIntune = value
}
// SetOwnership sets the ownership property value. The user experience work from anywhere device's ownership.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetOwnership(value *string)() {
    m.ownership = value
}
// SetProcessor64BitCheckFailed sets the processor64BitCheckFailed property value. The user experience work from anywhere device, Is processor hardware 64-bit architecture check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessor64BitCheckFailed(value *bool)() {
    m.processor64BitCheckFailed = value
}
// SetProcessorCoreCountCheckFailed sets the processorCoreCountCheckFailed property value. The user experience work from anywhere device, Is processor hardware core count check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorCoreCountCheckFailed(value *bool)() {
    m.processorCoreCountCheckFailed = value
}
// SetProcessorFamilyCheckFailed sets the processorFamilyCheckFailed property value. The user experience work from anywhere device, Is processor hardware family check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorFamilyCheckFailed(value *bool)() {
    m.processorFamilyCheckFailed = value
}
// SetProcessorSpeedCheckFailed sets the processorSpeedCheckFailed property value. The user experience work from anywhere device, Is processor hardware speed check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetProcessorSpeedCheckFailed(value *bool)() {
    m.processorSpeedCheckFailed = value
}
// SetRamCheckFailed sets the ramCheckFailed property value. Is the user experience analytics work from anywhere device RAM hardware check failed for device to upgrade to the latest version of windows
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetRamCheckFailed(value *bool)() {
    m.ramCheckFailed = value
}
// SetSecureBootCheckFailed sets the secureBootCheckFailed property value. The user experience work from anywhere device, Is secure boot hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSecureBootCheckFailed(value *bool)() {
    m.secureBootCheckFailed = value
}
// SetSerialNumber sets the serialNumber property value. The user experience work from anywhere device's serial number.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetStorageCheckFailed sets the storageCheckFailed property value. The user experience work from anywhere device, Is storage hardware check failed for device to upgrade to the latest version of windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetStorageCheckFailed(value *bool)() {
    m.storageCheckFailed = value
}
// SetTenantAttached sets the tenantAttached property value. The user experience work from anywhere device's tenantAttached.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTenantAttached(value *bool)() {
    m.tenantAttached = value
}
// SetTpmCheckFailed sets the tpmCheckFailed property value. The user experience work from anywhere device, Is Trusted Platform Module (TPM) hardware check failed for device to the latest version of upgrade to windows.
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetTpmCheckFailed(value *bool)() {
    m.tpmCheckFailed = value
}
// SetUpgradeEligibility sets the upgradeEligibility property value. Work From Anywhere windows device upgrade eligibility status
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetUpgradeEligibility(value *OperatingSystemUpgradeEligibility)() {
    m.upgradeEligibility = value
}
// SetWindowsScore sets the windowsScore property value. The user experience work from anywhere per device windows score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWindowsScore(value *float64)() {
    m.windowsScore = value
}
// SetWorkFromAnywhereScore sets the workFromAnywhereScore property value. The user experience work from anywhere per device overall score. Valid values -1.79769313486232E+308 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsWorkFromAnywhereDevice) SetWorkFromAnywhereScore(value *float64)() {
    m.workFromAnywhereScore = value
}

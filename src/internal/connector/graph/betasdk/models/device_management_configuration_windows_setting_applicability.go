package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationWindowsSettingApplicability 
type DeviceManagementConfigurationWindowsSettingApplicability struct {
    DeviceManagementConfigurationSettingApplicability
    // Version of CSP setting is a part of
    configurationServiceProviderVersion *string
    // Maximum supported version of Windows
    maximumSupportedVersion *string
    // Minimum supported version of Windows
    minimumSupportedVersion *string
    // Required AAD Trust Type
    requiredAzureAdTrustType *DeviceManagementConfigurationAzureAdTrustType
    // AzureAD setting requirement
    requiresAzureAd *bool
    // List of Windows SKUs that the setting is applicable for
    windowsSkus []DeviceManagementConfigurationWindowsSkus
}
// NewDeviceManagementConfigurationWindowsSettingApplicability instantiates a new DeviceManagementConfigurationWindowsSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationWindowsSettingApplicability()(*DeviceManagementConfigurationWindowsSettingApplicability) {
    m := &DeviceManagementConfigurationWindowsSettingApplicability{
        DeviceManagementConfigurationSettingApplicability: *NewDeviceManagementConfigurationSettingApplicability(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationWindowsSettingApplicability";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationWindowsSettingApplicabilityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationWindowsSettingApplicabilityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationWindowsSettingApplicability(), nil
}
// GetConfigurationServiceProviderVersion gets the configurationServiceProviderVersion property value. Version of CSP setting is a part of
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetConfigurationServiceProviderVersion()(*string) {
    return m.configurationServiceProviderVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingApplicability.GetFieldDeserializers()
    res["configurationServiceProviderVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationServiceProviderVersion(val)
        }
        return nil
    }
    res["maximumSupportedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumSupportedVersion(val)
        }
        return nil
    }
    res["minimumSupportedVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedVersion(val)
        }
        return nil
    }
    res["requiredAzureAdTrustType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationAzureAdTrustType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiredAzureAdTrustType(val.(*DeviceManagementConfigurationAzureAdTrustType))
        }
        return nil
    }
    res["requiresAzureAd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresAzureAd(val)
        }
        return nil
    }
    res["windowsSkus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDeviceManagementConfigurationWindowsSkus)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationWindowsSkus, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementConfigurationWindowsSkus))
            }
            m.SetWindowsSkus(res)
        }
        return nil
    }
    return res
}
// GetMaximumSupportedVersion gets the maximumSupportedVersion property value. Maximum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetMaximumSupportedVersion()(*string) {
    return m.maximumSupportedVersion
}
// GetMinimumSupportedVersion gets the minimumSupportedVersion property value. Minimum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetMinimumSupportedVersion()(*string) {
    return m.minimumSupportedVersion
}
// GetRequiredAzureAdTrustType gets the requiredAzureAdTrustType property value. Required AAD Trust Type
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetRequiredAzureAdTrustType()(*DeviceManagementConfigurationAzureAdTrustType) {
    return m.requiredAzureAdTrustType
}
// GetRequiresAzureAd gets the requiresAzureAd property value. AzureAD setting requirement
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetRequiresAzureAd()(*bool) {
    return m.requiresAzureAd
}
// GetWindowsSkus gets the windowsSkus property value. List of Windows SKUs that the setting is applicable for
func (m *DeviceManagementConfigurationWindowsSettingApplicability) GetWindowsSkus()([]DeviceManagementConfigurationWindowsSkus) {
    return m.windowsSkus
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationWindowsSettingApplicability) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingApplicability.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("configurationServiceProviderVersion", m.GetConfigurationServiceProviderVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumSupportedVersion", m.GetMaximumSupportedVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumSupportedVersion", m.GetMinimumSupportedVersion())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredAzureAdTrustType() != nil {
        cast := (*m.GetRequiredAzureAdTrustType()).String()
        err = writer.WriteStringValue("requiredAzureAdTrustType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresAzureAd", m.GetRequiresAzureAd())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsSkus() != nil {
        err = writer.WriteCollectionOfStringValues("windowsSkus", SerializeDeviceManagementConfigurationWindowsSkus(m.GetWindowsSkus()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationServiceProviderVersion sets the configurationServiceProviderVersion property value. Version of CSP setting is a part of
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetConfigurationServiceProviderVersion(value *string)() {
    m.configurationServiceProviderVersion = value
}
// SetMaximumSupportedVersion sets the maximumSupportedVersion property value. Maximum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetMaximumSupportedVersion(value *string)() {
    m.maximumSupportedVersion = value
}
// SetMinimumSupportedVersion sets the minimumSupportedVersion property value. Minimum supported version of Windows
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetMinimumSupportedVersion(value *string)() {
    m.minimumSupportedVersion = value
}
// SetRequiredAzureAdTrustType sets the requiredAzureAdTrustType property value. Required AAD Trust Type
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetRequiredAzureAdTrustType(value *DeviceManagementConfigurationAzureAdTrustType)() {
    m.requiredAzureAdTrustType = value
}
// SetRequiresAzureAd sets the requiresAzureAd property value. AzureAD setting requirement
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetRequiresAzureAd(value *bool)() {
    m.requiresAzureAd = value
}
// SetWindowsSkus sets the windowsSkus property value. List of Windows SKUs that the setting is applicable for
func (m *DeviceManagementConfigurationWindowsSettingApplicability) SetWindowsSkus(value []DeviceManagementConfigurationWindowsSkus)() {
    m.windowsSkus = value
}

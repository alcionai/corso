package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationWindowsSettingApplicabilityable 
type DeviceManagementConfigurationWindowsSettingApplicabilityable interface {
    DeviceManagementConfigurationSettingApplicabilityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationServiceProviderVersion()(*string)
    GetMaximumSupportedVersion()(*string)
    GetMinimumSupportedVersion()(*string)
    GetRequiredAzureAdTrustType()(*DeviceManagementConfigurationAzureAdTrustType)
    GetRequiresAzureAd()(*bool)
    GetWindowsSkus()([]DeviceManagementConfigurationWindowsSkus)
    SetConfigurationServiceProviderVersion(value *string)()
    SetMaximumSupportedVersion(value *string)()
    SetMinimumSupportedVersion(value *string)()
    SetRequiredAzureAdTrustType(value *DeviceManagementConfigurationAzureAdTrustType)()
    SetRequiresAzureAd(value *bool)()
    SetWindowsSkus(value []DeviceManagementConfigurationWindowsSkus)()
}

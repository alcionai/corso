package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingsable 
type DeviceManagementSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidDeviceAdministratorEnrollmentEnabled()(*bool)
    GetDerivedCredentialProvider()(*DerivedCredentialProviderType)
    GetDerivedCredentialUrl()(*string)
    GetDeviceComplianceCheckinThresholdDays()(*int32)
    GetDeviceInactivityBeforeRetirementInDay()(*int32)
    GetEnableAutopilotDiagnostics()(*bool)
    GetEnableDeviceGroupMembershipReport()(*bool)
    GetEnableEnhancedTroubleshootingExperience()(*bool)
    GetEnableLogCollection()(*bool)
    GetEnhancedJailBreak()(*bool)
    GetIgnoreDevicesForUnsupportedSettingsEnabled()(*bool)
    GetIsScheduledActionEnabled()(*bool)
    GetOdataType()(*string)
    GetSecureByDefault()(*bool)
    SetAndroidDeviceAdministratorEnrollmentEnabled(value *bool)()
    SetDerivedCredentialProvider(value *DerivedCredentialProviderType)()
    SetDerivedCredentialUrl(value *string)()
    SetDeviceComplianceCheckinThresholdDays(value *int32)()
    SetDeviceInactivityBeforeRetirementInDay(value *int32)()
    SetEnableAutopilotDiagnostics(value *bool)()
    SetEnableDeviceGroupMembershipReport(value *bool)()
    SetEnableEnhancedTroubleshootingExperience(value *bool)()
    SetEnableLogCollection(value *bool)()
    SetEnhancedJailBreak(value *bool)()
    SetIgnoreDevicesForUnsupportedSettingsEnabled(value *bool)()
    SetIsScheduledActionEnabled(value *bool)()
    SetOdataType(value *string)()
    SetSecureByDefault(value *bool)()
}

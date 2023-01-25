package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppProtectionable 
type IosManagedAppProtectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TargetedManagedAppProtectionable
    GetAllowedIosDeviceModels()(*string)
    GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserProtocol()(*string)
    GetCustomDialerAppProtocol()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableProtectionOfManagedOutboundOpenInData()(*bool)
    GetExemptedAppProtocols()([]KeyValuePairable)
    GetExemptedUniversalLinks()([]string)
    GetFaceIdBlocked()(*bool)
    GetFilterOpenInToOnlyManagedApps()(*bool)
    GetManagedUniversalLinks()([]string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningSdkVersion()(*string)
    GetMinimumWipeSdkVersion()(*string)
    GetProtectInboundDataFromUnknownSources()(*bool)
    GetThirdPartyKeyboardsBlocked()(*bool)
    SetAllowedIosDeviceModels(value *string)()
    SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserProtocol(value *string)()
    SetCustomDialerAppProtocol(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableProtectionOfManagedOutboundOpenInData(value *bool)()
    SetExemptedAppProtocols(value []KeyValuePairable)()
    SetExemptedUniversalLinks(value []string)()
    SetFaceIdBlocked(value *bool)()
    SetFilterOpenInToOnlyManagedApps(value *bool)()
    SetManagedUniversalLinks(value []string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningSdkVersion(value *string)()
    SetMinimumWipeSdkVersion(value *string)()
    SetProtectInboundDataFromUnknownSources(value *bool)()
    SetThirdPartyKeyboardsBlocked(value *bool)()
}

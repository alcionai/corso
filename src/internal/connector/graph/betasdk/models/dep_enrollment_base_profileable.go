package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepEnrollmentBaseProfileable 
type DepEnrollmentBaseProfileable interface {
    EnrollmentProfileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleIdDisabled()(*bool)
    GetApplePayDisabled()(*bool)
    GetConfigurationWebUrl()(*bool)
    GetDeviceNameTemplate()(*string)
    GetDiagnosticsDisabled()(*bool)
    GetDisplayToneSetupDisabled()(*bool)
    GetEnabledSkipKeys()([]string)
    GetIsDefault()(*bool)
    GetIsMandatory()(*bool)
    GetLocationDisabled()(*bool)
    GetPrivacyPaneDisabled()(*bool)
    GetProfileRemovalDisabled()(*bool)
    GetRestoreBlocked()(*bool)
    GetScreenTimeScreenDisabled()(*bool)
    GetSiriDisabled()(*bool)
    GetSupervisedModeEnabled()(*bool)
    GetSupportDepartment()(*string)
    GetSupportPhoneNumber()(*string)
    GetTermsAndConditionsDisabled()(*bool)
    GetTouchIdDisabled()(*bool)
    SetAppleIdDisabled(value *bool)()
    SetApplePayDisabled(value *bool)()
    SetConfigurationWebUrl(value *bool)()
    SetDeviceNameTemplate(value *string)()
    SetDiagnosticsDisabled(value *bool)()
    SetDisplayToneSetupDisabled(value *bool)()
    SetEnabledSkipKeys(value []string)()
    SetIsDefault(value *bool)()
    SetIsMandatory(value *bool)()
    SetLocationDisabled(value *bool)()
    SetPrivacyPaneDisabled(value *bool)()
    SetProfileRemovalDisabled(value *bool)()
    SetRestoreBlocked(value *bool)()
    SetScreenTimeScreenDisabled(value *bool)()
    SetSiriDisabled(value *bool)()
    SetSupervisedModeEnabled(value *bool)()
    SetSupportDepartment(value *string)()
    SetSupportPhoneNumber(value *string)()
    SetTermsAndConditionsDisabled(value *bool)()
    SetTouchIdDisabled(value *bool)()
}

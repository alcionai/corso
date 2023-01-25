package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfigurationable 
type Windows10EnrollmentCompletionPageConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeviceResetOnInstallFailure()(*bool)
    GetAllowDeviceUseOnInstallFailure()(*bool)
    GetAllowLogCollectionOnInstallFailure()(*bool)
    GetAllowNonBlockingAppInstallation()(*bool)
    GetBlockDeviceSetupRetryByUser()(*bool)
    GetCustomErrorMessage()(*string)
    GetDisableUserStatusTrackingAfterFirstUser()(*bool)
    GetInstallProgressTimeoutInMinutes()(*int32)
    GetInstallQualityUpdates()(*bool)
    GetSelectedMobileAppIds()([]string)
    GetShowInstallationProgress()(*bool)
    GetTrackInstallProgressForAutopilotOnly()(*bool)
    SetAllowDeviceResetOnInstallFailure(value *bool)()
    SetAllowDeviceUseOnInstallFailure(value *bool)()
    SetAllowLogCollectionOnInstallFailure(value *bool)()
    SetAllowNonBlockingAppInstallation(value *bool)()
    SetBlockDeviceSetupRetryByUser(value *bool)()
    SetCustomErrorMessage(value *string)()
    SetDisableUserStatusTrackingAfterFirstUser(value *bool)()
    SetInstallProgressTimeoutInMinutes(value *int32)()
    SetInstallQualityUpdates(value *bool)()
    SetSelectedMobileAppIds(value []string)()
    SetShowInstallationProgress(value *bool)()
    SetTrackInstallProgressForAutopilotOnly(value *bool)()
}

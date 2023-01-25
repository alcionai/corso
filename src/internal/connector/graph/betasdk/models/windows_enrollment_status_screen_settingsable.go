package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsEnrollmentStatusScreenSettingsable 
type WindowsEnrollmentStatusScreenSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDeviceUseBeforeProfileAndAppInstallComplete()(*bool)
    GetAllowDeviceUseOnInstallFailure()(*bool)
    GetAllowLogCollectionOnInstallFailure()(*bool)
    GetBlockDeviceSetupRetryByUser()(*bool)
    GetCustomErrorMessage()(*string)
    GetHideInstallationProgress()(*bool)
    GetInstallProgressTimeoutInMinutes()(*int32)
    GetOdataType()(*string)
    SetAllowDeviceUseBeforeProfileAndAppInstallComplete(value *bool)()
    SetAllowDeviceUseOnInstallFailure(value *bool)()
    SetAllowLogCollectionOnInstallFailure(value *bool)()
    SetBlockDeviceSetupRetryByUser(value *bool)()
    SetCustomErrorMessage(value *string)()
    SetHideInstallationProgress(value *bool)()
    SetInstallProgressTimeoutInMinutes(value *int32)()
    SetOdataType(value *string)()
}

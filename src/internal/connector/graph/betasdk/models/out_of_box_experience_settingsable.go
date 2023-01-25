package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutOfBoxExperienceSettingsable 
type OutOfBoxExperienceSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceUsageType()(*WindowsDeviceUsageType)
    GetHideEscapeLink()(*bool)
    GetHideEULA()(*bool)
    GetHidePrivacySettings()(*bool)
    GetOdataType()(*string)
    GetSkipKeyboardSelectionPage()(*bool)
    GetUserType()(*WindowsUserType)
    SetDeviceUsageType(value *WindowsDeviceUsageType)()
    SetHideEscapeLink(value *bool)()
    SetHideEULA(value *bool)()
    SetHidePrivacySettings(value *bool)()
    SetOdataType(value *string)()
    SetSkipKeyboardSelectionPage(value *bool)()
    SetUserType(value *WindowsUserType)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LoginPageTextVisibilitySettingsable 
type LoginPageTextVisibilitySettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHideAccountResetCredentials()(*bool)
    GetHideCannotAccessYourAccount()(*bool)
    GetHideForgotMyPassword()(*bool)
    GetHidePrivacyAndCookies()(*bool)
    GetHideResetItNow()(*bool)
    GetHideTermsOfUse()(*bool)
    GetOdataType()(*string)
    SetHideAccountResetCredentials(value *bool)()
    SetHideCannotAccessYourAccount(value *bool)()
    SetHideForgotMyPassword(value *bool)()
    SetHidePrivacyAndCookies(value *bool)()
    SetHideResetItNow(value *bool)()
    SetHideTermsOfUse(value *bool)()
    SetOdataType(value *string)()
}

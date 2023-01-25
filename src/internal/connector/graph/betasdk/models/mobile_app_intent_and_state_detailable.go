package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppIntentAndStateDetailable 
type MobileAppIntentAndStateDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetInstallState()(*ResultantAppState)
    GetMobileAppIntent()(*MobileAppIntent)
    GetOdataType()(*string)
    GetSupportedDeviceTypes()([]MobileAppSupportedDeviceTypeable)
    SetApplicationId(value *string)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetInstallState(value *ResultantAppState)()
    SetMobileAppIntent(value *MobileAppIntent)()
    SetOdataType(value *string)()
    SetSupportedDeviceTypes(value []MobileAppSupportedDeviceTypeable)()
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDisplayScreenConfigurationable 
type TeamworkDisplayScreenConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBacklightBrightness()(*int32)
    GetBacklightTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsHighContrastEnabled()(*bool)
    GetIsScreensaverEnabled()(*bool)
    GetOdataType()(*string)
    GetScreensaverTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetBacklightBrightness(value *int32)()
    SetBacklightTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsHighContrastEnabled(value *bool)()
    SetIsScreensaverEnabled(value *bool)()
    SetOdataType(value *string)()
    SetScreensaverTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}

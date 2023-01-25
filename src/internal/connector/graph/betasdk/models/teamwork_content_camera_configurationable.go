package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkContentCameraConfigurationable 
type TeamworkContentCameraConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsContentCameraInverted()(*bool)
    GetIsContentCameraOptional()(*bool)
    GetIsContentEnhancementEnabled()(*bool)
    GetOdataType()(*string)
    SetIsContentCameraInverted(value *bool)()
    SetIsContentCameraOptional(value *bool)()
    SetIsContentEnhancementEnabled(value *bool)()
    SetOdataType(value *string)()
}

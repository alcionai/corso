package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkMicrophoneConfigurationable 
type TeamworkMicrophoneConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultMicrophone()(TeamworkPeripheralable)
    GetIsMicrophoneOptional()(*bool)
    GetMicrophones()([]TeamworkPeripheralable)
    GetOdataType()(*string)
    SetDefaultMicrophone(value TeamworkPeripheralable)()
    SetIsMicrophoneOptional(value *bool)()
    SetMicrophones(value []TeamworkPeripheralable)()
    SetOdataType(value *string)()
}

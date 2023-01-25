package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSpeakerConfigurationable 
type TeamworkSpeakerConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultCommunicationSpeaker()(TeamworkPeripheralable)
    GetDefaultSpeaker()(TeamworkPeripheralable)
    GetIsCommunicationSpeakerOptional()(*bool)
    GetIsSpeakerOptional()(*bool)
    GetOdataType()(*string)
    GetSpeakers()([]TeamworkPeripheralable)
    SetDefaultCommunicationSpeaker(value TeamworkPeripheralable)()
    SetDefaultSpeaker(value TeamworkPeripheralable)()
    SetIsCommunicationSpeakerOptional(value *bool)()
    SetIsSpeakerOptional(value *bool)()
    SetOdataType(value *string)()
    SetSpeakers(value []TeamworkPeripheralable)()
}

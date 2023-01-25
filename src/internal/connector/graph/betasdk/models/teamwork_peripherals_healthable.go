package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkPeripheralsHealthable 
type TeamworkPeripheralsHealthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable)
    GetContentCameraHealth()(TeamworkPeripheralHealthable)
    GetDisplayHealthCollection()([]TeamworkPeripheralHealthable)
    GetMicrophoneHealth()(TeamworkPeripheralHealthable)
    GetOdataType()(*string)
    GetRoomCameraHealth()(TeamworkPeripheralHealthable)
    GetSpeakerHealth()(TeamworkPeripheralHealthable)
    SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)()
    SetContentCameraHealth(value TeamworkPeripheralHealthable)()
    SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)()
    SetMicrophoneHealth(value TeamworkPeripheralHealthable)()
    SetOdataType(value *string)()
    SetRoomCameraHealth(value TeamworkPeripheralHealthable)()
    SetSpeakerHealth(value TeamworkPeripheralHealthable)()
}

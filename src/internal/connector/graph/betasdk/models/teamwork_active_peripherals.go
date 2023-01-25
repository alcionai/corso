package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkActivePeripherals 
type TeamworkActivePeripherals struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The communicationSpeaker property
    communicationSpeaker TeamworkPeripheralable
    // The contentCamera property
    contentCamera TeamworkPeripheralable
    // The microphone property
    microphone TeamworkPeripheralable
    // The OdataType property
    odataType *string
    // The roomCamera property
    roomCamera TeamworkPeripheralable
    // The speaker property
    speaker TeamworkPeripheralable
}
// NewTeamworkActivePeripherals instantiates a new teamworkActivePeripherals and sets the default values.
func NewTeamworkActivePeripherals()(*TeamworkActivePeripherals) {
    m := &TeamworkActivePeripherals{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkActivePeripheralsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkActivePeripheralsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkActivePeripherals(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkActivePeripherals) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCommunicationSpeaker gets the communicationSpeaker property value. The communicationSpeaker property
func (m *TeamworkActivePeripherals) GetCommunicationSpeaker()(TeamworkPeripheralable) {
    return m.communicationSpeaker
}
// GetContentCamera gets the contentCamera property value. The contentCamera property
func (m *TeamworkActivePeripherals) GetContentCamera()(TeamworkPeripheralable) {
    return m.contentCamera
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkActivePeripherals) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["communicationSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["contentCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["microphone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["roomCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["speaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    return res
}
// GetMicrophone gets the microphone property value. The microphone property
func (m *TeamworkActivePeripherals) GetMicrophone()(TeamworkPeripheralable) {
    return m.microphone
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkActivePeripherals) GetOdataType()(*string) {
    return m.odataType
}
// GetRoomCamera gets the roomCamera property value. The roomCamera property
func (m *TeamworkActivePeripherals) GetRoomCamera()(TeamworkPeripheralable) {
    return m.roomCamera
}
// GetSpeaker gets the speaker property value. The speaker property
func (m *TeamworkActivePeripherals) GetSpeaker()(TeamworkPeripheralable) {
    return m.speaker
}
// Serialize serializes information the current object
func (m *TeamworkActivePeripherals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeaker", m.GetCommunicationSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCamera", m.GetContentCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphone", m.GetMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("roomCamera", m.GetRoomCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speaker", m.GetSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkActivePeripherals) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCommunicationSpeaker sets the communicationSpeaker property value. The communicationSpeaker property
func (m *TeamworkActivePeripherals) SetCommunicationSpeaker(value TeamworkPeripheralable)() {
    m.communicationSpeaker = value
}
// SetContentCamera sets the contentCamera property value. The contentCamera property
func (m *TeamworkActivePeripherals) SetContentCamera(value TeamworkPeripheralable)() {
    m.contentCamera = value
}
// SetMicrophone sets the microphone property value. The microphone property
func (m *TeamworkActivePeripherals) SetMicrophone(value TeamworkPeripheralable)() {
    m.microphone = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkActivePeripherals) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoomCamera sets the roomCamera property value. The roomCamera property
func (m *TeamworkActivePeripherals) SetRoomCamera(value TeamworkPeripheralable)() {
    m.roomCamera = value
}
// SetSpeaker sets the speaker property value. The speaker property
func (m *TeamworkActivePeripherals) SetSpeaker(value TeamworkPeripheralable)() {
    m.speaker = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkPeripheralsHealth 
type TeamworkPeripheralsHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The health details about the communication speaker.
    communicationSpeakerHealth TeamworkPeripheralHealthable
    // The health details about the content camera.
    contentCameraHealth TeamworkPeripheralHealthable
    // The health details about displays.
    displayHealthCollection []TeamworkPeripheralHealthable
    // The health details about the microphone.
    microphoneHealth TeamworkPeripheralHealthable
    // The OdataType property
    odataType *string
    // The health details about the room camera.
    roomCameraHealth TeamworkPeripheralHealthable
    // The health details about the speaker.
    speakerHealth TeamworkPeripheralHealthable
}
// NewTeamworkPeripheralsHealth instantiates a new teamworkPeripheralsHealth and sets the default values.
func NewTeamworkPeripheralsHealth()(*TeamworkPeripheralsHealth) {
    m := &TeamworkPeripheralsHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkPeripheralsHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkPeripheralsHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkPeripheralsHealth(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralsHealth) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCommunicationSpeakerHealth gets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable) {
    return m.communicationSpeakerHealth
}
// GetContentCameraHealth gets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) GetContentCameraHealth()(TeamworkPeripheralHealthable) {
    return m.contentCameraHealth
}
// GetDisplayHealthCollection gets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) GetDisplayHealthCollection()([]TeamworkPeripheralHealthable) {
    return m.displayHealthCollection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralsHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["communicationSpeakerHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["contentCameraHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["displayHealthCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralHealthable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkPeripheralHealthable)
            }
            m.SetDisplayHealthCollection(res)
        }
        return nil
    }
    res["microphoneHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneHealth(val.(TeamworkPeripheralHealthable))
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
    res["roomCameraHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["speakerHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    return res
}
// GetMicrophoneHealth gets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) GetMicrophoneHealth()(TeamworkPeripheralHealthable) {
    return m.microphoneHealth
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralsHealth) GetOdataType()(*string) {
    return m.odataType
}
// GetRoomCameraHealth gets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) GetRoomCameraHealth()(TeamworkPeripheralHealthable) {
    return m.roomCameraHealth
}
// GetSpeakerHealth gets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) GetSpeakerHealth()(TeamworkPeripheralHealthable) {
    return m.speakerHealth
}
// Serialize serializes information the current object
func (m *TeamworkPeripheralsHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeakerHealth", m.GetCommunicationSpeakerHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCameraHealth", m.GetContentCameraHealth())
        if err != nil {
            return err
        }
    }
    if m.GetDisplayHealthCollection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDisplayHealthCollection()))
        for i, v := range m.GetDisplayHealthCollection() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("displayHealthCollection", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphoneHealth", m.GetMicrophoneHealth())
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
        err := writer.WriteObjectValue("roomCameraHealth", m.GetRoomCameraHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speakerHealth", m.GetSpeakerHealth())
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
func (m *TeamworkPeripheralsHealth) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCommunicationSpeakerHealth sets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)() {
    m.communicationSpeakerHealth = value
}
// SetContentCameraHealth sets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) SetContentCameraHealth(value TeamworkPeripheralHealthable)() {
    m.contentCameraHealth = value
}
// SetDisplayHealthCollection sets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)() {
    m.displayHealthCollection = value
}
// SetMicrophoneHealth sets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) SetMicrophoneHealth(value TeamworkPeripheralHealthable)() {
    m.microphoneHealth = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralsHealth) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRoomCameraHealth sets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) SetRoomCameraHealth(value TeamworkPeripheralHealthable)() {
    m.roomCameraHealth = value
}
// SetSpeakerHealth sets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) SetSpeakerHealth(value TeamworkPeripheralHealthable)() {
    m.speakerHealth = value
}

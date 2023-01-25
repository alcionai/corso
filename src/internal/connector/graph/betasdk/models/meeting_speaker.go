package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingSpeaker 
type MeetingSpeaker struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Bio of the speaker.
    bio *string
    // Display name of the speaker.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewMeetingSpeaker instantiates a new meetingSpeaker and sets the default values.
func NewMeetingSpeaker()(*MeetingSpeaker) {
    m := &MeetingSpeaker{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingSpeakerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingSpeakerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingSpeaker(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingSpeaker) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBio gets the bio property value. Bio of the speaker.
func (m *MeetingSpeaker) GetBio()(*string) {
    return m.bio
}
// GetDisplayName gets the displayName property value. Display name of the speaker.
func (m *MeetingSpeaker) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingSpeaker) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBio(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingSpeaker) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MeetingSpeaker) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bio", m.GetBio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingSpeaker) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBio sets the bio property value. Bio of the speaker.
func (m *MeetingSpeaker) SetBio(value *string)() {
    m.bio = value
}
// SetDisplayName sets the displayName property value. Display name of the speaker.
func (m *MeetingSpeaker) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingSpeaker) SetOdataType(value *string)() {
    m.odataType = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Media 
type Media struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
    isTranscriptionShown *bool
    // Information about the source of media. Read-only.
    mediaSource MediaSourceable
    // The OdataType property
    odataType *string
}
// NewMedia instantiates a new media and sets the default values.
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMedia(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Media) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isTranscriptionShown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranscriptionShown(val)
        }
        return nil
    }
    res["mediaSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSource(val.(MediaSourceable))
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
// GetIsTranscriptionShown gets the isTranscriptionShown property value. If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
func (m *Media) GetIsTranscriptionShown()(*bool) {
    return m.isTranscriptionShown
}
// GetMediaSource gets the mediaSource property value. Information about the source of media. Read-only.
func (m *Media) GetMediaSource()(MediaSourceable) {
    return m.mediaSource
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Media) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Media) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isTranscriptionShown", m.GetIsTranscriptionShown())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaSource", m.GetMediaSource())
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
func (m *Media) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsTranscriptionShown sets the isTranscriptionShown property value. If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
func (m *Media) SetIsTranscriptionShown(value *bool)() {
    m.isTranscriptionShown = value
}
// SetMediaSource sets the mediaSource property value. Information about the source of media. Read-only.
func (m *Media) SetMediaSource(value MediaSourceable)() {
    m.mediaSource = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Media) SetOdataType(value *string)() {
    m.odataType = value
}

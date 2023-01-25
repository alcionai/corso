package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttachmentContentProperties 
type AttachmentContentProperties struct {
    ContentProperties
    // The currentLabel property
    currentLabel CurrentLabelable
}
// NewAttachmentContentProperties instantiates a new AttachmentContentProperties and sets the default values.
func NewAttachmentContentProperties()(*AttachmentContentProperties) {
    m := &AttachmentContentProperties{
        ContentProperties: *NewContentProperties(),
    }
    odataTypeValue := "#microsoft.graph.attachmentContentProperties";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAttachmentContentPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttachmentContentPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttachmentContentProperties(), nil
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *AttachmentContentProperties) GetCurrentLabel()(CurrentLabelable) {
    return m.currentLabel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttachmentContentProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ContentProperties.GetFieldDeserializers()
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(CurrentLabelable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AttachmentContentProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ContentProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *AttachmentContentProperties) SetCurrentLabel(value CurrentLabelable)() {
    m.currentLabel = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MentionsPreview 
type MentionsPreview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
    isMentioned *bool
    // The OdataType property
    odataType *string
}
// NewMentionsPreview instantiates a new mentionsPreview and sets the default values.
func NewMentionsPreview()(*MentionsPreview) {
    m := &MentionsPreview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMentionsPreviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMentionsPreviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMentionsPreview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MentionsPreview) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MentionsPreview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isMentioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMentioned(val)
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
// GetIsMentioned gets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
func (m *MentionsPreview) GetIsMentioned()(*bool) {
    return m.isMentioned
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MentionsPreview) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MentionsPreview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isMentioned", m.GetIsMentioned())
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
func (m *MentionsPreview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsMentioned sets the isMentioned property value. True if the signed-in user is mentioned in the parent resource instance. Read-only. Supports filter.
func (m *MentionsPreview) SetIsMentioned(value *bool)() {
    m.isMentioned = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MentionsPreview) SetOdataType(value *string)() {
    m.odataType = value
}

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaSource 
type MediaSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Enumeration value that indicates the media content category.
    contentCategory *MediaSourceContentCategory
    // The OdataType property
    odataType *string
}
// NewMediaSource instantiates a new mediaSource and sets the default values.
func NewMediaSource()(*MediaSource) {
    m := &MediaSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMediaSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaSource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaSource) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentCategory gets the contentCategory property value. Enumeration value that indicates the media content category.
func (m *MediaSource) GetContentCategory()(*MediaSourceContentCategory) {
    return m.contentCategory
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaSourceContentCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCategory(val.(*MediaSourceContentCategory))
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
func (m *MediaSource) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MediaSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetContentCategory() != nil {
        cast := (*m.GetContentCategory()).String()
        err := writer.WriteStringValue("contentCategory", &cast)
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
func (m *MediaSource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentCategory sets the contentCategory property value. Enumeration value that indicates the media content category.
func (m *MediaSource) SetContentCategory(value *MediaSourceContentCategory)() {
    m.contentCategory = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaSource) SetOdataType(value *string)() {
    m.odataType = value
}

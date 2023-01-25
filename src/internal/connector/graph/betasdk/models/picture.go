package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Picture provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Picture struct {
    Entity
    // The content property
    content []byte
    // The contentType property
    contentType *string
    // The height property
    height *int32
    // The width property
    width *int32
}
// NewPicture instantiates a new picture and sets the default values.
func NewPicture()(*Picture) {
    m := &Picture{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePictureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePictureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPicture(), nil
}
// GetContent gets the content property value. The content property
func (m *Picture) GetContent()([]byte) {
    return m.content
}
// GetContentType gets the contentType property value. The contentType property
func (m *Picture) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Picture) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. The height property
func (m *Picture) GetHeight()(*int32) {
    return m.height
}
// GetWidth gets the width property value. The width property
func (m *Picture) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *Picture) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *Picture) SetContent(value []byte)() {
    m.content = value
}
// SetContentType sets the contentType property value. The contentType property
func (m *Picture) SetContentType(value *string)() {
    m.contentType = value
}
// SetHeight sets the height property value. The height property
func (m *Picture) SetHeight(value *int32)() {
    m.height = value
}
// SetWidth sets the width property value. The width property
func (m *Picture) SetWidth(value *int32)() {
    m.width = value
}

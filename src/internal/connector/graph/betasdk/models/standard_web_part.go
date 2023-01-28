package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StandardWebPart 
type StandardWebPart struct {
    WebPart
    // Data of the webPart.
    data WebPartDataable
    // A Guid which indicates the type of the webParts
    webPartType *string
}
// NewStandardWebPart instantiates a new StandardWebPart and sets the default values.
func NewStandardWebPart()(*StandardWebPart) {
    m := &StandardWebPart{
        WebPart: *NewWebPart(),
    }
    odataTypeValue := "#microsoft.graph.standardWebPart";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateStandardWebPartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStandardWebPartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStandardWebPart(), nil
}
// GetData gets the data property value. Data of the webPart.
func (m *StandardWebPart) GetData()(WebPartDataable) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StandardWebPart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WebPart.GetFieldDeserializers()
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebPartDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val.(WebPartDataable))
        }
        return nil
    }
    res["webPartType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebPartType(val)
        }
        return nil
    }
    return res
}
// GetWebPartType gets the webPartType property value. A Guid which indicates the type of the webParts
func (m *StandardWebPart) GetWebPartType()(*string) {
    return m.webPartType
}
// Serialize serializes information the current object
func (m *StandardWebPart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WebPart.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webPartType", m.GetWebPartType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetData sets the data property value. Data of the webPart.
func (m *StandardWebPart) SetData(value WebPartDataable)() {
    m.data = value
}
// SetWebPartType sets the webPartType property value. A Guid which indicates the type of the webParts
func (m *StandardWebPart) SetWebPartType(value *string)() {
    m.webPartType = value
}

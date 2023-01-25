package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddWatermark 
type AddWatermark struct {
    MarkContent
    // The orientation property
    orientation *PageOrientation
}
// NewAddWatermark instantiates a new AddWatermark and sets the default values.
func NewAddWatermark()(*AddWatermark) {
    m := &AddWatermark{
        MarkContent: *NewMarkContent(),
    }
    odataTypeValue := "#microsoft.graph.addWatermark";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAddWatermarkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddWatermarkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddWatermark(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddWatermark) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MarkContent.GetFieldDeserializers()
    res["orientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePageOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrientation(val.(*PageOrientation))
        }
        return nil
    }
    return res
}
// GetOrientation gets the orientation property value. The orientation property
func (m *AddWatermark) GetOrientation()(*PageOrientation) {
    return m.orientation
}
// Serialize serializes information the current object
func (m *AddWatermark) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MarkContent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOrientation() != nil {
        cast := (*m.GetOrientation()).String()
        err = writer.WriteStringValue("orientation", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrientation sets the orientation property value. The orientation property
func (m *AddWatermark) SetOrientation(value *PageOrientation)() {
    m.orientation = value
}

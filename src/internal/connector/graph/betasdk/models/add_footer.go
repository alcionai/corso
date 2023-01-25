package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddFooter 
type AddFooter struct {
    MarkContent
    // The alignment property
    alignment *Alignment
    // The margin property
    margin *int32
}
// NewAddFooter instantiates a new AddFooter and sets the default values.
func NewAddFooter()(*AddFooter) {
    m := &AddFooter{
        MarkContent: *NewMarkContent(),
    }
    odataTypeValue := "#microsoft.graph.addFooter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAddFooterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddFooterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAddFooter(), nil
}
// GetAlignment gets the alignment property value. The alignment property
func (m *AddFooter) GetAlignment()(*Alignment) {
    return m.alignment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddFooter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MarkContent.GetFieldDeserializers()
    res["alignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlignment)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlignment(val.(*Alignment))
        }
        return nil
    }
    res["margin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMargin(val)
        }
        return nil
    }
    return res
}
// GetMargin gets the margin property value. The margin property
func (m *AddFooter) GetMargin()(*int32) {
    return m.margin
}
// Serialize serializes information the current object
func (m *AddFooter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MarkContent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlignment() != nil {
        cast := (*m.GetAlignment()).String()
        err = writer.WriteStringValue("alignment", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("margin", m.GetMargin())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlignment sets the alignment property value. The alignment property
func (m *AddFooter) SetAlignment(value *Alignment)() {
    m.alignment = value
}
// SetMargin sets the margin property value. The margin property
func (m *AddFooter) SetMargin(value *int32)() {
    m.margin = value
}

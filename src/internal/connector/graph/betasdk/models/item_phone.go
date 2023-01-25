package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPhone 
type ItemPhone struct {
    ItemFacet
    // Friendly name the user has assigned this phone number.
    displayName *string
    // Phone number provided by the user.
    number *string
    // The type property
    type_escaped *PhoneType
}
// NewItemPhone instantiates a new ItemPhone and sets the default values.
func NewItemPhone()(*ItemPhone) {
    m := &ItemPhone{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.itemPhone";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemPhoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPhoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPhone(), nil
}
// GetDisplayName gets the displayName property value. Friendly name the user has assigned this phone number.
func (m *ItemPhone) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPhone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
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
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*PhoneType))
        }
        return nil
    }
    return res
}
// GetNumber gets the number property value. Phone number provided by the user.
func (m *ItemPhone) GetNumber()(*string) {
    return m.number
}
// GetType gets the type property value. The type property
func (m *ItemPhone) GetType()(*PhoneType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *ItemPhone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Friendly name the user has assigned this phone number.
func (m *ItemPhone) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNumber sets the number property value. Phone number provided by the user.
func (m *ItemPhone) SetNumber(value *string)() {
    m.number = value
}
// SetType sets the type property value. The type property
func (m *ItemPhone) SetType(value *PhoneType)() {
    m.type_escaped = value
}

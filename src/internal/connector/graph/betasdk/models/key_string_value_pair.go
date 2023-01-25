package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeyStringValuePair 
type KeyStringValuePair struct {
    KeyTypedValuePair
    // The string value of the key-value pair.
    value *string
}
// NewKeyStringValuePair instantiates a new KeyStringValuePair and sets the default values.
func NewKeyStringValuePair()(*KeyStringValuePair) {
    m := &KeyStringValuePair{
        KeyTypedValuePair: *NewKeyTypedValuePair(),
    }
    odataTypeValue := "#microsoft.graph.keyStringValuePair";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateKeyStringValuePairFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeyStringValuePairFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeyStringValuePair(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeyStringValuePair) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.KeyTypedValuePair.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The string value of the key-value pair.
func (m *KeyStringValuePair) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *KeyStringValuePair) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.KeyTypedValuePair.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The string value of the key-value pair.
func (m *KeyStringValuePair) SetValue(value *string)() {
    m.value = value
}

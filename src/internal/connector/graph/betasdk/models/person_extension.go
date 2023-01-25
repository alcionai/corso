package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonExtension 
type PersonExtension struct {
    Extension
}
// NewPersonExtension instantiates a new PersonExtension and sets the default values.
func NewPersonExtension()(*PersonExtension) {
    m := &PersonExtension{
        Extension: *NewExtension(),
    }
    odataTypeValue := "#microsoft.graph.personExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Extension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *PersonExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Extension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}

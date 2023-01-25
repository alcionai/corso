package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserConfiguration provides operations to manage the collection of site entities.
type UserConfiguration struct {
    Entity
    // The binaryData property
    binaryData []byte
}
// NewUserConfiguration instantiates a new userConfiguration and sets the default values.
func NewUserConfiguration()(*UserConfiguration) {
    m := &UserConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserConfiguration(), nil
}
// GetBinaryData gets the binaryData property value. The binaryData property
func (m *UserConfiguration) GetBinaryData()([]byte) {
    return m.binaryData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["binaryData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBinaryData(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("binaryData", m.GetBinaryData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBinaryData sets the binaryData property value. The binaryData property
func (m *UserConfiguration) SetBinaryData(value []byte)() {
    m.binaryData = value
}

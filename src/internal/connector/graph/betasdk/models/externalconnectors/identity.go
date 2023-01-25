package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Identity provides operations to call the add method.
type Identity struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The type property
    type_escaped *IdentityType
}
// NewIdentity instantiates a new identity and sets the default values.
func NewIdentity()(*Identity) {
    m := &Identity{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Identity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*IdentityType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type property
func (m *Identity) GetType()(*IdentityType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *Identity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
// SetType sets the type property value. The type property
func (m *Identity) SetType(value *IdentityType)() {
    m.type_escaped = value
}
